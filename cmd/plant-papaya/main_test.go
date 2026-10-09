package main

import (
	"strings"
	"testing"
)

func TestPapayaSequenceLineUsesOnlyReviewedForms(t *testing.T) {
	tests := []struct {
		id   string
		line int
		raw  string
		want string
	}{
		{"CYP51G1", 20, "9521  MDVXXKFFNAXFLLVATLLVAKLISALIIPRS  9342", "MDVXXKFFNAXFLLVATLLVAKLISALIIPRS"},
		{"CYP72A73", 388, "SRTAFGSSHEEGKRIFQLMDELAILIQQLVQNVYIPGWR ()", "SRTAFGSSHEEGKRIFQLMDELAILIQQLVQNVYIPGWR"},
		{"CYP71AN8P", 298, "CFKQTSREFDAFLVQVIKKHQTNDDELTDCRNNFVHAVLQLQ*TNSL--DFKLTQHKMKAILLH", "CFKQTSREFDAFLVQVIKKHQTNDDELTDCRNNFVHAVLQLQ*TNSL--DFKLTQHKMKAILLH"},
		{"CYP87A10", 1437, "GVEINGASKNFMAFGGGMRFCVGTDFTKVQMAVFLHCLVTKYR (20", "GVEINGASKNFMAFGGGMRFCVGTDFTKVQMAVFLHCLVTKYR"},
		{"CYP715A7", 2591, "313893 TLLVLAMHPEWQEQLREEIRQVVGEKEVDATMLARL314009", "TLLVLAMHPEWQEQLREEIRQVVGEKEVDATMLARL"},
		{"CYP727A8", 2770, "29201 SHLFTKEEPCGNIMVVMFHGCLTTAGLIGNILARLATHPEIQDS  ()29332", "SHLFTKEEPCGNIMVVMFHGCLTTAGLIGNILARLATHPEIQDS"},
		{"CYP71B56P", 184, "xxxxxxxxxxxxxxxxxx", "XXXXXXXXXXXXXXXXXX"},
		{"CYP73A80", 609, "AF368378", ""},
		{"CYP93A12P", 1795, "first part", ""},
		{"CYP712A10P", 2544, "GS_ORF_8_from_ supercontig_234:168237..173712 (+ strand)", ""},
	}
	for _, tt := range tests {
		got, _ := papayaSequenceLine(tt.id, tt.line, tt.raw)
		if got != tt.want {
			t.Errorf("papayaSequenceLine(%q, %d, %q)=%q want %q", tt.id, tt.line, tt.raw, got, tt.want)
		}
	}
}

func TestParsePapayaBlocksKeepsDuplicatesAndExcludesVitis(t *testing.T) {
	text := strings.Join([]string{
		">CYP51G1", "annotation", "MPEPTIDEACDEFGHIKLMNPQRSTVWY*",
		">CYP51G1", "partial fragment", "MPEPTIDE*WITH*STAPS*",
		">CYP51G1 AM475390.2 Vitis vinifera", "MPEPTIDEACDEFGHIKLMNPQRSTVWY*",
	}, "\n")
	records, excluded, err := parsePapayaText(text)
	if err != nil {
		t.Fatal(err)
	}
	if len(records) != 2 || len(excluded) != 1 {
		t.Fatalf("records=%d excluded=%d", len(records), len(excluded))
	}
	if records[0].RecordKey == records[1].RecordKey || records[0].ID != records[1].ID {
		t.Fatalf("duplicate identity handling: %#v %#v", records[0], records[1])
	}
	if records[0].Sequence != "MPEPTIDEACDEFGHIKLMNPQRSTVWY" || records[1].Sequence != "MPEPTIDE*WITH*STAPS" {
		t.Fatalf("literal sequences=%q/%q", records[0].Sequence, records[1].Sequence)
	}
}

func TestParsePapayaRejectsUnknownInterstitialText(t *testing.T) {
	_, _, err := parsePapayaText(">CYP51G1\nMPEPTIDE\nunreviewed annotation\nMOREPEPTIDE*\n")
	if err == nil {
		t.Fatal("expected interruption error")
	}
}
