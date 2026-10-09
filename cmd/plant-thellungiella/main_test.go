package main

import (
	"strings"
	"testing"
)

func TestReviewThellungiellaRowsUsesExactWorkbookBoundaries(t *testing.T) {
	rows := map[int]map[string]string{
		1: {"H": "seq ID", "I": "best hit", "J": "%ID", "K": "CYP name", "L": ""},
	}
	for row := 2; row <= 213; row++ {
		rows[row] = map[string]string{"H": "Thelparv", "I": "CYP71A1", "K": "CYP71A", "L": strings.Repeat("A", 500)}
	}
	for row := 209; row <= 213; row++ {
		rows[row]["I"] = ""
		rows[row]["K"] = ""
	}
	rows[40]["I"] = "CYP71B4"
	rows[40]["K"] = "CYP71B pseudo"
	rows[40]["L"] = strings.Repeat("A", 332) + "O"
	rows[80]["L"] = strings.Repeat("A", 301) + "X"
	rows[179]["L"] = strings.Repeat("A", 462) + "-"

	accepted, excluded, err := reviewThellungiellaRows(rows)
	if err != nil {
		t.Fatal(err)
	}
	if len(accepted) != 207 || len(excluded) != 5 {
		t.Fatalf("accepted=%d excluded=%d", len(accepted), len(excluded))
	}
	if accepted[0].Row != 2 || accepted[103].Row != 105 || accepted[len(accepted)-1].Row != 208 {
		t.Fatalf("first/middle/last rows=%d/%d/%d", accepted[0].Row, accepted[103].Row, accepted[len(accepted)-1].Row)
	}
	if excluded[0].Row != 209 || excluded[len(excluded)-1].Row != 213 {
		t.Fatalf("excluded boundary=%d/%d", excluded[0].Row, excluded[len(excluded)-1].Row)
	}
	if accepted[38].Status != "source-pseudogene-label;nonstandard-O;short-sequence" {
		t.Fatalf("pseudo status=%q", accepted[38].Status)
	}
	if accepted[78].Status != "ambiguous-X;short-sequence" {
		t.Fatalf("X status=%q", accepted[78].Status)
	}
	if accepted[177].Status != "source-gap" {
		t.Fatalf("gap status=%q", accepted[177].Status)
	}
}

func TestReviewThellungiellaRowsRejectsCapsicumHeader(t *testing.T) {
	rows := map[int]map[string]string{1: {"H": "seq. ID", "I": "best hit", "J": "%ID", "K": "CYP name"}}
	if _, _, err := reviewThellungiellaRows(rows); err == nil {
		t.Fatal("expected exact Thellungiella header validation error")
	}
}
