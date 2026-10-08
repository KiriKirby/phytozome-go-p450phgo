package main

import "testing"

func TestReviewedCAld5HTableSpeciesRelationships(t *testing.T) {
	rows, err := readReviewedSources("../../sources")
	if err != nil {
		t.Fatal(err)
	}
	want := map[string][]string{"Oryza sativa": {"CYP84A6", "CYP84A7"}, "Arabidopsis thaliana": {"CYP84A1"}, "Liquidambar styraciflua": {"CYP84A3"}, "Populus trichocarpa": {"CYP84A10", "CYP84A11"}, "Eucalyptus globulus": {"CYP84A-like"}, "Medicago sativa": {"CYP84A20"}, "Setaria italica": {"CYP84A-like"}, "Zea mays": {"CYP84A33", "CYP84A34"}, "Sorghum bicolor": {"CYP84A-like1"}, "Brachypodium distachyon": {"CYP84A5"}, "Panicum virgatum": {"CYP84A-like1", "CYP84A-like2"}}
	got := map[string]map[string]bool{}
	for _, r := range rows {
		if got[r.Species] == nil {
			got[r.Species] = map[string]bool{}
		}
		got[r.Species][r.Symbol] = true
	}
	for species, symbols := range want {
		for _, symbol := range symbols {
			if !got[species][symbol] {
				t.Errorf("missing %s -> %s", species, symbol)
			}
		}
	}
}

func TestMergeRecordsDoesNotCrossAssignSpecies(t *testing.T) {
	rows := mergeRecords([]record{{Category: "plants", Species: "Oryza sativa", Symbol: "CYP84A6"}, {Category: "plants", Species: "Arabidopsis thaliana", Symbol: "CYP84A1"}})
	if len(rows) != 2 {
		t.Fatalf("rows=%d", len(rows))
	}
	if rows[0].Species == rows[1].Species {
		t.Fatal("species relationships collapsed")
	}
}
