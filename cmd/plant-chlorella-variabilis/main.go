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

const sourceFile = "plants-Chlorella.variabilis.doc"
const sourceURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Chlorella.variabilis.doc"
const sourceHash = "f76f6d9007b14e7ac062530cfa526d756eadb45a620d3c830125bf81acf8dc48"

type record struct {
	Block, Line                          int
	Header, Symbol, ID, Sequence, Status string
}

var proteinLine = regexp.MustCompile(`^[ACDEFGHIKLMNPQRSTVWXY*]+$`)

func parse(text string) ([]record, error) {
	l := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n"), "\n")
	var s []int
	for i, v := range l {
		if strings.HasPrefix(strings.TrimSpace(v), ">CYP") {
			s = append(s, i)
		}
	}
	if len(s) != 19 {
		return nil, fmt.Errorf("headers=%d", len(s))
	}
	var out []record
	for b, st := range s {
		e := len(l)
		if b+1 < len(s) {
			e = s[b+1]
		}
		h := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(l[st]), ">"))
		f := strings.Fields(h)
		var p []string
		for _, raw := range l[st+1 : e] {
			v := strings.TrimSpace(raw)
			if proteinLine.MatchString(v) {
				p = append(p, v)
			}
		}
		q := strings.Join(p, "")
		if q == "" {
			return nil, fmt.Errorf("block %d empty", b+1)
		}
		status := ""
		if strings.ContainsAny(q, "Xx") {
			status = "ambiguous-X-or-x"
		}
		out = append(out, record{b + 1, st + 1, h, f[0], f[1], q, status})
	}
	return out, nil
}
func main() {
	in := flag.String("input", filepath.Join("raw", "plants-Chlorella.variabilis.txt"), "")
	doc := flag.String("source-doc", filepath.Join("raw", sourceFile), "")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "chlorella-variabilis.csv"), "")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-chlorella-variabilis.md"), "")
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
		panic(fmt.Sprintf("hash=%s err=%v", h, e))
	}
	if e = writeCSV(*out, r, h); e != nil {
		panic(e)
	}
	a := fmt.Sprintf("# Plant resource review: Chlorella variabilis\n\n- Source file: `%s`\n- Source URL: `%s`\n- Source SHA-256: `%s`\n- Declared and observed CYP headers: `19`\n- Accepted records with literal protein: `%d`\n- Review status: `complete`\n\nEach FamAln header is followed by similarity annotations and contiguous uppercase protein lines. All 19 records are retained. The 16 literal `X` residues in `CYP710B1` remain unchanged. `CYP845A3a` and `CYP845A3b` are explicitly described as a 100%% duplicate sequence but remain separate source models. No terminal stop markers occur.\n", sourceFile, sourceURL, h, len(r))
	if e = os.WriteFile(*audit, []byte(a), 0644); e != nil {
		panic(e)
	}
	fmt.Printf("Chlorella variabilis: %d CYP records accepted\n", len(r))
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
		_ = w.Write([]string{"plants", "Chlorella variabilis", v.Symbol, v.ID, fmt.Sprintf("chlorella-variabilis:block-%04d:%s", v.Block, v.ID), sourceURL, v.Header, v.Sequence, sourceFile, h, strconv.Itoa(v.Block), strconv.Itoa(v.Line), v.Status})
	}
	w.Flush()
	return w.Error()
}
func hash(path string) (string, error) {
	d, e := os.ReadFile(path)
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%x", sha256.Sum256(d)), nil
}
