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

const sourceFile = "plants-Chlorella.vulgaris.doc"
const sourceURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Chlorella.vulgaris.doc"
const sourceHash = "82535cdaab58b282329dfd83d981f0ec472b5cf048fd9993d45c4ff6a4ccf424"

type record struct {
	Block, Line                          int
	Header, Symbol, ID, Sequence, Status string
}

var proteinLine = regexp.MustCompile(`^[ACDEFGHIKLMNPQRSTVWXY*]+$`)

func parse(text string) ([]record, error) {
	lines := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n"), "\n")
	var starts []int
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), ">CYP") {
			starts = append(starts, i)
		}
	}
	if len(starts) != 37 {
		return nil, fmt.Errorf("headers=%d, want 37", len(starts))
	}
	var out []record
	for block, start := range starts {
		end := len(lines)
		if block+1 < len(starts) {
			end = starts[block+1]
		}
		header := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(lines[start]), ">"))
		fields := strings.Fields(header)
		symbol, id := fields[0], fields[1]
		var pieces []string
		for _, raw := range lines[start+1 : end] {
			v := strings.TrimSpace(raw)
			if v == "This part not P450 seq" {
				break
			}
			if proteinLine.MatchString(v) {
				pieces = append(pieces, v)
			}
		}
		sequence := strings.Join(pieces, "")
		if sequence == "" {
			return nil, fmt.Errorf("block %d %s has no protein", block+1, symbol)
		}
		status := ""
		if strings.HasSuffix(symbol, "P") || strings.Contains(symbol, "fragment") {
			status = "source-pseudogene-or-fragment-label"
		}
		out = append(out, record{block + 1, start + 1, header, symbol, id, sequence, status})
	}
	return out, nil
}

func main() {
	input := flag.String("input", filepath.Join("raw", "plants-Chlorella.vulgaris.txt"), "")
	doc := flag.String("source-doc", filepath.Join("raw", sourceFile), "")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "chlorella-vulgaris.csv"), "")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-chlorella-vulgaris.md"), "")
	flag.Parse()
	data, err := os.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	records, err := parse(string(data))
	if err != nil {
		panic(err)
	}
	h, err := hash(*doc)
	if err != nil || h != sourceHash {
		panic(fmt.Sprintf("hash=%s err=%v", h, err))
	}
	if err := writeCSV(*out, records, h); err != nil {
		panic(err)
	}
	auditText := fmt.Sprintf("# Plant resource review: Chlorella vulgaris\n\n- Source file: `%s`\n- Source URL: `%s`\n- Source SHA-256: `%s`\n- Declared and observed CYP headers: `37`\n- Accepted records with literal protein: `%d`\n- Review status: `complete`\n\nThe FamAln report has one CYP header per record, followed by similarity annotations and contiguous uppercase protein lines. All 37 source records are retained in order. `CYP855C1P` contributes only the two protein lines before the exact source sentence `This part not P450 seq`; the following non-P450 character string is excluded. `CYP863-fragment1` remains a literal short fragment. No terminal stop markers occur.\n", sourceFile, sourceURL, h, len(records))
	if err := os.WriteFile(*audit, []byte(auditText), 0644); err != nil {
		panic(err)
	}
	fmt.Printf("Chlorella vulgaris: %d CYP records accepted\n", len(records))
}

func writeCSV(path string, records []record, h string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	_ = w.Write([]string{"category", "species", "symbol", "id", "record_key", "source_url", "source_note", "sequence", "source_file", "source_sha256", "source_block", "source_line", "review_status"})
	for _, r := range records {
		_ = w.Write([]string{"plants", "Chlorella vulgaris", r.Symbol, r.ID, fmt.Sprintf("chlorella-vulgaris:block-%04d:%s", r.Block, r.ID), sourceURL, r.Header, r.Sequence, sourceFile, h, strconv.Itoa(r.Block), strconv.Itoa(r.Line), r.Status})
	}
	w.Flush()
	return w.Error()
}
func hash(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha256.Sum256(data)), nil
}
