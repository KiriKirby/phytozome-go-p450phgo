package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestOstreococcus(t *testing.T) {
	d, err := os.ReadFile(filepath.Join("..", "..", "raw", "plants-ostreococcus.txt"))
	if err != nil {
		t.Fatal(err)
	}
	r, err := parse(string(d))
	if err != nil {
		t.Fatal(err)
	}
	if len(r) != 30 {
		t.Fatal(len(r))
	}
	if r[0].Species != "Ostreococcus tauri" || r[1].Species != "Ostreococcus lucimarinus" || r[21].Species != "Ostreococcus RCC809" {
		t.Fatalf("species split: %q %q %q", r[0].Species, r[1].Species, r[21].Species)
	}
	if r[0].Symbol != "CYP51G1" || !strings.HasPrefix(r[0].Sequence, "MTIILAIF") || !strings.HasSuffix(r[0].Sequence, "KPLGK") {
		t.Fatalf("first %+v", r[0])
	}
	if !strings.Contains(r[9].Sequence, "X") || !strings.Contains(r[9].Status, "ambiguous-X-or-x") {
		t.Fatalf("X record %+v", r[9])
	}
	if r[2].Symbol != "CYP97A15" || r[3].Symbol != "CYP97A15" || r[4].Symbol != "CYP97A15" || r[3].ID == r[4].ID {
		t.Fatalf("alternate CYP97A15 models not retained")
	}
	if r[29].Symbol != "CYP97C18" || !strings.HasSuffix(r[29].Sequence, "ATVKPRAMREAPVSSAGKCPMGH") {
		t.Fatalf("last %+v", r[29])
	}
}
