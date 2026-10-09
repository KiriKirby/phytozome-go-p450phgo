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

const (
	sourceFile = "plants-soybean.doc"
	sourceURL  = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/soybean.doc"
)

type record struct {
	Block, SourceLine, FirstSequenceLine, LastSequenceLine int
	Header, ID, Symbol, Sequence                           string
	Status                                                 []string
}

var (
	leadingCoordinate  = regexp.MustCompile(`^\d+\s+`)
	trailingCoordinate = regexp.MustCompile(`\s+\d+\s*$`)
	phaseMarker        = regexp.MustCompile(`\((?:[012?]|[012]\?)?\)`)
	digitP             = regexp.MustCompile(`(?i)\dP$`)
)

func main() {
	input := flag.String("input", filepath.Join("raw", "plants-soybean.txt"), "Word-normalized soybean source text")
	sourceDoc := flag.String("source-doc", filepath.Join("raw", sourceFile), "original downloaded Word document")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "soybean.csv"), "reviewed CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-soybean.md"), "review ledger")
	flag.Parse()
	data, err := os.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	records, excludedESTs, lineCount, err := parseText(string(data), true)
	if err != nil {
		panic(err)
	}
	hash, err := fileHash(*sourceDoc)
	if err != nil {
		panic(err)
	}
	if err := writeCSV(*out, records, hash); err != nil {
		panic(err)
	}
	if err := writeAudit(*audit, records, excludedESTs, lineCount, hash); err != nil {
		panic(err)
	}
	fmt.Printf("soybean: %d named CYP blocks accepted, %d with literal protein; %d EST DNA blocks excluded\n", len(records), countSequences(records), excludedESTs)
}

func parseText(text string, enforceLayout bool) ([]record, int, int, error) {
	text = strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n")
	lines := strings.Split(text, "\n")
	if len(lines) < 5 || strings.TrimSpace(lines[0]) != "Glycine max (soybean) P450s" || !strings.Contains(strings.Join(lines[:min(80, len(lines))], "\n"), "All soybean P450 sequences") {
		return nil, 0, len(lines), fmt.Errorf("soybean preamble changed")
	}
	estMarker := -1
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "955 CYTOCHROME P450 EST FOR Glycine max") {
			estMarker = i
			break
		}
	}
	if estMarker < 0 {
		return nil, 0, len(lines), fmt.Errorf("soybean EST appendix marker missing")
	}
	var headers []int
	for i, line := range lines[:estMarker] {
		if strings.HasPrefix(strings.TrimSpace(line), ">CYP") {
			headers = append(headers, i)
		}
	}
	var estHeaders int
	for _, line := range lines[estMarker+1:] {
		if strings.HasPrefix(strings.TrimSpace(line), ">gi|") {
			estHeaders++
		}
	}
	if enforceLayout && (len(lines) != 10783 || len(headers) != 171 || headers[0]+1 != 84 || headers[len(headers)-1]+1 != 2259 || estMarker+1 != 2271 || estHeaders != 953) {
		return nil, estHeaders, len(lines), fmt.Errorf("soybean layout changed: lines=%d CYP=%d first=%d last=%d EST-marker=%d EST-headers=%d", len(lines), len(headers), headers[0]+1, headers[len(headers)-1]+1, estMarker+1, estHeaders)
	}
	records := make([]record, 0, len(headers))
	for block, start := range headers {
		end := estMarker
		if block+1 < len(headers) {
			end = headers[block+1]
		}
		header := strings.TrimSpace(lines[start])
		fields := strings.Fields(strings.TrimPrefix(header, ">"))
		if len(fields) == 0 || !strings.HasPrefix(fields[0], "CYP") {
			return nil, estHeaders, len(lines), fmt.Errorf("soybean block %d invalid header %q", block+1, header)
		}
		r := record{Block: block + 1, SourceLine: start + 1, Header: header, ID: strings.Trim(fields[0], ",;|"), Symbol: strings.Trim(fields[0], ",;|")}
		blockText := strings.ToLower(strings.Join(lines[start:end], "\n"))
		flags := map[string]bool{}
		if strings.Contains(blockText, "pseudogene") || digitP.MatchString(r.Symbol) {
			flags["source-pseudogene"] = true
		}
		if strings.Contains(strings.ToLower(header), "frag") || strings.Contains(blockText, "fragment") || strings.Contains(blockText, "missing ") || strings.Contains(blockText, "partial") {
			flags["source-fragment-or-missing-region"] = true
		}
		if strings.Contains(blockText, "frameshift") || strings.Contains(blockText, "&") {
			flags["source-frameshift-marker"] = true
		}
		var parts []string
		for i := start + 1; i < end; i++ {
			if part, ok := sequenceLine(lines[i]); ok {
				if r.FirstSequenceLine == 0 {
					r.FirstSequenceLine = i + 1
				}
				r.LastSequenceLine = i + 1
				parts = append(parts, part)
			}
		}
		raw := strings.Join(parts, "")
		if strings.HasSuffix(raw, "*") {
			raw = strings.TrimRight(raw, "*")
			flags["source-terminal-stop"] = true
		}
		if strings.Contains(raw, "*") {
			flags["internal-stop"] = true
		}
		if strings.Contains(raw, "?") {
			flags["source-question-mark-residue"] = true
		}
		if raw == "" {
			flags["sequence-missing"] = true
		} else if len(raw) < 350 {
			flags["short-sequence"] = true
		}
		if len(raw) > 650 {
			flags["unusually-long-sequence"] = true
		}
		for _, name := range []string{"source-pseudogene", "source-fragment-or-missing-region", "source-frameshift-marker", "source-terminal-stop", "internal-stop", "source-question-mark-residue", "sequence-missing", "short-sequence", "unusually-long-sequence"} {
			if flags[name] {
				r.Status = append(r.Status, name)
			}
		}
		r.Sequence = raw
		records = append(records, r)
	}
	return records, estHeaders, len(lines), nil
}

func sequenceLine(line string) (string, bool) {
	value := strings.TrimSpace(line)
	value = leadingCoordinate.ReplaceAllString(value, "")
	value = trailingCoordinate.ReplaceAllString(value, "")
	value = phaseMarker.ReplaceAllString(value, "")
	value = strings.ReplaceAll(value, "&", "")
	value = strings.Join(strings.Fields(value), "")
	if len(value) < 8 {
		return "", false
	}
	for _, aa := range value {
		if !strings.ContainsRune("ACDEFGHIKLMNOPQRSTVWXY*?", aa) {
			return "", false
		}
	}
	return value, true
}

func writeCSV(path string, records []record, hash string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
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
		note := fmt.Sprintf("Soybean Word CYP block %d at normalized source line %d; literal protein lines %d-%d; header=%s", r.Block, r.SourceLine, r.FirstSequenceLine, r.LastSequenceLine, r.Header)
		_ = w.Write([]string{"plants", "Glycine max", r.Symbol, r.ID, fmt.Sprintf("soybean:block-%04d:%s", r.Block, r.ID), sourceURL, note, r.Sequence, sourceFile, hash, strconv.Itoa(r.Block), strconv.Itoa(r.SourceLine), strings.Join(r.Status, ";")})
	}
	w.Flush()
	return w.Error()
}

func writeAudit(path string, records []record, excludedESTs, lineCount int, hash string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	counts := map[string]int{}
	headers, sequences := map[string]int{}, map[string]int{}
	for _, r := range records {
		headers[r.Header]++
		if r.Sequence != "" {
			sequences[r.Sequence]++
		}
		for _, status := range r.Status {
			counts[status]++
		}
	}
	duplicateHeaders, duplicateSequences := 0, 0
	for _, count := range headers {
		if count > 1 {
			duplicateHeaders++
		}
	}
	for _, count := range sequences {
		if count > 1 {
			duplicateSequences++
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: soybean\n\n- Source file: `%s`\n- URL: %s\n- Source SHA-256: `%s`\n- Container inspected: legacy Word document, 10,775 paragraphs, no tables, 199 pages\n- Normalized text lines: `%d`\n- Accepted named CYP blocks: `%d`\n- Accepted blocks with literal protein sequence: `%d`\n- Accepted blocks without literal protein sequence: `%d`\n- Excluded EST nucleotide headers after the explicit appendix marker: `%d`\n- Duplicate exact-header groups retained: `%d`\n- Duplicate literal-sequence groups retained: `%d`\n- Review status: `complete`\n\n", sourceFile, sourceURL, hash, lineCount, len(records), countSequences(records), len(records)-countSequences(records), excludedESTs, duplicateHeaders, duplicateSequences)
	b.WriteString("## Resource-specific interpretation\n\nThe document is divided at its explicit `955 CYTOCHROME P450 EST FOR Glycine max` heading. Before that heading, 171 `>CYP...` annotation blocks are accepted in source order; 115 contain literal protein paragraphs and 56 provide names and annotations without a protein sequence. After the heading, all `>gi|...` records are nucleotide EST data and are excluded rather than translated. Protein paragraphs are accepted only when the entire paragraph consists of source amino-acid characters plus reviewed coordinate, phase, ampersand, or question-mark notation. Coordinate numbers, parenthesized phase markers, whitespace, and ampersand frameshift separators are layout annotations, not residues. The one `??` residue segment and the one internal stop remain literal. Only terminal stop markers are removed. Pseudogene, fragment, missing-region, and frameshift evidence remains in review status. Duplicate names and sequences remain block-distinct. No sequence is translated, repaired, completed, or obtained externally.\n\n## Status counts\n\n| Status | Records |\n|---|---:|\n")
	names := make([]string, 0, len(counts))
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Fprintf(&b, "| %s | %d |\n", name, counts[name])
	}
	b.WriteString("\n## Representative blocks\n\n| Block | Source line | Symbol | Sequence | Status | Header |\n|---:|---:|---|---:|---|---|\n")
	for _, index := range []int{0, 85, 170} {
		r := records[index]
		fmt.Fprintf(&b, "| %d | %d | %s | %d aa | %s | %s |\n", r.Block, r.SourceLine, md(r.Symbol), len(r.Sequence), md(strings.Join(r.Status, ";")), md(r.Header))
	}
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func countSequences(records []record) int {
	count := 0
	for _, r := range records {
		if r.Sequence != "" {
			count++
		}
	}
	return count
}

func fileHash(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha256.Sum256(data)), nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func md(value string) string { return strings.ReplaceAll(strings.TrimSpace(value), "|", "\\|") }
