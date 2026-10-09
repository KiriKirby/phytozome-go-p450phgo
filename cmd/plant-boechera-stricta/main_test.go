package main

import (
	"strings"
	"testing"
)

func TestReviewRowsUsesBoecheraStrictaLayout(t *testing.T) {
	rows := map[int]map[string]string{1: {"H": "seq ID", "I": "best hit", "J": "%ID", "K": "CYP name"}}
	for row := 2; row <= 234; row++ {
		rows[row] = map[string]string{"H": "Boecstri", "I": "CYP71A1", "K": "CYP71A", "L": strings.Repeat("A", 500)}
	}
	for row := 229; row <= 234; row++ {
		rows[row]["I"] = ""
		rows[row]["K"] = ""
	}
	rows[69]["I"] = "CYP76C8P"
	rows[14]["L"] = strings.Repeat("X", 96) + strings.Repeat("A", 281)
	rows[45]["L"] = strings.Repeat("A", 292) + "-"
	rows[36]["L"] = "O" + strings.Repeat("A", 420)

	accepted, excluded, err := reviewRows(rows, "A1:L234")
	if err != nil {
		t.Fatal(err)
	}
	if len(accepted) != 227 || len(excluded) != 6 {
		t.Fatalf("accepted=%d excluded=%d", len(accepted), len(excluded))
	}
	if accepted[0].Row != 2 || accepted[113].Row != 115 || accepted[226].Row != 228 || excluded[0].Row != 229 || excluded[5].Row != 234 {
		t.Fatal("Boechera stricta row boundary changed")
	}
	if accepted[67].Status != "source-pseudogene-label" || accepted[12].Status != "ambiguous-X" {
		t.Fatalf("statuses=%q/%q", accepted[67].Status, accepted[12].Status)
	}
	if accepted[43].Status != "source-gap;short-sequence" || accepted[34].Status != "nonstandard-O" {
		t.Fatalf("exception statuses=%q/%q", accepted[43].Status, accepted[34].Status)
	}
}

func TestReviewRowsRequiresExactBoecheraShape(t *testing.T) {
	rows := map[int]map[string]string{1: {"H": "seq ID", "I": "best hit", "J": "%ID", "K": "CYP name"}}
	if _, _, err := reviewRows(rows, "A1:L235"); err == nil {
		t.Fatal("expected exact Boechera stricta used-range error")
	}
	if _, _, err := reviewRows(map[int]map[string]string{1: {"H": "Seq. ID", "I": "best hit", "J": "%ID", "K": "CYP name"}}, "A1:L234"); err == nil {
		t.Fatal("expected exact Boechera stricta header error")
	}
}
