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
	sourceFile = "plants-Jatropha.P450s.2012.doc"
	sourceURL  = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Jatropha.P450s.2012.doc"
	sourceHash = "ab0a6347736fdd21093c00a5597fbdf52b9b040e487daee9e2f16c0f2b85c9d3"
)

type record struct {
	Block, SourceLine, FirstSequenceLine, LastSequenceLine int
	Header, ID, Symbol, RecordKey, Sequence                string
	Status, Annotations                                    []string
}

type excludedRecord struct {
	Block, SourceLine int
	Header, Reason    string
}

var (
	leadingCoordinate  = regexp.MustCompile(`^\d+\s+`)
	trailingCoordinate = regexp.MustCompile(`\s+\d+\s*&?\s*$`)
	phaseMarker        = regexp.MustCompile(`\((?:0|1|2|\?)?\)`)
	whitespace         = regexp.MustCompile(`\s+`)
	headerID           = regexp.MustCompile(`\b(?:Jc[A-Za-z0-9.]+|JHS[A-Za-z0-9.]+|BABX[A-Za-z0-9.]+)\b`)
	headerCYP          = regexp.MustCompile(`\bCYP[0-9A-Za-z]+`)
	nonKey             = regexp.MustCompile(`[^A-Za-z0-9._-]+`)
)

func main() {
	input := flag.String("input", filepath.Join("raw", "plants-Jatropha.P450s.2012.txt"), "Word-normalized Jatropha source text")
	sourceDoc := flag.String("source-doc", filepath.Join("raw", sourceFile), "original downloaded Word document")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "jatropha-curcas.csv"), "reviewed CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-jatropha-curcas.md"), "review ledger")
	flag.Parse()

	data, err := os.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	records, excluded, err := parseJatrophaText(string(data))
	if err != nil {
		panic(err)
	}
	if len(records) != 481 || countSequences(records) != 481 || len(excluded) != 56 {
		panic(fmt.Errorf("Jatropha reviewed block counts changed: accepted=%d sequences=%d excluded=%d", len(records), countSequences(records), len(excluded)))
	}
	hash, err := fileSHA256(*sourceDoc)
	if err != nil {
		panic(err)
	}
	if hash != sourceHash {
		panic(fmt.Errorf("Jatropha source hash changed: %s", hash))
	}
	if err := writeCSV(*out, records, hash); err != nil {
		panic(err)
	}
	if err := writeAudit(*audit, records, excluded, hash); err != nil {
		panic(err)
	}
	fmt.Printf("Jatropha curcas: %d literal sequence records accepted, %d foreign/structural/false-positive blocks excluded\n", len(records), len(excluded))
}

func parseJatrophaText(text string) ([]record, []excludedRecord, error) {
	text = strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n")
	lines := strings.Split(text, "\n")
	var headers []int
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), ">") {
			headers = append(headers, i)
		}
	}
	if len(headers) != 537 {
		return nil, nil, fmt.Errorf("Jatropha header count changed: %d", len(headers))
	}
	falsePositiveHeading := -1
	for i, line := range lines {
		if strings.TrimSpace(line) == "20 False positive hits (3.7%)" {
			falsePositiveHeading = i
			break
		}
	}
	if falsePositiveHeading < 0 {
		return nil, nil, fmt.Errorf("Jatropha false-positive heading missing")
	}

	var accepted []record
	var excluded []excludedRecord
	for block, start := range headers {
		end := len(lines)
		if block+1 < len(headers) {
			end = headers[block+1]
		}
		header := strings.TrimSpace(lines[start])
		plain := strings.TrimSpace(strings.TrimPrefix(header, ">"))
		reason := ""
		switch {
		case start > falsePositiveHeading:
			reason = "block is below the document's explicit `20 False positive hits` heading"
		case strings.Contains(header, "Ricinus communis"):
			reason = "explicit Ricinus communis assembly/comparison block"
		case strings.Contains(header, "Populus trichocarpa"):
			reason = "explicit Populus trichocarpa comparison block"
		case plain == "CYP735A22":
			reason = "name-only preface; the immediately following BABX01044566.1 block carries its literal assembled sequence and inherits this source name"
		case plain == "CYP90C (one sequence)" || plain == "CYP90D (one sequence plus one pseudogene)":
			reason = "family count heading formatted with `>`; the following named blocks carry the records and sequences"
		}
		if reason != "" {
			excluded = append(excluded, excludedRecord{Block: block + 1, SourceLine: start + 1, Header: plain, Reason: reason})
			continue
		}

		r := record{Block: block + 1, SourceLine: start + 1, Header: plain}
		if match := headerID.FindString(plain); match != "" {
			r.ID = match
		}
		if match := headerCYP.FindString(plain); match != "" {
			r.Symbol = match
		}
		// The document gives the CYP735A22 name on a dedicated `>` preface,
		// followed immediately by the accession-bearing block with the sequence.
		if plain == "BABX01044566.1" {
			r.Symbol = "CYP735A22"
		}
		if r.ID == "" {
			r.ID = r.Symbol
		}
		if r.ID == "" {
			fields := strings.Fields(plain)
			if len(fields) > 0 {
				r.ID = strings.Trim(fields[0], "|,;")
			}
		}
		if r.ID == "" {
			return nil, nil, fmt.Errorf("Jatropha block %d has no literal source identifier", block+1)
		}
		r.RecordKey = fmt.Sprintf("jatropha-curcas:block-%04d:%s", block+1, keyPart(r.ID))

		blockText := strings.Join(lines[start:end], "\n")
		var parts []string
		for i := start + 1; i < end; i++ {
			sequence, ok := jatrophaSequenceLine(lines[i])
			if ok {
				if r.FirstSequenceLine == 0 {
					r.FirstSequenceLine = i + 1
				}
				r.LastSequenceLine = i + 1
				parts = append(parts, sequence)
				continue
			}
			if value := strings.TrimSpace(lines[i]); value != "" && len(r.Annotations) < 8 {
				r.Annotations = append(r.Annotations, value)
			}
		}
		rawSequence := strings.Join(parts, "")
		if strings.HasSuffix(rawSequence, "*") {
			rawSequence = strings.TrimSuffix(rawSequence, "*")
			r.Status = append(r.Status, "source-terminal-stop")
		}
		r.Sequence = rawSequence
		if r.Sequence == "" {
			return nil, nil, fmt.Errorf("Jatropha block %d (%s) has no literal sequence", block+1, plain)
		}
		lowerBlock := strings.ToLower(blockText)
		flags := map[string]bool{}
		if strings.Contains(lowerBlock, "pseudo") {
			flags["source-pseudogene"] = true
		}
		for _, phrase := range []string{"partial", "short", "n-term", "nterm", "c-term", "cterm", "fragment", "runs off", "missing", "exon 1", "exon 2", "mid region", "i-helix"} {
			if strings.Contains(lowerBlock, phrase) {
				flags["source-fragment-or-missing-region"] = true
			}
		}
		if strings.Contains(lowerBlock, "frameshift") {
			flags["source-frameshift"] = true
		}
		if strings.Contains(blockText, "&") {
			flags["source-joined-piece"] = true
		}
		if strings.Contains(lowerBlock, "gap") {
			flags["source-gap-annotation"] = true
		}
		if strings.Contains(r.Sequence, "*") {
			flags["internal-stop"] = true
		}
		if strings.ContainsAny(r.Sequence, "Xx") {
			flags["ambiguous-X-or-x"] = true
		}
		if strings.Contains(r.Sequence, "O") {
			flags["nonstandard-O"] = true
		}
		if len(r.Sequence) < 350 {
			flags["short-sequence"] = true
		}
		for _, residue := range r.Sequence {
			if !strings.ContainsRune("ACDEFGHIKLMNOPQRSTVWXYx*-", residue) {
				return nil, nil, fmt.Errorf("Jatropha block %d has unexpected residue %q", block+1, residue)
			}
		}
		for _, name := range []string{"source-pseudogene", "source-fragment-or-missing-region", "source-frameshift", "source-joined-piece", "source-gap-annotation", "internal-stop", "ambiguous-X-or-x", "nonstandard-O", "short-sequence"} {
			if flags[name] {
				r.Status = append(r.Status, name)
			}
		}
		accepted = append(accepted, r)
	}
	return accepted, excluded, nil
}

// jatrophaSequenceLine implements only the line forms observed in this Word
// resource. COMPLETE and FINISHED are explicit annotations even though their
// letters happen to be amino-acid codes. Other lowercase prose is rejected;
// lowercase x remains literal source uncertainty. Coordinates, phase markers,
// whitespace, and source `&` join marks are metadata around literal residues.
func jatrophaSequenceLine(raw string) (string, bool) {
	value := strings.TrimSpace(raw)
	if value == "" || value == "COMPLETE" || value == "FINISHED" {
		return "", false
	}
	for _, r := range value {
		if r >= 'a' && r <= 'w' || r == 'y' || r == 'z' {
			return "", false
		}
	}
	value = leadingCoordinate.ReplaceAllString(value, "")
	value = trailingCoordinate.ReplaceAllString(value, "")
	value = phaseMarker.ReplaceAllString(value, "")
	value = strings.ReplaceAll(value, "&", "")
	value = whitespace.ReplaceAllString(value, "")
	if value == "" {
		return "", false
	}
	for _, residue := range value {
		if !strings.ContainsRune("ACDEFGHIKLMNOPQRSTVWXYx*-", residue) {
			return "", false
		}
	}
	return value, true
}

func writeCSV(path string, records []record, hash string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	w := csv.NewWriter(file)
	_ = w.Write([]string{"category", "species", "symbol", "id", "record_key", "source_url", "source_note", "sequence", "source_file", "source_sha256", "source_block", "source_line", "source_header", "review_status"})
	for _, r := range records {
		note := fmt.Sprintf("Jatropha Word block %d at normalized source line %d; literal protein lines %d-%d; header=%s", r.Block, r.SourceLine, r.FirstSequenceLine, r.LastSequenceLine, r.Header)
		if len(r.Annotations) > 0 {
			note += "; annotations: " + strings.Join(r.Annotations, " / ")
		}
		_ = w.Write([]string{"plants", "Jatropha curcas", r.Symbol, r.ID, r.RecordKey, sourceURL, note, r.Sequence, sourceFile, hash, strconv.Itoa(r.Block), strconv.Itoa(r.SourceLine), r.Header, strings.Join(r.Status, ";")})
	}
	w.Flush()
	return w.Error()
}

func writeAudit(path string, records []record, excluded []excludedRecord, hash string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	counts := map[string]int{}
	for _, r := range records {
		for _, status := range r.Status {
			counts[status]++
		}
	}
	names := make([]string, 0, len(counts))
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	excludedCounts := map[string]int{}
	for _, r := range excluded {
		switch {
		case strings.Contains(r.Reason, "False positive"):
			excludedCounts["explicit false positives"]++
		case strings.Contains(r.Reason, "Ricinus"):
			excludedCounts["Ricinus comparison/assembly blocks"]++
		case strings.Contains(r.Reason, "Populus"):
			excludedCounts["Populus comparison blocks"]++
		default:
			excludedCounts["structural/name-only `>` headings"]++
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Jatropha curcas\n\n- Source file: `%s`\n- URL: %s\n- Source SHA-256: `%s`\n- Container inspected: legacy Word document, 7,107 paragraphs, no tables, 129 pages\n- Normalized split lines: `7,108`\n- Source `>` blocks: `537`\n- Accepted Jatropha records with literal sequence: `%d`\n- Excluded foreign, structural, or false-positive blocks: `%d`\n- Review status: `complete`\n\n", sourceFile, sourceURL, hash, len(records), len(excluded))
	b.WriteString("## Resource-specific interpretation\n\nThe title page says 484 Jatropha sequences after subtracting 31 Ricinus helpers and 20 false positives. Complete block inspection finds 537 `>` blocks: 32 headers explicitly identify Ricinus communis, one identifies Populus trichocarpa, and 20 occur below the document's exact `20 False positive hits (3.7%)` heading. Three additional `>` lines are not independent sequence records: `CYP90C (one sequence)` and `CYP90D (one sequence plus one pseudogene)` are family-count headings, while the name-only `CYP735A22` preface applies to the immediately following `BABX01044566.1` sequence block. Therefore the literal file contains 481 accepted Jatropha sequence records; the source's 484 claim is retained here as an explicit discrepancy rather than manufactured as empty or duplicate records.\n\nFor this document only, protein lines may contain separated leading/trailing genomic coordinates, `(0)`/`(1)`/`(2)`/`(?)`/`()` phase markers, and `&` joins. Those markers are removed while the residue text on the same line is retained. Lowercase `x`, uppercase `X`, `O`, and internal stops remain literal. Only an explicit final `*` on a record is removed; 175 records carry that marker. The words `COMPLETE` and `FINISHED` are annotations even though every letter is also an amino-acid code; they are explicitly rejected. Other lowercase prose cannot become sequence. No fragment is translated, repaired, extended, merged across `>` blocks, or filled from Ricinus or another database.\n\nAccession-only Jatropha headers remain independent records with an empty CYP symbol, except `BABX01044566.1`, whose source name is supplied by the immediately preceding name-only `CYP735A22` preface. Duplicate accessions and adjacent-gene pieces remain block-distinct through unique `RecordKey` values.\n\n## Status counts\n\n| Status | Records |\n|---|---:|\n")
	for _, name := range names {
		fmt.Fprintf(&b, "| %s | %d |\n", name, counts[name])
	}
	b.WriteString("\n## Exclusion counts\n\n| Reason | Blocks |\n|---|---:|\n")
	for _, name := range []string{"Ricinus comparison/assembly blocks", "Populus comparison blocks", "structural/name-only `>` headings", "explicit false positives"} {
		fmt.Fprintf(&b, "| %s | %d |\n", name, excludedCounts[name])
	}
	b.WriteString("\n## Representative records\n\n| Accepted index | Source block | Source line | ID | CYP symbol | Sequence | Status | Header |\n|---:|---:|---:|---|---|---:|---|---|\n")
	for _, index := range []int{0, len(records) / 2, len(records) - 1} {
		r := records[index]
		fmt.Fprintf(&b, "| %d | %d | %d | %s | %s | %d aa | %s | %s |\n", index+1, r.Block, r.SourceLine, md(r.ID), md(r.Symbol), len(r.Sequence), md(strings.Join(r.Status, ";")), md(r.Header))
	}
	b.WriteString("\n## Excluded blocks\n\n| Source block | Source line | Header | Reason |\n|---:|---:|---|---|\n")
	for _, r := range excluded {
		fmt.Fprintf(&b, "| %d | %d | %s | %s |\n", r.Block, r.SourceLine, md(r.Header), md(r.Reason))
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

func fileSHA256(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha256.Sum256(data)), nil
}

func keyPart(value string) string {
	value = nonKey.ReplaceAllString(strings.TrimSpace(value), "-")
	return strings.ToLower(strings.Trim(value, "-"))
}

func md(value string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(value), "|", "\\|"), "\n", " ")
}
