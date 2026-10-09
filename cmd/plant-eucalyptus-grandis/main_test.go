package main

import (
	"strings"
	"testing"
)

func TestReviewRowsUsesEucalyptusLayout(t *testing.T) {
	rows := map[int]map[string]string{1: {"A": "Gotoh's seq ID", "I": "blast ID", "J": "best hit", "K": "%ID", "L": "sequence"}}
	for r := 2; r <= 770; r++ {
		rows[r] = map[string]string{"A": ">Eucagran", "I": "Eucagran", "J": "CYP71A1", "L": strings.Repeat("A", 500)}
	}
	for _, r := range []int{755, 757, 758, 759, 768, 769, 770} {
		rows[r]["J"] = "EucagranModel"
	}
	rows[2]["L"] = strings.Repeat("A", 250) + "x" + strings.Repeat("A", 249)
	rows[3]["L"] = strings.Repeat("A", 500) + "**"
	rows[4]["L"] = strings.Repeat("A", 250) + "*" + strings.Repeat("A", 249)
	rows[772] = map[string]string{"A": "green <55%"}
	rows[773] = map[string]string{"A": "yellow <39%"}
	a, x, e := reviewRows(rows, "A1:L773")
	if e != nil {
		t.Fatal(e)
	}
	if len(a) != 762 || len(x) != 7 {
		t.Fatalf("%d/%d", len(a), len(x))
	}
	if a[0].Row != 2 || a[len(a)-1].Row != 767 {
		t.Fatal("boundary")
	}
	if a[0].Status != "ambiguous-X-or-x" || !strings.Contains(a[0].Sequence, "x") {
		t.Fatalf("lowercase x not preserved: %#v", a[0])
	}
	if a[1].Status != "source-terminal-stop" || strings.Contains(a[1].Sequence, "*") || len(a[1].Sequence) != 500 {
		t.Fatalf("terminal stop handling changed: %#v", a[1])
	}
	if a[2].Status != "internal-stop" || !strings.Contains(a[2].Sequence, "*") {
		t.Fatalf("internal stop not preserved: %#v", a[2])
	}
}
func TestReviewRowsRequiresExactEucalyptusHeader(t *testing.T) {
	if _, _, e := reviewRows(map[int]map[string]string{1: {"A": "seq ID", "I": "blast ID", "J": "best hit", "K": "%ID", "L": "sequence"}}, "A1:L773"); e == nil {
		t.Fatal("expected header error")
	}
}
