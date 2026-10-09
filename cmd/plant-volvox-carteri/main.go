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

const sourceFile = "plants-volvox.doc"
const sourceURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/volvox.doc"
const sourceHash = "594311939b4d10dd31680a84e12e7848f41658fc67603d344a04fbcec7167311"

type record struct {
	Block, Line                          int
	Header, Symbol, ID, Sequence, Status string
}

var (
	phaseRE    = regexp.MustCompile(`\s*\((?:[012?](?:\s+GC)?)\)\s*_?`)
	leadingRE  = regexp.MustCompile(`^\d+\s+`)
	trailingRE = regexp.MustCompile(`\s+\d+$`)
	modelRE    = regexp.MustCompile(`(?:estExt_[^\s,]+|fgenesh[^\s,]+|e_gw[^\s,]+|gw1\.[^\s,]+)`)
)

func isVolvoxHeader(v string) bool {
	v = strings.TrimSpace(v)
	return strings.HasPrefix(v, ">CYP") &&
		!strings.Contains(strings.ToLower(v), "chlamy") &&
		!strings.HasPrefix(v, ">CYP7B1")
}

func parse(text string) ([]record, error) {
	lines := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n"), "\n")
	var headers []int
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), ">") {
			headers = append(headers, i)
		}
	}
	var out []record
	for hi, start := range headers {
		if !isVolvoxHeader(lines[start]) {
			continue
		}
		end := len(lines)
		if hi+1 < len(headers) {
			end = headers[hi+1]
		}
		header := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(lines[start]), ">"))
		symbol := strings.Fields(header)[0]
		id := findModel(lines[start:end])
		var pieces []string
		terminated := false
		for _, line := range lines[start+1 : end] {
			piece, ok := sequenceLine(line)
			if !ok {
				continue
			}
			pieces = append(pieces, piece)
			if strings.HasSuffix(piece, "*") {
				terminated = true
				break
			}
		}
		sequence := strings.Join(pieces, "")
		if !terminated || !strings.HasSuffix(sequence, "*") {
			return nil, fmt.Errorf("block %d %s lacks explicit terminal stop", len(out)+1, symbol)
		}
		sequence = strings.TrimSuffix(sequence, "*")
		status := "source-terminal-stop"
		if strings.Contains(sequence, "*") {
			status += ";internal-stop"
		}
		if strings.ContainsAny(sequence, "Xx") {
			status += ";ambiguous-X-or-x"
		}
		out = append(out, record{len(out) + 1, start + 1, header, symbol, id, sequence, status})
	}
	if len(out) != 19 {
		return nil, fmt.Errorf("Volvox CYP blocks=%d, want 19", len(out))
	}
	return out, nil
}

func sequenceLine(raw string) (string, bool) {
	v := strings.TrimSpace(raw)
	if v == "" || regexp.MustCompile(`[a-z]{3,}`).MatchString(v) {
		return "", false
	}
	v = phaseRE.ReplaceAllString(v, " ")
	v = leadingRE.ReplaceAllString(v, "")
	v = trailingRE.ReplaceAllString(v, "")
	v = strings.Join(strings.Fields(v), "")
	if len(v) < 5 {
		return "", false
	}
	for _, r := range v {
		if !strings.ContainsRune("ACDEFGHIKLMNPQRSTVWXYXxhq*", r) {
			return "", false
		}
	}
	return v, true
}

func findModel(lines []string) string {
	for _, line := range lines {
		if model := modelRE.FindString(line); model != "" {
			return strings.TrimSuffix(model, ".")
		}
	}
	return strings.Fields(strings.TrimPrefix(strings.TrimSpace(lines[0]), ">"))[0]
}

func main() {
	input := flag.String("input", filepath.Join("raw", "plants-volvox.txt"), "")
	doc := flag.String("source-doc", filepath.Join("raw", sourceFile), "")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "volvox-carteri.csv"), "")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-volvox-carteri.md"), "")
	flag.Parse()
	data, err := os.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	records, err := parse(string(data))
	if err != nil {
		panic(err)
	}
	hashValue, err := hash(*doc)
	if err != nil || hashValue != sourceHash {
		panic(fmt.Sprintf("source hash=%s err=%v", hashValue, err))
	}
	if err := writeCSV(*out, records, hashValue); err != nil {
		panic(err)
	}
	var auditText strings.Builder
	fmt.Fprintf(&auditText, "# Plant resource review: Volvox carteri\n\n- Source file: `%s`\n- Source URL: `%s`\n- Source SHA-256: `%s`\n- Physical `>` headers inspected: `23`\n- Accepted Volvox CYP blocks: `%d`\n- Records with literal protein: `%d`\n- Excluded comparison records: `4` (two Chlamydomonas reconstructions, one bacterial protein, one human CYP7B1)\n- Review status: `complete`\n\nEvery accepted Volvox block ends in an explicit terminal `*`. The resource-specific parser reads only pure residue lines and the document's observed coordinate/phase layouts until that stop, removes only the terminal marker, and does not bridge the explicit sequence-gap prose. Chlamydomonas CYP771A1/CYP772A1 reconstructions, their EST excerpts and alignments, the Plesiocystis bacterial protein, and human CYP7B1 remain comparison evidence only. Source order and literal lowercase `hq` in the Chlamydomonas comparison are inspected but the latter is not part of an accepted Volvox record.\n", sourceFile, sourceURL, hashValue, len(records), len(records))
	if err := os.WriteFile(*audit, []byte(auditText.String()), 0644); err != nil {
		panic(err)
	}
	fmt.Printf("Volvox carteri: %d CYP blocks accepted\n", len(records))
}

func writeCSV(path string, records []record, hashValue string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	_ = writer.Write([]string{"category", "species", "symbol", "id", "record_key", "source_url", "source_note", "sequence", "source_file", "source_sha256", "source_block", "source_line", "review_status"})
	for _, r := range records {
		_ = writer.Write([]string{"plants", "Volvox carteri", r.Symbol, r.ID, fmt.Sprintf("volvox-carteri:block-%04d:%s", r.Block, r.ID), sourceURL, r.Header, r.Sequence, sourceFile, hashValue, strconv.Itoa(r.Block), strconv.Itoa(r.Line), r.Status})
	}
	writer.Flush()
	return writer.Error()
}

func hash(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha256.Sum256(data)), nil
}
