package main

import (
	"strings"
	"testing"
)

func TestReviewRowsUsesBrassicaRapaLayout(t *testing.T) {
	rows := map[int]map[string]string{1: {"H": "seq ID", "I": "best hit", "J": "%ID", "K": "CYP name"}}
	for row := 2; row <= 383; row++ {
		rows[row] = map[string]string{"H": "Brasrapa", "I": "CYP71A1", "K": "CYP71A", "L": strings.Repeat("A", 500)}
	}
	for row := 379; row <= 383; row++ {
		rows[row]["I"] = ""
		rows[row]["K"] = ""
	}
	rows[30]["I"] = "CYP71B18P"
	rows[8]["L"] = strings.Repeat("X", 226) + strings.Repeat("A", 207)
	rows[29]["L"] = strings.Repeat("A", 253) + "-"
	accepted, excluded, err := reviewRows(rows)
	if err != nil {
		t.Fatal(err)
	}
	if len(accepted) != 377 || len(excluded) != 5 {
		t.Fatalf("accepted=%d excluded=%d", len(accepted), len(excluded))
	}
	if accepted[0].Row != 2 || accepted[188].Row != 190 || accepted[376].Row != 378 || excluded[0].Row != 379 || excluded[4].Row != 383 {
		t.Fatal("Brassica rapa row boundary changed")
	}
	if accepted[28].Status != "source-pseudogene-label" || accepted[6].Status != "ambiguous-X" {
		t.Fatalf("statuses=%q/%q", accepted[28].Status, accepted[6].Status)
	}
	if accepted[27].Status != "source-gap;short-sequence" {
		t.Fatalf("gap status=%q", accepted[27].Status)
	}
}

func TestReviewRowsRequiresExactHeader(t *testing.T) {
	if _, _, err := reviewRows(map[int]map[string]string{1: {"H": "Seq. ID", "I": "best hit", "J": "%ID", "K": "CYP name"}}); err == nil {
		t.Fatal("expected exact Brassica rapa header error")
	}
}
