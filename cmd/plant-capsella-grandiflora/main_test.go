package main

import (
	"strings"
	"testing"
)

func TestReviewRowsUsesCapsellaGrandifloraColumns(t *testing.T) {
	rows := map[int]map[string]string{1: {"H": "best hit", "I": "%ID", "J": "CYP name"}}
	for row := 2; row <= 227; row++ {
		rows[row] = map[string]string{"G": "Capsgran", "H": "CYP71A1", "J": "CYP71A", "K": strings.Repeat("A", 500)}
	}
	for row := 222; row <= 227; row++ {
		rows[row]["H"] = ""
		rows[row]["J"] = ""
	}
	rows[4]["K"] = strings.Repeat("A", 439) + strings.Repeat("X", 58)
	rows[16]["H"] = "CYP71A27P"
	accepted, excluded, err := reviewRows(rows)
	if err != nil {
		t.Fatal(err)
	}
	if len(accepted) != 220 || len(excluded) != 6 {
		t.Fatalf("accepted=%d excluded=%d", len(accepted), len(excluded))
	}
	if accepted[0].Row != 2 || accepted[110].Row != 112 || accepted[219].Row != 221 || excluded[0].Row != 222 || excluded[5].Row != 227 {
		t.Fatal("Capsella grandiflora row boundary changed")
	}
	if accepted[2].Status != "ambiguous-X" || accepted[14].Status != "source-pseudogene-label" {
		t.Fatalf("statuses=%q/%q", accepted[2].Status, accepted[14].Status)
	}
}

func TestReviewRowsRejectsSequenceInWrongColumn(t *testing.T) {
	rows := map[int]map[string]string{1: {"H": "best hit", "I": "%ID", "J": "CYP name", "K": "sequence"}}
	if _, _, err := reviewRows(rows); err == nil {
		t.Fatal("expected exact blank K header validation error")
	}
}
