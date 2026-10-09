package main

import (
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
)

const sourceFile = "plants-Brachypodium.FASTA.doc"
const sourceURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Brachypodium.FASTA.doc"
const sourceHash = "d10a7ca25a08dc4f98e8a86a925b401b6f31337f59d37f89107f78e6f0894d9a"

type record struct {
	Block, SourceLine, FirstSequenceLine, LastSequenceLine int
	Header, ID, Symbol, RecordKey, Sequence                string
	Status, Annotations                                    []string
}
type excluded struct {
	Block, SourceLine int
	Header, Reason    string
}

var leadingCoordinate = regexp.MustCompile(`^\d+\s+`)
var trailingCoordinate = regexp.MustCompile(`\s+\d+\s*&?\s*$`)
var phaseMarker = regexp.MustCompile(`\((?:0|1|2|\?)?\)`)
var whitespace = regexp.MustCompile(`\s+`)
var cypRE = regexp.MustCompile(`^CYP[^\s]+`)
var modelRE = regexp.MustCompile(`\bBradi[0-9A-Za-z.]+\b`)
var nonKey = regexp.MustCompile(`[^A-Za-z0-9._-]+`)

func main() {
	input := flag.String("input", filepath.Join("raw", "plants-Brachypodium.FASTA.txt"), "Word-normalized Brachypodium text")
	doc := flag.String("source-doc", filepath.Join("raw", sourceFile), "original Word file")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "brachypodium-distachyon.csv"), "reviewed CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-brachypodium-distachyon.md"), "audit")
	flag.Parse()
	d, err := os.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	rows, x, err := parseBrachypodium(string(d))
	if err != nil {
		panic(err)
	}
	h, err := fileHash(*doc)
	if err != nil {
		panic(err)
	}
	if h != sourceHash {
		panic(fmt.Errorf("Brachypodium hash=%s", h))
	}
	if err = writeCSV(*out, rows, h); err != nil {
		panic(err)
	}
	if err = writeAudit(*audit, rows, x, h); err != nil {
		panic(err)
	}
	fmt.Printf("Brachypodium distachyon: %d literal sequence blocks accepted, %d foreign block excluded\n", len(rows), len(x))
}

func parseBrachypodium(text string) ([]record, []excluded, error) {
	lines := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n"), "\n")
	var hs []int
	for i, l := range lines {
		if strings.HasPrefix(strings.TrimSpace(l), ">") {
			hs = append(hs, i)
		}
	}
	if len(hs) != 279 {
		return nil, nil, fmt.Errorf("Brachypodium headers=%d", len(hs))
	}
	var rows []record
	var x []excluded
	for bi, start := range hs {
		end := len(lines)
		if bi+1 < len(hs) {
			end = hs[bi+1]
		}
		header := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(lines[start]), ">"))
		if header == "CYP71AM1 Sorghum bicolor XM_002451987" {
			x = append(x, excluded{bi + 1, start + 1, header, "explicit Sorghum bicolor comparison sequence"})
			continue
		}
		symbol := cypRE.FindString(header)
		if symbol == "" {
			return nil, nil, fmt.Errorf("block %d missing CYP", bi+1)
		}
		id := modelRE.FindString(header)
		if id == "" {
			id = symbol
		}
		r := record{Block: bi + 1, SourceLine: start + 1, Header: header, ID: id, Symbol: symbol, RecordKey: fmt.Sprintf("brachypodium-distachyon:block-%04d:%s", bi+1, keyPart(id))}
		var parts []string
		for i := start + 1; i < end; i++ {
			if seq, ok := brachypodiumSequenceLine(lines[i]); ok {
				if r.FirstSequenceLine == 0 {
					r.FirstSequenceLine = i + 1
				}
				r.LastSequenceLine = i + 1
				parts = append(parts, seq)
			} else if v := strings.TrimSpace(lines[i]); v != "" && len(r.Annotations) < 12 {
				r.Annotations = append(r.Annotations, v)
			}
		}
		r.Sequence = strings.Join(parts, "")
		if strings.HasSuffix(r.Sequence, "*") {
			r.Sequence = strings.TrimSuffix(r.Sequence, "*")
			r.Status = append(r.Status, "source-terminal-stop")
		}
		if r.Sequence == "" {
			return nil, nil, fmt.Errorf("block %d has no sequence", bi+1)
		}
		blockText := strings.ToLower(strings.Join(lines[start:end], "\n"))
		flags := map[string]bool{}
		if strings.Contains(blockText, "pseudo") || strings.HasSuffix(strings.Split(symbol, "-")[0], "P") {
			flags["source-pseudogene"] = true
		}
		for _, p := range []string{"fragment", "missing", "deletion", "no gene model", "n-term", "c-term"} {
			if strings.Contains(blockText, p) {
				flags["source-fragment-or-missing-region"] = true
			}
		}
		if strings.Contains(strings.Join(lines[start:end], "\n"), "&") {
			flags["source-joined-piece"] = true
		}
		if strings.Contains(r.Sequence, "*") {
			flags["internal-stop"] = true
		}
		if strings.ContainsAny(r.Sequence, "Xx") {
			flags["ambiguous-X-or-x"] = true
		}
		if strings.Contains(r.Sequence, "-") {
			flags["source-gap"] = true
		}
		if len(r.Sequence) < 350 {
			flags["short-sequence"] = true
		}
		for _, aa := range r.Sequence {
			if !strings.ContainsRune("ACDEFGHIKLMNOPQRSTVWXYx*-", aa) {
				return nil, nil, fmt.Errorf("block %d residue %q", bi+1, aa)
			}
		}
		for _, n := range []string{"source-pseudogene", "source-fragment-or-missing-region", "source-joined-piece", "internal-stop", "ambiguous-X-or-x", "source-gap", "short-sequence"} {
			if flags[n] {
				r.Status = append(r.Status, n)
			}
		}
		rows = append(rows, r)
	}
	if len(rows) != 278 || len(x) != 1 {
		return nil, nil, fmt.Errorf("accepted=%d excluded=%d", len(rows), len(x))
	}
	return rows, x, nil
}

func brachypodiumSequenceLine(raw string) (string, bool) {
	v := strings.TrimSpace(raw)
	if v == "" {
		return "", false
	}
	for _, r := range v {
		if (r >= 'a' && r <= 'w') || r == 'y' || r == 'z' {
			return "", false
		}
	}
	v = leadingCoordinate.ReplaceAllString(v, "")
	v = trailingCoordinate.ReplaceAllString(v, "")
	v = phaseMarker.ReplaceAllString(v, "")
	v = strings.ReplaceAll(v, "&", "")
	v = whitespace.ReplaceAllString(v, "")
	if v == "" {
		return "", false
	}
	for _, aa := range v {
		if !strings.ContainsRune("ACDEFGHIKLMNOPQRSTVWXYx*-", aa) {
			return "", false
		}
	}
	return v, true
}
func writeCSV(path string, rows []record, hash string) error {
	if e := os.MkdirAll(filepath.Dir(path), 0755); e != nil {
		return e
	}
	f, e := os.Create(path)
	if e != nil {
		return e
	}
	defer f.Close()
	w := csv.NewWriter(f)
	_ = w.Write([]string{"category", "species", "symbol", "id", "record_key", "source_url", "source_note", "sequence", "source_file", "source_sha256", "source_block", "source_line", "source_header", "review_status"})
	for _, r := range rows {
		note := fmt.Sprintf("Brachypodium Word block %d at normalized line %d; literal protein lines %d-%d; header=%s", r.Block, r.SourceLine, r.FirstSequenceLine, r.LastSequenceLine, r.Header)
		if len(r.Annotations) > 0 {
			note += "; annotations: " + strings.Join(r.Annotations, " / ")
		}
		_ = w.Write([]string{"plants", "Brachypodium distachyon", r.Symbol, r.ID, r.RecordKey, sourceURL, note, r.Sequence, sourceFile, hash, strconv.Itoa(r.Block), strconv.Itoa(r.SourceLine), r.Header, strings.Join(r.Status, ";")})
	}
	w.Flush()
	return w.Error()
}
func writeAudit(path string, rows []record, x []excluded, hash string) error {
	counts := map[string]int{}
	seqs := map[string]int{}
	for _, r := range rows {
		seqs[r.Sequence]++
		for _, s := range r.Status {
			counts[s]++
		}
	}
	names := []string{}
	for n := range counts {
		names = append(names, n)
	}
	sort.Strings(names)
	dupes := 0
	for _, n := range seqs {
		if n > 1 {
			dupes++
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Brachypodium distachyon\n\n- Source file: `%s`\n- URL: %s\n- SHA-256: `%s`\n- Container inspected: legacy Word, 2,652 paragraphs, 48 pages, no tables\n- Source `>` blocks: `279`\n- Accepted Brachypodium literal sequence blocks: `%d`\n- Explicit foreign blocks excluded: `%d`\n- Duplicate sequence groups retained: `%d`\n- Review status: `complete`\n\nAll header-delimited blocks were inspected. Block 74 is explicitly `Sorghum bicolor` and is excluded; the other 278 CYP blocks belong to the Brachypodium resource, including no-model fragments and pseudogenes. This working annotation document embeds genomic coordinates, phase markers, `&` joins and deletion prose among literal residue lines. Only those reviewed layout tokens are removed. X/x, gaps and internal stops remain literal, and only a final `*` is removed. Nothing is repaired, translated, merged, or externally completed.\n\n## Status counts\n\n| Status | Records |\n|---|---:|\n", sourceFile, sourceURL, hash, len(rows), len(x), dupes)
	for _, n := range names {
		fmt.Fprintf(&b, "| %s | %d |\n", n, counts[n])
	}
	b.WriteString("\n## Representative records\n\n| Accepted index | Source block | Line | ID | CYP | Length |\n|---:|---:|---:|---|---|---:|\n")
	for _, i := range []int{0, len(rows) / 2, len(rows) - 1} {
		r := rows[i]
		fmt.Fprintf(&b, "| %d | %d | %d | %s | %s | %d |\n", i+1, r.Block, r.SourceLine, r.ID, r.Symbol, len(r.Sequence))
	}
	fmt.Fprintf(&b, "\nExcluded block %d at line %d: `%s` — %s.\n", x[0].Block, x[0].SourceLine, x[0].Header, x[0].Reason)
	return os.WriteFile(path, []byte(b.String()), 0644)
}
func fileHash(path string) (string, error) {
	d, e := os.ReadFile(path)
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%x", sha256.Sum256(d)), nil
}
func keyPart(v string) string {
	return strings.ToLower(strings.Trim(nonKey.ReplaceAllString(v, "-"), "-"))
}
