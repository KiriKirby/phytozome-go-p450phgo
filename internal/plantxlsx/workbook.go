package plantxlsx

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Workbook struct {
	Dimension     string
	Rows          map[int]map[string]string
	HiddenRows    []int
	HiddenColumns []string
	MergedCells   []string
	TableParts    int
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
	Columns []struct {
		Min, Max int `xml:",attr"`
		Hidden   int `xml:"hidden,attr"`
	} `xml:"cols>col"`
	Rows []struct {
		Number int `xml:"r,attr"`
		Hidden int `xml:"hidden,attr"`
		Cells  []struct {
			Ref       string `xml:"r,attr"`
			Type      string `xml:"t,attr"`
			Value     string `xml:"v"`
			Inline    string `xml:"is>t"`
			InlineRun []struct {
				Text string `xml:"t"`
			} `xml:"is>r"`
		} `xml:"c"`
	} `xml:"sheetData>row"`
	Merged []struct {
		Ref string `xml:"ref,attr"`
	} `xml:"mergeCells>mergeCell"`
	TableParts []struct{} `xml:"tableParts>tablePart"`
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

func Read(path, expectedSheet string) (*Workbook, error) {
	return ReadSheet(path, expectedSheet, true)
}

// ReadSheet reads one explicitly named sheet. requireOnlySheet preserves the
// strict single-sheet contract used by existing resource parsers; a parser for
// a reviewed multi-sheet workbook can select one sheet after validating names.
func ReadSheet(path, expectedSheet string, requireOnlySheet bool) (*Workbook, error) {
	z, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
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
		r, err := f.Open()
		if err != nil {
			return err
		}
		defer r.Close()
		return xml.NewDecoder(io.LimitReader(r, 128<<20)).Decode(dst)
	}
	var wb workbookXML
	if err := read("xl/workbook.xml", &wb); err != nil {
		return nil, err
	}
	selected := -1
	for i := range wb.Sheets {
		if wb.Sheets[i].Name == expectedSheet {
			selected = i
			break
		}
	}
	if selected < 0 || (requireOnlySheet && len(wb.Sheets) != 1) {
		return nil, fmt.Errorf("workbook sheet layout changed: got %v, expected sheet %q (only=%t)", sheetNames(wb), expectedSheet, requireOnlySheet)
	}
	var rel relationshipsXML
	if err := read("xl/_rels/workbook.xml.rels", &rel); err != nil {
		return nil, err
	}
	targets := map[string]string{}
	for _, item := range rel.Relationships {
		targets[item.ID] = "xl/" + strings.TrimPrefix(filepath.ToSlash(item.Target), "/")
	}
	shared := []string{}
	if files["xl/sharedStrings.xml"] != nil {
		var ss sharedStringsXML
		if err := read("xl/sharedStrings.xml", &ss); err != nil {
			return nil, err
		}
		shared = make([]string, len(ss.Items))
		for i, item := range ss.Items {
			shared[i] = strings.Join(item.Text, "")
			for _, run := range item.Runs {
				shared[i] += run.Text
			}
		}
	}
	var ws worksheetXML
	if err := read(targets[wb.Sheets[selected].RID], &ws); err != nil {
		return nil, err
	}
	out := &Workbook{Dimension: ws.Dimension.Ref, Rows: map[int]map[string]string{}, TableParts: len(ws.TableParts)}
	for _, col := range ws.Columns {
		if col.Hidden == 1 {
			out.HiddenColumns = append(out.HiddenColumns, fmt.Sprintf("%d-%d", col.Min, col.Max))
		}
	}
	for _, merged := range ws.Merged {
		out.MergedCells = append(out.MergedCells, merged.Ref)
	}
	for _, row := range ws.Rows {
		if row.Hidden == 1 {
			out.HiddenRows = append(out.HiddenRows, row.Number)
		}
		values := map[string]string{}
		for _, cell := range row.Cells {
			value := cell.Value
			switch cell.Type {
			case "s":
				index, err := strconv.Atoi(value)
				if err != nil || index < 0 || index >= len(shared) {
					return nil, fmt.Errorf("invalid shared string %s", cell.Ref)
				}
				value = shared[index]
			case "inlineStr":
				value = cell.Inline
				for _, run := range cell.InlineRun {
					value += run.Text
				}
			}
			values[columnRE.FindString(cell.Ref)] = value
		}
		out.Rows[row.Number] = values
	}
	return out, nil
}

func sheetNames(wb workbookXML) []string {
	names := make([]string, len(wb.Sheets))
	for i, sheet := range wb.Sheets {
		names[i] = sheet.Name
	}
	return names
}

// SheetNames exposes the complete workbook sheet list for resource-specific
// structure assertions.
func SheetNames(path string) ([]string, error) {
	z, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer z.Close()
	for _, f := range z.File {
		if filepath.ToSlash(f.Name) != "xl/workbook.xml" {
			continue
		}
		r, err := f.Open()
		if err != nil {
			return nil, err
		}
		defer r.Close()
		var wb workbookXML
		if err := xml.NewDecoder(io.LimitReader(r, 16<<20)).Decode(&wb); err != nil {
			return nil, err
		}
		return sheetNames(wb), nil
	}
	return nil, fmt.Errorf("missing workbook part xl/workbook.xml")
}
