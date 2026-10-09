package plantxlsx

import "testing"

func TestSheetNames(t *testing.T) {
	wb := workbookXML{}
	wb.Sheets = append(wb.Sheets, struct {
		Name string `xml:"name,attr"`
		RID  string `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
	}{Name: "reviewed sheet", RID: "rId1"})
	got := sheetNames(wb)
	if len(got) != 1 || got[0] != "reviewed sheet" {
		t.Fatalf("sheetNames=%v", got)
	}
}
