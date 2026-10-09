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

const sourceFile = "plants-moss.doc"
const sourceURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/moss.doc"
const sourceHash = "7eb6937402e7a6f714470791aecc8f92c328bd209159cb672c3069c85c05cc9b"

type record struct {
	Block, Line                          int
	Header, Symbol, ID, Sequence, Status string
}

var cypRE = regexp.MustCompile(`\b(CYP[0-9A-Za-z?_-]+)\b`)
var excluded = map[int]string{16: "Panax ginseng", 47: "Cycas rumphii", 61: "Picea glauca", 63: "Pinus taeda", 71: "Lycopersicon esculentum", 72: "Cucumis sativus", 73: "Lycopersicon esculentum", 82: "Picea sitchensis", 93: "Ceratopteris richardii", 102: "Ceratodon purpureus", 104: "Chlamydomonas", 112: "Lotus japonicus", 113: "Medicago truncatula", 114: "Picea glauca", 115: "Helicosporidium sp."}
var bacterial = map[int]bool{109: true, 110: true, 111: true}
var alignmentOnly = map[int]bool{105: true}

func parse(text string) ([]record, []record, error) {
	l := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n"), "\n")
	var st []int
	for i, v := range l {
		if strings.HasPrefix(strings.TrimSpace(v), ">") {
			st = append(st, i)
		}
	}
	if len(st) != 115 {
		return nil, nil, fmt.Errorf("moss headers=%d", len(st))
	}
	var a, x []record
	for bi, s := range st {
		e := len(l)
		if bi+1 < len(st) {
			e = st[bi+1]
		}
		h := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(l[s]), ">"))
		sym := ""
		if m := cypRE.FindStringSubmatch(h); len(m) == 2 {
			sym = m[1]
		}
		id := strings.Fields(h)[0]
		var p []string
		for i := s + 1; i < e; i++ {
			if q, ok := seqline(l[i]); ok {
				p = append(p, q)
			}
		}
		q := strings.Join(p, "")
		if q == "" && alignmentOnly[bi+1] {
			x = append(x, record{bi + 1, s + 1, h, sym, id, "", "excluded: pairwise alignment only; no standalone source sequence"})
			continue
		}
		if q == "" {
			return nil, nil, fmt.Errorf("block %d empty", bi+1)
		}
		term := strings.HasSuffix(q, "*")
		if term {
			q = strings.TrimSuffix(q, "*")
		}
		status := ""
		if term {
			status = "source-terminal-stop"
		}
		if strings.Contains(q, "*") {
			status += ";internal-stop"
		}
		if strings.ContainsAny(q, "Xx") {
			status += ";ambiguous-X-or-x"
		}
		r := record{bi + 1, s + 1, h, sym, id, q, strings.Trim(status, ";")}
		if sp, ok := excluded[bi+1]; ok {
			r.Status = "excluded: explicit foreign " + sp
			x = append(x, r)
		} else if bacterial[bi+1] {
			r.Status = "excluded: source-declared bacterial contamination"
			x = append(x, r)
		} else {
			a = append(a, r)
		}
	}
	if len(a) != 96 || len(x) != 19 {
		return nil, nil, fmt.Errorf("moss counts %d/%d", len(a), len(x))
	}
	return a, x, nil
}
func seqline(raw string) (string, bool) {
	v := strings.TrimSpace(raw)
	if v == "" || strings.HasPrefix(v, "$") {
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
func main() {
	in := flag.String("input", filepath.Join("raw", "plants-moss.txt"), "")
	doc := flag.String("source-doc", filepath.Join("raw", sourceFile), "")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "physcomitrella-patens.csv"), "")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-physcomitrella-patens.md"), "")
	flag.Parse()
	d, e := os.ReadFile(*in)
	if e != nil {
		panic(e)
	}
	a, x, e := parse(string(d))
	if e != nil {
		panic(e)
	}
	h, e := hash(*doc)
	if e != nil || h != sourceHash {
		panic(h)
	}
	if e = write(*out, a, h); e != nil {
		panic(e)
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Physcomitrella patens\n\n- Source SHA-256: `%s`\n- Explicit blocks: `115`\n- Accepted moss blocks: `%d`\n- Excluded foreign helper blocks: `15`\n- Excluded bacterial contaminants: `3`\n- Excluded alignment-only comparison blocks: `1`\n- Review status: `complete`\n\nThe preamble declares 98 moss sequence pieces, including three bacterial contaminants, plus 15 non-Physcomitrella assembly helpers. The exact foreign and bacterial blocks are excluded by explicit source labels. Block 105 is only a pairwise alignment of its two neighboring records and has no standalone source sequence, so it is not promoted to a record. Wrapped, coordinate/phase and terminal-stop layouts are handled block-locally; X/x and internal stops remain literal.\n", h, len(a))
	for _, v := range x {
		fmt.Fprintf(&b, "\n- block %d `%s`: %s", v.Block, v.ID, v.Status)
	}
	if e = os.WriteFile(*audit, []byte(b.String()), 0644); e != nil {
		panic(e)
	}
	fmt.Printf("Physcomitrella patens: %d accepted, %d excluded\n", len(a), len(x))
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
		_ = w.Write([]string{"plants", "Physcomitrella patens", v.Symbol, v.ID, fmt.Sprintf("physcomitrella-patens:block-%04d:%s", v.Block, v.ID), sourceURL, v.Header, v.Sequence, sourceFile, h, strconv.Itoa(v.Block), strconv.Itoa(v.Line), v.Status})
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
