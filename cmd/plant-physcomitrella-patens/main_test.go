package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMoss(t *testing.T) {
	d, e := os.ReadFile(filepath.Join("..", "..", "raw", "plants-moss.txt"))
	if e != nil {
		t.Fatal(e)
	}
	a, x, e := parse(string(d))
	if e != nil {
		t.Fatal(e)
	}
	if len(a) != 96 || len(x) != 19 {
		t.Fatal(len(a), len(x))
	}
	if a[0].Symbol != "CYP51G1" || len(a[0].Sequence) != 502 || a[len(a)-1].Block != 108 || len(a[len(a)-1].Sequence) != 527 {
		t.Fatalf("representatives changed %+v %+v", a[0], a[len(a)-1])
	}
}
