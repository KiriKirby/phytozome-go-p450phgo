package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMicromonas(t *testing.T) {
	d, e := os.ReadFile(filepath.Join("..", "..", "raw", "plants-micromonas.txt"))
	if e != nil {
		t.Fatal(e)
	}
	r, e := parse(string(d))
	if e != nil {
		t.Fatal(e)
	}
	if len(r) != 32 {
		t.Fatal(len(r))
	}
	if r[0].Species != "Micromonas sp. RCC299" || r[1].Species != "Micromonas pusilla CCMP1545" || r[28].Species != "Micromonas sp. CCMP490 EST" {
		t.Fatalf("species %q %q %q", r[0].Species, r[1].Species, r[28].Species)
	}
	if !strings.HasSuffix(r[0].Sequence, "HPCTVRYKRRKL") || strings.HasSuffix(r[16].Sequence, "*") {
		t.Fatal("stop handling")
	}
	if r[2].Sequence == r[3].Sequence {
		t.Fatal("alternate models unexpectedly merged")
	}
	if r[19].Sequence == "" || r[19].Status != "" {
		t.Fatalf("partial block %+v", r[19])
	}
	if !strings.Contains(r[16].Sequence, "X") {
		t.Fatal("X fragment lost")
	}
}
