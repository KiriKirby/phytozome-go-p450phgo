package main

import (
	"strings"
	"testing"
)

func TestReviewRowsUsesMimulusLayout(t *testing.T) {
	rows := map[int]map[string]string{1: {"H": "seq. ID", "I": "best hit", "J": "%ID", "K": "CYP name"}}
	for row := 2; row <= 383; row++ {
		rows[row] = map[string]string{"H": "Mimugutt", "I": "CYP71A1", "K": "CYP71A", "L": strings.Repeat("A", 500)}
	}
	for row := 370; row <= 383; row++ {
		rows[row]["I"] = ""
		rows[row]["K"] = ""
	}
	rows[20]["I"] = "CYP71AP14"
	rows[30]["I"] = "CYP71AU36P"
	rows[16]["L"] = strings.Repeat("X", 15) + strings.Repeat("A", 473)
	rows[45]["L"] = "O" + strings.Repeat("A", 367)
	rows[73]["L"] = strings.Repeat("A", 431) + "-"
	accepted, excluded, err := reviewRows(rows, "A1:L383")
	if err != nil {
		t.Fatal(err)
	}
	if len(accepted) != 368 || len(excluded) != 14 {
		t.Fatalf("accepted=%d excluded=%d", len(accepted), len(excluded))
	}
	if accepted[0].Row != 2 || accepted[183].Row != 185 || accepted[367].Row != 369 || excluded[0].Row != 370 || excluded[13].Row != 383 {
		t.Fatal("Mimulus row boundary changed")
	}
	if strings.Contains(accepted[18].Status, "pseudogene") {
		t.Fatalf("CYP71AP14 false pseudogene=%q", accepted[18].Status)
	}
	if accepted[28].Status != "source-pseudogene-label" || accepted[14].Status != "ambiguous-X" {
		t.Fatalf("statuses=%q/%q", accepted[28].Status, accepted[14].Status)
	}
	if accepted[43].Status != "nonstandard-O" || accepted[71].Status != "source-gap" {
		t.Fatalf("exceptions=%q/%q", accepted[43].Status, accepted[71].Status)
	}
}

func TestReviewRowsRequiresExactMimulusHeader(t *testing.T) {
	if _, _, err := reviewRows(map[int]map[string]string{1: {"H": "seq ID", "I": "best hit", "J": "%ID", "K": "CYP name"}}, "A1:L383"); err == nil {
		t.Fatal("expected exact Mimulus header error")
	}
}
