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
	sourceFile = "plants-Ricinus.communis.doc"
	sourceURL  = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Ricinus.communis.doc"
	sourceHash = "6778128f2a9c5fafdf97cc10c0d4c00ec384a1fb1e5bacda19e6b9d0867926c0"
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
	xpID               = regexp.MustCompile(`\bXP_[0-9.]+`)
	gbID               = regexp.MustCompile(`\bEEF[0-9.]+`)
	nonKey             = regexp.MustCompile(`[^A-Za-z0-9._-]+`)
)

func main() {
	input := flag.String("input", filepath.Join("raw", "plants-Ricinus.communis.txt"), "Word-normalized Ricinus source text")
	sourceDoc := flag.String("source-doc", filepath.Join("raw", sourceFile), "original downloaded Word document")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "ricinus-communis.csv"), "reviewed CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-ricinus-communis.md"), "review ledger")
	flag.Parse()

	data, err := os.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	records, excluded, err := parseRicinusText(string(data))
	if err != nil {
		panic(err)
	}
	if len(records) != 263 || countSequences(records) != 263 || len(excluded) != 1 {
		panic(fmt.Errorf("Ricinus reviewed block counts changed: accepted=%d sequences=%d excluded=%d", len(records), countSequences(records), len(excluded)))
	}
	hash, err := fileSHA256(*sourceDoc)
	if err != nil {
		panic(err)
	}
	if hash != sourceHash {
		panic(fmt.Errorf("Ricinus source hash changed: %s", hash))
	}
	if err := writeCSV(*out, records, hash); err != nil {
		panic(err)
	}
	if err := writeAudit(*audit, records, excluded, hash); err != nil {
		panic(err)
	}
	fmt.Printf("Ricinus communis: %d literal sequence blocks accepted, %d explicit bacterial contamination block excluded\n", len(records), len(excluded))
}

func parseRicinusText(text string) ([]record, []excludedRecord, error) {
	text = strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n")
	lines := strings.Split(text, "\n")
	var headers []int
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), ">") {
			headers = append(headers, i)
		}
	}
	if len(headers) != 264 {
		return nil, nil, fmt.Errorf("Ricinus header count changed: %d", len(headers))
	}
	var records []record
	var excluded []excludedRecord
	for block, start := range headers {
		end := len(lines)
		if block+1 < len(headers) {
			end = headers[block+1]
		}
		header := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(lines[start]), ">"))
		if strings.HasPrefix(header, "gi|255594353|ref|XP_002536077.1") {
			excluded = append(excluded, excludedRecord{Block: block + 1, SourceLine: start + 1, Header: header, Reason: "source section and annotations explicitly identify this block as bacterial contamination"})
			continue
		}
		fields := strings.Fields(header)
		if len(fields) == 0 || !strings.HasPrefix(fields[0], "CYP") {
			return nil, nil, fmt.Errorf("Ricinus block %d has unexpected non-CYP header %q", block+1, header)
		}
		r := record{Block: block + 1, SourceLine: start + 1, Header: header, Symbol: strings.Trim(fields[0], "|,;")}
		r.ID = xpID.FindString(header)
		if r.ID == "" {
			r.ID = gbID.FindString(header)
		}
		if r.ID == "" {
			r.ID = r.Symbol
		}
		r.RecordKey = fmt.Sprintf("ricinus-communis:block-%04d:%s", block+1, keyPart(r.ID))

		var parts []string
		for i := start + 1; i < end; i++ {
			sequence, ok := ricinusSequenceLine(lines[i])
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
			return nil, nil, fmt.Errorf("Ricinus block %d (%s) has no literal sequence", block+1, header)
		}
		lowerBlock := strings.ToLower(strings.Join(lines[start:end], "\n"))
		flags := map[string]bool{}
		if strings.HasSuffix(strings.ToUpper(r.Symbol), "P") || strings.Contains(lowerBlock, "pseudogene") {
			flags["source-pseudogene"] = true
		}
		for _, phrase := range []string{"fragment", "incomplete", "missing", "top part only", "n-term", "c-term", "runs off", "exon 1", "exon 2"} {
			if strings.Contains(lowerBlock, phrase) {
				flags["source-fragment-or-missing-region"] = true
			}
		}
		if strings.Contains(lowerBlock, "frameshift") {
			flags["source-frameshift"] = true
		}
		if strings.Contains(lowerBlock, "join") || strings.Contains(strings.Join(lines[start:end], "\n"), "&") {
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
		if len(r.Sequence) < 350 {
			flags["short-sequence"] = true
		}
		for _, residue := range r.Sequence {
			if !strings.ContainsRune("ACDEFGHIKLMNOPQRSTVWXYx*-", residue) {
				return nil, nil, fmt.Errorf("Ricinus block %d has unexpected residue %q", block+1, residue)
			}
		}
		for _, name := range []string{"source-pseudogene", "source-fragment-or-missing-region", "source-frameshift", "source-joined-piece", "source-gap-annotation", "internal-stop", "ambiguous-X-or-x", "short-sequence"} {
			if flags[name] {
				r.Status = append(r.Status, name)
			}
		}
		records = append(records, r)
	}
	return records, excluded, nil
}

func ricinusSequenceLine(raw string) (string, bool) {
	value := strings.TrimSpace(raw)
	if value == "" || value == "EXON 2" {
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
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	_ = w.Write([]string{"category", "species", "symbol", "id", "record_key", "source_url", "source_note", "sequence", "source_file", "source_sha256", "source_block", "source_line", "source_header", "review_status"})
	for _, r := range records {
		note := fmt.Sprintf("Ricinus Word block %d at normalized source line %d; literal protein lines %d-%d; header=%s", r.Block, r.SourceLine, r.FirstSequenceLine, r.LastSequenceLine, r.Header)
		if len(r.Annotations) > 0 {
			note += "; annotations: " + strings.Join(r.Annotations, " / ")
		}
		_ = w.Write([]string{"plants", "Ricinus communis", r.Symbol, r.ID, r.RecordKey, sourceURL, note, r.Sequence, sourceFile, hash, strconv.Itoa(r.Block), strconv.Itoa(r.SourceLine), r.Header, strings.Join(r.Status, ";")})
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
	duplicateSequences := 0
	seenSequence := map[string]int{}
	for _, r := range records {
		seenSequence[r.Sequence]++
	}
	for _, count := range seenSequence {
		if count > 1 {
			duplicateSequences++
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Ricinus communis\n\n- Source file: `%s`\n- URL: %s\n- Source SHA-256: `%s`\n- Container inspected: legacy Word document, 3,733 paragraphs, no tables, 80 pages\n- Normalized split lines: `3,733`\n- Source `>` blocks: `264`\n- Accepted Ricinus CYP blocks with literal sequence: `%d`\n- Explicit bacterial contamination blocks excluded: `%d`\n- Duplicate literal-sequence groups retained: `%d`\n- Review status: `complete`\n\n", sourceFile, sourceURL, hash, len(records), len(excluded), duplicateSequences)
	b.WriteString("## Resource-specific interpretation\n\nThe title page says `261 sequences`, after combining many RefSeq duplicates. The literal document contains 264 `>` blocks: 263 CYP-labeled Ricinus blocks and one final accession-only block under the explicit `Contamination` heading whose own annotations say its top BLAST hits are bacteria. The contamination block is excluded. All 263 CYP blocks are retained in source order, including alternate entries and the two CYP71B64 blocks whose literal protein sequences are identical. The source's 261 claim is recorded as a discrepancy; records are not silently merged to force that count.\n\nProtein lines in this document may carry leading/trailing coordinates, phase markers, and `&` joins. Those reviewed markers are removed while literal residue text remains. Lowercase `x`, uppercase `X`, and internal stops are retained. Only a final `*` is removed. The exact all-uppercase text `EXON 2` is an annotation, not four residues, and is explicitly rejected. Other lowercase prose cannot become sequence. No fragment is translated, repaired, extended, or filled from another database.\n\nThe preferred record ID is the source's XP accession, then its EEF accession, otherwise the CYP symbol. Duplicate names and sequences remain independent through the source-block number in `RecordKey`.\n\n## Status counts\n\n| Status | Records |\n|---|---:|\n")
	for _, name := range names {
		fmt.Fprintf(&b, "| %s | %d |\n", name, counts[name])
	}
	b.WriteString("\n## Representative records\n\n| Accepted index | Source block | Source line | ID | CYP symbol | Sequence | Status | Header |\n|---:|---:|---:|---|---|---:|---|---|\n")
	for _, index := range []int{0, len(records) / 2, len(records) - 1} {
		r := records[index]
		fmt.Fprintf(&b, "| %d | %d | %d | %s | %s | %d aa | %s | %s |\n", index+1, r.Block, r.SourceLine, md(r.ID), md(r.Symbol), len(r.Sequence), md(strings.Join(r.Status, ";")), md(r.Header))
	}
	b.WriteString("\n## Excluded block\n\n| Source block | Source line | Header | Reason |\n|---:|---:|---|---|\n")
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
