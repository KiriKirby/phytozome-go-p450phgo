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
	capsicumSourceFile = "plants-Capsicum.annuum.xlsx"
	capsicumSourceURL  = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Capsicum.annuum.xlsx"
	capsicumSheetName  = "Sorted by CYP name"
)

type capsicumRecord struct {
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
	input := flag.String("input", filepath.Join("raw", capsicumSourceFile), "downloaded Capsicum workbook")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "capsicum-annuum.csv"), "reviewed structured CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-capsicum-annum.md"), "Capsicum review ledger")
	flag.Parse()

	rows, dimension, hiddenRows, err := readCapsicumWorkbook(*input)
	if err != nil {
		panic(err)
	}
	records, excluded, err := reviewCapsicumRows(rows)
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
	fmt.Printf("Capsicum annuum: %d named CYP rows with sequence, %d unnamed candidates excluded\n", len(records), len(excluded))
}

func readCapsicumWorkbook(path string) (map[int]map[string]string, string, []int, error) {
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
	var relationships relationshipsXML
	if err := read("xl/_rels/workbook.xml.rels", &relationships); err != nil {
		return nil, "", nil, err
	}
	relTargets := map[string]string{}
	for _, relationship := range relationships.Relationships {
		relTargets[relationship.ID] = "xl/" + strings.TrimPrefix(filepath.ToSlash(relationship.Target), "/")
	}
	sheetPath := ""
	for _, sheet := range workbook.Sheets {
		if sheet.Name == capsicumSheetName {
			sheetPath = relTargets[sheet.RID]
			break
		}
	}
	if sheetPath == "" {
		return nil, "", nil, fmt.Errorf("required sheet %q not found", capsicumSheetName)
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

func reviewCapsicumRows(rows map[int]map[string]string) ([]capsicumRecord, []capsicumRecord, error) {
	header := rows[1]
	if strings.TrimSpace(header["H"]) != "seq. ID" || strings.TrimSpace(header["I"]) != "best hit" || strings.TrimSpace(header["J"]) != "%ID" || strings.TrimSpace(header["K"]) != "CYP name" || strings.TrimSpace(header["L"]) != "" {
		return nil, nil, fmt.Errorf("Capsicum header layout changed: H=%q I=%q J=%q K=%q L=%q", header["H"], header["I"], header["J"], header["K"], header["L"])
	}
	var accepted, excluded []capsicumRecord
	for row := 2; row <= 650; row++ {
		values := rows[row]
		record := capsicumRecord{Row: row, SequenceID: strings.TrimSpace(values["H"]), BestHit: strings.TrimSpace(values["I"]), CYPName: strings.TrimSpace(values["K"]), Sequence: strings.TrimSpace(values["L"])}
		record.SourceLength = len(record.Sequence)
		if record.SequenceID == "" || record.Sequence == "" {
			return nil, nil, fmt.Errorf("Capsicum row %d is missing seq. ID or sequence", row)
		}
		if record.CYPName == "" {
			record.Status = "excluded: CYP name column is empty"
			excluded = append(excluded, record)
			continue
		}
		if record.BestHit == "" || !strings.HasPrefix(strings.ToUpper(record.CYPName), "CYP") {
			return nil, nil, fmt.Errorf("Capsicum named row %d has invalid best hit/family", row)
		}
		var statuses []string
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
				return nil, nil, fmt.Errorf("Capsicum row %d contains unexpected residue %q", row, residue)
			}
		}
		record.Status = strings.Join(statuses, ";")
		accepted = append(accepted, record)
	}
	return accepted, excluded, nil
}

func writeCSV(path string, records []capsicumRecord, sourceHash string) error {
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
		note := fmt.Sprintf("Capsicum workbook row %d; seq. ID=%s; best hit=%s; CYP name=%s; sequence column=L (source header is blank)", record.Row, record.SequenceID, record.BestHit, record.CYPName)
		_ = w.Write([]string{"plants", "Capsicum annuum", record.BestHit, record.SequenceID, fmt.Sprintf("capsicum-annuum:row-%04d:%s", record.Row, record.SequenceID), capsicumSourceURL, note, record.Sequence, capsicumSourceFile, sourceHash, capsicumSheetName, strconv.Itoa(record.Row), record.CYPName, record.Status})
	}
	w.Flush()
	return w.Error()
}

func writeAudit(path string, records, excluded []capsicumRecord, dimension string, hiddenRows []int, sourceHash string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	counts := map[string]int{}
	for _, record := range records {
		for _, status := range strings.Split(record.Status, ";") {
			if status != "" {
				counts[status]++
			}
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Capsicum annuum\n\n- Source file: `%s`\n- URL: %s\n- Source SHA-256: `%s`\n- Workbook sheet: `%s`\n- Used range: `%s`\n- Hidden rows: `%v`\n- Merged cells: `none`\n- Named CYP rows accepted: `%d`\n- Accepted rows with literal sequence: `%d`\n- Unnamed candidate rows excluded: `%d`\n- Review status: `complete`\n\n", capsicumSourceFile, capsicumSourceURL, sourceHash, capsicumSheetName, dimension, hiddenRows, len(records), len(records), len(excluded))
	b.WriteString("## Resource-specific interpretation\n\nThis workbook has one sheet and no hidden rows, hidden columns, merged cells, or native table part. The header occupies row 1. For this workbook only, column H is `seq. ID`, column I is `best hit`, column J is `%ID`, column K is `CYP name`, and column L is the full source protein sequence even though its header cell is blank. Rows 2–618 have a named CYP family in K and are accepted. Rows 619–650 have an empty K value and remain excluded unnamed candidates.\n\nThe displayed CYP symbol is the exact I-column best hit; H remains the gene/sequence ID and row-level key. Literal `X`, `O`, and `-` characters are retained and flagged. No sequence character is repaired or removed.\n\n## Status counts\n\n| Status | Records |\n|---|---:|\n")
	statusNames := make([]string, 0, len(counts))
	for status := range counts {
		statusNames = append(statusNames, status)
	}
	sort.Strings(statusNames)
	for _, status := range statusNames {
		fmt.Fprintf(&b, "| %s | %d |\n", status, counts[status])
	}
	b.WriteString("\n## Boundary and representative rows\n\n| Source row | seq. ID | Best hit | CYP name | Sequence | Status |\n|---:|---|---|---|---:|---|\n")
	for _, index := range []int{0, len(records) / 2, len(records) - 1} {
		record := records[index]
		fmt.Fprintf(&b, "| %d | %s | %s | %s | %d aa | %s |\n", record.Row, md(record.SequenceID), md(record.BestHit), md(record.CYPName), len(record.Sequence), md(record.Status))
	}
	for _, record := range excluded {
		if record.Row == 619 || record.Row == 650 {
			fmt.Fprintf(&b, "| %d | %s | %s | *(empty)* | %d aa (excluded) | %s |\n", record.Row, md(record.SequenceID), md(record.BestHit), len(record.Sequence), md(record.Status))
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
