package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExactSorghum(t *testing.T) {
	d, e := os.ReadFile(filepath.Join("..", "..", "raw", "plants-sorghum.txt"))
	if e != nil {
		t.Fatal(e)
	}
	r, e := parse(string(d))
	if e != nil {
		t.Fatal(e)
	}
	if len(r) != 628 {
		t.Fatal(len(r))
	}
	for _, c := range []struct {
		i, b, l int
		id, s   string
	}{{0, 1, 142, "Sb01g035160", "CYP71K10"}, {313, 314, 525, "Sb07g005120", ""}, {627, 628, 496, "fgenesh1_pm.C_chr_2000675", "CYP734A5"}} {
		v := r[c.i]
		if v.Block != c.b || v.ID != c.id || v.Symbol != c.s || len(v.Sequence) != c.l {
			t.Errorf("case=%+v got=%+v", c, v)
		}
	}
}
