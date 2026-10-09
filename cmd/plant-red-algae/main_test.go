package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRedAlgae(t *testing.T) {
	ct, e := os.ReadFile(filepath.Join("..", "..", "raw", "plants-redalgae.txt"))
	if e != nil {
		t.Fatal(e)
	}
	it, e := os.ReadFile(filepath.Join("..", "..", "raw", "plants-red.algae.index.txt"))
	if e != nil {
		t.Fatal(e)
	}
	c, e := parseCyan(string(ct))
	if e != nil {
		t.Fatal(e)
	}
	a, e := parseAlignment(string(it))
	if e != nil {
		t.Fatal(e)
	}
	r, e := combine(c, a)
	if e != nil {
		t.Fatal(e)
	}
	if len(r) != 13 {
		t.Fatal(len(r))
	}
	if r[0].ID != "201" || !strings.HasPrefix(r[0].Sequence, "MILARNLL") || r[4].ID != "444" {
		t.Fatalf("cyan endpoints %+v %+v", r[0], r[4])
	}
	if r[5].ID != "CYP51con10" || len(r[5].Sequence) != 628 || !strings.Contains(r[5].Sequence, "-") {
		t.Fatalf("Galdiera first %+v", r[5])
	}
	if r[12].ID != "con1041" || len(r[12].Sequence) != 628 {
		t.Fatalf("last %+v", r[12])
	}
}
