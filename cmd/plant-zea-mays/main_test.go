package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestZeaSequenceLine(t *testing.T) {
	for _, x := range []struct {
		in, out string
		ok      bool
	}{
		{"274 TTEGEVTGLLIAALFAGQHT 293", "TTEGEVTGLLIAALFAGQHT", true},
		{"    LADMDVLYRCIKEALRLHP", "LADMDVLYRCIKEALRLHP", true},
		{"71C2", "", false}, {"ESTs sorted alphabetically", "", false},
	} {
		got, ok := zeaSequenceLine(x.in)
		if got != x.out || ok != x.ok {
			t.Errorf("%q=(%q,%v)", x.in, got, ok)
		}
	}
}

func TestParseZea(t *testing.T) {
	d, err := os.ReadFile(filepath.Join("..", "..", "raw", "plants-zea.txt"))
	if err != nil {
		t.Fatal(err)
	}
	text, err := decodeZeaText(d)
	if err != nil {
		t.Fatal(err)
	}
	rows, index, err := parseZea(text)
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 26 || countSequences(rows) != 18 || len(index) != 42 {
		t.Fatalf("rows=%d sequences=%d index=%d", len(rows), countSequences(rows), len(index))
	}
	for _, c := range []struct {
		i, block, length           int
		id, symbol, prefix, suffix string
	}{
		{0, 1, 219, "AI770623", "CYP51", "TTEGEVTG", "MVNYKRRKLVVDN"},
		{3, 4, 0, "AI714669", "CYP71C2", "", ""},
		{11, 12, 268, "AI855377", "CYP72A5", "LLANGLVN", "PYXVITLHP"},
		{22, 23, 282, "AI734373", "CYP98A1", "KIGASLSI", "TFMATPLQAVATPRL"},
		{25, 26, 127, "AI947887", "CYP714A2", "LRQLKILT", "SKFSFSVSPGYQHS"},
	} {
		r := rows[c.i]
		if r.Block != c.block || r.ID != c.id || r.Symbol != c.symbol || len(r.Sequence) != c.length || !strings.HasPrefix(r.Sequence, c.prefix) || !strings.HasSuffix(r.Sequence, c.suffix) {
			t.Errorf("case %+v got block=%d id=%s symbol=%s len=%d seq=%s", c, r.Block, r.ID, r.Symbol, len(r.Sequence), r.Sequence)
		}
	}
	if !strings.Contains(rows[1].Sequence, "X") || !strings.Contains(rows[21].Sequence, "*") {
		t.Fatal("literal X or internal stop lost")
	}
	if strings.HasSuffix(rows[0].Sequence, "*") {
		t.Fatal("terminal stop retained")
	}
}
