package main

import (
	"strings"
	"testing"
)

func TestReviewRowsUsesGossypiumLayout(t *testing.T) {
	rows := map[int]map[string]string{1: {"H": "seq. ID", "I": "best hit", "J": "%ID", "K": "CYP name"}}
	for row := 2; row <= 462; row++ {
		rows[row] = map[string]string{"H": "Gossraim", "I": "CYP71A1", "K": "CYP71A", "L": strings.Repeat("A", 500)}
	}
	for row := 451; row <= 462; row++ {
		rows[row]["I"] = ""
		rows[row]["K"] = ""
	}
	rows[40]["I"] = "CYP71BE8P"
	rows[161]["L"] = strings.Repeat("A", 1520)
	rows[23]["L"] = strings.Repeat("X", 169) + strings.Repeat("A", 338)
	rows[55]["L"] = strings.Repeat("A", 200) + "-"
	rows[9]["L"] = "O" + strings.Repeat("A", 505)
	accepted, excluded, err := reviewRows(rows, "A1:L462")
	if err != nil {
		t.Fatal(err)
	}
	if len(accepted) != 449 || len(excluded) != 12 {
		t.Fatalf("accepted=%d excluded=%d", len(accepted), len(excluded))
	}
	if accepted[0].Row != 2 || accepted[224].Row != 226 || accepted[448].Row != 450 || excluded[0].Row != 451 || excluded[11].Row != 462 {
		t.Fatal("Gossypium row boundary changed")
	}
	if accepted[38].Status != "source-pseudogene-label" || accepted[159].Status != "unusually-long-sequence" {
		t.Fatalf("statuses=%q/%q", accepted[38].Status, accepted[159].Status)
	}
	if accepted[21].Status != "ambiguous-X" || accepted[53].Status != "source-gap;short-sequence" || accepted[7].Status != "nonstandard-O" {
		t.Fatalf("exceptions=%q/%q/%q", accepted[21].Status, accepted[53].Status, accepted[7].Status)
	}
}

func TestReviewRowsRequiresExactGossypiumShape(t *testing.T) {
	if _, _, err := reviewRows(map[int]map[string]string{1: {"H": "seq ID", "I": "best hit", "J": "%ID", "K": "CYP name"}}, "A1:L462"); err == nil {
		t.Fatal("expected exact Gossypium header error")
	}
}
