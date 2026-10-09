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
	sourceFile = "plants-Prunus.persica.P450s.doc"
	sourceURL  = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Prunus.persica.P450s.doc"
)

type record struct {
	Block, SourceLine                    int
	Header, PeptideAccession, GeneID, ID string
	Sequence                             string
	Status                               []string
}

var headerRE = regexp.MustCompile(`^>([0-9]+)_peptide\|Ppersica\|([^|]+)\|([^|]+)$`)

func main() {
	input := flag.String("input", filepath.Join("raw", "plants-Prunus.persica.P450s.txt"), "Word-normalized older peach source text")
	sourceDoc := flag.String("source-doc", filepath.Join("raw", sourceFile), "original downloaded Word document")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "prunus-persica-old.csv"), "reviewed CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-prunus-persica-peach-older-version.md"), "review ledger")
	flag.Parse()
	data, err := os.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	records, lineCount, err := parseText(string(data))
	if err != nil {
		panic(err)
	}
	if len(records) != 306 {
		panic(fmt.Errorf("older Prunus persica block count changed: %d", len(records)))
	}
	hash, err := fileHash(*sourceDoc)
	if err != nil {
		panic(err)
	}
	if err := writeCSV(*out, records, hash); err != nil {
		panic(err)
	}
	if err := writeAudit(*audit, records, hash, lineCount); err != nil {
		panic(err)
	}
	fmt.Printf("Prunus persica older Word resource: %d literal sequence blocks retained\n", len(records))
}

func parseText(text string) ([]record, int, error) {
	return parseTextWithLayout(text, true)
}

func parseTextWithLayout(text string, enforceLayout bool) ([]record, int, error) {
	text = strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n")
	lines := strings.Split(text, "\n")
	if len(lines) < 4 || strings.TrimSpace(lines[0]) != "Prunus Persica P450s" || !strings.Contains(strings.Join(lines[:min(11, len(lines))], "\n"), "305 sequences retrieved") || !strings.Contains(strings.Join(lines[:min(11, len(lines))], "\n"), "have not been annotated manually") {
		return nil, len(lines), fmt.Errorf("older Prunus persica preamble changed")
	}
	var headers []int
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), ">") {
			headers = append(headers, i)
		}
	}
	if enforceLayout && (len(headers) != 306 || headers[0]+1 != 12 || headers[len(headers)-1]+1 != 3058) {
		return nil, len(lines), fmt.Errorf("older Prunus persica header boundary changed: count=%d first=%d last=%d", len(headers), headers[0]+1, headers[len(headers)-1]+1)
	}
	records := make([]record, 0, len(headers))
	for block, start := range headers {
		end := len(lines)
		if block+1 < len(headers) {
			end = headers[block+1]
		}
		header := strings.TrimSpace(lines[start])
		match := headerRE.FindStringSubmatch(header)
		if match == nil {
			return nil, len(lines), fmt.Errorf("older Prunus persica block %d invalid header %q", block+1, header)
		}
		r := record{Block: block + 1, SourceLine: start + 1, Header: header, PeptideAccession: match[1], GeneID: match[2], ID: match[3]}
		var parts []string
		for i := start + 1; i < end; i++ {
			value := strings.TrimSpace(lines[i])
			if value == "" {
				continue
			}
			for _, aa := range value {
				if !strings.ContainsRune("ACDEFGHIKLMNPQRSTVWY*", aa) {
					return nil, len(lines), fmt.Errorf("older Prunus persica block %d line %d unexpected residue %q", block+1, i+1, aa)
				}
			}
			parts = append(parts, value)
		}
		raw := strings.Join(parts, "")
		if raw == "" {
			return nil, len(lines), fmt.Errorf("older Prunus persica block %d has no literal sequence", block+1)
		}
		r.Status = append(r.Status, "source-unannotated")
		if strings.HasSuffix(raw, "*") {
			raw = strings.TrimRight(raw, "*")
			r.Status = append(r.Status, "source-terminal-stop")
		}
		if strings.Contains(raw, "*") {
			r.Status = append(r.Status, "internal-stop")
		}
		if len(raw) < 350 {
			r.Status = append(r.Status, "short-sequence")
		}
		if len(raw) > 650 {
			r.Status = append(r.Status, "unusually-long-sequence")
		}
		r.Sequence = raw
		records = append(records, r)
	}
	return records, len(lines), nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
		note := fmt.Sprintf("Older Prunus persica Word block %d at normalized source line %d; peptide accession=%s; gene ID=%s; header=%s", r.Block, r.SourceLine, r.PeptideAccession, r.GeneID, r.Header)
		_ = w.Write([]string{"plants", "Prunus persica", "", r.ID, fmt.Sprintf("prunus-persica-old:block-%04d:%s", r.Block, r.ID), sourceURL, note, r.Sequence, sourceFile, hash, strconv.Itoa(r.Block), strconv.Itoa(r.SourceLine), strings.Join(r.Status, ";")})
	}
	w.Flush()
	return w.Error()
}

func writeAudit(path string, records []record, hash string, lineCount int) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	counts := map[string]int{}
	headers := map[string][]int{}
	sequences := map[string][]int{}
	for _, r := range records {
		for _, status := range r.Status {
			counts[status]++
		}
		headers[r.Header] = append(headers[r.Header], r.Block)
		sequences[r.Sequence] = append(sequences[r.Sequence], r.Block)
	}
	duplicateHeaders, duplicateSequences := 0, 0
	for _, blocks := range headers {
		if len(blocks) > 1 {
			duplicateHeaders++
		}
	}
	for _, blocks := range sequences {
		if len(blocks) > 1 {
			duplicateSequences++
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Prunus persica older version\n\n- Source file: `%s`\n- URL: %s\n- Source SHA-256: `%s`\n- Container inspected: legacy Word document, 3,066 paragraphs, no tables, 49 pages\n- Normalized text lines: `%d`\n- Header-delimited sequence blocks retained: `%d`\n- Source title claim: `305 sequences`\n- Duplicate exact-header groups retained: `%d`\n- Duplicate literal-sequence groups retained: `%d`\n- Review status: `complete`\n\n", sourceFile, sourceURL, hash, lineCount, len(records), duplicateHeaders, duplicateSequences)
	b.WriteString("## Resource-specific interpretation\n\nThe document states that 305 sequences were retrieved by a Phytozome Biomart P450-keyword search and explicitly says they were not checked for accuracy or annotated manually. The body actually contains 306 complete `>accession_peptide|Ppersica|gene|transcript` blocks, each followed only by literal uppercase protein lines. Blocks 54/55 and 156/157 are two exact duplicate header-and-sequence pairs; all four blocks remain source-order records with distinct `RecordKey` values. Because this old resource provides no CYP assignment, `symbol` remains empty and is not inferred from the updated peach workbook. The transcript field is used as ID, while peptide accession and gene ID remain in provenance. Explicit terminal `*` markers are removed; all other source sequence content is retained. No sequence is repaired, annotated, merged, or deduplicated.\n\n## Status counts\n\n| Status | Records |\n|---|---:|\n")
	names := make([]string, 0, len(counts))
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Fprintf(&b, "| %s | %d |\n", name, counts[name])
	}
	b.WriteString("\n## Representative blocks\n\n| Block | Source line | Transcript ID | Sequence | Status | Header |\n|---:|---:|---|---:|---|---|\n")
	for _, index := range []int{0, 153, 305} {
		r := records[index]
		fmt.Fprintf(&b, "| %d | %d | %s | %d aa | %s | %s |\n", r.Block, r.SourceLine, md(r.ID), len(r.Sequence), md(strings.Join(r.Status, ";")), md(r.Header))
	}
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func fileHash(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha256.Sum256(data)), nil
}

func md(value string) string { return strings.ReplaceAll(strings.TrimSpace(value), "|", "\\|") }
