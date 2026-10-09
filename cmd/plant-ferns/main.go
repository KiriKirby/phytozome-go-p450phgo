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

const sourceFile = "plants-ferns.doc"
const sourceURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/ferns.doc"
const sourceHash = "0e030a3b69e8812adfd0cdf264c4866dee0ea9bafc9231eafcda9626aa1be321"

type record struct {
	Block, Line                                   int
	Species, Symbol, ID, Header, Sequence, Status string
}

var headerRE = regexp.MustCompile(`^(CYP\S*|CYP)\s+(Ceratopteris richardii|Adiantum capillus-veneris)\b`)
var accessionRE = regexp.MustCompile(`\b([A-Z]{1,2}\d{6}(?:\.1)?)\b`)
var coordinateLineRE = regexp.MustCompile(`^\d+\s+([A-Z]+)(?:\s+\d+)?(?:\s+([A-Z]+))?(?:\s+\d+)?$`)

func parse(text string) ([]record, error) {
	lines := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n"), "\n")
	var starts []int
	for i, raw := range lines {
		if headerRE.MatchString(strings.TrimSpace(raw)) {
			starts = append(starts, i)
		}
	}
	if len(starts) != 18 {
		return nil, fmt.Errorf("fern headers=%d", len(starts))
	}
	var out []record
	for bi, start := range starts {
		end := len(lines)
		if bi+1 < len(starts) {
			end = starts[bi+1]
		}
		header := strings.TrimSpace(lines[start])
		hm := headerRE.FindStringSubmatch(header)
		species := hm[2]
		symbol := hm[1]
		id := ""
		var parts []string
		terminal := false
		for i := start + 1; i < end; i++ {
			line := strings.TrimSpace(lines[i])
			if id == "" {
				if m := accessionRE.FindStringSubmatch(line); len(m) == 2 {
					id = m[1]
				}
			}
			seq := ""
			if line != "" && only(line, "ACDEFGHIKLMNPQRSTVWXY*") {
				seq = line
			} else if m := coordinateLineRE.FindStringSubmatch(line); len(m) > 0 {
				seq = m[1] + m[2]
			}
			if seq != "" {
				seq = strings.ReplaceAll(seq, " ", "")
				if strings.HasSuffix(seq, "*") {
					terminal = true
					seq = strings.TrimSuffix(seq, "*")
				}
				parts = append(parts, seq)
			}
		}
		if id == "" {
			return nil, fmt.Errorf("fern block %d lacks accession", bi+1)
		}
		seq := strings.Join(parts, "")
		if seq == "" {
			return nil, fmt.Errorf("fern block %d lacks literal sequence", bi+1)
		}
		status := "literal source fragment"
		if terminal {
			status += ";source-terminal-stop"
		}
		if strings.ContainsAny(seq, "X") {
			status += ";ambiguous-X"
		}
		out = append(out, record{bi + 1, start + 1, species, symbol, id, header, seq, status})
	}
	return out, nil
}
func only(s, alphabet string) bool {
	for _, r := range strings.ReplaceAll(s, " ", "") {
		if !strings.ContainsRune(alphabet, r) {
			return false
		}
	}
	return strings.TrimSpace(s) != ""
}

func main() {
	in := flag.String("input", filepath.Join("raw", "plants-ferns.txt"), "")
	doc := flag.String("source-doc", filepath.Join("raw", sourceFile), "")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "ferns.csv"), "")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-ferns.md"), "")
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
		panic(fmt.Errorf("fern hash %s %v", h, e))
	}
	if e = writeCSV(*out, r, h); e != nil {
		panic(e)
	}
	if e = writeAudit(*audit, r, h); e != nil {
		panic(e)
	}
	fmt.Printf("ferns: %d literal source fragments accepted\n", len(r))
}
func writeCSV(path string, r []record, h string) error {
	if e := os.MkdirAll(filepath.Dir(path), 0755); e != nil {
		return e
	}
	f, e := os.Create(path)
	if e != nil {
		return e
	}
	defer f.Close()
	w := csv.NewWriter(f)
	_ = w.Write([]string{"category", "species", "symbol", "id", "record_key", "source_url", "source_note", "sequence", "source_file", "source_sha256", "source_block", "source_line", "review_status"})
	for _, v := range r {
		_ = w.Write([]string{"plants", v.Species, v.Symbol, v.ID, fmt.Sprintf("ferns:block-%04d:%s", v.Block, v.ID), sourceURL, v.Header, v.Sequence, sourceFile, h, strconv.Itoa(v.Block), strconv.Itoa(v.Line), v.Status})
	}
	w.Flush()
	return w.Error()
}
func writeAudit(path string, r []record, h string) error {
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: ferns\n\n- Source file: `%s`\n- URL: %s\n- Source SHA-256: `%s`\n- Accepted blocks: `%d`\n- Accepted literal sequences: `%d`\n- Species: `Ceratopteris richardii`, `Adiantum capillus-veneris`\n- Review status: `complete`\n\nEach of the 18 CYP-heading blocks contains one accession and one literal protein fragment. Ordinary wrapped proteins, terminal-stop records and coordinate-wrapped fragments are parsed only within their own block. Numeric coordinates are removed; literal X remains. Broad source labels such as `CYP`, `CYP72clan`, and question-mark family labels are preserved rather than made more specific from similarity annotations.\n\n| Block | Species | Symbol | ID | Sequence |\n|---:|---|---|---|---:|\n", sourceFile, sourceURL, h, len(r), len(r))
	for _, i := range []int{0, 8, 17} {
		v := r[i]
		fmt.Fprintf(&b, "| %d | %s | %s | %s | %d aa |\n", v.Block, v.Species, v.Symbol, v.ID, len(v.Sequence))
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
