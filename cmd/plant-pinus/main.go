package main

import (
	"crypto/sha256"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const sourceFile = "plants-Pinus.P450.doc"
const sourceURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Pinus.P450.doc"
const sourceHash = "2c081245df38f4ee2bc6e6c00d4cb0610a1207b930c7b70ec3b5e40377ed0298"

type record struct {
	Index, Line                         int
	Species, Symbol, ID, Note, Sequence string
}

var accessionLine = regexp.MustCompile(`^([A-Z]{1,2}\d{5,6}(?:\.1)?)\s+(.+)$`)

func parse(text string) ([]record, error) {
	lines := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n"), "\n")
	if len(lines) < 140 || strings.TrimSpace(lines[8]) != "74 P450 ESTs from Pinus" || strings.TrimSpace(lines[12]) != "4 full length sequences are now available from mRNAs or assembled from ESTs" {
		return nil, fmt.Errorf("Pinus preamble changed")
	}
	type fullBlock struct {
		start, end          int
		species, symbol, id string
	}
	blocks := []fullBlock{
		{14, 24, "Pinus taeda", "CYP73A20", "AF096998"},
		{25, 36, "Pinus taeda", "CYP73A23", "CYP73A23-assembled-ESTs"},
		{37, 48, "Pinus radiata", "CYP78A4", "AF049067"},
		{49, 60, "Pinus taeda", "CYP98A15", "AY064170"},
	}
	var out []record
	for _, b := range blocks {
		header := strings.TrimSpace(lines[b.start])
		if !strings.HasPrefix(header, b.symbol) {
			return nil, fmt.Errorf("full-sequence header line %d changed: %q", b.start+1, header)
		}
		var parts []string
		firstSequenceLine := b.start + 1
		if b.symbol == "CYP73A23" {
			firstSequenceLine++ // source header wraps its accession list once
		}
		for i := firstSequenceLine; i < b.end; i++ {
			v := strings.Join(strings.Fields(lines[i]), "")
			if v == "" {
				continue
			}
			for _, aa := range v {
				if !strings.ContainsRune("ACDEFGHIKLMNPQRSTVWXY", aa) {
					return nil, fmt.Errorf("full sequence line %d contains %q", i+1, aa)
				}
			}
			parts = append(parts, v)
		}
		seq := strings.Join(parts, "")
		if seq == "" {
			return nil, fmt.Errorf("empty full sequence line %d", b.start+1)
		}
		out = append(out, record{len(out) + 1, b.start + 1, b.species, b.symbol, b.id, header, seq})
	}

	// The source's exact 74-line EST catalogue follows the four complete
	// proteins. It contains accession/similarity annotations, not standalone
	// proteins. Later BLAST Sbjct rows are alignment fragments and are never
	// promoted to sequences for these catalogue entries.
	for i := 62; i <= 138; i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" || strings.HasPrefix(line, "all four ") || strings.HasPrefix(line, "have the first ") || strings.HasPrefix(line, "in the middle") {
			continue
		}
		m := accessionLine.FindStringSubmatch(line)
		if len(m) != 3 {
			return nil, fmt.Errorf("EST catalogue line %d changed: %q", i+1, line)
		}
		out = append(out, record{len(out) + 1, i + 1, "Pinus", "", m[1], m[2], ""})
	}
	if len(out) != 78 {
		return nil, fmt.Errorf("Pinus records=%d, want 78", len(out))
	}
	return out, nil
}

func main() {
	in := flag.String("input", filepath.Join("raw", "plants-Pinus.P450.txt"), "normalized Pinus source")
	doc := flag.String("source-doc", filepath.Join("raw", sourceFile), "Word source")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "pinus.csv"), "reviewed CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-pinus.md"), "audit")
	flag.Parse()
	d, err := os.ReadFile(*in)
	if err != nil {
		panic(err)
	}
	records, err := parse(string(d))
	if err != nil {
		panic(err)
	}
	hash, err := fileHash(*doc)
	if err != nil || hash != sourceHash {
		panic(fmt.Errorf("Pinus source hash changed: %s: %v", hash, err))
	}
	if err = writeCSV(*out, records, hash); err != nil {
		panic(err)
	}
	if err = writeAudit(*audit, records, hash); err != nil {
		panic(err)
	}
	fmt.Printf("Pinus: %d records accepted (%d literal full proteins)\n", len(records), sequenceCount(records))
}

func writeCSV(path string, records []record, hash string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	_ = w.Write([]string{"category", "species", "symbol", "id", "record_key", "source_url", "source_note", "sequence", "source_file", "source_sha256", "source_record", "source_line", "review_status"})
	for _, r := range records {
		status := "sequence-missing: source EST catalogue supplies annotation only; BLAST alignment snippets rejected"
		if r.Sequence != "" {
			status = "literal full protein"
			if strings.ContainsAny(r.Sequence, "X") {
				status += ";ambiguous-X"
			}
		}
		_ = w.Write([]string{"plants", r.Species, r.Symbol, r.ID, fmt.Sprintf("pinus:record-%04d:%s", r.Index, r.ID), sourceURL, r.Note, r.Sequence, sourceFile, hash, strconv.Itoa(r.Index), strconv.Itoa(r.Line), status})
	}
	w.Flush()
	return w.Error()
}

func writeAudit(path string, records []record, hash string) error {
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Pinus\n\n- Source file: `%s`\n- URL: %s\n- Source SHA-256: `%s`\n- Source declaration: `74 P450 ESTs from Pinus`; `4 full length sequences`\n- Accepted source records: `%d`\n- Records with literal complete protein: `%d`\n- Sequence-missing EST catalogue records: `%d`\n- Review status: `complete`\n\n", sourceFile, sourceURL, hash, len(records), sequenceCount(records), len(records)-sequenceCount(records))
	b.WriteString("## Resource-specific interpretation\n\nThis is explicitly the old Pinus collection. Its first four named blocks are complete literal proteins: three *Pinus taeda* proteins and one *Pinus radiata* protein. The following exact 74-line catalogue lists EST accessions and similarity annotations but no standalone protein for each accession, so those source records remain independently searchable with empty sequence. The long remainder consists of BLAST Query/Sbjct alignments and repeats of the four reference proteins. Alignment fragments are not FASTA proteins and are not concatenated, inferred, translated, or imported. The repeated CYP98A/73A20/73A23/78A4 search references do not create duplicates.\n\n| Source record | Species | Symbol | ID | Sequence |\n|---:|---|---|---|---:|\n")
	for _, i := range []int{0, 1, 2, 3, 4, 40, 77} {
		r := records[i]
		fmt.Fprintf(&b, "| %d | %s | %s | %s | %d aa |\n", r.Index, r.Species, r.Symbol, r.ID, len(r.Sequence))
	}
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func sequenceCount(records []record) int {
	n := 0
	for _, r := range records {
		if r.Sequence != "" {
			n++
		}
	}
	return n
}
func fileHash(path string) (string, error) {
	d, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha256.Sum256(d)), nil
}
