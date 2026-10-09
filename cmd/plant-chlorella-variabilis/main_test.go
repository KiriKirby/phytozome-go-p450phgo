package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestChlorellaVariabilis(t *testing.T) {
	d, e := os.ReadFile(filepath.Join("..", "..", "raw", "plants-Chlorella.variabilis.txt"))
	if e != nil {
		t.Fatal(e)
	}
	r, e := parse(string(d))
	if e != nil {
		t.Fatal(e)
	}
	if len(r) != 19 {
		t.Fatal(len(r))
	}
	if r[0].Symbol != "CYP51G1" || len(r[0].Sequence) != 494 || !strings.HasSuffix(r[0].Sequence, "CRVRYRRKKLVA") {
		t.Fatalf("first %+v len=%d", r[0], len(r[0].Sequence))
	}
	if r[4].Symbol != "CYP710B1" || !strings.Contains(r[4].Sequence, "XXXXXXXXXXXXXXXX") {
		t.Fatalf("X record %+v", r[4])
	}
	if r[10].Sequence != r[11].Sequence || r[10].ID == r[11].ID {
		t.Fatalf("duplicate pair not retained: %+v %+v", r[10], r[11])
	}
	if r[18].Symbol != "CYP851A1" || !strings.HasSuffix(r[18].Sequence, "RLTLKPRQ") {
		t.Fatalf("last %+v", r[18])
	}
}
