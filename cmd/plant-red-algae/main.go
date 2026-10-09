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

const indexFile = "plants-red.algae.index.doc"
const indexURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/red.algae.index.doc"
const indexHash = "2fff7205cfacd8ffaa44d78c3686acca7cfc5f245d1510881724dc0cd1ae3089"
const cyanFile = "plants-redalgae.doc"
const cyanURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/redalgae.doc"
const cyanHash = "688a40386047000757b4a1ac4ec65af697dfe4dac04992b298cb36903e783626"

type record struct {
	Block, Line                                          int
	Species, Symbol, ID, Sequence, Note, File, Hash, URL string
}

var proteinLine = regexp.MustCompile(`^[ACDEFGHIKLMNPQRSTVWXY*]+$`)

func parseCyan(text string) ([]record, error) {
	l := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n"), "\n")
	var out []record
	for i := 0; i < len(l); i++ {
		v := strings.TrimSpace(l[i])
		if !regexp.MustCompile(`^#\d+$`).MatchString(v) {
			continue
		}
		id := strings.TrimPrefix(v, "#")
		j := i + 1
		symbol := ""
		for ; j < len(l) && !strings.EqualFold(strings.TrimSpace(l[j]), "Deduced sequence"); j++ {
		}
		if j == len(l) {
			return nil, fmt.Errorf("%s missing Deduced sequence", id)
		}
		var p []string
		for j = j + 1; j < len(l); j++ {
			q := strings.TrimSpace(l[j])
			if q == "" || strings.HasPrefix(q, "#") {
				break
			}
			if !proteinLine.MatchString(q) {
				return nil, fmt.Errorf("%s unexpected protein line %q", id, q)
			}
			p = append(p, q)
		}
		seq := strings.Join(p, "")
		if seq == "" {
			return nil, fmt.Errorf("%s empty", id)
		}
		switch id {
		case "331":
			symbol = "CYP51"
		case "4211", "4765":
			symbol = "CYP710"
		}
		out = append(out, record{len(out) + 1, i + 1, "Cyanidioschyzon merolae", symbol, id, seq, "deduced sequence block #" + id, cyanFile, cyanHash, cyanURL})
	}
	if len(out) != 5 {
		return nil, fmt.Errorf("Cyanidioschyzon records=%d", len(out))
	}
	return out, nil
}

func parseAlignment(text string) (map[string]string, error) {
	l := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n"), "\n")
	start := -1
	for i, v := range l {
		if strings.TrimSpace(v) == "13    628" {
			start = i + 1
			break
		}
	}
	if start < 0 {
		return nil, fmt.Errorf("Phylip header missing")
	}
	labels := []string{"331", "CYP51con10", "4211", "4765", "710B_gen", "710B1_like", "con981", "con989", "con454", "con1062", "444", "201", "con1041"}
	seqs := map[string]string{}
	line := start
	for block := 0; block < 13; block++ {
		for line < len(l) && strings.TrimSpace(l[line]) == "" {
			line++
		}
		for row := 0; row < 13; row++ {
			if line >= len(l) {
				return nil, fmt.Errorf("alignment ended block %d row %d", block, row)
			}
			raw := l[line]
			line++
			var part string
			if block == 0 {
				fields := strings.Fields(raw)
				if len(fields) < 2 || fields[0] != labels[row] {
					return nil, fmt.Errorf("alignment label block %d row %d: %q", block, row, raw)
				}
				part = strings.Join(fields[1:], "")
			} else {
				part = strings.Join(strings.Fields(raw), "")
			}
			if !validAligned(part) {
				return nil, fmt.Errorf("bad alignment part %q", part)
			}
			seqs[labels[row]] += part
		}
	}
	for _, label := range labels {
		if len(seqs[label]) != 628 {
			return nil, fmt.Errorf("%s alignment length=%d", label, len(seqs[label]))
		}
	}
	return seqs, nil
}
func validAligned(v string) bool {
	if v == "" {
		return false
	}
	for _, r := range v {
		if !strings.ContainsRune("ACDEFGHIKLMNPQRSTVWXY-", r) {
			return false
		}
	}
	return true
}

func combine(cyan []record, aligned map[string]string) ([]record, error) {
	alignID := map[string]string{"201": "201", "4765": "4765", "4211": "4211", "331": "331", "444": "444"}
	for _, r := range cyan {
		if strings.ReplaceAll(aligned[alignID[r.ID]], "-", "") != r.Sequence {
			return nil, fmt.Errorf("alignment cross-check mismatch for %s", r.ID)
		}
	}
	gLabels := []string{"CYP51con10", "710B_gen", "710B1_like", "con981", "con989", "con454", "con1062", "con1041"}
	out := append([]record{}, cyan...)
	for i, id := range gLabels {
		symbol := ""
		if id == "CYP51con10" {
			symbol = "CYP51"
		}
		if id == "710B_gen" || id == "710B1_like" {
			symbol = "CYP710"
		}
		out = append(out, record{6 + i, 14 + i, "Galdieria sulphuraria", symbol, id, aligned[id], "literal 628-column Phylip alignment row " + id, indexFile, indexHash, indexURL})
	}
	return out, nil
}

func main() {
	cyanText := flag.String("cyan-text", filepath.Join("raw", "plants-redalgae.txt"), "")
	indexText := flag.String("index-text", filepath.Join("raw", "plants-red.algae.index.txt"), "")
	cyanDoc := flag.String("cyan-doc", filepath.Join("raw", cyanFile), "")
	indexDoc := flag.String("index-doc", filepath.Join("raw", indexFile), "")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "red-algae.csv"), "")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-red-algae.md"), "")
	flag.Parse()
	ct, e := os.ReadFile(*cyanText)
	if e != nil {
		panic(e)
	}
	it, e := os.ReadFile(*indexText)
	if e != nil {
		panic(e)
	}
	cyan, e := parseCyan(string(ct))
	if e != nil {
		panic(e)
	}
	aln, e := parseAlignment(string(it))
	if e != nil {
		panic(e)
	}
	records, e := combine(cyan, aln)
	if e != nil {
		panic(e)
	}
	ch, e := hash(*cyanDoc)
	if e != nil || ch != cyanHash {
		panic(fmt.Sprintf("cyan hash=%s err=%v", ch, e))
	}
	ih, e := hash(*indexDoc)
	if e != nil || ih != indexHash {
		panic(fmt.Sprintf("index hash=%s err=%v", ih, e))
	}
	if e = writeCSV(*out, records); e != nil {
		panic(e)
	}
	a := fmt.Sprintf("# Plant resource review: Red Algae\n\n- Index source: `%s`, SHA-256 `%s`\n- Cyanidioschyzon child source: `%s`, SHA-256 `%s`\n- Accepted Cyanidioschyzon merolae deduced proteins: `5`\n- Accepted Galdieria sulphuraria Phylip rows: `8`\n- Total records with literal sequence: `%d`\n- Review status: `complete`\n\nThe index has two historical relative hyperlinks, `redalgae.htm` and `Galdiera.htm`. The migrated Cyanidioschyzon child document is available as `redalgae.doc` and supplies five complete deduced proteins. The Galdiera child URL is no longer available, so its eight records are taken only from the index's explicit complete `13    628` Phylip alignment. Alignment gaps are retained literally; no unalignment, repair or external completion is performed. The five Cyanidioschyzon child proteins exactly match their index rows after removing alignment gaps, providing a source-internal cross-check.\n", indexFile, indexHash, cyanFile, cyanHash, len(records))
	if e = os.WriteFile(*audit, []byte(a), 0644); e != nil {
		panic(e)
	}
	fmt.Printf("Red Algae: %d records accepted\n", len(records))
}
func writeCSV(path string, r []record) error {
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
		_ = w.Write([]string{"plants", v.Species, v.Symbol, v.ID, fmt.Sprintf("red-algae:block-%04d:%s", v.Block, v.ID), v.URL, v.Note, v.Sequence, v.File, v.Hash, strconv.Itoa(v.Block), strconv.Itoa(v.Line), ""})
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
