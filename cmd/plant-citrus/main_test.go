package main

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"testing"
)

func TestCitrusReviewedResourcesHavePairedSequenceSources(t *testing.T) {
	for _, name := range []string{"citrus-clementina-71", "citrus-clementina-other", "citrus-sinensis-71", "citrus-sinensis-other"} {
		expectedFile := map[string]string{
			"citrus-clementina-71":    "plants-C.clementina.71clan.htm",
			"citrus-clementina-other": "plants-C.clementina.other.htm",
			"citrus-sinensis-71":      "plants-C_sinensis.71clan.htm",
			"citrus-sinensis-other":   "plants-C_sinensis.other.htm",
		}[name]
		path := filepath.Join("..", "..", "sources", "reviewed", "plants", name+".csv")
		f, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}
		rows, err := csv.NewReader(f).ReadAll()
		f.Close()
		if err != nil {
			t.Fatal(err)
		}
		if len(rows) < 2 {
			t.Fatalf("%s has no reviewed rows", name)
		}
		if rows[1][7] == "" {
			t.Fatalf("%s lost its literal sequence", name)
		}
		if rows[1][5] == "" || rows[1][8] != expectedFile {
			t.Fatalf("%s provenance mismatch: %#v", name, rows[1])
		}
	}
}
