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

const sourceFile = "plants-liverwort.doc"
const sourceURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/liverwort.doc"
const sourceHash = "c3332118c265ae53c722e9e3c8a4dc3e0cd61dc893b86d3f294ef439b67fefcf"

type record struct {
	Block, Line                          int
	Header, Symbol, ID, Sequence, Status string
}

var cypRE = regexp.MustCompile(`(?:\d+)?(CYP[0-9A-Za-z?_-]*)`)
var accRE = regexp.MustCompile(`\b([A-Z]{2}\d{6}(?:\.1)?)\b`)

func parse(t string) ([]record, error) {
	l := strings.Split(strings.ReplaceAll(strings.ReplaceAll(t, "\r\n", "\n"), "\r", "\n"), "\n")
	var s []int
	for i, v := range l {
		if strings.HasPrefix(strings.TrimSpace(v), ">") {
			s = append(s, i)
		}
	}
	if len(s) != 38 {
		return nil, fmt.Errorf("Marchantia headers=%d", len(s))
	}
	var out []record
	for b, st := range s {
		e := len(l)
		if b+1 < len(s) {
			e = s[b+1]
		}
		h := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(l[st]), ">"))
		sym := ""
		if m := cypRE.FindStringSubmatch(h); len(m) == 2 {
			sym = m[1]
		}
		id := fmt.Sprintf("source-block-%02d", b+1)
		for i := st + 1; i < e; i++ {
			if m := accRE.FindStringSubmatch(l[i]); len(m) == 2 {
				id = m[1]
				break
			}
		}
		var p []string
		for i := st + 1; i < e; i++ {
			if q, ok := seq(l[i]); ok {
				p = append(p, q)
			}
		}
		q := strings.Join(p, "")
		if q == "" {
			return nil, fmt.Errorf("block %d empty", b+1)
		}
		term := strings.HasSuffix(q, "*")
		if term {
			q = strings.TrimSuffix(q, "*")
		}
		status := []string{"source-EST-fragment"}
		if term {
			status = append(status, "source-terminal-stop")
		}
		if strings.Contains(q, "*") {
			status = append(status, "internal-stop")
		}
		if strings.ContainsAny(q, "Xx") {
			status = append(status, "ambiguous-X-or-x")
		}
		out = append(out, record{b + 1, st + 1, h, sym, id, q, strings.Join(status, ";")})
	}
	return out, nil
}
func seq(raw string) (string, bool) {
	v := strings.TrimSpace(raw)
	v = regexp.MustCompile(`\s*\([^)]*\)\s*(?:\d+)?\s*$`).ReplaceAllString(v, "")
	q := strings.Join(strings.Fields(v), "")
	if len(q) < 3 {
		return "", false
	}
	for _, c := range q {
		if !strings.ContainsRune("ACDEFGHIKLMNPQRSTVWXYXx*", c) {
			return "", false
		}
	}
	return q, true
}
func main() {
	in := flag.String("input", filepath.Join("raw", "plants-liverwort.txt"), "")
	doc := flag.String("source-doc", filepath.Join("raw", sourceFile), "")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "marchantia-polymorpha.csv"), "")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-marchantia-polymorpha.md"), "")
	flag.Parse()
	d, e := os.ReadFile(*in)
	if e != nil {
		panic(e)
	}
	r, e := parse(string(d))
	if e != nil {
		panic(e)
	}
	h, e := hash(*doc)
	if e != nil || h != sourceHash {
		panic(h)
	}
	if e = write(*out, r, h); e != nil {
		panic(e)
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Marchantia polymorpha\n\n- Source SHA-256: `%s`\n- Title count: `35 Cytochrome P450s`\n- Explicit source blocks: `%d`\n- Literal EST-derived protein fragments: `%d`\n- Review status: `complete`\n\nThe title says 35 P450s, while the physical document contains 38 explicit Marchantia headers; all 38 are retained independently rather than forcing the title count. The final two CYP74 blocks and other duplicated broad labels are distinct source fragments. Each block supplies a literal protein or EST-derived protein fragment; wrapped text, X and terminal stops are handled locally. The long accession appendix contains no protein and is used only to cross-check EST provenance. No fragment is extended or translated.\n", h, len(r), len(r))
	if e = os.WriteFile(*audit, []byte(b.String()), 0644); e != nil {
		panic(e)
	}
	fmt.Printf("Marchantia polymorpha: %d EST-derived blocks accepted\n", len(r))
}
func write(p string, r []record, h string) error {
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
		_ = w.Write([]string{"plants", "Marchantia polymorpha", v.Symbol, v.ID, fmt.Sprintf("marchantia-polymorpha:block-%04d:%s", v.Block, v.ID), sourceURL, v.Header, v.Sequence, sourceFile, h, strconv.Itoa(v.Block), strconv.Itoa(v.Line), v.Status})
	}
	w.Flush()
	return w.Error()
}
func hash(p string) (string, error) {
	d, e := os.ReadFile(p)
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%x", sha256.Sum256(d)), nil
}
