package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNelumboSequenceLineReviewedForms(t *testing.T) {
	tests := []struct {
		line, want string
		ok         bool
	}{{"8439395", "", false}, {"MDLKENKFLSVGLLIVATMVVAKFLSAFLM", "MDLKENKFLSVGLLIVATMVVAKFLSAFLM", true}, {"1925414 43582 GYKGDKEGYPIPAGTDLF (0) 43526 1925361", "GYKGDKEGYPIPAGTDLF", true}, {"GREY-TQR*RDVSYLFHDLYNGM &", "GREY-TQR*RDVSYLFHDLYNGM", true}, {"81% to CYP51G1 Arabidopsis", "", false}}
	for _, tt := range tests {
		got, ok := nelumboSequenceLine(tt.line)
		if got != tt.want || ok != tt.ok {
			t.Errorf("line %q=(%q,%v), want (%q,%v)", tt.line, got, ok, tt.want, tt.ok)
		}
	}
}

func TestParseReviewedNelumboSources(t *testing.T) {
	root := filepath.Join("..", "..")
	annotations, _, err := reviewWorkbook(filepath.Join(root, "raw", workbookFile))
	if err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(root, "raw", "plants-Lotus.P450.set.txt"))
	if err != nil {
		t.Fatal(err)
	}
	records, excluded, err := parseNelumboText(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if len(annotations) != 372 || len(records) != 364 || countSequences(records) != 364 || len(excluded) != 2 {
		t.Fatalf("annotations=%d records=%d sequences=%d excluded=%d", len(annotations), len(records), countSequences(records), len(excluded))
	}
	if excluded[0].Block != 276 || excluded[1].Block != 277 || !strings.Contains(excluded[0].Header, "Aquilegia") || !strings.Contains(excluded[1].Header, "Aquilegia") {
		t.Fatalf("excluded=%#v", excluded)
	}
	checks := []struct {
		index, block, length   int
		symbol, prefix, suffix string
	}{{0, 1, 487, "CYP51G1a", "MDLKENKF", "KGKVMVRYKRRRLSVE"}, {len(records) / 2, 183, 101, "CYP89A84P", "LTKKEIVS", "SHEVTEDVTIDGYLVP"}, {len(records) - 1, 366, 501, "CYP736A100P", "VVVIVSVL", "LLVIPTFRLKNNFGSF"}}
	for _, c := range checks {
		r := records[c.index]
		if r.Block != c.block || r.Symbol != c.symbol || len(r.Sequence) != c.length || !strings.HasPrefix(r.Sequence, c.prefix) || !strings.HasSuffix(r.Sequence, c.suffix) {
			t.Errorf("record[%d] block=%d symbol=%s len=%d prefix/suffix=%s/%s", c.index, r.Block, r.Symbol, len(r.Sequence), r.Sequence[:min(12, len(r.Sequence))], r.Sequence[max(0, len(r.Sequence)-14):])
		}
	}
}
