package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestVitisSequenceLineUsesReviewedLiteralForms(t *testing.T) {
	tests := []struct {
		line, want string
		ok         bool
	}{
		{"190429  MDVDNKFFNVALLIVATVVVAKLISALLIPKSRKRLP  190256", "MDVDNKFFNVALLIVATVVVAKLISALLIPKSRKRLP", true},
		{"176205  DLFLAGVDTGAITLTWAMTELARNPRIMKKAQ (0) 175594", "DLFLAGVDTGAITLTWAMTELARNPRIMKKAQ", true},
		{"QEVREAxxxx 4046", "QEVREAxxxx", true},
		{"C  L  Q  P  L  E  N  I  Y*", "CLQPLENIY*", true},
		{"$$$$", "", false}, {"&&&&&", "", false}, {"(GAP)", "", false},
		{"CYP72 family", "", false}, {"same seq as CAAP02000057.1", "", false},
	}
	for _, tt := range tests {
		got, ok := vitisSequenceLine(tt.line)
		if got != tt.want || ok != tt.ok {
			t.Errorf("vitisSequenceLine(%q)=(%q,%v), want (%q,%v)", tt.line, got, ok, tt.want, tt.ok)
		}
	}
}

func TestVitisForeignReferenceLedger(t *testing.T) {
	if len(foreignReferenceBlocks) != 30 {
		t.Fatalf("foreign reference blocks=%d", len(foreignReferenceBlocks))
	}
	for _, block := range []int{5, 74, 338, 365, 401, 593, 606} {
		if _, ok := foreignReferenceBlocks[block]; !ok {
			t.Errorf("missing reviewed foreign block %d", block)
		}
	}
	for _, native := range []int{8, 455, 520, 579, 693} {
		if _, ok := foreignReferenceBlocks[native]; ok {
			t.Errorf("Vitis comparison block %d incorrectly excluded", native)
		}
	}
}

func TestParseReviewedVitisSource(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "..", "raw", "plants-vitis.txt"))
	if err != nil {
		t.Fatal(err)
	}
	records, excluded, err := parseVitisText(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if len(records) != 672 || len(excluded) != 30 || countSequences(records) != 672 {
		t.Fatalf("records=%d sequences=%d excluded=%d", len(records), countSequences(records), len(excluded))
	}
	checks := []struct {
		index, block, length   int
		symbol, prefix, suffix string
	}{
		{0, 1, 486, "CYP51G6", "MDVDNKFF", "VMVRYKRRVLPVD"},
		{len(records) / 2, 348, 526, "CYP82D10", "MYFLLQYLNITT", "VLISPRLSSCSLYN"},
		{len(records) - 1, 702, 492, "CYP736A27", "MAVWTWTA", "HLVAIPTYRLRQ"},
	}
	for _, c := range checks {
		r := records[c.index]
		if r.Block != c.block || r.Symbol != c.symbol || len(r.Sequence) != c.length || !strings.HasPrefix(r.Sequence, c.prefix) || !strings.HasSuffix(r.Sequence, c.suffix) {
			t.Errorf("record[%d]=block %d symbol %s len %d prefix/suffix %.12s/%.14s", c.index, r.Block, r.Symbol, len(r.Sequence), r.Sequence, r.Sequence[max(0, len(r.Sequence)-14):])
		}
	}
	for _, r := range records {
		if r.Block == 693 && (len(r.Sequence) != 500 || !strings.HasSuffix(r.Sequence, "NHLYAIPTYRLLI")) {
			t.Errorf("block 693 boundary wrong: len=%d suffix=%s", len(r.Sequence), r.Sequence[max(0, len(r.Sequence)-20):])
		}
	}
}
