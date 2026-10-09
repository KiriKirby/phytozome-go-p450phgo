package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseTextKeepsUnannotatedDuplicateBlocks(t *testing.T) {
	preamble := "Prunus Persica P450s\n305 sequences retrieved\nThese sequences have not been annotated manually.\n"
	var b strings.Builder
	b.WriteString(preamble)
	for i := 1; i <= 3; i++ {
		fmtID := fmt.Sprintf("ppa%06dm", i)
		b.WriteString(fmt.Sprintf(">%d_peptide|Ppersica|%s.g|%s\n", 10000000+i, fmtID, fmtID))
		sequence := strings.Repeat("A", 400)
		if i <= 2 {
			sequence += "*"
		}
		b.WriteString(sequence + "\n")
	}
	records, _, err := parseTextWithLayout(b.String(), false)
	if err != nil {
		t.Fatal(err)
	}
	if len(records) != 3 || records[0].ID != "ppa000001m" || len(records[0].Sequence) != 400 {
		t.Fatalf("records=%#v", records)
	}
	if strings.Join(records[0].Status, ";") != "source-unannotated;source-terminal-stop" || strings.Join(records[2].Status, ";") != "source-unannotated" {
		t.Fatalf("statuses=%v / %v", records[0].Status, records[2].Status)
	}
}

func TestHeaderPattern(t *testing.T) {
	match := headerRE.FindStringSubmatch(">17641348_peptide|Ppersica|ppa004078m.g|ppa004078m")
	if match == nil || match[1] != "17641348" || match[2] != "ppa004078m.g" || match[3] != "ppa004078m" {
		t.Fatalf("match=%v", match)
	}
}
