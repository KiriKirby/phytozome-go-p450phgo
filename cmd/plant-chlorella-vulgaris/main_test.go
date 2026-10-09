package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestChlorellaVulgaris(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "..", "raw", "plants-Chlorella.vulgaris.txt"))
	if err != nil {
		t.Fatal(err)
	}
	r, err := parse(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if len(r) != 37 {
		t.Fatal(len(r))
	}
	checks := []struct {
		i                      int
		symbol, prefix, suffix string
		length                 int
	}{{0, "CYP51G1", "MFDLSELR", "CRVRYRRRKLTL", 491}, {21, "CYP855C1P", "MSSSHTAD", "VPKGTGVMVRLS", 120}, {36, "CYP863-fragment1", "MLEQMPYT", "PGRVVKYQA", 59}}
	for _, c := range checks {
		x := r[c.i]
		if x.Symbol != c.symbol || !strings.HasPrefix(x.Sequence, c.prefix) || !strings.HasSuffix(x.Sequence, c.suffix) || len(x.Sequence) != c.length {
			t.Fatalf("check=%+v record=%+v length=%d", c, x, len(x.Sequence))
		}
	}
}
