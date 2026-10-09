package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestOldSequenceLine(t *testing.T) {
	for _, x := range []struct {
		in, out string
		ok      bool
	}{{"MDMENT TQMT*", "MDMENTTQMT*", true}, {"CYP71 family", "", false}, {"85% to CYP51G1", "", false}} {
		g, o := oldSequenceLine(x.in)
		if g != x.out || o != x.ok {
			t.Errorf("%q=(%q,%v)", x.in, g, o)
		}
	}
}
func TestParseOldAquilegia(t *testing.T) {
	d, e := os.ReadFile(filepath.Join("..", "..", "raw", "plants-Aquilegia.P450s.txt"))
	if e != nil {
		t.Fatal(e)
	}
	r, e := parseAquilegiaOld(string(d))
	if e != nil {
		t.Fatal(e)
	}
	if len(r) != 472 {
		t.Fatal(len(r))
	}
	for _, c := range []struct {
		i, b, l     int
		id, s, p, x string
	}{{0, 1, 496, "22061829", "CYP51G1", "MDMENTTQ", "MVRFKRRQLSID"}, {236, 237, 504, "22033044", "CYP81", "MEIFSCFI", "TRSSVLNFISQL"}, {471, 472, 112, "22037341", "CYP749", "MKIEATKT", "FQIEKLKVCLLL"}} {
		x := r[c.i]
		if x.Block != c.b || x.ID != c.id || x.Symbol != c.s || len(x.Sequence) != c.l || !strings.HasPrefix(x.Sequence, c.p) || !strings.HasSuffix(x.Sequence, c.x) {
			t.Errorf("i=%d block=%d id=%s symbol=%s len=%d pre/suf=%s/%s", c.i, x.Block, x.ID, x.Symbol, len(x.Sequence), x.Sequence[:min(8, len(x.Sequence))], x.Sequence[max(0, len(x.Sequence)-12):])
		}
	}
}
