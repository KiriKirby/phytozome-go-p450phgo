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
	sourceFile = "plants-medicago.seqs.doc"
	sourceURL  = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/medicago.seqs.doc"
)

type record struct {
	Block, SourceLine, FirstSequenceLine, LastSequenceLine int
	Header, ID, Symbol, Sequence                           string
	Status                                                 []string
}

type exclusion struct {
	HeaderPrefix string
	Reason       string
}

var (
	coordinateToken = regexp.MustCompile(`^\d+$`)
	phaseToken      = regexp.MustCompile(`^\((?:[012?]|[012]\?)?\)$`)
	proteinToken    = regexp.MustCompile(`^[ACDEFGHIKLMNOPQRSTVWXYacdefghiklmnopqrstvwyx*?]+$`)
	digitP          = regexp.MustCompile(`(?i)\dP(?:$|[-/])`)
	cypID           = regexp.MustCompile(`^CYP\d`)

	// The source explicitly says that some of its 376 sequence pieces are
	// other-species assembly helpers. These 24 blocks were reviewed one by
	// one and are excluded from the Medicago resource. The final three blocks
	// are under the document's literal "False positives" heading.
	excludedBlocks = map[int]exclusion{
		64:  {">CYP73A3 Medicago sativa", "foreign Medicago sativa comparison"},
		69:  {">CYP76E1 From Chris Steele", "foreign Medicago sativa/alfalfa comparison"},
		72:  {">CYP76E2 TC107627", "foreign Medicago sativa assembly helper"},
		73:  {">76F5 C Steele", "foreign Medicago sativa/alfalfa comparison"},
		76:  {">76F6 Length", "foreign Medicago sativa/alfalfa comparison"},
		128: {">CYP82D1 C Steele", "foreign Medicago sativa/alfalfa comparison"},
		132: {">CYP83E1 C Steele", "foreign Medicago sativa/alfalfa comparison"},
		151: {">CYP84A19 CYP84Ms1", "foreign Medicago sativa comparison"},
		152: {">CYP84A20 CYP84Ms2", "foreign Medicago sativa comparison"},
		179: {">CYP93C7v1 AF195801", "foreign Medicago sativa comparison"},
		180: {">CYP93C8 AF195800", "foreign Medicago sativa comparison"},
		181: {">CYP93C AF195802", "foreign Medicago sativa comparison"},
		182: {">CYP93C5 Glycine max", "foreign Glycine max comparison"},
		192: {">CYP706A11 AP006082.1", "foreign Lotus japonicus comparison"},
		232: {">AJ410089.1 Medicago sativa", "foreign Medicago sativa assembly helper"},
		255: {">CYP74B4v1  Medicago sativa", "foreign Medicago sativa comparison"},
		278: {">CYP707A16 Glycine max", "foreign Glycine max comparison"},
		287: {">CYP716D4 Stevia rebaudiana", "foreign Stevia rebaudiana comparison"},
		294: {">CYP724 DT014285.1 Vitis vinifera", "foreign Vitis vinifera comparison"},
		297: {">CX704924.1 Glycine max", "foreign Glycine max assembly helper"},
		300: {">Zinnia elegans CYP733", "foreign Zinnia elegans comparison"},
		301: {">CG816619.1 Glycine max", "foreign Glycine max assembly helper"},
		372: {">Soybean CYP727", "foreign soybean comparison after no-member CYP727 note"},
		373: {">Lotus japonicus CYP727", "foreign Lotus japonicus comparison after no-member CYP727 note"},
		374: {">CR331796.1", "explicit false positive: ubiquitin ligase"},
		375: {">CG954156.1", "explicit false positive"},
		376: {">CG969307.1", "explicit false positive"},
	}
)

func main() {
	input := flag.String("input", filepath.Join("raw", "plants-medicago.seqs.txt"), "Word-normalized Medicago source text")
	sourceDoc := flag.String("source-doc", filepath.Join("raw", sourceFile), "original downloaded Word document")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "medicago-truncatula.csv"), "reviewed CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-medicago-truncatula.md"), "review ledger")
	flag.Parse()
	data, err := os.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	records, excluded, lineCount, err := parseText(string(data), true)
	if err != nil {
		panic(err)
	}
	hash, err := fileHash(*sourceDoc)
	if err != nil {
		panic(err)
	}
	if hash != "0d4967190bb35bf639018a90473e402d5ae25b09ea5d8fa3d5944aec904e0199" {
		panic(fmt.Errorf("Medicago source hash changed: %s", hash))
	}
	if err := writeCSV(*out, records, hash); err != nil {
		panic(err)
	}
	if err := writeAudit(*audit, records, excluded, lineCount, hash); err != nil {
		panic(err)
	}
	fmt.Printf("Medicago truncatula: %d source pieces accepted with %d literal protein sequences; %d foreign/false-positive pieces excluded\n", len(records), countSequences(records), len(excluded))
}

func parseText(text string, enforceLayout bool) ([]record, map[int]exclusion, int, error) {
	text = strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n")
	lines := strings.Split(text, "\n")
	if len(lines) < 10 || strings.TrimSpace(lines[0]) != "Medicago truncatula (model legume species)" || strings.TrimSpace(lines[2]) != "There are 376 sequence pieces here.  Some are duplicates." || strings.TrimSpace(lines[5]) != "They are sorted by clan" {
		return nil, nil, len(lines), fmt.Errorf("Medicago preamble changed")
	}
	var headers []int
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), ">") {
			headers = append(headers, i)
		}
	}
	if enforceLayout && (len(lines) != 3951 || len(headers) != 376 || headers[0]+1 != 14 || headers[len(headers)-1]+1 != 3947) {
		return nil, nil, len(lines), fmt.Errorf("Medicago layout changed: lines=%d headers=%d first=%d last=%d", len(lines), len(headers), headers[0]+1, headers[len(headers)-1]+1)
	}
	if enforceLayout {
		for block, want := range excludedBlocks {
			if block < 1 || block > len(headers) || !strings.HasPrefix(strings.TrimSpace(lines[headers[block-1]]), want.HeaderPrefix) {
				return nil, nil, len(lines), fmt.Errorf("Medicago exclusion block %d changed; expected prefix %q", block, want.HeaderPrefix)
			}
		}
	}

	capacity := len(headers)
	if enforceLayout {
		capacity -= len(excludedBlocks)
	}
	records := make([]record, 0, capacity)
	for blockIndex, start := range headers {
		block := blockIndex + 1
		if _, skip := excludedBlocks[block]; enforceLayout && skip {
			continue
		}
		end := len(lines)
		if blockIndex+1 < len(headers) {
			end = headers[blockIndex+1]
		}
		header := strings.TrimSpace(lines[start])
		fields := strings.Fields(strings.TrimPrefix(header, ">"))
		if len(fields) == 0 {
			return nil, nil, len(lines), fmt.Errorf("Medicago block %d has empty header", block)
		}
		id := strings.Trim(fields[0], ",;|")
		symbol := ""
		if cypID.MatchString(id) {
			symbol = id
		}
		r := record{Block: block, SourceLine: start + 1, Header: header, ID: id, Symbol: symbol}
		flags := map[string]bool{}
		blockText := strings.ToLower(strings.Join(lines[start:end], "\n"))
		if symbol == "" {
			flags["source-accession-piece"] = true
		}
		if strings.Contains(blockText, "pseudogene") || digitP.MatchString(symbol) {
			flags["source-pseudogene"] = true
		}
		if strings.Contains(blockText, "fragment") || strings.Contains(blockText, "partial") || strings.Contains(blockText, "missing") || strings.Contains(blockText, "runs off") || strings.Contains(blockText, "piece") {
			flags["source-fragment-or-missing-region"] = true
		}
		var parts []string
		for i := start + 1; i < end; i++ {
			part, ok := medicagoSequenceLine(lines[i], block, i+1)
			if !ok {
				continue
			}
			if r.FirstSequenceLine == 0 {
				r.FirstSequenceLine = i + 1
			}
			r.LastSequenceLine = i + 1
			parts = append(parts, part)
		}
		sequence := strings.Join(parts, "")
		if strings.HasSuffix(sequence, "*") {
			sequence = strings.TrimRight(sequence, "*")
			flags["source-terminal-stop"] = true
		}
		if strings.Contains(sequence, "*") {
			flags["internal-stop"] = true
		}
		if strings.Contains(sequence, "?") {
			flags["source-question-mark-residue"] = true
		}
		if sequence == "" {
			flags["sequence-missing"] = true
		} else if len(sequence) < 350 {
			flags["short-sequence"] = true
		}
		if len(sequence) > 650 {
			flags["unusually-long-sequence"] = true
		}
		for _, name := range []string{"source-accession-piece", "source-pseudogene", "source-fragment-or-missing-region", "source-terminal-stop", "internal-stop", "source-question-mark-residue", "sequence-missing", "short-sequence", "unusually-long-sequence"} {
			if flags[name] {
				r.Status = append(r.Status, name)
			}
		}
		r.Sequence = sequence
		records = append(records, r)
	}
	if enforceLayout && (len(records) != 349 || countSequences(records) != 349) {
		return nil, nil, len(lines), fmt.Errorf("Medicago accepted layout changed: records=%d sequences=%d", len(records), countSequences(records))
	}
	return records, excludedBlocks, len(lines), nil
}

func medicagoSequenceLine(line string, block, sourceLine int) (string, bool) {
	// ABE79358 is the document's one header whose literal protein starts after
	// a same-line identity/species annotation rather than in a clean paragraph.
	if block == 361 && sourceLine == 3763 {
		const marker = "[Medicago truncatula] "
		at := strings.Index(line, marker)
		if at < 0 {
			return "", false
		}
		value := strings.Join(strings.Fields(line[at+len(marker):]), "")
		if value == "" || !proteinToken.MatchString(value) {
			return "", false
		}
		return value, true
	}

	raw := strings.TrimSpace(line)
	if raw == "" {
		return "", false
	}
	hasCoordinate := false
	var parts []string
	for _, token := range strings.Fields(raw) {
		token = strings.Trim(token, `"`)
		if coordinateToken.MatchString(token) || phaseToken.MatchString(token) {
			hasCoordinate = true
			continue
		}
		if !proteinToken.MatchString(token) {
			return "", false
		}
		parts = append(parts, token)
	}
	value := strings.Join(parts, "")
	if value == "" {
		return "", false
	}
	upper := 0
	for _, ch := range value {
		if strings.ContainsRune("ACDEFGHIKLMNOPQRSTVWXY", ch) {
			upper++
		}
	}
	// Coordinate-bearing source lines may contain a one-residue exon. Clean
	// sequence-only paragraphs need at least eight uppercase residues; this
	// prevents prose annotations made only from amino-acid-shaped letters.
	if (hasCoordinate && upper == 0) || (!hasCoordinate && upper < 8) {
		return "", false
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
		note := fmt.Sprintf("Medicago Word sequence piece %d at normalized source line %d; literal protein lines %d-%d; header=%s", r.Block, r.SourceLine, r.FirstSequenceLine, r.LastSequenceLine, r.Header)
		_ = w.Write([]string{"plants", "Medicago truncatula", r.Symbol, r.ID, fmt.Sprintf("medicago-truncatula:piece-%04d:%s", r.Block, r.ID), sourceURL, note, r.Sequence, sourceFile, hash, strconv.Itoa(r.Block), strconv.Itoa(r.SourceLine), strings.Join(r.Status, ";")})
	}
	w.Flush()
	return w.Error()
}

func writeAudit(path string, records []record, excluded map[int]exclusion, lineCount int, hash string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	counts := map[string]int{}
	headers, sequences := map[string]int{}, map[string]int{}
	for _, r := range records {
		headers[r.Header]++
		sequences[r.Sequence]++
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
	foreign, falsePositive := 0, 0
	for _, item := range excluded {
		if strings.Contains(item.Reason, "false positive") {
			falsePositive++
		} else {
			foreign++
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Medicago truncatula\n\n- Source file: `%s`\n- URL: %s\n- Source SHA-256: `%s`\n- Container inspected: legacy Word document, 3,950 paragraphs, no tables, 64 pages\n- Normalized split lines: `%d`\n- Source `>` sequence pieces: 376\n- Accepted Medicago P450/candidate-P450 pieces: %d\n- Accepted pieces with literal protein sequence: %d\n- Explicit other-species assembly/comparison pieces excluded: %d\n- Explicit false-positive pieces excluded: %d\n- Duplicate exact-header groups retained: %d\n- Duplicate literal-sequence groups retained: %d\n- Review status: `complete`\n\n", sourceFile, sourceURL, hash, lineCount, len(records), countSequences(records), foreign, falsePositive, duplicateHeaders, duplicateSequences)
	b.WriteString("## Resource-specific interpretation\n\nThe preamble calls the entries `376 sequence pieces`, warns that some are duplicates, and says some other-species pieces are present for gene assembly. Each Medicago header-delimited piece is therefore retained independently in source order; accession-only genomic/cDNA pieces are not merged into adjacent named CYPs, and duplicate or alternate pieces are not collapsed. Twenty-four blocks explicitly belonging to Medicago sativa/alfalfa, Glycine max/soybean, Lotus japonicus, Stevia rebaudiana, Vitis vinifera, or Zinnia elegans are comparison/assembly helpers and are excluded from the Medicago species. The three blocks below the literal `False positives` heading are excluded. All remaining 349 blocks contain literal protein. Coordinates, standalone phase markers, whitespace, and wrapping quotes are document layout, not residues. The ABE79358 block has its protein after the exact same-line `[Medicago truncatula]` marker and is handled only by that reviewed layout. Internal stops, X, question marks, short pieces, pseudogenes, and fragments remain literal; only explicit terminal stops are removed. Nothing is translated, repaired, completed, or fetched elsewhere.\n\n## Status counts\n\n| Status | Records |\n|---|---:|\n")
	names := make([]string, 0, len(counts))
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Fprintf(&b, "| %s | %d |\n", name, counts[name])
	}
	b.WriteString("\n## Excluded source pieces\n\n| Block | Header prefix | Reason |\n|---:|---|---|\n")
	blocks := make([]int, 0, len(excluded))
	for block := range excluded {
		blocks = append(blocks, block)
	}
	sort.Ints(blocks)
	for _, block := range blocks {
		item := excluded[block]
		fmt.Fprintf(&b, "| %d | %s | %s |\n", block, md(item.HeaderPrefix), md(item.Reason))
	}
	b.WriteString("\n## Representative accepted pieces\n\n| Accepted index | Source block | Source line | ID | Symbol | Sequence | Status | Header |\n|---:|---:|---:|---|---|---:|---|---|\n")
	for _, index := range []int{0, len(records) / 2, len(records) - 1} {
		r := records[index]
		fmt.Fprintf(&b, "| %d | %d | %d | %s | %s | %d aa | %s | %s |\n", index+1, r.Block, r.SourceLine, md(r.ID), md(r.Symbol), len(r.Sequence), md(strings.Join(r.Status, ";")), md(r.Header))
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

func md(value string) string { return strings.ReplaceAll(strings.TrimSpace(value), "|", "\\|") }
