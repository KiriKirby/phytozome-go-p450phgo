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
	sourceFile = "plants-papaya.doc"
	sourceURL  = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/papaya.doc"
)

type record struct {
	Block, SourceLine               int
	Header, ID, RecordKey, Sequence string
	Status, Annotations             []string
}

var (
	leadingCoordinate  = regexp.MustCompile(`^\d+\s+`)
	trailingCoordinate = regexp.MustCompile(`\s+\d+$`)
	whitespace         = regexp.MustCompile(`\s+`)
	nonKey             = regexp.MustCompile(`[^A-Za-z0-9._-]+`)
)

func main() {
	input := flag.String("input", filepath.Join("raw", "plants-papaya.txt"), "Word-normalized papaya source text")
	sourceDoc := flag.String("source-doc", filepath.Join("raw", sourceFile), "original downloaded Word document")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "papaya.csv"), "reviewed CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-carica-papaya.md"), "review ledger")
	flag.Parse()

	data, err := os.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	records, excluded, err := parsePapayaText(string(data))
	if err != nil {
		panic(err)
	}
	if len(records) != 182 || len(excluded) != 43 {
		panic(fmt.Errorf("papaya reviewed block counts changed: accepted=%d excluded=%d", len(records), len(excluded)))
	}
	hash, err := fileSHA256(*sourceDoc)
	if err != nil {
		panic(err)
	}
	if err := writeCSV(*out, records, hash); err != nil {
		panic(err)
	}
	if err := writeAudit(*audit, records, excluded, hash); err != nil {
		panic(err)
	}
	fmt.Printf("Carica papaya: %d source blocks with sequence, %d explicit Vitis comparison blocks excluded\n", len(records), len(excluded))
}

func parsePapayaText(text string) ([]record, []record, error) {
	text = strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n")
	lines := strings.Split(text, "\n")
	var headers []int
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), ">") {
			headers = append(headers, i)
		}
	}
	var accepted, excluded []record
	for block, start := range headers {
		end := len(lines)
		if block+1 < len(headers) {
			end = headers[block+1]
		}
		header := strings.TrimSpace(lines[start])
		plain := strings.TrimSpace(strings.TrimPrefix(header, ">"))
		fields := strings.Fields(plain)
		if len(fields) == 0 || !strings.HasPrefix(fields[0], "CYP") {
			return nil, nil, fmt.Errorf("papaya block %d has invalid header %q", block+1, header)
		}
		r := record{Block: block + 1, SourceLine: start + 1, Header: plain, ID: strings.Trim(fields[0], ",;|")}
		r.RecordKey = fmt.Sprintf("carica-papaya:%s:block-%04d", keyPart(r.ID), r.Block)
		if strings.Contains(strings.ToLower(header), "vitis vinifera") {
			r.Status = []string{"excluded", "explicit Vitis vinifera comparison block"}
			excluded = append(excluded, r)
			continue
		}

		var parts []string
		sequenceStarted := false
		blockText := strings.ToLower(strings.Join(lines[start:end], "\n"))
		for i := start + 1; i < end; i++ {
			lineNumber := i + 1
			raw := lines[i]
			if papayaInterstitialAnnotation(r.ID, lineNumber, raw) {
				if len(r.Annotations) < 6 {
					r.Annotations = append(r.Annotations, strings.TrimSpace(raw))
				}
				continue
			}
			sequence, ok := papayaSequenceLine(r.ID, lineNumber, raw)
			if ok {
				parts = append(parts, sequence)
				sequenceStarted = true
				continue
			}
			trimmed := strings.TrimSpace(raw)
			if trimmed == "" || trimmed == "(0)" || trimmed == "(1)" || trimmed == "(2)" {
				continue
			}
			if sequenceStarted {
				return nil, nil, fmt.Errorf("papaya %s source line %d interrupts literal sequence: %q", r.ID, lineNumber, raw)
			}
			if len(r.Annotations) < 6 {
				r.Annotations = append(r.Annotations, trimmed)
			}
		}
		r.Sequence = strings.TrimRight(strings.Join(parts, ""), "*")
		if r.Sequence == "" {
			return nil, nil, fmt.Errorf("papaya %s block %d has no literal sequence", r.ID, r.Block)
		}
		flags := map[string]bool{}
		if strings.HasSuffix(strings.ToUpper(r.ID), "P") {
			flags["source-pseudogene-label"] = true
		}
		if strings.Contains(blockText, "frameshift") {
			flags["source-frameshift"] = true
		}
		for _, phrase := range []string{"partial", "fragment", "missing", "c-term", "n-term", "seq. gap", "seq gap", "sequence gap", "runs off contig", "not recovered", "exon 1 only"} {
			if strings.Contains(blockText, phrase) {
				flags["source-partial-or-gap"] = true
			}
		}
		if strings.Contains(r.Sequence, "*") {
			flags["internal-stop"] = true
		}
		if strings.Contains(r.Sequence, "-") {
			flags["source-gap"] = true
		}
		if strings.Contains(r.Sequence, "X") {
			flags["ambiguous-X"] = true
		}
		if len(r.Sequence) < 350 {
			flags["short-sequence"] = true
		}
		if len(r.Sequence) > 650 {
			flags["unusually-long-sequence"] = true
		}
		for _, aa := range r.Sequence {
			if !strings.ContainsRune("ACDEFGHIKLMNPQRSTVWXY*-", aa) {
				return nil, nil, fmt.Errorf("papaya %s source line %d has unexpected residue %q", r.ID, r.SourceLine, aa)
			}
		}
		for _, name := range []string{"source-pseudogene-label", "source-frameshift", "source-partial-or-gap", "internal-stop", "source-gap", "ambiguous-X", "short-sequence", "unusually-long-sequence"} {
			if flags[name] {
				r.Status = append(r.Status, name)
			}
		}
		accepted = append(accepted, r)
	}
	return accepted, excluded, nil
}

// papayaSequenceLine accepts only the literal line forms observed in papaya.doc:
// uppercase residues, lowercase x gap placeholders, optional separated genomic
// coordinates, and the source's phase markers. Three malformed source lines are
// handled by exact record and line number rather than by a permissive rule.
func papayaSequenceLine(id string, lineNumber int, raw string) (string, bool) {
	value := strings.TrimSpace(raw)
	if value == "" {
		return "", false
	}
	hasCoordinate := leadingCoordinate.MatchString(value) || trailingCoordinate.MatchString(value)
	hasPhase := strings.Contains(value, "(0)") || strings.Contains(value, "(1)") || strings.Contains(value, "(2)") || strings.Contains(value, "()")
	exactMalformed := (id == "CYP87A10" && lineNumber == 1437) || (id == "CYP715A7" && lineNumber == 2591) || (id == "CYP727A8" && lineNumber == 2770)
	if strings.IndexFunc(value, func(r rune) bool { return r >= '0' && r <= '9' }) >= 0 && !hasCoordinate && !hasPhase && !exactMalformed {
		return "", false
	}
	value = leadingCoordinate.ReplaceAllString(value, "")
	value = trailingCoordinate.ReplaceAllString(value, "")
	value = strings.NewReplacer("(0)", "", "(1)", "", "(2)", "", "()", "").Replace(value)
	switch {
	case id == "CYP87A10" && lineNumber == 1437:
		value = strings.TrimSuffix(value, "(20")
	case id == "CYP715A7" && lineNumber == 2591:
		value = strings.TrimSuffix(value, "314009")
	case id == "CYP727A8" && lineNumber == 2770:
		value = strings.TrimSuffix(value, "29332")
	}
	value = whitespace.ReplaceAllString(value, "")
	value = strings.ReplaceAll(value, "x", "X")
	if value == "" {
		return "", false
	}
	for _, aa := range value {
		if !strings.ContainsRune("ACDEFGHIKLMNPQRSTVWXY*-", aa) {
			return "", false
		}
	}
	return value, true
}

func papayaInterstitialAnnotation(id string, lineNumber int, raw string) bool {
	if id != "CYP712A10P" {
		return false
	}
	trimmed := strings.TrimSpace(raw)
	return (lineNumber == 2544 && trimmed == "GS_ORF_8_from_ supercontig_234:168237..173712 (+ strand)") ||
		(lineNumber == 2545 && trimmed == "48% to 705A25")
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
		note := fmt.Sprintf("Carica papaya Word block %d at normalized source line %d; header=%s", r.Block, r.SourceLine, r.Header)
		if len(r.Annotations) > 0 {
			note += "; annotations=" + strings.Join(r.Annotations, " / ")
		}
		_ = w.Write([]string{"plants", "Carica papaya", r.ID, r.ID, r.RecordKey, sourceURL, note, r.Sequence, sourceFile, hash, strconv.Itoa(r.Block), strconv.Itoa(r.SourceLine), strings.Join(r.Status, ";")})
	}
	w.Flush()
	return w.Error()
}

func writeAudit(path string, records, excluded []record, hash string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	counts := map[string]int{}
	names := map[string][]int{}
	sequences := map[string][]int{}
	for _, r := range records {
		for _, status := range r.Status {
			counts[status]++
		}
		names[r.ID] = append(names[r.ID], r.SourceLine)
		sequences[r.Sequence] = append(sequences[r.Sequence], r.SourceLine)
	}
	duplicateNames, duplicateSequences := 0, 0
	for _, lines := range names {
		if len(lines) > 1 {
			duplicateNames++
		}
	}
	for _, lines := range sequences {
		if len(lines) > 1 {
			duplicateSequences++
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Carica papaya\n\n- Source file: `%s`\n- URL: %s\n- Source SHA-256: `%s`\n- Container inspected: legacy Word document, 3,112 paragraphs, no tables\n- Header-delimited blocks: `%d`\n- Papaya blocks accepted with literal sequence: `%d`\n- Explicit Vitis vinifera comparison blocks excluded: `%d`\n- Source-labeled pseudogenes: `%d`\n- Duplicate CYP-name groups retained: `%d`\n- Duplicate literal-sequence groups: `%d`\n- Review status: `complete`\n\n", sourceFile, sourceURL, hash, len(records)+len(excluded), len(records), len(excluded), counts["source-pseudogene-label"], duplicateNames, duplicateSequences)
	b.WriteString("## Resource-specific interpretation\n\nThe document declares 182 sequences: 143 genes and 39 pseudogenes. It contains 225 `>` blocks in total. Exactly 43 headers explicitly identify `Vitis vinifera`; these are the comparison sequences mentioned on the title page and are excluded. The remaining 182 blocks are retained in source order and all contain literal protein sequence.\n\nFor this document only, sequence lines are uppercase amino-acid text (with lowercase `x` normalized to literal `X`), optionally bounded by separated genomic coordinates and `(0)`/`(1)`/`(2)` phase markers. One literal `--` gap, every `X`, every internal `*`, fragments, frameshifts, and pseudogene blocks are retained. Only terminal `*` markers are removed. Three malformed coordinate/phase lines and the two annotation lines separating the strands of CYP712A10P are handled by exact source line and record ID. Any other non-empty line after a sequence begins is an error, so the parser cannot silently jump over a new annotation.\n\nCYP729A17 occurs twice as source versions v1 and v2 and remains two records with distinct `RecordKey` values. No sequence is repaired, translated, filled, or deduplicated.\n\n## Status counts\n\n| Status | Records |\n|---|---:|\n")
	statusNames := make([]string, 0, len(counts))
	for name := range counts {
		statusNames = append(statusNames, name)
	}
	sort.Strings(statusNames)
	for _, name := range statusNames {
		fmt.Fprintf(&b, "| %s | %d |\n", name, counts[name])
	}
	b.WriteString("\n## Representative and boundary blocks\n\n| Block | Source line | CYP name | Sequence | Status | Header |\n|---:|---:|---|---:|---|---|\n")
	for _, index := range []int{0, len(records) / 2, len(records) - 2, len(records) - 1} {
		r := records[index]
		fmt.Fprintf(&b, "| %d | %d | %s | %d aa | %s | %s |\n", r.Block, r.SourceLine, md(r.ID), len(r.Sequence), md(strings.Join(r.Status, ";")), md(r.Header))
	}
	b.WriteString("\n## Excluded comparison blocks\n\n")
	for _, r := range excluded {
		fmt.Fprintf(&b, "- Block %d, line %d: `%s`\n", r.Block, r.SourceLine, md(r.Header))
	}
	return os.WriteFile(path, []byte(b.String()), 0o644)
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

func md(value string) string { return strings.ReplaceAll(strings.TrimSpace(value), "|", "\\|") }
