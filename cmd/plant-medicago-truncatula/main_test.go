package main

import (
	"strings"
	"testing"
)

func TestMedicagoSequenceLineReviewedLayouts(t *testing.T) {
	tests := []struct {
		line       string
		block, row int
		want       string
		ok         bool
	}{
		{"117162  PSPPRLPIIGNYLQLGTLSHRSFQSLSQKYGPLMMLHLGQLPVLVVSSIHMAKEVMQTHG  117341", 2, 35, "PSPPRLPIIGNYLQLGTLSHRSFQSLSQKYGPLMMLHLGQLPVLVVSSIHMAKEVMQTHG", true},
		{"31168  LFDGTTIEGF  31139 31137 DADTTIKATML (0) 31105", 120, 1325, "LFDGTTIEGFDADTTIKATML", true},
		{"8    E  6", 89, 999, "E", true},
		{"mDHQTLLLVISFVSATIL", 71, 793, "mDHQTLLLVISFVSATIL", true},
		{"pseudogene same seq as CR932040.2b", 48, 520, "", false},
		{"only 434 aa a little short", 251, 2699, "", false},
		{"DEFINITION Medicago truncatula cytochrome P450", 1, 15, "", false},
		{"89% to 704G9 [Medicago truncatula] MEFIDFLFAMKPLFPILIAIGLAGFIIKIHGIRNFDKKRKY", 361, 3763, "MEFIDFLFAMKPLFPILIAIGLAGFIIKIHGIRNFDKKRKY", true},
	}
	for _, tt := range tests {
		got, ok := medicagoSequenceLine(tt.line, tt.block, tt.row)
		if got != tt.want || ok != tt.ok {
			t.Errorf("line %d got (%q,%t), want (%q,%t)", tt.row, got, ok, tt.want, tt.ok)
		}
	}
}

func TestParseMedicagoSeparatesSourcePiecesAndExclusions(t *testing.T) {
	text := strings.Join([]string{
		"Medicago truncatula (model legume species)",
		"",
		"There are 376 sequence pieces here.  Some are duplicates.",
		"Some are from other species for use in assembling  genes",
		"",
		"They are sorted by clan",
		"",
		">CYP51G1 Medicago truncatula",
		"MNVFDGNKFLNTLLLLITTLIAAKLISSFIIPKSKKRL*",
		">CG922887.1 Medicago truncatula genomic clone",
		"31 PFGSGRRICPGLPLAMRMLHMMLGSLLISFDWKLENDM 150",
	}, "\n")
	records, _, _, err := parseText(text, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(records) != 2 {
		t.Fatalf("records=%d", len(records))
	}
	if records[0].ID != "CYP51G1" || records[0].Symbol != "CYP51G1" || records[0].Sequence != "MNVFDGNKFLNTLLLLITTLIAAKLISSFIIPKSKKRL" || !strings.Contains(strings.Join(records[0].Status, ";"), "source-terminal-stop") {
		t.Fatalf("named=%#v", records[0])
	}
	if records[1].ID != "CG922887.1" || records[1].Symbol != "" || records[1].Sequence != "PFGSGRRICPGLPLAMRMLHMMLGSLLISFDWKLENDM" || !strings.Contains(strings.Join(records[1].Status, ";"), "source-accession-piece") {
		t.Fatalf("accession=%#v", records[1])
	}
}

func TestMedicagoExclusionListIsExact(t *testing.T) {
	if len(excludedBlocks) != 27 {
		t.Fatalf("excluded blocks=%d", len(excludedBlocks))
	}
	foreign, falsePositive := 0, 0
	for _, item := range excludedBlocks {
		if strings.Contains(item.Reason, "false positive") {
			falsePositive++
		} else {
			foreign++
		}
	}
	if foreign != 24 || falsePositive != 3 {
		t.Fatalf("foreign=%d false-positive=%d", foreign, falsePositive)
	}
}
