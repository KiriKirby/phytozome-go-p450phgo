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

const (
	potatoSourceFile = "plants-potato.P450s.doc"
	potatoSourceURL  = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/potato.P450s.doc"
)

type potatoRecord struct {
	Block       int
	SourceLine  int
	Header      string
	ID          string
	Symbol      string
	RecordKey   string
	Sequence    string
	Status      []string
	Annotations []string
}

var (
	leadingCoordinate  = regexp.MustCompile(`^\s*\d+\s+`)
	trailingCoordinate = regexp.MustCompile(`\s+\d+\s*$`)
	assemblyMarker     = regexp.MustCompile(`(?i)\((?:0|1|\?|0\?|1\?|sequence gap)\)`)
	whitespace         = regexp.MustCompile(`\s+`)
	nonKey             = regexp.MustCompile(`[^A-Za-z0-9._-]+`)
)

func main() {
	input := flag.String("input", filepath.Join("raw", "plants-potato.P450s.txt"), "Word-normalized potato source text")
	sourceDoc := flag.String("source-doc", filepath.Join("raw", potatoSourceFile), "original downloaded Word document")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "potato.csv"), "reviewed structured CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-potato.md"), "potato review ledger")
	flag.Parse()

	data, err := os.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	records, excluded := parsePotatoText(string(data))
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
	fmt.Printf("potato: %d accepted blocks, %d with sequence, %d excluded comparison blocks\n", len(records), countSequences(records), len(excluded))
}

func parsePotatoText(text string) ([]potatoRecord, []potatoRecord) {
	text = strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\r", "\n")
	lines := strings.Split(text, "\n")
	var headers []int
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), ">") {
			headers = append(headers, i)
		}
	}
	var accepted, excluded []potatoRecord
	for block, start := range headers {
		end := len(lines)
		if block+1 < len(headers) {
			end = headers[block+1]
		}
		header := strings.TrimSpace(lines[start])
		record := potatoRecord{Block: block + 1, SourceLine: start + 1, Header: strings.TrimPrefix(header, ">")}
		if reason := potatoComparisonReason(header); reason != "" {
			record.Status = []string{"excluded", reason}
			excluded = append(excluded, record)
			continue
		}
		if !potatoRecordHeader(header) {
			record.Status = []string{"excluded", "not a potato CYP record header"}
			excluded = append(excluded, record)
			continue
		}
		record.ID, record.Symbol = potatoIdentity(header)
		record.RecordKey = fmt.Sprintf("potato:%s:block-%04d", keyPart(record.Symbol), block+1)
		var sequenceParts []string
		flags := map[string]bool{}
		for _, raw := range lines[start+1 : end] {
			line := strings.TrimSpace(raw)
			if line == "" || line == "$" {
				continue
			}
			lower := strings.ToLower(line)
			if strings.Contains(lower, "pseudogene") {
				flags["pseudogene"] = true
			}
			if strings.Contains(lower, "partial") || strings.Contains(lower, "c-term") || strings.Contains(lower, "n-term") || strings.Contains(lower, "missing") || strings.Contains(lower, "est") {
				flags["partial"] = true
			}
			if strings.Contains(lower, "frameshift") {
				flags["frameshift"] = true
			}
			if strings.Contains(lower, "sequence gap") {
				flags["sequence-gap"] = true
			}
			if strings.ContainsAny(line, "?") && assemblyMarker.MatchString(line) {
				flags["uncertain-boundary"] = true
			}
			if len(record.Annotations) < 4 && !isPotatoSequenceLine(line) {
				record.Annotations = append(record.Annotations, line)
			}
			if sequence, ok := potatoSequenceLine(line); ok {
				sequenceParts = append(sequenceParts, sequence)
			}
		}
		record.Sequence = strings.TrimRight(strings.Join(sequenceParts, ""), "*")
		if strings.Contains(record.Sequence, "*") {
			flags["internal-stop"] = true
		}
		if strings.ContainsAny(record.Sequence, "BXZ") {
			flags["ambiguous-residue"] = true
		}
		if len(record.Sequence) > 0 && len(record.Sequence) < 350 {
			flags["short-sequence"] = true
		}
		if len(record.Sequence) > 650 {
			flags["unusually-long-sequence"] = true
		}
		if len(record.Sequence) == 0 {
			flags["sequence-missing"] = true
		}
		for _, name := range []string{"pseudogene", "partial", "frameshift", "sequence-gap", "uncertain-boundary", "internal-stop", "ambiguous-residue", "short-sequence", "unusually-long-sequence", "sequence-missing"} {
			if flags[name] {
				record.Status = append(record.Status, name)
			}
		}
		accepted = append(accepted, record)
	}
	return accepted, excluded
}

func potatoRecordHeader(header string) bool {
	lower := strings.ToLower(strings.TrimSpace(strings.TrimPrefix(header, ">")))
	return strings.HasPrefix(lower, "cyp") || strings.HasPrefix(lower, "sgn-u270131 solanum tuberosum")
}

// potatoComparisonReason is deliberately specific to potato.P450s.doc. These
// blocks are alignments/reference sequences embedded for comparison, not
// potato records, even when their headers contain a CYP name.
func potatoComparisonReason(header string) string {
	lower := strings.ToLower(strings.TrimSpace(strings.TrimPrefix(header, ">")))
	switch {
	case strings.HasPrefix(lower, "fi033502.1 tobacco"):
		return "tobacco comparison sequence"
	case strings.HasPrefix(lower, "ft251464.1 tomato"):
		return "tomato comparison sequence"
	case strings.HasPrefix(lower, "tobacco fi058991.1"):
		return "tobacco comparison sequence"
	case strings.HasPrefix(lower, "scaffold02618 tomato"):
		return "tomato comparison scaffold"
	case strings.HasPrefix(lower, "cyp80f5 tobacco"):
		return "tobacco comparison sequence"
	case strings.HasPrefix(lower, "cyp80n1 ortholog from eggplant"):
		return "eggplant BLAST comparison"
	case strings.HasPrefix(lower, "cyp80n1 ortholog from tobacco"):
		return "tobacco BLAST comparison"
	case strings.HasPrefix(lower, "cyp80n1 ortholog from nicotiana"):
		return "Nicotiana BLAST comparison"
	case strings.HasPrefix(lower, "cyp80n1 ortholog solanum phureja"):
		return "Solanum phureja comparison"
	case strings.HasPrefix(lower, "cyp704b capsicum annuum"):
		return "Capsicum annuum comparison sequence"
	default:
		return ""
	}
}

func potatoIdentity(header string) (string, string) {
	plain := strings.TrimSpace(strings.TrimPrefix(header, ">"))
	if strings.HasPrefix(strings.ToLower(plain), "sgn-u270131 ") {
		return "SGN-U270131", "SGN-U270131"
	}
	fields := strings.Fields(plain)
	if len(fields) == 0 {
		return "", ""
	}
	name := strings.Trim(fields[0], "|,;")
	if strings.EqualFold(name, "CYP51") && len(fields) >= 3 && strings.EqualFold(fields[1], "pseudogene") {
		name += "-pseudogene-" + strings.Trim(fields[2], "|,;")
	}
	return name, name
}

func isPotatoSequenceLine(line string) bool {
	_, ok := potatoSequenceLine(line)
	return ok
}

func potatoSequenceLine(line string) (string, bool) {
	trimmed := strings.TrimSpace(line)
	lower := strings.ToLower(trimmed)
	for _, prefix := range []string{"query", "sbjct", "score", "identities", "positives", "frame", "length", "method", "strand"} {
		if strings.HasPrefix(lower, prefix) {
			return "", false
		}
	}
	if strings.Contains(trimmed, "+") || strings.Contains(trimmed, "%") || strings.Contains(trimmed, "=") {
		return "", false
	}
	value := leadingCoordinate.ReplaceAllString(trimmed, "")
	value = trailingCoordinate.ReplaceAllString(value, "")
	value = assemblyMarker.ReplaceAllString(value, "")
	value = strings.ReplaceAll(value, "&", "")
	value = whitespace.ReplaceAllString(value, "")
	value = strings.ToUpper(value)
	if len(value) < 8 {
		return "", false
	}
	for _, residue := range value {
		if !strings.ContainsRune("ACDEFGHIKLMNPQRSTVWYBXZ*", residue) {
			return "", false
		}
	}
	return value, true
}

func keyPart(value string) string {
	value = nonKey.ReplaceAllString(strings.TrimSpace(value), "-")
	value = strings.Trim(value, "-")
	return strings.ToLower(value)
}

func writeCSV(path string, records []potatoRecord, sourceHash string) error {
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
	for _, record := range records {
		note := "Potato resource block " + strconv.Itoa(record.Block) + " at normalized source line " + strconv.Itoa(record.SourceLine) + "; header: " + record.Header
		if len(record.Annotations) > 0 {
			note += "; annotations: " + strings.Join(record.Annotations, " / ")
		}
		_ = w.Write([]string{"plants", "potato", record.Symbol, record.ID, record.RecordKey, potatoSourceURL, note, record.Sequence, potatoSourceFile, sourceHash, strconv.Itoa(record.Block), strconv.Itoa(record.SourceLine), record.Header, strings.Join(record.Status, ";")})
	}
	w.Flush()
	return w.Error()
}

func writeAudit(path string, records, excluded []potatoRecord, sourceHash string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	statusCounts := map[string]int{}
	for _, record := range records {
		for _, status := range record.Status {
			statusCounts[status]++
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: potato\n\n- Source file: `%s`\n- URL: %s\n- Source SHA-256: `%s`\n- Container inspected: legacy Word document, 13,474 paragraphs, no tables\n- Accepted source blocks: `%d`\n- Blocks with literal protein sequence: `%d`\n- Blocks without accepted sequence: `%d`\n- Explicit foreign/comparison blocks excluded: `%d`\n- Review status: `complete`\n\n", potatoSourceFile, potatoSourceURL, sourceHash, len(records), countSequences(records), len(records)-countSequences(records), len(excluded))
	b.WriteString("## Resource-specific interpretation\n\nThis file is not ordinary FASTA. Each `>` block may contain a potato gene model, genomic coordinates, exon-boundary markers such as `(0)`/`(1)`/`(?)`, joined fragments marked with `&`, pseudogenes, EST fragments, alternate versions, or an explicitly foreign comparison sequence. The parser is exclusive to `potato.P450s.doc`. It preserves every accepted potato block in source order, does not merge equal CYP names, strips only a terminal `*`, and retains internal stops and literal `X` residues. `Query`/`Sbjct` alignments and the explicitly named tobacco, tomato, eggplant, Nicotiana, Solanum phureja, and Capsicum comparison blocks are excluded.\n\n")
	b.WriteString("Known source uncertainty remains visible in `review_status`; no residue, splice boundary, frameshift, gap, or missing terminus is repaired.\n\n## Status counts\n\n| Status | Records |\n|---|---:|\n")
	for _, status := range []string{"pseudogene", "partial", "frameshift", "sequence-gap", "uncertain-boundary", "internal-stop", "ambiguous-residue", "short-sequence", "unusually-long-sequence", "sequence-missing"} {
		fmt.Fprintf(&b, "| %s | %d |\n", status, statusCounts[status])
	}
	b.WriteString("\n## Excluded comparison blocks\n\n| Source line | Header | Reason |\n|---:|---|---|\n")
	for _, record := range excluded {
		reason := "excluded"
		if len(record.Status) > 1 {
			reason = record.Status[1]
		}
		fmt.Fprintf(&b, "| %d | %s | %s |\n", record.SourceLine, markdown(record.Header), markdown(reason))
	}
	b.WriteString("\n## Accepted records\n\n| Block | Source line | ID / symbol | Sequence | Status | Header |\n|---:|---:|---|---:|---|---|\n")
	for _, record := range records {
		seq := "missing"
		if record.Sequence != "" {
			seq = fmt.Sprintf("%d aa", len(record.Sequence))
		}
		fmt.Fprintf(&b, "| %d | %d | %s | %s | %s | %s |\n", record.Block, record.SourceLine, markdown(record.Symbol), seq, markdown(strings.Join(record.Status, "; ")), markdown(record.Header))
	}
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func countSequences(records []potatoRecord) int {
	count := 0
	for _, record := range records {
		if record.Sequence != "" {
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

func markdown(value string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(value), "|", "\\|"), "\n", " ")
}
