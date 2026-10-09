package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPinusExactSourceLayout(t *testing.T) {
	d, err := os.ReadFile(filepath.Join("..", "..", "raw", "plants-Pinus.P450.txt"))
	if err != nil {
		t.Fatal(err)
	}
	r, err := parse(string(d))
	if err != nil {
		t.Fatal(err)
	}
	if len(r) != 78 || sequenceCount(r) != 4 {
		t.Fatalf("counts=%d/%d", len(r), sequenceCount(r))
	}
	for _, c := range []struct {
		i, length           int
		species, symbol, id string
	}{
		{0, 506, "Pinus taeda", "CYP73A20", "AF096998"},
		{1, 512, "Pinus taeda", "CYP73A23", "CYP73A23-assembled-ESTs"},
		{2, 553, "Pinus radiata", "CYP78A4", "AF049067"},
		{3, 512, "Pinus taeda", "CYP98A15", "AY064170"},
		{4, 0, "Pinus", "", "BF610290.1"},
		{77, 0, "Pinus", "", "AW754540.1"},
	} {
		v := r[c.i]
		if v.Species != c.species || v.Symbol != c.symbol || v.ID != c.id || len(v.Sequence) != c.length {
			t.Errorf("case %+v got %+v len=%d", c, v, len(v.Sequence))
		}
	}
}
