package main

import "testing"

func TestReviewCapsicumRowsUsesWorkbookSpecificColumns(t *testing.T) {
	rows := map[int]map[string]string{
		1: {"H": "seq. ID", "I": "best hit", "J": "%ID", "K": "CYP name", "L": ""},
	}
	for row := 2; row <= 650; row++ {
		rows[row] = map[string]string{"H": "Capsannu", "I": "CYP71D1", "K": "CYP71D", "L": "M" + string(make([]byte, 0)) + "ACDEFGHIKLMNPQRSTVWY"}
	}
	rows[618]["L"] = "MXXXXACDEFGHIKLMNPQRSTVWY-"
	for row := 619; row <= 650; row++ {
		rows[row]["K"] = ""
	}
	accepted, excluded, err := reviewCapsicumRows(rows)
	if err != nil {
		t.Fatal(err)
	}
	if len(accepted) != 617 || len(excluded) != 32 {
		t.Fatalf("accepted=%d excluded=%d", len(accepted), len(excluded))
	}
	if accepted[len(accepted)-1].Row != 618 || excluded[0].Row != 619 {
		t.Fatalf("boundary accepted=%d excluded=%d", accepted[len(accepted)-1].Row, excluded[0].Row)
	}
}
