package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMarchantia(t *testing.T) {
	d, e := os.ReadFile(filepath.Join("..", "..", "raw", "plants-liverwort.txt"))
	if e != nil {
		t.Fatal(e)
	}
	r, e := parse(string(d))
	if e != nil {
		t.Fatal(e)
	}
	for _, c := range []struct {
		i, l    int
		sym, id string
	}{{0, 120, "CYP51G1", "BJ856245"}, {18, 197, "CYP", "BJ846270"}, {37, 60, "CYP74", "BJ840718"}} {
		v := r[c.i]
		if len(v.Sequence) != c.l || v.Symbol != c.sym || v.ID != c.id {
			t.Errorf("case=%+v got=%+v len=%d", c, v, len(v.Sequence))
		}
	}
}
