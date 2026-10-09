package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFerns(t *testing.T) {
	d, e := os.ReadFile(filepath.Join("..", "..", "raw", "plants-ferns.txt"))
	if e != nil {
		t.Fatal(e)
	}
	r, e := parse(string(d))
	if e != nil {
		t.Fatal(e)
	}
	for _, c := range []struct {
		i, l        int
		sp, sym, id string
	}{{0, 484, "Ceratopteris richardii", "CYP51G1", "CV734775.1"}, {8, 153, "Adiantum capillus-veneris", "CYP75?", "BP921525"}, {17, 138, "Ceratopteris richardii", "CYP", "BQ087157"}} {
		v := r[c.i]
		if len(v.Sequence) != c.l || v.Species != c.sp || v.Symbol != c.sym || v.ID != c.id {
			t.Errorf("case=%+v got=%+v len=%d", c, v, len(v.Sequence))
		}
	}
}
