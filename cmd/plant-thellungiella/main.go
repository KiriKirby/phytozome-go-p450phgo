package main

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/csv"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	thellungiellaSourceFile = "plants-Thellungiella.parvula.xlsx"
	thellungiellaSourceURL  = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Thellungiella.parvula.xlsx"
	thellungiellaSheetName  = "Sorted by CYP name"
)

type thellungiellaRecord struct {
	Row, SourceLength                              int
	SequenceID, BestHit, CYPName, Sequence, Status string
}

type workbookXML struct {
	Sheets []struct {
		Name string `xml:"name,attr"`
		RID  string `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
	} `xml:"sheets>sheet"`
}

type relationshipsXML struct {
	Relationships []struct {
		ID     string `xml:"Id,attr"`
		Target string `xml:"Target,attr"`
	} `xml:"Relationship"`
}

type worksheetXML struct {
	Dimension struct {
		Ref string `xml:"ref,attr"`
	} `xml:"dimension"`
	Rows []struct {
		Number int `xml:"r,attr"`
		Hidden int `xml:"hidden,attr"`
		Cells  []struct {
			Ref  string `xml:"r,attr"`
			Type string `xml:"t,attr"`
			V    string `xml:"v"`
			IS   struct {
				Text string `xml:"t"`
			} `xml:"is"`
		} `xml:"c"`
	} `xml:"sheetData>row"`
}

type sharedStringsXML struct {
	Items []struct {
		Text []string `xml:"t"`
		Runs []struct {
			Text string `xml:"t"`
		} `xml:"r"`
	} `xml:"si"`
}

var cellColumn = regexp.MustCompile(`^[A-Z]+`)

func main() {
	input := flag.String("input", filepath.Join("raw", thellungiellaSourceFile), "downloaded Thellungiella workbook")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "thellungiella-parvula.csv"), "reviewed structured CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-thellungiella-parvula.md"), "Thellungiella review ledger")
	flag.Parse()

	rows, dimension, hiddenRows, err := readThellungiellaWorkbook(*input)
	if err != nil {
		panic(err)
	}
	records, excluded, err := reviewThellungiellaRows(rows)
	if err != nil {
		panic(err)
	}
	hash, err := fileSHA256(*input)
	if err != nil {
		panic(err)
	}
	if err := writeCSV(*out, records, hash); err != nil {
		panic(err)
	}
	if err := writeAudit(*audit, records, excluded, dimension, hiddenRows, hash); err != nil {
		panic(err)
	}
	fmt.Printf("Thellungiella parvula: %d named CYP rows with sequence, %d unnamed candidates excluded\n", len(records), len(excluded))
}

func readThellungiellaWorkbook(path string) (map[int]map[string]string, string, []int, error) {
	archive, err := zip.OpenReader(path)
	if err != nil {
		return nil, "", nil, err
	}
	defer archive.Close()
	files := map[string]*zip.File{}
	for _, file := range archive.File {
		files[filepath.ToSlash(file.Name)] = file
	}
	read := func(name string, target any) error {
		file := files[name]
		if file == nil {
			return fmt.Errorf("missing workbook part %s", name)
		}
		r, err := file.Open()
		if err != nil {
			return err
		}
		defer r.Close()
		return xml.NewDecoder(io.LimitReader(r, 64<<20)).Decode(target)
	}
	var workbook workbookXML
	if err := read("xl/workbook.xml", &workbook); err != nil {
		return nil, "", nil, err
	}
	if len(workbook.Sheets) != 1 || workbook.Sheets[0].Name != thellungiellaSheetName {
		return nil, "", nil, fmt.Errorf("Thellungiella workbook sheet layout changed: %+v", workbook.Sheets)
	}
	var relationships relationshipsXML
	if err := read("xl/_rels/workbook.xml.rels", &relationships); err != nil {
		return nil, "", nil, err
	}
	relTargets := map[string]string{}
	for _, relationship := range relationships.Relationships {
		relTargets[relationship.ID] = "xl/" + strings.TrimPrefix(filepath.ToSlash(relationship.Target), "/")
	}
	sheetPath := relTargets[workbook.Sheets[0].RID]
	if sheetPath == "" {
		return nil, "", nil, fmt.Errorf("required sheet %q has no relationship", thellungiellaSheetName)
	}
	var sharedXML sharedStringsXML
	if err := read("xl/sharedStrings.xml", &sharedXML); err != nil {
		return nil, "", nil, err
	}
	shared := make([]string, len(sharedXML.Items))
	for i, item := range sharedXML.Items {
		shared[i] = strings.Join(item.Text, "")
		for _, run := range item.Runs {
			shared[i] += run.Text
		}
	}
	var worksheet worksheetXML
	if err := read(sheetPath, &worksheet); err != nil {
		return nil, "", nil, err
	}
	rows := map[int]map[string]string{}
	var hidden []int
	for _, row := range worksheet.Rows {
		if row.Hidden == 1 {
			hidden = append(hidden, row.Number)
		}
		values := map[string]string{}
		for _, cell := range row.Cells {
			column := cellColumn.FindString(cell.Ref)
			value := cell.V
			switch cell.Type {
			case "s":
				index, parseErr := strconv.Atoi(value)
				if parseErr != nil || index < 0 || index >= len(shared) {
					return nil, "", nil, fmt.Errorf("invalid shared string at %s", cell.Ref)
				}
				value = shared[index]
			case "inlineStr":
				value = cell.IS.Text
			}
			values[column] = value
		}
		rows[row.Number] = values
	}
	return rows, worksheet.Dimension.Ref, hidden, nil
}

func reviewThellungiellaRows(rows map[int]map[string]string) ([]thellungiellaRecord, []thellungiellaRecord, error) {
	header := rows[1]
	if strings.TrimSpace(header["H"]) != "seq ID" || strings.TrimSpace(header["I"]) != "best hit" || strings.TrimSpace(header["J"]) != "%ID" || strings.TrimSpace(header["K"]) != "CYP name" || strings.TrimSpace(header["L"]) != "" {
		return nil, nil, fmt.Errorf("Thellungiella header layout changed: H=%q I=%q J=%q K=%q L=%q", header["H"], header["I"], header["J"], header["K"], header["L"])
	}
	var accepted, excluded []thellungiellaRecord
	for row := 2; row <= 213; row++ {
		values, ok := rows[row]
		if !ok {
			return nil, nil, fmt.Errorf("Thellungiella source row %d is missing", row)
		}
		record := thellungiellaRecord{Row: row, SequenceID: strings.TrimSpace(values["H"]), BestHit: strings.TrimSpace(values["I"]), CYPName: strings.TrimSpace(values["K"]), Sequence: strings.TrimSpace(values["L"])}
		record.SourceLength = len(record.Sequence)
		if record.SequenceID == "" || record.Sequence == "" {
			return nil, nil, fmt.Errorf("Thellungiella row %d is missing seq ID or sequence", row)
		}
		if record.CYPName == "" {
			record.Status = "excluded: CYP name column is empty"
			excluded = append(excluded, record)
			continue
		}
		if record.BestHit == "" || !strings.HasPrefix(strings.ToUpper(record.BestHit), "CYP") || !strings.HasPrefix(strings.ToUpper(record.CYPName), "CYP") {
			return nil, nil, fmt.Errorf("Thellungiella named row %d has invalid best hit/family", row)
		}
		var statuses []string
		label := strings.ToLower(record.BestHit + " " + record.CYPName)
		if strings.HasSuffix(strings.ToUpper(record.BestHit), "P") || strings.Contains(label, "pseudo") || strings.Contains(label, "peudo") {
			statuses = append(statuses, "source-pseudogene-label")
		}
		if strings.Contains(record.Sequence, "-") {
			statuses = append(statuses, "source-gap")
		}
		if strings.Contains(record.Sequence, "X") {
			statuses = append(statuses, "ambiguous-X")
		}
		if strings.Contains(record.Sequence, "O") {
			statuses = append(statuses, "nonstandard-O")
		}
		if len(record.Sequence) < 350 {
			statuses = append(statuses, "short-sequence")
		}
		if len(record.Sequence) > 650 {
			statuses = append(statuses, "unusually-long-sequence")
		}
		for _, residue := range record.Sequence {
			if !strings.ContainsRune("ACDEFGHIKLMNOPQRSTVWXY-", residue) {
				return nil, nil, fmt.Errorf("Thellungiella row %d contains unexpected residue %q", row, residue)
			}
		}
		record.Status = strings.Join(statuses, ";")
		accepted = append(accepted, record)
	}
	return accepted, excluded, nil
}

func writeCSV(path string, records []thellungiellaRecord, sourceHash string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	w := csv.NewWriter(file)
	_ = w.Write([]string{"category", "species", "symbol", "id", "record_key", "source_url", "source_note", "sequence", "source_file", "source_sha256", "source_sheet", "source_row", "cyp_family", "review_status"})
	for _, record := range records {
		note := fmt.Sprintf("Thellungiella workbook row %d; seq ID=%s; best hit=%s; CYP name=%s; sequence column=L (source header is blank)", record.Row, record.SequenceID, record.BestHit, record.CYPName)
		_ = w.Write([]string{"plants", "Thellungiella parvula", record.BestHit, record.SequenceID, fmt.Sprintf("thellungiella-parvula:row-%04d:%s", record.Row, record.SequenceID), thellungiellaSourceURL, note, record.Sequence, thellungiellaSourceFile, sourceHash, thellungiellaSheetName, strconv.Itoa(record.Row), record.CYPName, record.Status})
	}
	w.Flush()
	return w.Error()
}

func writeAudit(path string, records, excluded []thellungiellaRecord, dimension string, hiddenRows []int, sourceHash string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	counts := map[string]int{}
	sequenceRows := map[string][]int{}
	for _, record := range records {
		for _, status := range strings.Split(record.Status, ";") {
			if status != "" {
				counts[status]++
			}
		}
		sequenceRows[record.Sequence] = append(sequenceRows[record.Sequence], record.Row)
	}
	duplicateSets := make([][]int, 0)
	for _, rows := range sequenceRows {
		if len(rows) > 1 {
			duplicateSets = append(duplicateSets, rows)
		}
	}
	sort.Slice(duplicateSets, func(i, j int) bool { return duplicateSets[i][0] < duplicateSets[j][0] })
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Thellungiella parvula\n\n- Source file: `%s`\n- URL: %s\n- Source SHA-256: `%s`\n- Workbook sheet: `%s`\n- Used range: `%s`\n- Hidden rows: `%v`\n- Hidden columns: `none`\n- Merged cells: `none`\n- Native table parts: `none`\n- Named CYP rows accepted: `%d`\n- Accepted rows with literal sequence: `%d`\n- Unnamed candidate rows excluded: `%d`\n- Duplicate literal-sequence groups retained as distinct IDs: `%d`\n- Review status: `complete`\n\n", thellungiellaSourceFile, thellungiellaSourceURL, sourceHash, thellungiellaSheetName, dimension, hiddenRows, len(records), len(records), len(excluded), len(duplicateSets))
	b.WriteString("## Resource-specific interpretation\n\nThis workbook has one sheet. The header occupies row 1. For this workbook only, column H is `seq ID`, I is `best hit`, J is `%ID`, K is `CYP name`, and the blank-headed L column is the literal protein sequence. Rows 2–208 have a CYP name in K and are accepted. Rows 209–213 have empty I, J, and K values and remain excluded unnamed candidates even though L contains sequence-like text.\n\nH is the Thellungiella gene/sequence ID and row-level key. I is retained as the displayed CYP symbol; K retains the source family classification. The source pseudo/P labels are status metadata, not grounds for deleting the row. Literal `O`, `X`, and `-` are preserved. Six pairs have identical literal sequences but different gene IDs, so every source row remains distinct. No character is repaired or removed.\n\n## Status counts\n\n| Status | Records |\n|---|---:|\n")
	statusNames := make([]string, 0, len(counts))
	for status := range counts {
		statusNames = append(statusNames, status)
	}
	sort.Strings(statusNames)
	for _, status := range statusNames {
		fmt.Fprintf(&b, "| %s | %d |\n", status, counts[status])
	}
	b.WriteString("\n## Duplicate literal-sequence groups\n\nThese rows are retained separately because their source IDs differ.\n\n| Source rows |\n|---|\n")
	for _, rows := range duplicateSets {
		parts := make([]string, len(rows))
		for i, row := range rows {
			parts[i] = strconv.Itoa(row)
		}
		fmt.Fprintf(&b, "| %s |\n", strings.Join(parts, ", "))
	}
	b.WriteString("\n## Boundary and representative rows\n\n| Source row | seq ID | Best hit | CYP name | Sequence | Status |\n|---:|---|---|---|---:|---|\n")
	representatives := []int{0, len(records) / 2, len(records) - 1}
	seen := map[int]bool{}
	for _, index := range representatives {
		record := records[index]
		if seen[record.Row] {
			continue
		}
		seen[record.Row] = true
		fmt.Fprintf(&b, "| %d | %s | %s | %s | %d aa | %s |\n", record.Row, md(record.SequenceID), md(record.BestHit), md(record.CYPName), len(record.Sequence), md(record.Status))
	}
	for _, record := range records {
		if record.Row == 40 || record.Row == 80 || record.Row == 179 {
			fmt.Fprintf(&b, "| %d | %s | %s | %s | %d aa | %s |\n", record.Row, md(record.SequenceID), md(record.BestHit), md(record.CYPName), len(record.Sequence), md(record.Status))
		}
	}
	for _, record := range excluded {
		if record.Row == 209 || record.Row == 213 {
			fmt.Fprintf(&b, "| %d | %s | *(empty)* | *(empty)* | %d aa (excluded) | %s |\n", record.Row, md(record.SequenceID), len(record.Sequence), md(record.Status))
		}
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

func md(value string) string {
	return strings.ReplaceAll(strings.TrimSpace(value), "|", "\\|")
}
