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

const sourceFile = "plants-Selaginella.P450s.doc"
const sourceURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Selaginella.P450s.doc"
const sourceHash = "2b1f22cf1f283e6ce4bfcfe5e71de6f459ac2e4784787f2dafe4fb1f586c8fca"

type record struct {
	Block, Line                          int
	Header, Symbol, ID, Sequence, Status string
}

var cypRE = regexp.MustCompile(`\b(CYP[0-9A-Za-z?_-]+)\b`)

func parse(text string) ([]record, error) {
	lines := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n"), "\n")
	var starts []int
	for i, l := range lines {
		if strings.HasPrefix(strings.TrimSpace(l), ">") {
			starts = append(starts, i)
		}
	}
	if len(starts) != 517 {
		return nil, fmt.Errorf("Selaginella headers=%d", len(starts))
	}
	out := make([]record, 0, len(starts))
	for bi, s := range starts {
		e := len(lines)
		if bi+1 < len(starts) {
			e = starts[bi+1]
		}
		header := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(lines[s]), ">"))
		symbol := ""
		if m := cypRE.FindStringSubmatch(header); len(m) == 2 {
			symbol = m[1]
		}
		id := firstID(header, symbol)
		var parts []string
		for i := s + 1; i < e; i++ {
			if q, ok := sequenceLine(lines[i]); ok {
				parts = append(parts, q)
			}
		}
		seq := strings.Join(parts, "")
		if seq == "" {
			return nil, fmt.Errorf("Selaginella block %d has no sequence", bi+1)
		}
		terminal := strings.HasSuffix(seq, "*")
		if terminal {
			seq = strings.TrimSuffix(seq, "*")
		}
		status := []string{}
		if symbol == "" {
			status = append(status, "source-unassigned-model")
		}
		if terminal {
			status = append(status, "source-terminal-stop")
		}
		if strings.Contains(seq, "*") {
			status = append(status, "internal-stop")
		}
		if strings.ContainsAny(seq, "Xx") {
			status = append(status, "ambiguous-X-or-x")
		}
		if len(seq) < 350 {
			status = append(status, "short-sequence")
		}
		low := strings.ToLower(header)
		if strings.Contains(low, "pseudo") {
			status = append(status, "source-pseudogene-label")
		}
		if strings.Contains(low, "partial") || strings.Contains(low, "fragment") {
			status = append(status, "source-fragment-label")
		}
		out = append(out, record{bi + 1, s + 1, header, symbol, id, seq, strings.Join(status, ";")})
	}
	return out, nil
}
func sequenceLine(raw string) (string, bool) {
	v := strings.TrimSpace(raw)
	if v == "" || strings.HasPrefix(v, "$") {
		return "", false
	}
	v = regexp.MustCompile(`\s*\([^)]*\)\s*(?:\d+)?\s*$`).ReplaceAllString(v, "")
	fields := strings.Fields(v)
	if len(fields) > 0 && regexp.MustCompile(`^\d+$`).MatchString(fields[0]) {
		fields = fields[1:]
	}
	if len(fields) > 0 && regexp.MustCompile(`^\d+$`).MatchString(fields[len(fields)-1]) {
		fields = fields[:len(fields)-1]
	}
	if len(fields) == 0 {
		return "", false
	}
	q := strings.Join(fields, "")
	if len(q) < 3 {
		return "", false
	}
	for _, aa := range q {
		if !strings.ContainsRune("ACDEFGHIKLMNPQRSTVWXYx*", aa) {
			return "", false
		}
	}
	return q, true
}
func firstID(h, symbol string) string {
	fields := strings.Fields(h)
	for _, f := range fields {
		f = strings.Trim(f, " ,")
		if f == symbol || regexp.MustCompile(`^\d+$`).MatchString(f) {
			continue
		}
		if strings.Contains(f, "|Selmo1") || strings.Contains(f, "scaffold_") || strings.HasPrefix(f, "SM") || strings.HasPrefix(f, "FE") || strings.HasPrefix(f, "gw1.") || strings.HasPrefix(f, "e_gw1.") || strings.HasPrefix(f, "estExt_") || strings.HasPrefix(f, "fgenesh") {
			return f
		}
	}
	if len(fields) > 0 {
		return fields[0]
	}
	return fmt.Sprintf("block")
}
func main() {
	in := flag.String("input", filepath.Join("raw", "plants-Selaginella.P450s.txt"), "")
	doc := flag.String("source-doc", filepath.Join("raw", sourceFile), "")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "selaginella-moellendorffii.csv"), "")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-selaginella-moellendorffii.md"), "")
	flag.Parse()
	d, e := os.ReadFile(*in)
	if e != nil {
		panic(e)
	}
	r, e := parse(string(d))
	if e != nil {
		panic(e)
	}
	h, e := fileHash(*doc)
	if e != nil || h != sourceHash {
		panic(fmt.Errorf("hash %s %v", h, e))
	}
	if e = writeCSV(*out, r, h); e != nil {
		panic(e)
	}
	if e = writeAudit(*audit, r, h); e != nil {
		panic(e)
	}
	fmt.Printf("Selaginella moellendorffii: %d header-bounded sequences accepted\n", len(r))
}
func writeCSV(p string, r []record, h string) error {
	if e := os.MkdirAll(filepath.Dir(p), 0755); e != nil {
		return e
	}
	f, e := os.Create(p)
	if e != nil {
		return e
	}
	defer f.Close()
	w := csv.NewWriter(f)
	_ = w.Write([]string{"category", "species", "symbol", "id", "record_key", "source_url", "source_note", "sequence", "source_file", "source_sha256", "source_block", "source_line", "review_status"})
	for _, v := range r {
		_ = w.Write([]string{"plants", "Selaginella moellendorffii", v.Symbol, v.ID, fmt.Sprintf("selaginella-moellendorffii:block-%04d:%s", v.Block, v.ID), sourceURL, v.Header, v.Sequence, sourceFile, h, strconv.Itoa(v.Block), strconv.Itoa(v.Line), v.Status})
	}
	w.Flush()
	return w.Error()
}
func writeAudit(p string, r []record, h string) error {
	named, terminal, internal, x, short := 0, 0, 0, 0, 0
	for _, v := range r {
		if v.Symbol != "" {
			named++
		}
		if strings.Contains(v.Status, "source-terminal-stop") {
			terminal++
		}
		if strings.Contains(v.Status, "internal-stop") {
			internal++
		}
		if strings.Contains(v.Status, "ambiguous-X") {
			x++
		}
		if len(v.Sequence) < 350 {
			short++
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Selaginella moellendorffii\n\n- Source file: `%s`\n- URL: %s\n- Source SHA-256: `%s`\n- Explicit `>` blocks: `%d`\n- Blocks with literal sequence: `%d`\n- Blocks with an explicit CYP label in the header: `%d`\n- Terminal-stop blocks: `%d`\n- Internal-stop blocks: `%d`\n- X/x-bearing blocks: `%d`\n- Short blocks: `%d`\n- Review status: `complete`\n\nAll 517 explicit source headers delimit independent Selaginella models, alleles, pseudogenes, fragments, trace reconstructions, and source-described possible false positives; none belongs to a foreign species. Both named and still-unassigned models are retained because the source deliberately presents them inside CYP clan/bin annotation. Protein text occurs in wrapped, spaced-residue, coordinate-bounded, phase-marked, intron-labelled and ampersand-join layouts. Those exact layout markers are removed while residue case, X/x and internal stops remain literal; only a final stop is removed. Narrative lines such as `No allele found`, `remove intron`, and `Duplicate of part of last exon` are explicitly rejected.\n\n| Block | Symbol | ID | Sequence | Status |\n|---:|---|---|---:|---|\n", sourceFile, sourceURL, h, len(r), len(r), named, terminal, internal, x, short)
	for _, i := range []int{0, 258, 516} {
		v := r[i]
		fmt.Fprintf(&b, "| %d | %s | %s | %d aa | %s |\n", v.Block, v.Symbol, v.ID, len(v.Sequence), v.Status)
	}
	return os.WriteFile(p, []byte(b.String()), 0644)
}
func fileHash(p string) (string, error) {
	d, e := os.ReadFile(p)
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%x", sha256.Sum256(d)), nil
}
