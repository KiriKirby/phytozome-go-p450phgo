package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBrachypodiumLineForms(t *testing.T) {
	for _, x := range []struct {
		in, out string
		ok      bool
	}{{"541 NIFLGGVDTGAIVLVWAMAELVRNP 29 &", "NIFLGGVDTGAIVLVWAMAELVRNP", true}, {"(?) DYFEKWGQEGIIDLKHELDQV", "DYFEKWGQEGIIDLKHELDQV", true}, {"(deletion)", "", false}, {"86% to CYP710A8 rice", "", false}} {
		g, o := brachypodiumSequenceLine(x.in)
		if g != x.out || o != x.ok {
			t.Errorf("%q=(%q,%v)", x.in, g, o)
		}
	}
}
func TestParseBrachypodium(t *testing.T) {
	d, e := os.ReadFile(filepath.Join("..", "..", "raw", "plants-Brachypodium.FASTA.txt"))
	if e != nil {
		t.Fatal(e)
	}
	r, x, e := parseBrachypodium(string(d))
	if e != nil {
		t.Fatal(e)
	}
	if len(r) != 278 || len(x) != 1 || x[0].Block != 74 {
		t.Fatalf("records=%d excluded=%v", len(r), x)
	}
	for _, c := range []struct {
		i, b, l     int
		id, s, p, z string
	}{{0, 1, 490, "Bradi4g25930", "CYP51G1", "MDLLAAEP", "VNYKRRKLIVEN"}, {139, 141, 532, "Bradi3g47750", "CYP86E1", "MAAAAGWT", "NVHYSTTVAAAAEE"}, {277, 279, 529, "Bradi4g29860.1", "CYP735A4", "MAMATASA", "PKHGVPVHLRPLRP"}} {
		v := r[c.i]
		if v.Block != c.b || v.ID != c.id || v.Symbol != c.s || len(v.Sequence) != c.l || !strings.HasPrefix(v.Sequence, c.p) || !strings.HasSuffix(v.Sequence, c.z) {
			t.Errorf("case=%+v got block=%d id=%s symbol=%s len=%d", c, v.Block, v.ID, v.Symbol, len(v.Sequence))
		}
	}
}
