package main

import (
	"strings"
	"testing"
)

func TestReviewRowsUsesCapsellaRubellaLayout(t *testing.T) {
	rows := map[int]map[string]string{1: {"H": "seq ID", "I": "best hit", "J": "%ID", "K": "CYP name"}}
	for row := 2; row <= 248; row++ {
		rows[row] = map[string]string{"H": "Capsrube", "I": "CYP71A1", "K": "CYP71A", "L": strings.Repeat("A", 500)}
	}
	for row := 247; row <= 248; row++ {
		rows[row]["I"] = ""
		rows[row]["K"] = ""
	}
	rows[16]["I"] = "CYP71A27P"
	rows[232]["L"] = strings.Repeat("A", 416) + strings.Repeat("X", 96)
	accepted, excluded, err := reviewRows(rows)
	if err != nil {
		t.Fatal(err)
	}
	if len(accepted) != 245 || len(excluded) != 2 {
		t.Fatalf("accepted=%d excluded=%d", len(accepted), len(excluded))
	}
	if accepted[0].Row != 2 || accepted[122].Row != 124 || accepted[244].Row != 246 || excluded[0].Row != 247 || excluded[1].Row != 248 {
		t.Fatal("Capsella rubella row boundary changed")
	}
	if accepted[14].Status != "source-pseudogene-label" || accepted[230].Status != "ambiguous-X" {
		t.Fatalf("statuses=%q/%q", accepted[14].Status, accepted[230].Status)
	}
}

func TestReviewRowsRequiresUppercaseSheetSpecificHeader(t *testing.T) {
	if _, _, err := reviewRows(map[int]map[string]string{1: {"H": "Seq. ID", "I": "best hit", "J": "%ID", "K": "CYP name"}}); err == nil {
		t.Fatal("expected exact Capsella rubella header error")
	}
}
