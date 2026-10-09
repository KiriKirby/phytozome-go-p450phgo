package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestVolvoxCarteri(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "..", "raw", "plants-volvox.txt"))
	if err != nil {
		t.Fatal(err)
	}
	records, err := parse(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if len(records) != 19 {
		t.Fatalf("records=%d", len(records))
	}
	wants := []struct {
		index, line int
		symbol      string
		prefix      string
		suffix      string
	}{
		{0, 8, "CYP51G1", "MADLTAEL", "CRVKYTRRKLL"},
		{9, 159, "CYP744D1", "MVGSSALA", "AVWLQLHSRNTAPIVAV"},
		{18, 377, "CYP772A1", "MFVTDLLA", "PPCDLRRLVGVKVPRKPCWVQLGRIA"},
	}
	for _, want := range wants {
		r := records[want.index]
		if r.Line != want.line || r.Symbol != want.symbol || !strings.HasPrefix(r.Sequence, want.prefix) || !strings.HasSuffix(r.Sequence, want.suffix) {
			t.Fatalf("record %d mismatch: %+v", want.index+1, r)
		}
	}
	for _, r := range records {
		if r.Sequence == "" || strings.HasSuffix(r.Sequence, "*") {
			t.Fatalf("invalid sequence: %+v", r)
		}
	}
}
