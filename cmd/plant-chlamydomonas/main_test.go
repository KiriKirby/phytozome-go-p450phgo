package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestChlamydomonas(t *testing.T) {
	d, e := os.ReadFile(filepath.Join("..", "..", "raw", "plants-chlamydomonas.txt"))
	if e != nil {
		t.Fatal(e)
	}
	r, e := parse(string(d))
	if e != nil {
		t.Fatal(e)
	}
	if len(r) != 41 {
		t.Fatal(len(r))
	}
	for _, i := range []int{0, 20, 40} {
		if r[i].Sequence == "" {
			t.Fatalf("empty representative %+v", r[i])
		}
	}
}
