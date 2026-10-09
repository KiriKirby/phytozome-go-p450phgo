package main

import (
	"crypto/sha256"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const sourceFile = "plants-micromonas.doc"
const sourceURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/micromonas.doc"
const sourceHash = "df99108a1b9cf8cc4270647573de75fa1340e54d37bbfea0746f984c7f656d74"

type record struct {
	Block, Line                                   int
	Header, Species, Symbol, ID, Sequence, Status string
}

func parse(text string) ([]record, error) {
	l := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n"), "\n")
	var s []int
	for i, v := range l {
		if strings.HasPrefix(strings.TrimSpace(v), ">CYP") {
			s = append(s, i)
		}
	}
	if len(s) != 32 {
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
		if len(f) < 2 {
			return nil, fmt.Errorf("bad header %d", b+1)
		}
		sp := species(h)
		var p []string
		for _, raw := range l[st+1 : e] {
			v := strings.Join(strings.Fields(strings.TrimSpace(raw)), "")
			if !valid(v) {
				continue
			}
			p = append(p, v)
		}
		q := strings.Join(p, "")
		if q == "" {
			return nil, fmt.Errorf("empty block %d", b+1)
		}
		status := ""
		if strings.HasSuffix(q, "*") {
			q = strings.TrimSuffix(q, "*")
			status = "source-terminal-stop"
		}
		if strings.ContainsAny(q, "Xx") {
			if status != "" {
				status += ";"
			}
			status += "ambiguous-X-or-x"
		}
		out = append(out, record{b + 1, st + 1, h, sp, f[0], f[1], q, status})
	}
	return out, nil
}
func valid(v string) bool {
	if len(v) < 2 {
		return false
	}
	for _, r := range v {
		if !strings.ContainsRune("ACDEFGHIKLMNPQRSTVWXYXx*", r) {
			return false
		}
	}
	return true
}
func species(h string) string {
	if strings.Contains(h, "CCMP1545") || strings.Contains(h, "MicpuC2") {
		return "Micromonas pusilla CCMP1545"
	}
	if strings.Contains(h, "CCMP490") || strings.Contains(h, "Chlorella sp.") {
		return "Micromonas sp. CCMP490 EST"
	}
	return "Micromonas sp. RCC299"
}
func main() {
	in := flag.String("input", filepath.Join("raw", "plants-micromonas.txt"), "")
	doc := flag.String("source-doc", filepath.Join("raw", sourceFile), "")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "micromonas.csv"), "")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-micromonas.md"), "")
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
	a := fmt.Sprintf("# Plant resource review: Micromonas\n\n- Source file: `%s`\n- Source URL: `%s`\n- Source SHA-256: `%s`\n- Physical `>CYP` headers: `32`\n- Accepted records: `%d`\n- Species/source groups: RCC299, CCMP1545 and explicit CCMP490/Chlorella EST fragments\n- Review status: `complete`\n\nEvery header-bounded block is retained, including alternate models, deleted/partial CYP746A1, the Chlorella EST CYP747A1 fragment, and CCMP490 EST fragments. Only literal residue lines are accepted; annotations such as `yellow regions`, similarity prose and `same sequence` are not sequence. Explicit terminal `*` markers alone are removed; X/x and incomplete no-stop fragments remain literal.\n", sourceFile, sourceURL, h, len(r))
	if e = os.WriteFile(*audit, []byte(a), 0644); e != nil {
		panic(e)
	}
	fmt.Printf("Micromonas: %d records accepted\n", len(r))
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
		_ = w.Write([]string{"plants", v.Species, v.Symbol, v.ID, fmt.Sprintf("micromonas:block-%04d:%s", v.Block, v.ID), sourceURL, v.Header, v.Sequence, sourceFile, h, strconv.Itoa(v.Block), strconv.Itoa(v.Line), v.Status})
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
