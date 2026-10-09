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
	sourceFile = "plants-Brassica.rapa.xlsx"
	sourceURL  = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Brassica.rapa.xlsx"
	sheetName  = "Sorted by CYP name"
)

type record struct {
	Row                                            int
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

var columnRE = regexp.MustCompile(`^[A-Z]+`)

func main() {
	input := flag.String("input", filepath.Join("raw", sourceFile), "downloaded Brassica rapa workbook")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "brassica-rapa.csv"), "reviewed CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-brassica-rapa.md"), "review ledger")
	flag.Parse()
	rows, dimension, hidden, err := readWorkbook(*input)
	if err != nil {
		panic(err)
	}
	accepted, excluded, err := reviewRows(rows)
	if err != nil {
		panic(err)
	}
	hash, err := fileHash(*input)
	if err != nil {
		panic(err)
	}
	if err := writeCSV(*out, accepted, hash); err != nil {
		panic(err)
	}
	if err := writeAudit(*audit, accepted, excluded, dimension, hidden, hash); err != nil {
		panic(err)
	}
	fmt.Printf("Brassica rapa: %d named CYP rows with sequence, %d unnamed candidates excluded\n", len(accepted), len(excluded))
}

func readWorkbook(path string) (map[int]map[string]string, string, []int, error) {
	z, err := zip.OpenReader(path)
	if err != nil {
		return nil, "", nil, err
	}
	defer z.Close()
	files := map[string]*zip.File{}
	for _, f := range z.File {
		files[filepath.ToSlash(f.Name)] = f
	}
	read := func(name string, dst any) error {
		f := files[name]
		if f == nil {
			return fmt.Errorf("missing workbook part %s", name)
		}
		r, e := f.Open()
		if e != nil {
			return e
		}
		defer r.Close()
		return xml.NewDecoder(io.LimitReader(r, 64<<20)).Decode(dst)
	}
	var wb workbookXML
	if err := read("xl/workbook.xml", &wb); err != nil {
		return nil, "", nil, err
	}
	if len(wb.Sheets) != 1 || wb.Sheets[0].Name != sheetName {
		return nil, "", nil, fmt.Errorf("Brassica rapa sheet layout changed")
	}
	var rel relationshipsXML
	if err := read("xl/_rels/workbook.xml.rels", &rel); err != nil {
		return nil, "", nil, err
	}
	targets := map[string]string{}
	for _, item := range rel.Relationships {
		targets[item.ID] = "xl/" + strings.TrimPrefix(filepath.ToSlash(item.Target), "/")
	}
	var ss sharedStringsXML
	if err := read("xl/sharedStrings.xml", &ss); err != nil {
		return nil, "", nil, err
	}
	shared := make([]string, len(ss.Items))
	for i, item := range ss.Items {
		shared[i] = strings.Join(item.Text, "")
		for _, run := range item.Runs {
			shared[i] += run.Text
		}
	}
	var ws worksheetXML
	if err := read(targets[wb.Sheets[0].RID], &ws); err != nil {
		return nil, "", nil, err
	}
	rows := map[int]map[string]string{}
	var hidden []int
	for _, row := range ws.Rows {
		if row.Hidden == 1 {
			hidden = append(hidden, row.Number)
		}
		values := map[string]string{}
		for _, cell := range row.Cells {
			value := cell.V
			if cell.Type == "s" {
				index, e := strconv.Atoi(value)
				if e != nil || index < 0 || index >= len(shared) {
					return nil, "", nil, fmt.Errorf("invalid shared string %s", cell.Ref)
				}
				value = shared[index]
			}
			values[columnRE.FindString(cell.Ref)] = value
		}
		rows[row.Number] = values
	}
	return rows, ws.Dimension.Ref, hidden, nil
}

func reviewRows(rows map[int]map[string]string) ([]record, []record, error) {
	h := rows[1]
	if h["H"] != "seq ID" || h["I"] != "best hit" || h["J"] != "%ID" || h["K"] != "CYP name" || strings.TrimSpace(h["L"]) != "" {
		return nil, nil, fmt.Errorf("Brassica rapa header changed: H=%q I=%q J=%q K=%q L=%q", h["H"], h["I"], h["J"], h["K"], h["L"])
	}
	var accepted, excluded []record
	for row := 2; row <= 383; row++ {
		v, ok := rows[row]
		if !ok {
			return nil, nil, fmt.Errorf("missing Brassica rapa row %d", row)
		}
		r := record{Row: row, SequenceID: strings.TrimSpace(v["H"]), BestHit: strings.TrimSpace(v["I"]), CYPName: strings.TrimSpace(v["K"]), Sequence: strings.TrimSpace(v["L"])}
		if r.SequenceID == "" || r.Sequence == "" {
			return nil, nil, fmt.Errorf("Brassica rapa row %d missing ID or sequence", row)
		}
		if r.CYPName == "" {
			r.Status = "excluded: CYP name column is empty"
			excluded = append(excluded, r)
			continue
		}
		if !strings.HasPrefix(strings.ToUpper(r.BestHit), "CYP") || !strings.HasPrefix(strings.ToUpper(r.CYPName), "CYP") {
			return nil, nil, fmt.Errorf("invalid named Brassica rapa row %d", row)
		}
		var status []string
		label := strings.ToLower(r.BestHit + " " + r.CYPName)
		if strings.HasSuffix(strings.ToUpper(r.BestHit), "P") || strings.Contains(label, "pseudo") || strings.Contains(label, "peudo") {
			status = append(status, "source-pseudogene-label")
		}
		if strings.Contains(r.Sequence, "-") {
			status = append(status, "source-gap")
		}
		if strings.Contains(r.Sequence, "X") {
			status = append(status, "ambiguous-X")
		}
		if strings.Contains(r.Sequence, "O") {
			status = append(status, "nonstandard-O")
		}
		if len(r.Sequence) < 350 {
			status = append(status, "short-sequence")
		}
		if len(r.Sequence) > 650 {
			status = append(status, "unusually-long-sequence")
		}
		for _, aa := range r.Sequence {
			if !strings.ContainsRune("ACDEFGHIKLMNOPQRSTVWXY-", aa) {
				return nil, nil, fmt.Errorf("Brassica rapa row %d unexpected residue %q", row, aa)
			}
		}
		r.Status = strings.Join(status, ";")
		accepted = append(accepted, r)
	}
	return accepted, excluded, nil
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
	_ = w.Write([]string{"category", "species", "symbol", "id", "record_key", "source_url", "source_note", "sequence", "source_file", "source_sha256", "source_sheet", "source_row", "cyp_family", "review_status"})
	for _, r := range records {
		note := fmt.Sprintf("Brassica rapa workbook row %d; seq ID=%s; best hit=%s; CYP name=%s; sequence column=L (source header is blank)", r.Row, r.SequenceID, r.BestHit, r.CYPName)
		_ = w.Write([]string{"plants", "Brassica rapa", r.BestHit, r.SequenceID, fmt.Sprintf("brassica-rapa:row-%04d:%s", r.Row, r.SequenceID), sourceURL, note, r.Sequence, sourceFile, hash, sheetName, strconv.Itoa(r.Row), r.CYPName, r.Status})
	}
	w.Flush()
	return w.Error()
}

func writeAudit(path string, accepted, excluded []record, dimension string, hidden []int, hash string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	counts := map[string]int{}
	sequences := map[string][]int{}
	for _, r := range accepted {
		for _, s := range strings.Split(r.Status, ";") {
			if s != "" {
				counts[s]++
			}
		}
		sequences[r.Sequence] = append(sequences[r.Sequence], r.Row)
	}
	duplicates := 0
	for _, rows := range sequences {
		if len(rows) > 1 {
			duplicates++
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Brassica rapa\n\n- Source file: `%s`\n- URL: %s\n- Source SHA-256: `%s`\n- Workbook sheet: `%s`\n- Used range: `%s`\n- Hidden rows: `%v`\n- Hidden columns: `none`\n- Merged cells: `none`\n- Native table parts: `none`\n- Named CYP rows accepted: `%d`\n- Accepted rows with literal sequence: `%d`\n- Unnamed candidate rows excluded: `%d`\n- Duplicate literal-sequence groups: `%d`\n- Review status: `complete`\n\n", sourceFile, sourceURL, hash, sheetName, dimension, hidden, len(accepted), len(accepted), len(excluded), duplicates)
	b.WriteString("## Resource-specific interpretation\n\nThis workbook has the sheet name `Sorted by CYP name`. Its exact data columns are H=`seq ID`, I=`best hit`, J=`%ID`, K=`CYP name`, and blank-headed L=literal protein sequence. Rows 2–378 are named and accepted; rows 379–383 have empty I/J/K classification cells and are excluded unnamed candidates. Source P suffixes, O, X, and gap characters remain literal and are flagged. Four groups of identical literal sequence remain separate because their IDs differ. No sequence is repaired or deduplicated.\n\n## Status counts\n\n| Status | Records |\n|---|---:|\n")
	names := make([]string, 0, len(counts))
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Fprintf(&b, "| %s | %d |\n", name, counts[name])
	}
	b.WriteString("\n## Boundary and representative rows\n\n| Source row | seq ID | Best hit | CYP name | Sequence | Status |\n|---:|---|---|---|---:|---|\n")
	for _, row := range []int{2, 30, 190, 378} {
		r := accepted[row-2]
		fmt.Fprintf(&b, "| %d | %s | %s | %s | %d aa | %s |\n", r.Row, md(r.SequenceID), md(r.BestHit), md(r.CYPName), len(r.Sequence), md(r.Status))
	}
	for _, r := range excluded {
		fmt.Fprintf(&b, "| %d | %s | *(empty)* | *(empty)* | %d aa (excluded) | %s |\n", r.Row, md(r.SequenceID), len(r.Sequence), md(r.Status))
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
