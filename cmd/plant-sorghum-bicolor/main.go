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

const sourceFile = "plants-sorghum.doc"
const sourceURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/sorghum.doc"
const sourceHash = "d67e794a5671b6aa4b32d726d83803faad913dc0bb1190569bf0f020d81708e2"

type record struct {
	Block, Line                  int
	Header, ID, Symbol, Sequence string
	Status                       []string
}

var cyp = regexp.MustCompile(`\bCYP[0-9A-Za-z]+`)

func main() {
	in := flag.String("input", filepath.Join("raw", "plants-sorghum.txt"), "normalized Sorghum text")
	doc := flag.String("source-doc", filepath.Join("raw", sourceFile), "Word source")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "sorghum-bicolor.csv"), "CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-sorghum-bicolor.md"), "audit")
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
	if e != nil {
		panic(e)
	}
	if h != sourceHash {
		panic(h)
	}
	if e = writeCSV(*out, r, h); e != nil {
		panic(e)
	}
	if e = writeAudit(*audit, r, h); e != nil {
		panic(e)
	}
	fmt.Printf("Sorghum bicolor: %d literal source models accepted\n", len(r))
}
func parse(text string) ([]record, error) {
	text = strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n")
	for _, s := range []string{"AGEELHRLTDLNYGMRAMAINLPGFAFHKAFKARKKLVSVLQGVL>Sb01g034320|Sorbi1", "MCVLAKKDVPLFRLRFSSAEVVVAASARVAAQFLRTHDANFSNRPPNSGAEH>Sb07g000493|Sorbi1"} {
		if strings.Count(text, s) != 1 {
			return nil, fmt.Errorf("missing exact joined header %q", s)
		}
		text = strings.Replace(text, s, strings.Replace(s, ">", "\n>", 1), 1)
	}
	lines := strings.Split(text, "\n")
	hs := []int{}
	for i, l := range lines {
		if strings.HasPrefix(strings.TrimSpace(l), ">") {
			hs = append(hs, i)
		}
	}
	if len(hs) != 628 {
		return nil, fmt.Errorf("headers=%d", len(hs))
	}
	r := []record{}
	for b, st := range hs {
		en := len(lines)
		if b+1 < len(hs) {
			en = hs[b+1]
		}
		head := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(lines[st]), ">"))
		id := strings.Split(head, "|")[0]
		sym := cyp.FindString(head)
		parts := []string{}
		for i := st + 1; i < en; i++ {
			v := strings.TrimSpace(lines[i])
			if v == "" || v == "&" {
				continue
			}
			ok := true
			for _, x := range v {
				if !strings.ContainsRune("ACDEFGHIKLMNPQRSTVWY*", x) {
					ok = false
					break
				}
			}
			if ok {
				parts = append(parts, v)
			}
		}
		seq := strings.Join(parts, "")
		stt := []string{}
		if strings.HasSuffix(seq, "*") {
			seq = strings.TrimSuffix(seq, "*")
			stt = append(stt, "source-terminal-stop")
		}
		if seq == "" {
			return nil, fmt.Errorf("block %d no sequence", b+1)
		}
		if strings.Contains(seq, "*") {
			stt = append(stt, "internal-stop")
		}
		if len(seq) < 350 {
			stt = append(stt, "short-sequence")
		}
		if sym == "" {
			stt = append(stt, "source-CYP-assignment-not-present")
		}
		r = append(r, record{b + 1, st + 1, head, id, sym, seq, stt})
	}
	return r, nil
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
	_ = w.Write([]string{"category", "species", "symbol", "id", "record_key", "source_url", "source_note", "sequence", "source_file", "source_sha256", "source_block", "source_line", "source_header", "review_status"})
	for _, v := range r {
		_ = w.Write([]string{"plants", "Sorghum bicolor", v.Symbol, v.ID, fmt.Sprintf("sorghum-bicolor:block-%04d:%s", v.Block, v.ID), sourceURL, fmt.Sprintf("Sorghum gene-model block %d; normalized source line %d; complete header=%s", v.Block, v.Line, v.Header), v.Sequence, sourceFile, h, strconv.Itoa(v.Block), strconv.Itoa(v.Line), v.Header, strings.Join(v.Status, ";")})
	}
	w.Flush()
	return w.Error()
}
func writeAudit(p string, r []record, h string) error {
	cs := map[string]int{}
	seqs := map[string]int{}
	for _, v := range r {
		seqs[v.Sequence]++
		for _, s := range v.Status {
			cs[s]++
		}
	}
	ns := []string{}
	for n := range cs {
		ns = append(ns, n)
	}
	sort.Strings(ns)
	dg := 0
	for _, n := range seqs {
		if n > 1 {
			dg++
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Sorghum bicolor\n\n- Source file: `%s`\n- URL: %s\n- SHA-256: `%s`\n- Word layout: 5,101 paragraphs, 72 pages, no tables\n- Declared gene models: `628` in `372` ampersand-separated bins\n- Accepted literal sequence blocks: `%d`\n- Duplicate sequence groups retained: `%d`\n- Review status: `complete`\n\nThe normalized document initially exposes 626 line-start headers. Exactly two additional headers are attached to the preceding sequence text at source lines 1068 and 1077; splitting those two reviewed literal joins restores the title's 628 models. Ampersands separate gene bins, not records. Each model remains independent, including alternates and internal-stop-rich unrevised JGI models. The header's CYP comparison is retained when present; headers without a CYP comparison keep an empty symbol. Only a terminal stop is removed.\n\n## Status counts\n\n| Status | Records |\n|---|---:|\n", sourceFile, sourceURL, h, len(r), dg)
	for _, n := range ns {
		fmt.Fprintf(&b, "| %s | %d |\n", n, cs[n])
	}
	b.WriteString("\n## Representative records\n\n| Block | ID | CYP comparison | Length |\n|---:|---|---|---:|\n")
	for _, i := range []int{0, len(r) / 2, len(r) - 1} {
		v := r[i]
		fmt.Fprintf(&b, "| %d | %s | %s | %d |\n", v.Block, v.ID, v.Symbol, len(v.Sequence))
	}
	return os.WriteFile(p, []byte(b.String()), 0644)
}
func hash(p string) (string, error) {
	d, e := os.ReadFile(p)
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%x", sha256.Sum256(d)), nil
}
