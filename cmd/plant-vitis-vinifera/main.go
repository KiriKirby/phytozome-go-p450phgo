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
	sourceFile = "plants-vitis.doc"
	sourceURL  = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/vitis.doc"
	sourceHash = "37e8b56e9a9db299205b5ab80e0985697b55c4c21d14e2c5194fe8f255c4ec42"
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
	trailingCoordinate = regexp.MustCompile(`\s+\d+(?:\s+(?:aa|a)\s+\d+(?:-\d+)?)?\s*&?\s*$`)
	trailingAAComment  = regexp.MustCompile(`\s+(?:aa|a)\s+\d+(?:-\d+)?\s*$`)
	phaseMarker        = regexp.MustCompile(`\((?:0|1|2|\?|0\?|1\?|2\?)?\)`)
	whitespace         = regexp.MustCompile(`\s+`)
	accession          = regexp.MustCompile(`\b(?:CAAP|CAN|CAO|CAI|ABC|ABH|BAE|CAB|AAP)[0-9]+(?:\.[0-9]+)?\b|\bAM[0-9]+(?:\.[0-9]+)?\b`)
	headerCYP          = regexp.MustCompile(`\bCYP[0-9A-Za-z]+(?:[-.][0-9A-Za-z]+)*`)
	pseudogeneSymbol   = regexp.MustCompile(`(?i)^CYP[0-9A-Z]*P(?:V[0-9]+|X|[-.]|$)`)
	nonKey             = regexp.MustCompile(`[^A-Za-z0-9._-]+`)
)

// These are the exact source block numbers of reference sequences embedded to
// establish names for Vitis families. The list was reviewed block-by-block;
// comparisons merely mentioned in a Vitis header are intentionally not here.
var foreignReferenceBlocks = map[int]string{
	5: "Nicotiana tabacum naming reference", 6: "Nicotiana tabacum naming reference",
	7: "Glycine max naming reference", 10: "Gossypium raimondii naming reference",
	71: "Solanum tuberosum naming reference", 72: "Solanum lycopersicum naming reference",
	73: "Gossypium naming reference", 74: "Citrus hybrid naming reference",
	338: "Medicago sativa naming reference", 339: "Populus naming reference",
	340: "Coptis japonica naming reference", 365: "Ammi majus naming reference",
	401: "Populus naming reference", 402: "Populus naming reference",
	403: "Populus naming reference", 404: "Populus naming reference",
	593: "Arabidopsis naming reference", 594: "Populus naming reference",
	595: "rice naming reference", 596: "rice naming reference", 597: "Lolium naming reference",
	598: "rice naming reference", 599: "Medicago truncatula naming reference",
	600: "Populus naming reference", 601: "Medicago truncatula naming reference",
	602: "Populus naming reference", 603: "Populus naming reference",
	604: "Populus naming reference", 605: "Petunia hybrida naming reference",
	606: "Populus naming reference",
}

func main() {
	input := flag.String("input", filepath.Join("raw", "plants-vitis.txt"), "Word-normalized Vitis source text")
	sourceDoc := flag.String("source-doc", filepath.Join("raw", sourceFile), "original downloaded Word document")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "vitis-vinifera.csv"), "reviewed CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-vitis-vinifera.md"), "review ledger")
	flag.Parse()

	data, err := os.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	records, excluded, err := parseVitisText(string(data))
	if err != nil {
		panic(err)
	}
	if len(records) != 672 || len(excluded) != 30 {
		panic(fmt.Errorf("Vitis reviewed block counts changed: accepted=%d sequences=%d excluded=%d", len(records), countSequences(records), len(excluded)))
	}
	hash, err := fileSHA256(*sourceDoc)
	if err != nil {
		panic(err)
	}
	if hash != sourceHash {
		panic(fmt.Errorf("Vitis source hash changed: %s", hash))
	}
	if err := writeCSV(*out, records, hash); err != nil {
		panic(err)
	}
	if err := writeAudit(*audit, records, excluded, hash); err != nil {
		panic(err)
	}
	fmt.Printf("Vitis vinifera: %d source-native blocks accepted (%d with literal sequence), %d explicit foreign reference blocks excluded\n", len(records), countSequences(records), len(excluded))
}

func parseVitisText(text string) ([]record, []excludedRecord, error) {
	text = strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n")
	lines := strings.Split(text, "\n")
	var headers []int
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), ">") {
			headers = append(headers, i)
		}
	}
	if len(headers) != 702 {
		return nil, nil, fmt.Errorf("Vitis header count changed: %d", len(headers))
	}

	var records []record
	var excluded []excludedRecord
	for blockIndex, start := range headers {
		block := blockIndex + 1
		end := len(lines)
		if blockIndex+1 < len(headers) {
			end = headers[blockIndex+1]
		}
		header := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(lines[start]), ">"))
		if reason, ok := foreignReferenceBlocks[block]; ok {
			excluded = append(excluded, excludedRecord{Block: block, SourceLine: start + 1, Header: header, Reason: reason})
			continue
		}
		r := record{Block: block, SourceLine: start + 1, Header: header}
		r.Symbol = headerCYP.FindString(header)
		if r.Symbol == "" && strings.HasPrefix(header, "71A9/CYP71AH3") {
			r.Symbol = "CYP71AH3"
		}
		r.ID = accession.FindString(header)
		if r.ID == "" {
			r.ID = r.Symbol
		}
		if r.ID == "" {
			fields := strings.Fields(header)
			if len(fields) > 0 {
				r.ID = strings.Trim(fields[0], "|,;")
			}
		}
		if r.ID == "" {
			return nil, nil, fmt.Errorf("Vitis block %d has no source identifier", block)
		}
		r.RecordKey = fmt.Sprintf("vitis-vinifera:block-%04d:%s", block, keyPart(r.ID))

		var parts []string
		for i := start + 1; i < end; i++ {
			// Block 693 has a complete Vitis CYP sequence ending on line 8067,
			// followed by an unrelated unheaded protein before block 694.
			if block == 693 && i+1 >= 8069 {
				continue
			}
			sequence, ok := vitisSequenceLine(lines[i])
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
			return nil, nil, fmt.Errorf("Vitis block %d (%s) has no literal sequence", block, header)
		}
		blockText := strings.Join(lines[start:end], "\n")
		lowerBlock := strings.ToLower(blockText)
		flags := map[string]bool{}
		if pseudogeneSymbol.MatchString(r.Symbol) || strings.Contains(lowerBlock, "pseudogene") {
			flags["source-pseudogene"] = true
		}
		for _, phrase := range []string{"fragment", "partial", "missing", "runs off", "n-term", "c-term", "exon 1 only", "exon 2 only"} {
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
				return nil, nil, fmt.Errorf("Vitis block %d has unexpected residue %q", block, residue)
			}
		}
		for _, name := range []string{"source-pseudogene", "source-fragment-or-missing-region", "source-frameshift", "source-joined-piece", "source-gap-annotation", "internal-stop", "ambiguous-X-or-x", "nonstandard-O", "short-sequence"} {
			if flags[name] {
				r.Status = append(r.Status, name)
			}
		}
		records = append(records, r)
	}
	return records, excluded, nil
}

// vitisSequenceLine implements the exact line grammar observed in this Word
// source: optional genomic coordinates, literal residue groups, reviewed phase
// markers, whitespace, and source join marks. Lowercase prose and the source's
// separators are never residues.
func vitisSequenceLine(raw string) (string, bool) {
	value := strings.TrimSpace(raw)
	if value == "" || value == "$$$$" || strings.Trim(value, "&") == "" || value == "(GAP)" {
		return "", false
	}
	// A few fragment lines append exact `aa 324-339` / `a 369-377`
	// coordinate comments after the trailing genomic coordinate.
	value = trailingAAComment.ReplaceAllString(value, "")
	for _, r := range value {
		if (r >= 'a' && r <= 'w') || r == 'y' || r == 'z' {
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
		note := fmt.Sprintf("Vitis Word block %d at normalized source line %d; literal protein lines %d-%d; header=%s", r.Block, r.SourceLine, r.FirstSequenceLine, r.LastSequenceLine, r.Header)
		if len(r.Annotations) > 0 {
			note += "; annotations: " + strings.Join(r.Annotations, " / ")
		}
		_ = w.Write([]string{"plants", "Vitis vinifera", r.Symbol, r.ID, r.RecordKey, sourceURL, note, r.Sequence, sourceFile, hash, strconv.Itoa(r.Block), strconv.Itoa(r.SourceLine), r.Header, strings.Join(r.Status, ";")})
	}
	w.Flush()
	return w.Error()
}

func writeAudit(path string, records []record, excluded []excludedRecord, hash string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	counts := map[string]int{}
	seenSequence := map[string]int{}
	for _, r := range records {
		for _, s := range r.Status {
			counts[s]++
		}
		seenSequence[r.Sequence]++
	}
	names := make([]string, 0, len(counts))
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	duplicateGroups := 0
	for _, n := range seenSequence {
		if n > 1 {
			duplicateGroups++
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Vitis vinifera\n\n- Source file: `%s`\n- URL: %s\n- Source SHA-256: `%s`\n- Container inspected: legacy Word document, 8,443 paragraphs, 138 pages\n- Normalized split lines: `8,443`\n- Source `>` blocks: `702`\n- Accepted source-native Vitis blocks: `%d`\n- Accepted blocks with literal sequence: `%d`\n- Explicit foreign naming-reference blocks excluded: `%d`\n- Duplicate literal-sequence groups retained: `%d`\n- Review status: `complete`\n\n", sourceFile, sourceURL, hash, len(records), countSequences(records), len(excluded), duplicateGroups)
	b.WriteString("## Resource-specific interpretation\n\nThe document is a working 2007 curation file, not plain FASTA. It interleaves Vitis WGS assemblies, GenPept entries, alleles, duplicates, fragments, pseudogenes, and explicit sequences from other plant species used as family-naming references. All 702 headers were reviewed in source order. The 30 blocks listed below are explicit foreign reference sequences and are excluded; Vitis headers that merely compare identity to another species remain included.\n\nThe title-page statement `591 sequences are present below` predates later revisions and does not match the current 702 literal blocks. The current file contains many alternate genome-project, allele, duplicate, and revised records. They are retained independently instead of being silently merged to force the historical count. Block 693 has a complete terminal-star-bounded CYP736A21 sequence followed by an unrelated unheaded protein; only its CYP sequence through normalized line 8067 is accepted.\n\nReviewed protein lines may carry leading/trailing genomic coordinates, phase markers, whitespace, and `&` joins. Those layout markers are removed; literal X/x, O, gap characters, and internal stops are preserved. Only one final `*` is removed. `$$$$`, ampersand separators, `(GAP)`, and lowercase prose are annotations. Nothing is translated, repaired, or externally completed.\n\n## Status counts\n\n| Status | Records |\n|---|---:|\n")
	for _, name := range names {
		fmt.Fprintf(&b, "| %s | %d |\n", name, counts[name])
	}
	b.WriteString("\n## Representative records\n\n| Accepted index | Source block | Source line | ID | CYP symbol | Sequence | Status | Header |\n|---:|---:|---:|---|---|---:|---|---|\n")
	for _, index := range []int{0, len(records) / 2, len(records) - 1} {
		r := records[index]
		fmt.Fprintf(&b, "| %d | %d | %d | %s | %s | %d aa | %s | %s |\n", index+1, r.Block, r.SourceLine, md(r.ID), md(r.Symbol), len(r.Sequence), md(strings.Join(r.Status, ";")), md(r.Header))
	}
	b.WriteString("\n## Excluded foreign reference blocks\n\n| Source block | Source line | Header | Reason |\n|---:|---:|---|---|\n")
	for _, r := range excluded {
		fmt.Fprintf(&b, "| %d | %d | %s | %s |\n", r.Block, r.SourceLine, md(r.Header), md(r.Reason))
	}
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func countSequences(records []record) int {
	n := 0
	for _, r := range records {
		if r.Sequence != "" {
			n++
		}
	}
	return n
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
