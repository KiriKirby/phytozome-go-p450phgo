package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExactLolium(t *testing.T) {
	d, e := os.ReadFile(filepath.Join("..", "..", "raw", "plants-lolium.txt"))
	if e != nil {
		t.Fatal(e)
	}
	r, e := parse(string(d))
	if e != nil {
		t.Fatal(e)
	}
	for _, c := range []struct {
		i, l   int
		id, cl string
	}{{0, 525, "AF321870", "Lol-83"}, {8, 520, "AF321862", "Lol-31-j"}, {15, 517, "AF321855", "Lol-FHH-v"}} {
		v := r[c.i]
		if v.ID != c.id || v.Clone != c.cl || len(v.Sequence) != c.l {
			t.Errorf("case=%+v got=%+v", c, v)
		}
	}
}
