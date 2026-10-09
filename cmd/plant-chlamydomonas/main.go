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

const sourceFile = "plants-chlamydomonas.doc"
const sourceURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/chlamydomonas.doc"
const sourceHash = "5570e540d9e03dc6e33081775c5739185634226c39ad3fac9e976cac40347179"

type record struct {
	Block, Line                          int
	Header, Symbol, ID, Sequence, Status string
}

var symbolRE = regexp.MustCompile(`^(CYP\S+)`)

func isPrimary(v string) bool { return strings.HasPrefix(v, ">CYP") && !strings.HasPrefix(v, ">CYP4F") }
func parse(t string) ([]record, error) {
	l := strings.Split(strings.ReplaceAll(strings.ReplaceAll(t, "\r\n", "\n"), "\r", "\n"), "\n")
	var s []int
	for i, v := range l {
		if isPrimary(strings.TrimSpace(v)) {
			s = append(s, i)
		}
	}
	if len(s) != 41 {
		return nil, fmt.Errorf("Chlamydomonas primary headers=%d", len(s))
	}
	var out []record
	for b, st := range s {
		e := len(l)
		if b+1 < len(s) {
			e = s[b+1]
		}
		h := strings.TrimSpace(strings.TrimPrefix(l[st], ">"))
		sym := symbolRE.FindString(h)
		id := firstModel(h)
		begin := st + 1
		for i := begin; i < e; i++ {
			if strings.Contains(strings.ToLower(l[i]), "newest data:") {
				begin = i + 1
			}
		}
		var p []string
		for i := begin; i < e; i++ {
			if strings.HasPrefix(strings.TrimSpace(l[i]), ">") {
				// Some reviewed blocks begin with a nested model-ID header before
				// their protein. A later nested header after residues begins
				// comparison/trace evidence.
				if len(p) > 0 {
					break
				}
				continue
			}
			if q, ok := seq(l[i]); ok {
				p = append(p, q)
			}
		}
		q := strings.Join(p, "")
		if q == "" {
			return nil, fmt.Errorf("block %d %s has no selected protein", b+1, sym)
		}
		term := strings.HasSuffix(q, "*")
		if term {
			q = strings.TrimSuffix(q, "*")
		}
		status := []string{}
		if term {
			status = append(status, "source-terminal-stop")
		}
		if strings.Contains(q, "*") {
			status = append(status, "internal-stop")
		}
		if strings.ContainsAny(q, "Xx") {
			status = append(status, "ambiguous-X-or-x")
		}
		if strings.Contains(strings.ToLower(h), "pseudo") {
			status = append(status, "source-pseudogene-label")
		}
		out = append(out, record{b + 1, st + 1, h, sym, id, q, strings.Join(status, ";")})
	}
	return out, nil
}
func seq(raw string) (string, bool) {
	v := strings.TrimSpace(raw)
	if v == "" || strings.Contains(strings.ToLower(v), "missing") || strings.HasPrefix(v, "Query") || strings.HasPrefix(v, "Sbjct") {
		return "", false
	}
	v = regexp.MustCompile(`\s*\([^)]*\)\s*(?:\d+)?\s*$`).ReplaceAllString(v, "")
	f := strings.Fields(v)
	if len(f) > 0 && regexp.MustCompile(`^\d+$`).MatchString(f[0]) {
		f = f[1:]
	}
	if len(f) > 0 && regexp.MustCompile(`^\d+$`).MatchString(f[len(f)-1]) {
		f = f[:len(f)-1]
	}
	q := strings.Join(f, "")
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
func firstModel(h string) string {
	for _, f := range strings.Fields(h) {
		f = strings.Trim(f, " ,")
		if strings.HasPrefix(f, "C_") || strings.Contains(f, "scaffold_") || strings.Contains(f, "[Chlre3:") {
			return f
		}
	}
	return strings.Fields(h)[0]
}
func main() {
	in := flag.String("input", filepath.Join("raw", "plants-chlamydomonas.txt"), "")
	doc := flag.String("source-doc", filepath.Join("raw", sourceFile), "")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "chlamydomonas.csv"), "")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-chlamydomonas.md"), "")
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
	fmt.Fprintf(&b, "# Plant resource review: Chlamydomonas reinhardtii\n\n- Source SHA-256: `%s`\n- Source declaration: `39 named genes, 2 named pseudogenes, + one bacterial contaminant`\n- Accepted named/pseudogene CYP blocks: `%d`\n- Excluded bacterial model: `1`\n- Review status: `complete`\n\nThe 41 Chlamydomonas CYP headers match the declared 39 genes plus two pseudogenes. Each block often contains an old assembly followed by an explicitly labelled `newest data: version 3` reconstruction; when present, only that latest literal reconstruction is published. Nested Cycas, Volvox, Medicago, human CYP4F and trace/alignment headers are comparison evidence and do not create records or add residues. The separately labelled bacterial scaffold is excluded. Coordinate, phase and wrapping markers are removed while literal X and internal stops remain.\n", h, len(r))
	if e = os.WriteFile(*audit, []byte(b.String()), 0644); e != nil {
		panic(e)
	}
	fmt.Printf("Chlamydomonas reinhardtii: %d named/pseudogene blocks accepted\n", len(r))
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
		_ = w.Write([]string{"plants", "Chlamydomonas reinhardtii", v.Symbol, v.ID, fmt.Sprintf("chlamydomonas:block-%04d:%s", v.Block, v.ID), sourceURL, v.Header, v.Sequence, sourceFile, h, strconv.Itoa(v.Block), strconv.Itoa(v.Line), v.Status})
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
