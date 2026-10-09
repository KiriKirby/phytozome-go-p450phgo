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

const sourceFile = "plants-ostreococcus.doc"
const sourceURL = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/ostreococcus.doc"
const sourceHash = "12c744823cab1fd85471b10f024100b2438f4465e61e314f5f7e8367b06ef622"

type record struct {
	Block, Line                           int
	Header, Species, Symbol, ID, Sequence string
	Status                                string
}

var (
	leadingCoord  = regexp.MustCompile(`^\d+\s+`)
	trailingCoord = regexp.MustCompile(`\s+\d+$`)
)

func parse(text string) ([]record, error) {
	lines := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n"), "\n")
	var starts []int
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), ">CYP") {
			starts = append(starts, i)
		}
	}
	if len(starts) != 30 {
		return nil, fmt.Errorf("Ostreococcus headers=%d, want 30", len(starts))
	}
	rccStart := -1
	for i, line := range lines {
		if strings.TrimSpace(line) == "Ostreococcus RCC809 micro algae (from Dr. Gotoh? FamAln analysis)" {
			rccStart = i
			break
		}
	}
	var out []record
	for b, start := range starts {
		end := len(lines)
		if b+1 < len(starts) {
			end = starts[b+1]
		}
		header := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(lines[start]), ">"))
		fields := strings.Fields(header)
		if len(fields) < 2 {
			return nil, fmt.Errorf("block %d malformed header", b+1)
		}
		species := speciesFor(start, header, rccStart)
		symbol, id := fields[0], fields[1]
		var pieces []string
		for _, raw := range lines[start+1 : end] {
			v := strings.TrimSpace(raw)
			if v == "" {
				continue
			}
			v = leadingCoord.ReplaceAllString(v, "")
			v = trailingCoord.ReplaceAllString(v, "")
			v = strings.Join(strings.Fields(v), "")
			if v == "" {
				continue
			}
			if !validResidues(v) {
				continue
			}
			pieces = append(pieces, v)
		}
		sequence := strings.Join(pieces, "")
		if sequence == "" {
			return nil, fmt.Errorf("block %d %s has no sequence", b+1, symbol)
		}
		status := ""
		if strings.HasSuffix(sequence, "*") {
			sequence = strings.TrimSuffix(sequence, "*")
			status = "source-terminal-stop"
		}
		if strings.Contains(sequence, "*") {
			status = appendStatus(status, "internal-stop")
		}
		if strings.ContainsAny(sequence, "Xx") {
			status = appendStatus(status, "ambiguous-X-or-x")
		}
		out = append(out, record{b + 1, start + 1, header, species, symbol, id, sequence, status})
	}
	return out, nil
}

func validResidues(v string) bool {
	if len(v) < 2 && v != "*" {
		return false
	}
	for _, r := range v {
		if !strings.ContainsRune("ACDEFGHIKLMNPQRSTVWXYXx*", r) {
			return false
		}
	}
	return true
}

func speciesFor(line int, header string, rccStart int) string {
	if rccStart >= 0 && line >= rccStart {
		return "Ostreococcus RCC809"
	}
	if strings.Contains(header, "Ost9901_3") || strings.Contains(header, "eugene.") {
		return "Ostreococcus lucimarinus"
	}
	return "Ostreococcus tauri"
}

func appendStatus(current, next string) string {
	if current == "" {
		return next
	}
	return current + ";" + next
}

func main() {
	input := flag.String("input", filepath.Join("raw", "plants-ostreococcus.txt"), "")
	doc := flag.String("source-doc", filepath.Join("raw", sourceFile), "")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "ostreococcus.csv"), "")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-ostreococcus.md"), "")
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
		panic(fmt.Sprintf("source hash=%s err=%v", h, err))
	}
	if err := writeCSV(*out, records, h); err != nil {
		panic(err)
	}
	auditText := fmt.Sprintf("# Plant resource review: Ostreococcus\n\n- Source file: `%s`\n- Source URL: `%s`\n- Source SHA-256: `%s`\n- Physical `>CYP` headers: `30`\n- Accepted records: `%d`\n- Actual species split: `10` Ostreococcus tauri, `11` Ostreococcus lucimarinus, `9` Ostreococcus RCC809\n- Review status: `complete`\n\nThe preamble says each of the three species has ten P450s, but the physical headers do not have a 10/10/10 split: O. lucimarinus includes two explicit CYP97A15 model blocks while the RCC809 section contains nine headers. All 30 source blocks remain distinct rather than being merged or padded to force the prose count. Species are assigned from the source header identifiers and exact RCC809 section boundary, not inferred from similarity text. Each block is bounded by the next `>CYP` header; coordinate numbers in the two `CYP800A1` layouts are removed only as line-edge coordinates, while residue letters, X/x, internal stops and duplicate/alternate models remain literal. Only an explicit terminal `*` is removed.\n", sourceFile, sourceURL, h, len(records))
	if err := os.WriteFile(*audit, []byte(auditText), 0644); err != nil {
		panic(err)
	}
	fmt.Printf("Ostreococcus: %d records accepted\n", len(records))
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
		_ = w.Write([]string{"plants", r.Species, r.Symbol, r.ID, fmt.Sprintf("ostreococcus:block-%04d:%s", r.Block, r.ID), sourceURL, r.Header, r.Sequence, sourceFile, h, strconv.Itoa(r.Block), strconv.Itoa(r.Line), r.Status})
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
