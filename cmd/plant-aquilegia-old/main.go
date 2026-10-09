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

const sourceFile = "plants-Aquilegia.P450s.doc"
const sourceURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Aquilegia.P450s.doc"
const sourceHash = "f6de5f1ace0cfc95318675a8bb8e9e8b09f393fdb303bec281d509a9911cac58"

type record struct {
	Block, SourceLine, FirstSequenceLine, LastSequenceLine int
	Header, ID, Model, Symbol, RecordKey, Sequence         string
	Status, Annotations                                    []string
}

var headerID = regexp.MustCompile(`^([0-9]+)\s+`)
var headerModel = regexp.MustCompile(`\|([^|]+)\|([^|]+)$`)
var symbolRE = regexp.MustCompile(`\bCYP[0-9A-Za-z]+(?:[-.]?[A-Za-z0-9]+)*`)
var whitespace = regexp.MustCompile(`\s+`)
var phaseMarker = regexp.MustCompile(`\((?:0|1|2|\?)?\)`)
var nonKey = regexp.MustCompile(`[^A-Za-z0-9._-]+`)

func main() {
	input := flag.String("input", filepath.Join("raw", "plants-Aquilegia.P450s.txt"), "Word-normalized old Aquilegia text")
	doc := flag.String("source-doc", filepath.Join("raw", sourceFile), "original Word file")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "aquilegia-coerulea-old.csv"), "reviewed CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-aquilegia-coerulea-old.md"), "audit")
	flag.Parse()
	d, e := os.ReadFile(*input)
	if e != nil {
		panic(e)
	}
	r, e := parseAquilegiaOld(string(d))
	if e != nil {
		panic(e)
	}
	if len(r) != 472 {
		panic(fmt.Errorf("old Aquilegia count=%d", len(r)))
	}
	h, e := fileHash(*doc)
	if e != nil {
		panic(e)
	}
	if h != sourceHash {
		panic(fmt.Errorf("old Aquilegia hash=%s", h))
	}
	if e = writeCSV(*out, r, h); e != nil {
		panic(e)
	}
	if e = writeAudit(*audit, r, h); e != nil {
		panic(e)
	}
	fmt.Printf("Aquilegia coerulea old Word release: %d literal sequence blocks accepted\n", len(r))
}

func parseAquilegiaOld(text string) ([]record, error) {
	text = strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n")
	lines := strings.Split(text, "\n")
	var hs []int
	for i, l := range lines {
		if strings.HasPrefix(strings.TrimSpace(l), ">") {
			hs = append(hs, i)
		}
	}
	if len(hs) != 472 {
		return nil, fmt.Errorf("old Aquilegia headers=%d", len(hs))
	}
	rs := make([]record, 0, 472)
	for bi, start := range hs {
		end := len(lines)
		if bi+1 < len(hs) {
			end = hs[bi+1]
		}
		header := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(lines[start]), ">"))
		m := headerID.FindStringSubmatch(header)
		if len(m) != 2 || !strings.Contains(header, "Aquilegia coerulea|") {
			return nil, fmt.Errorf("block %d unexpected header %q", bi+1, header)
		}
		r := record{Block: bi + 1, SourceLine: start + 1, Header: header, ID: m[1]}
		if mm := headerModel.FindStringSubmatch(header); len(mm) == 3 {
			r.Model = mm[2]
		}
		r.RecordKey = fmt.Sprintf("aquilegia-coerulea-old:block-%04d:%s", r.Block, keyPart(r.ID))
		var parts []string
		for i := start + 1; i < end; i++ {
			if r.Symbol == "" {
				r.Symbol = symbolRE.FindString(strings.TrimSpace(lines[i]))
			}
			if seq, ok := oldSequenceLine(lines[i]); ok {
				if r.FirstSequenceLine == 0 {
					r.FirstSequenceLine = i + 1
				}
				r.LastSequenceLine = i + 1
				parts = append(parts, seq)
			} else if v := strings.TrimSpace(lines[i]); v != "" && len(r.Annotations) < 5 {
				r.Annotations = append(r.Annotations, v)
			}
		}
		raw := strings.Join(parts, "")
		if strings.HasSuffix(raw, "*") {
			raw = strings.TrimSuffix(raw, "*")
			r.Status = append(r.Status, "source-terminal-stop")
		}
		r.Sequence = raw
		if r.Symbol == "" {
			// Seven source annotation lines use the literal broad label
			// `like 76/80`; preserve it rather than inventing a family.
			if len(r.Annotations) > 0 && strings.EqualFold(r.Annotations[0], "like 76/80") {
				r.Symbol = "CYP76/80-like"
			}
			if len(r.Annotations) > 0 && strings.EqualFold(r.Annotations[0], "new cluster A like 716/718") {
				r.Symbol = "CYP716/718-like cluster A"
			}
		}
		if r.Symbol == "" || r.Sequence == "" {
			return nil, fmt.Errorf("block %d missing symbol/sequence", r.Block)
		}
		flags := map[string]bool{}
		if strings.Contains(strings.ToLower(strings.Join(lines[start:end], "\n")), "pseudo") {
			flags["source-pseudogene"] = true
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
				return nil, fmt.Errorf("block %d residue %q", r.Block, aa)
			}
		}
		for _, n := range []string{"source-pseudogene", "internal-stop", "ambiguous-X-or-x", "source-gap", "short-sequence"} {
			if flags[n] {
				r.Status = append(r.Status, n)
			}
		}
		rs = append(rs, r)
	}
	return rs, nil
}
func oldSequenceLine(raw string) (string, bool) {
	v := strings.TrimSpace(raw)
	if v == "" {
		return "", false
	}
	for _, r := range v {
		if (r >= 'a' && r <= 'w') || r == 'y' || r == 'z' {
			return "", false
		}
	}
	v = phaseMarker.ReplaceAllString(v, "")
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
func writeCSV(path string, rs []record, hash string) error {
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
	for _, r := range rs {
		note := fmt.Sprintf("older Aquilegia Word block %d at normalized line %d; Phytozome model=%s; literal protein lines %d-%d; header=%s", r.Block, r.SourceLine, r.Model, r.FirstSequenceLine, r.LastSequenceLine, r.Header)
		_ = w.Write([]string{"plants", "Aquilegia coerulea", r.Symbol, r.ID, r.RecordKey, sourceURL, note, r.Sequence, sourceFile, hash, strconv.Itoa(r.Block), strconv.Itoa(r.SourceLine), r.Header, strings.Join(r.Status, ";")})
	}
	w.Flush()
	return w.Error()
}
func writeAudit(path string, rs []record, hash string) error {
	counts := map[string]int{}
	seqs := map[string]int{}
	for _, r := range rs {
		seqs[r.Sequence]++
		for _, s := range r.Status {
			counts[s]++
		}
	}
	names := make([]string, 0, len(counts))
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
	fmt.Fprintf(&b, "# Plant resource review: Aquilegia coerulea old Word release\n\n- Source file: `%s`\n- URL: %s\n- SHA-256: `%s`\n- Container inspected: legacy Word, 2,088 paragraphs, 95 pages, no tables\n- Source `>` blocks: `%d`\n- Accepted Aquilegia blocks with literal sequence: `%d`\n- Duplicate literal-sequence groups retained: `%d`\n- Review status: `complete`\n\nThe source title declares 472 Aquilegia sequences, exactly matching its 472 explicit Aquilegia-labelled headers. Each header carries a numeric Phytozome ID and model path; the following annotation line supplies the source CYP family/subfamily, and subsequent uppercase residue lines supply the protein. All blocks are retained in source order. Whitespace and reviewed phase markers are layout; X/x, gaps and internal stops remain literal, and only a terminal `*` is removed. This older release remains distinct from the current workbook.\n\n## Status counts\n\n| Status | Records |\n|---|---:|\n", sourceFile, sourceURL, hash, len(rs), len(rs), dupes)
	for _, n := range names {
		fmt.Fprintf(&b, "| %s | %d |\n", n, counts[n])
	}
	b.WriteString("\n## Representative records\n\n| Index | Block | Line | ID | Model | CYP | Length |\n|---:|---:|---:|---|---|---|---:|\n")
	for _, i := range []int{0, len(rs) / 2, len(rs) - 1} {
		r := rs[i]
		fmt.Fprintf(&b, "| %d | %d | %d | %s | %s | %s | %d |\n", i+1, r.Block, r.SourceLine, r.ID, r.Model, r.Symbol, len(r.Sequence))
	}
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
	v = nonKey.ReplaceAllString(strings.TrimSpace(v), "-")
	return strings.ToLower(strings.Trim(v, "-"))
}
