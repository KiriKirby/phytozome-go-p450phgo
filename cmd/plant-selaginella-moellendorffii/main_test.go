package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSelaginella(t *testing.T) {
	d, e := os.ReadFile(filepath.Join("..", "..", "raw", "plants-Selaginella.P450s.txt"))
	if e != nil {
		t.Fatal(e)
	}
	r, e := parse(string(d))
	if e != nil {
		t.Fatal(e)
	}
	if len(r) != 517 {
		t.Fatal(len(r))
	}
	for _, c := range []struct {
		i, l    int
		sym, id string
	}{{0, 494, "CYP51G1v1", "estExt_Genewise1Plus.C_470465|Selmo1"}, {258, 123, "CYP791A2v1", "gw1.112.69.1|Selmo1"}, {516, 512, "CYP711A17v2", "e_gw1.19.137.1|Selmo1"}} {
		v := r[c.i]
		if len(v.Sequence) != c.l || v.Symbol != c.sym || v.ID != c.id {
			t.Errorf("case=%+v got=%+v len=%d", c, v, len(v.Sequence))
		}
	}
}
func TestSequenceLineRejectsNarrative(t *testing.T) {
	for _, s := range []string{"No allele found", "remove intron", "Duplicate of part of last exon", "This model is incorrect"} {
		if q, ok := sequenceLine(s); ok {
			t.Errorf("accepted %q as %q", s, q)
		}
	}
}
