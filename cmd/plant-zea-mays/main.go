package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode/utf16"
)

const sourceFile = "plants-zea.doc"
const sourceURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/zea.doc"
const sourceHash = "96f8a5f71ed55f34ce1dba75949c71c892d3115ec3a140b8c2ab42ec7302a9bd"

type record struct {
	Block, SourceLine, FirstSequenceLine, LastSequenceLine int
	Header, ID, Symbol, RecordKey, Sequence                string
	Status                                                 []string
}

var accessionRE = regexp.MustCompile(`\b(?:AI|T)\d+\b`)
var leadingCoordinateRE = regexp.MustCompile(`^\d+\s+`)
var trailingCoordinateRE = regexp.MustCompile(`\s+\d+$`)

// This table is deliberately specific to the 26 displayed Zea groups in this
// 1999 document. Compound similarity assignments stay compound; no accession
// is assigned a more specific CYP name than the source supplies.
var blockSymbols = []string{
	"CYP51", "CYP51", "CYP71A5/CYP71E1", "CYP71C2", "CYP71C3", "CYP71C3",
	"CYP71C3", "CYP71C3", "CYP71D10", "CYP71D7", "CYP71E1/CYP71B18", "CYP72A5",
	"CYP72A5", "CYP72A14", "CYP73A7", "CYP78A1", "CYP81A4", "CYP81A3",
	"CYP81A3", "CYP90A1", "CYP90C1", "CYP97B2", "CYP98A1", "CYP706A5",
	"CYP707A3", "CYP714A2",
}

func main() {
	input := flag.String("input", filepath.Join("raw", "plants-zea.txt"), "Word-normalized Zea text")
	doc := flag.String("source-doc", filepath.Join("raw", sourceFile), "original Word file")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "zea-mays.csv"), "reviewed CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-zea-mays.md"), "audit")
	flag.Parse()
	d, err := os.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	text, err := decodeZeaText(d)
	if err != nil {
		panic(err)
	}
	rows, indexIDs, err := parseZea(text)
	if err != nil {
		panic(err)
	}
	hash, err := fileHash(*doc)
	if err != nil {
		panic(err)
	}
	if hash != sourceHash {
		panic(fmt.Errorf("Zea source hash=%s", hash))
	}
	if err = writeCSV(*out, rows, hash); err != nil {
		panic(err)
	}
	if err = writeAudit(*audit, rows, indexIDs, hash); err != nil {
		panic(err)
	}
	fmt.Printf("Zea mays release: %d source groups, %d literal protein fragments, %d alphabetical accessions\n", len(rows), countSequences(rows), len(indexIDs))
}

func decodeZeaText(data []byte) (string, error) {
	if len(data) >= 2 && bytes.Equal(data[:2], []byte{0xff, 0xfe}) {
		if (len(data)-2)%2 != 0 {
			return "", fmt.Errorf("odd UTF-16LE byte count")
		}
		units := make([]uint16, 0, (len(data)-2)/2)
		for i := 2; i < len(data); i += 2 {
			units = append(units, uint16(data[i])|uint16(data[i+1])<<8)
		}
		return string(utf16.Decode(units)), nil
	}
	return string(bytes.TrimPrefix(data, []byte{0xef, 0xbb, 0xbf})), nil
}

func parseZea(text string) ([]record, []string, error) {
	lines := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n"), "\n")
	indexStart := -1
	var headers []int
	for i, line := range lines {
		if strings.TrimSpace(line) == "ESTs sorted alphabetically" {
			indexStart = i
			break
		}
		if strings.HasPrefix(strings.TrimSpace(line), ">") {
			headers = append(headers, i)
		}
	}
	if indexStart < 0 || len(headers) != 26 {
		return nil, nil, fmt.Errorf("Zea layout headers=%d index=%d", len(headers), indexStart+1)
	}
	rows := make([]record, 0, 26)
	for bi, start := range headers {
		end := indexStart
		if bi+1 < len(headers) {
			end = headers[bi+1]
		}
		header := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(lines[start]), ">"))
		ids := accessionRE.FindAllString(header, -1)
		if len(ids) == 0 {
			return nil, nil, fmt.Errorf("block %d has no accession", bi+1)
		}
		r := record{Block: bi + 1, SourceLine: start + 1, Header: header, ID: ids[0], Symbol: blockSymbols[bi]}
		r.RecordKey = fmt.Sprintf("zea-mays:block-%04d:%s", r.Block, strings.ToLower(r.ID))
		var parts []string
		for i := start + 1; i < end; i++ {
			if seq, ok := zeaSequenceLine(lines[i]); ok {
				if r.FirstSequenceLine == 0 {
					r.FirstSequenceLine = i + 1
				}
				r.LastSequenceLine = i + 1
				parts = append(parts, seq)
			}
		}
		r.Sequence = strings.Join(parts, "")
		if strings.HasSuffix(r.Sequence, "*") {
			r.Sequence = strings.TrimSuffix(r.Sequence, "*")
			r.Status = append(r.Status, "source-terminal-stop")
		}
		if r.Sequence == "" {
			r.Status = append(r.Status, "source-sequence-not-present")
		} else {
			r.Status = append(r.Status, "est-protein-fragment")
		}
		if strings.Contains(r.Sequence, "*") {
			r.Status = append(r.Status, "internal-stop")
		}
		if strings.Contains(r.Sequence, "X") {
			r.Status = append(r.Status, "ambiguous-X")
		}
		if len(ids) > 1 {
			r.Status = append(r.Status, "grouped-accessions")
		}
		rows = append(rows, r)
	}
	if countSequences(rows) != 18 {
		return nil, nil, fmt.Errorf("Zea literal sequence groups=%d", countSequences(rows))
	}
	var indexIDs []string
	seen := map[string]bool{}
	for _, line := range lines[indexStart+1:] {
		m := accessionRE.FindString(strings.TrimSpace(line))
		if m != "" && !seen[m] {
			seen[m] = true
			indexIDs = append(indexIDs, m)
		}
	}
	if len(indexIDs) != 42 {
		return nil, nil, fmt.Errorf("Zea alphabetical accessions=%d", len(indexIDs))
	}
	return rows, indexIDs, nil
}

func zeaSequenceLine(raw string) (string, bool) {
	v := strings.TrimSpace(raw)
	if v == "" {
		return "", false
	}
	v = leadingCoordinateRE.ReplaceAllString(v, "")
	v = trailingCoordinateRE.ReplaceAllString(v, "")
	v = strings.Join(strings.Fields(v), "")
	if v == "" {
		return "", false
	}
	for _, aa := range v {
		if !strings.ContainsRune("ACDEFGHIKLMNPQRSTVWYX*", aa) {
			return "", false
		}
	}
	return v, true
}

func writeCSV(path string, rows []record, hash string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	_ = w.Write([]string{"category", "species", "symbol", "id", "record_key", "source_url", "source_note", "sequence", "source_file", "source_sha256", "source_block", "source_line", "source_header", "review_status"})
	for _, r := range rows {
		span := "no literal protein in source group"
		if r.Sequence != "" {
			span = fmt.Sprintf("literal protein lines %d-%d", r.FirstSequenceLine, r.LastSequenceLine)
		}
		note := fmt.Sprintf("Zea 1999 EST group %d at normalized line %d; %s; complete source header=%s", r.Block, r.SourceLine, span, r.Header)
		_ = w.Write([]string{"plants", "Zea mays", r.Symbol, r.ID, r.RecordKey, sourceURL, note, r.Sequence, sourceFile, hash, strconv.Itoa(r.Block), strconv.Itoa(r.SourceLine), r.Header, strings.Join(r.Status, ";")})
	}
	w.Flush()
	return w.Error()
}

func writeAudit(path string, rows []record, indexIDs []string, hash string) error {
	counts := map[string]int{}
	for _, r := range rows {
		for _, s := range r.Status {
			counts[s]++
		}
	}
	names := make([]string, 0, len(counts))
	for n := range counts {
		names = append(names, n)
	}
	sort.Strings(names)
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Zea mays\n\n- Source file: `%s`\n- URL: %s\n- SHA-256: `%s`\n- Container inspected: legacy Word, 155 paragraphs, 3 pages, no tables\n- Source family-sorted `>` groups: `%d`\n- Accepted source groups: `%d`\n- Groups with literal protein fragments: `%d`\n- Groups without literal protein: `%d`\n- Distinct accessions in alphabetical appendix: `%d`\n- Review status: `complete`\n\nThe title says `39 Zea mays P450 ESTs`, but the document's own alphabetical appendix contains 42 distinct accession rows. The family-sorted body represents those accessions in 26 annotation groups (and mentions the same 42 distinct accessions, including T70647 in the AI622303 comparison). The release therefore preserves the 26 actual source groups instead of inventing one protein per accession. Eighteen groups contain literal coordinate-delimited protein fragments; eight contain annotation only and retain an empty sequence. The alphabetical appendix is an index and creates no duplicate records. X and internal stops remain literal; only a terminal `*` is removed.\n\n## Status counts\n\n| Status | Records |\n|---|---:|\n", sourceFile, sourceURL, hash, len(rows), len(rows), countSequences(rows), len(rows)-countSequences(rows), len(indexIDs))
	for _, n := range names {
		fmt.Fprintf(&b, "| %s | %d |\n", n, counts[n])
	}
	b.WriteString("\n## Representative groups\n\n| Block | Line | ID | CYP assignment | Length | Source header |\n|---:|---:|---|---|---:|---|\n")
	for _, i := range []int{0, 3, 11, 22, 25} {
		r := rows[i]
		fmt.Fprintf(&b, "| %d | %d | %s | %s | %d | %s |\n", r.Block, r.SourceLine, r.ID, r.Symbol, len(r.Sequence), strings.ReplaceAll(r.Header, "|", "\\|"))
	}
	return os.WriteFile(path, []byte(b.String()), 0644)
}

func countSequences(rows []record) int {
	n := 0
	for _, r := range rows {
		if r.Sequence != "" {
			n++
		}
	}
	return n
}
func fileHash(path string) (string, error) {
	d, e := os.ReadFile(path)
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%x", sha256.Sum256(d)), nil
}
