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

const sourceFile = "plants-lolium.doc"
const sourceURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/lolium.doc"
const sourceHash = "fde60d3892a9a642dae66ff5bbccb753ad1decabbf7338cd241c431d26a54603"

type record struct {
	Block, Line                 int
	Header, ID, Clone, Sequence string
}

var acc = regexp.MustCompile(`\bAF\d+\b`)
var clone = regexp.MustCompile(`clone\s+(\S+)`)

func main() {
	in := flag.String("input", filepath.Join("raw", "plants-lolium.txt"), "normalized Lolium text")
	doc := flag.String("source-doc", filepath.Join("raw", sourceFile), "Word source")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "lolium-rigidum.csv"), "CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-lolium-rigidum.md"), "audit")
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
	fmt.Printf("Lolium rigidum: %d source-declared unassigned sequences accepted\n", len(r))
}
func parse(text string) ([]record, error) {
	lines := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n"), "\n")
	hs := []int{}
	for i, l := range lines {
		if strings.HasPrefix(strings.TrimSpace(l), ">") {
			hs = append(hs, i)
		}
	}
	if len(hs) != 16 {
		return nil, fmt.Errorf("headers=%d", len(hs))
	}
	r := []record{}
	for b, st := range hs {
		en := len(lines)
		if b+1 < len(hs) {
			en = hs[b+1]
		}
		head := strings.TrimSpace(strings.TrimPrefix(lines[st], ">"))
		id := acc.FindString(head)
		m := clone.FindStringSubmatch(head)
		if id == "" || len(m) != 2 {
			return nil, fmt.Errorf("block %d header", b+1)
		}
		parts := []string{}
		for i := st + 1; i < en; i++ {
			v := strings.Join(strings.Fields(lines[i]), "")
			ok := v != ""
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
		if !strings.HasSuffix(seq, "*") {
			return nil, fmt.Errorf("block %d no terminal stop", b+1)
		}
		seq = strings.TrimSuffix(seq, "*")
		r = append(r, record{b + 1, st + 1, head, id, m[1], seq})
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
		_ = w.Write([]string{"plants", "Lolium rigidum", "", v.ID, fmt.Sprintf("lolium-rigidum:block-%04d:%s", v.Block, v.ID), sourceURL, fmt.Sprintf("Lolium source block %d; clone=%s; source explicitly says sequences are not named yet; header=%s", v.Block, v.Clone, v.Header), v.Sequence, sourceFile, h, strconv.Itoa(v.Block), strconv.Itoa(v.Line), v.Header, "source-unannotated;source-terminal-stop"})
	}
	w.Flush()
	return w.Error()
}
func writeAudit(p string, r []record, h string) error {
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Lolium rigidum\n\n- Source file: `%s`\n- URL: %s\n- SHA-256: `%s`\n- Word layout: 205 paragraphs, 4 pages, no tables\n- Declared sequences: `16`\n- Accepted literal sequences: `%d`\n- Review status: `complete`\n\nThe preamble explicitly says these 16 GenBank sequences `are not named yet`. Each of the 16 clone/accession headers has one terminal-stop-bounded literal protein. Similarity text such as `78%% to 72A18` is retained only in the source header and is not converted into an invented CYP assignment; `Symbol` therefore remains empty. Three comparison notes wrap onto a separate line and are annotations, not sequence.\n\n| Block | Clone | Accession | Length |\n|---:|---|---|---:|\n", sourceFile, sourceURL, h, len(r))
	for _, i := range []int{0, len(r) / 2, len(r) - 1} {
		v := r[i]
		fmt.Fprintf(&b, "| %d | %s | %s | %d |\n", v.Block, v.Clone, v.ID, len(v.Sequence))
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
