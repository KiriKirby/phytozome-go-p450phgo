package main

import (
	"strings"
	"testing"
)

func TestReviewRowsUsesEutremaLayoutAndBoundary(t *testing.T) {
	rows := map[int]map[string]string{1: {"H": "Seq. ID", "I": "best hit ", "J": "%ID", "K": "CYP name"}}
	for row := 2; row <= 231; row++ {
		rows[row] = map[string]string{"H": "Eutrsals", "I": "CYP71A1", "K": "CYP71A", "L": strings.Repeat("A", 500)}
	}
	for row := 228; row <= 231; row++ {
		rows[row]["I"] = ""
		rows[row]["K"] = ""
	}
	rows[40]["I"], rows[40]["K"] = "CYP71B11 pseudo", "CYP71B pseudo"
	rows[46]["I"], rows[46]["L"] = "CYP72A266P", strings.Repeat("A", 308)
	accepted, excluded, err := reviewRows(rows)
	if err != nil {
		t.Fatal(err)
	}
	if len(accepted) != 226 || len(excluded) != 4 {
		t.Fatalf("accepted=%d excluded=%d", len(accepted), len(excluded))
	}
	if accepted[0].Row != 2 || accepted[113].Row != 115 || accepted[225].Row != 227 || excluded[0].Row != 228 || excluded[3].Row != 231 {
		t.Fatal("Eutrema row boundary changed")
	}
	if accepted[38].Status != "source-pseudogene-label" || accepted[44].Status != "source-pseudogene-label;short-sequence" {
		t.Fatalf("statuses=%q/%q", accepted[38].Status, accepted[44].Status)
	}
}

func TestReviewRowsRejectsThellungiellaHeader(t *testing.T) {
	if _, _, err := reviewRows(map[int]map[string]string{1: {"H": "seq ID", "I": "best hit", "J": "%ID", "K": "CYP name"}}); err == nil {
		t.Fatal("expected exact Eutrema header error")
	}
}
