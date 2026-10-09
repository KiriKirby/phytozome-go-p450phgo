package main

import (
	"strings"
	"testing"
)

func TestReviewRowsUsesTheobromaLayout(t *testing.T) {
	rows := map[int]map[string]string{1: {"H": "Seq ID", "I": "best hit", "J": "%ID", "K": "CYP name"}}
	for row := 2; row <= 347; row++ {
		rows[row] = map[string]string{"H": "Theocaca", "I": "CYP71A1", "K": "CYP71A", "L": strings.Repeat("A", 500)}
	}
	for row := 338; row <= 347; row++ {
		rows[row]["I"] = ""
		rows[row]["K"] = ""
	}
	rows[17]["I"] = "CYP71AT5P"
	rows[93]["L"] = strings.Repeat("X", 72) + strings.Repeat("A", 457)
	rows[266]["L"] = strings.Repeat("A", 432) + "-"
	rows[15]["L"] = "O" + strings.Repeat("A", 316)
	accepted, excluded, err := reviewRows(rows, "A1:L347")
	if err != nil {
		t.Fatal(err)
	}
	if len(accepted) != 336 || len(excluded) != 10 {
		t.Fatalf("accepted=%d excluded=%d", len(accepted), len(excluded))
	}
	if accepted[0].Row != 2 || accepted[167].Row != 169 || accepted[335].Row != 337 || excluded[0].Row != 338 || excluded[9].Row != 347 {
		t.Fatal("Theobroma cacao row boundary changed")
	}
	if accepted[15].Status != "source-pseudogene-label" || accepted[91].Status != "ambiguous-X" {
		t.Fatalf("statuses=%q/%q", accepted[15].Status, accepted[91].Status)
	}
	if accepted[264].Status != "source-gap" || accepted[13].Status != "nonstandard-O;short-sequence" {
		t.Fatalf("exceptions=%q/%q", accepted[264].Status, accepted[13].Status)
	}
}

func TestReviewRowsRequiresExactTheobromaHeader(t *testing.T) {
	if _, _, err := reviewRows(map[int]map[string]string{1: {"H": "Seq. ID", "I": "best hit", "J": "%ID", "K": "CYP name"}}, "A1:L347"); err == nil {
		t.Fatal("expected exact Theobroma header error")
	}
}
