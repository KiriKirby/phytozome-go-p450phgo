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
	sourceFile = "plants-Eucalyptus.grandis.xlsx"
	sourceURL  = "https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Eucalyptus.grandis.xlsx"
	sheetName  = "sorted by CYP name"
)

type record struct {
	Row                                          int
	SourceID, BlastID, BestHit, Sequence, Status string
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
var pseudogeneRE = regexp.MustCompile(`(?i)\dP$`)

func main() {
	input := flag.String("input", filepath.Join("raw", sourceFile), "downloaded Eucalyptus grandis workbook")
	out := flag.String("out", filepath.Join("sources", "reviewed", "plants", "eucalyptus-grandis.csv"), "reviewed CSV")
	audit := flag.String("audit", filepath.Join("docs", "plants", "plant-eucalyptus-grandis.md"), "review ledger")
	flag.Parse()
	rows, dimension, hidden, err := readWorkbook(*input)
	if err != nil {
		panic(err)
	}
	accepted, excluded, err := reviewRows(rows, dimension)
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
	fmt.Printf("Eucalyptus grandis: %d CYP rows with sequence, %d non-CYP model rows excluded\n", len(accepted), len(excluded))
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
		return nil, "", nil, fmt.Errorf("Eucalyptus grandis sheet layout changed")
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

func reviewRows(rows map[int]map[string]string, dimension string) ([]record, []record, error) {
	if dimension != "A1:L773" {
		return nil, nil, fmt.Errorf("Eucalyptus grandis used range changed: %q", dimension)
	}
	h := rows[1]
	if h["A"] != "Gotoh's seq ID" || h["I"] != "blast ID" || h["J"] != "best hit" || h["K"] != "%ID" || h["L"] != "sequence" {
		return nil, nil, fmt.Errorf("Eucalyptus grandis header changed")
	}
	var accepted, excluded []record
	for row := 2; row <= 770; row++ {
		v := rows[row]
		if v == nil {
			return nil, nil, fmt.Errorf("missing Eucalyptus grandis data row %d", row)
		}
		rawSequence := strings.TrimSpace(v["L"])
		r := record{Row: row, SourceID: strings.TrimSpace(v["A"]), BlastID: strings.TrimSpace(v["I"]), BestHit: strings.TrimSpace(v["J"]), Sequence: rawSequence}
		if r.SourceID == "" || r.BlastID == "" || r.BestHit == "" || r.Sequence == "" {
			return nil, nil, fmt.Errorf("Eucalyptus grandis row %d missing required data", row)
		}
		if !strings.HasPrefix(strings.ToUpper(r.BestHit), "CYP") {
			r.Status = "excluded: non-CYP model/comparison row"
			excluded = append(excluded, r)
			continue
		}
		var status []string
		if pseudogeneRE.MatchString(r.BestHit) {
			status = append(status, "source-pseudogene-label")
		}
		if strings.HasSuffix(rawSequence, "*") {
			r.Sequence = strings.TrimRight(rawSequence, "*")
			status = append(status, "source-terminal-stop")
		}
		if strings.Contains(r.Sequence, "*") {
			status = append(status, "internal-stop")
		}
		if strings.Contains(r.Sequence, "-") {
			status = append(status, "source-gap")
		}
		if strings.ContainsAny(r.Sequence, "Xx") {
			status = append(status, "ambiguous-X-or-x")
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
			if !strings.ContainsRune("ACDEFGHIKLMNOPQRSTVWXY*-x", aa) {
				return nil, nil, fmt.Errorf("Eucalyptus grandis row %d unexpected residue %q", row, aa)
			}
		}
		r.Status = strings.Join(status, ";")
		accepted = append(accepted, r)
	}
	if rows[771] != nil {
		return nil, nil, fmt.Errorf("Eucalyptus grandis row 771 is no longer blank")
	}
	if strings.TrimSpace(rows[772]["A"]) != "green <55%" || strings.TrimSpace(rows[773]["A"]) != "yellow <39%" {
		return nil, nil, fmt.Errorf("Eucalyptus grandis legend rows changed")
	}
	if len(accepted) != 762 || len(excluded) != 7 {
		return nil, nil, fmt.Errorf("Eucalyptus grandis boundary changed: accepted=%d excluded=%d", len(accepted), len(excluded))
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
	_ = w.Write([]string{"category", "species", "symbol", "id", "record_key", "source_url", "source_note", "sequence", "source_file", "source_sha256", "source_sheet", "source_row", "review_status"})
	for _, r := range records {
		id := strings.TrimPrefix(r.SourceID, ">")
		note := fmt.Sprintf("Eucalyptus grandis workbook row %d; Gotoh source ID=%s; blast ID=%s; best hit=%s; sequence column=L", r.Row, r.SourceID, r.BlastID, r.BestHit)
		_ = w.Write([]string{"plants", "Eucalyptus grandis", r.BestHit, id, fmt.Sprintf("eucalyptus-grandis:row-%04d:%s", r.Row, id), sourceURL, note, r.Sequence, sourceFile, hash, sheetName, strconv.Itoa(r.Row), r.Status})
	}
	w.Flush()
	return w.Error()
}

func writeAudit(path string, accepted, excluded []record, dimension string, hidden []int, hash string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	counts := map[string]int{}
	for _, r := range accepted {
		for _, s := range strings.Split(r.Status, ";") {
			if s != "" {
				counts[s]++
			}
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Plant resource review: Eucalyptus grandis\n\n- Source file: `%s`\n- URL: %s\n- Source SHA-256: `%s`\n- Workbook sheet: `%s`\n- Used range: `%s`\n- Hidden rows: `%v`\n- Hidden columns: `none`\n- Merged cells: `none`\n- Native table parts: `none`\n- Accepted CYP rows: `%d`\n- Accepted rows with literal sequence: `%d`\n- Non-CYP model/comparison rows excluded: `%d`\n- Missing worksheet row: `771` (blank)\n- Legend rows excluded: `772-773`\n- Review status: `complete`\n\n", sourceFile, sourceURL, hash, sheetName, dimension, hidden, len(accepted), len(accepted), len(excluded))
	b.WriteString("## Resource-specific interpretation\n\nThis workbook uses A=`Gotoh's seq ID` (including a literal leading `>`), I=`blast ID`, J=`best hit`, K=`%ID`, and L=`sequence`. Rows 2-770 are data rows. A row is accepted only when J begins with `CYP`; seven rows whose J values are Eucalyptus model IDs are excluded. Worksheet row 771 is blank, and rows 772-773 are color-threshold legend text, not records. One or more explicit terminal `*` markers are removed and audited; internal `*`, O, uppercase X, lowercase x, and gaps are retained exactly and flagged. No sequence is otherwise repaired or deduplicated.\n\n## Status counts\n\n| Status | Records |\n|---|---:|\n")
	names := make([]string, 0, len(counts))
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Fprintf(&b, "| %s | %d |\n", name, counts[name])
	}
	b.WriteString("\n## Boundary and representative rows\n\n| Row | Source ID | blast ID | Best hit | Sequence | Status |\n|---:|---|---|---|---:|---|\n")
	for _, row := range []int{2, 37, 386, 763, 764} {
		for _, r := range accepted {
			if r.Row == row {
				fmt.Fprintf(&b, "| %d | %s | %s | %s | %d aa | %s |\n", r.Row, md(r.SourceID), md(r.BlastID), md(r.BestHit), len(r.Sequence), md(r.Status))
			}
		}
	}
	for _, r := range excluded {
		fmt.Fprintf(&b, "| %d | %s | %s | %s | %d aa (excluded) | %s |\n", r.Row, md(r.SourceID), md(r.BlastID), md(r.BestHit), len(r.Sequence), md(r.Status))
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
