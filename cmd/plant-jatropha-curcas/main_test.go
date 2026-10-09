package main

import (
	"strings"
	"testing"
)

func TestJatrophaSequenceLineUsesReviewedLiteralForms(t *testing.T) {
	tests := []struct {
		line, want string
		ok         bool
	}{
		{"541 NIFLGGVDTGAIVLVWAMAELVRNP 29 &", "NIFLGGVDTGAIVLVWAMAELVRNP", true},
		{"(0) xxIAGGTDTTTVTVTWGLALLLNHPI (2)", "xxIAGGTDTTTVTVTWGLALLLNHPI", true},
		{"MKGENGNEGSLVTAQCGP***FLRKLCITEF", "MKGENGNEGSLVTAQCGP***FLRKLCITEF", true},
		{"COMPLETE", "", false},
		{"FINISHED", "", false},
		{"complete", "", false},
		{"No ESTs", "", false},
		{"86% to CYP51G1 Ricinus communis", "", false},
	}
	for _, tt := range tests {
		got, ok := jatrophaSequenceLine(tt.line)
		if ok != tt.ok || got != tt.want {
			t.Errorf("jatrophaSequenceLine(%q)=(%q,%v), want (%q,%v)", tt.line, got, ok, tt.want, tt.ok)
		}
	}
}

func TestParseJatrophaExactExclusionsAndSpecialCYP735Name(t *testing.T) {
	var b strings.Builder
	for i := 1; i <= 537; i++ {
		header := ">CYP71A JcTEST" + string(rune(0x400+i))
		if i == 1 {
			header = ">CYP735A22"
		} else if i == 2 {
			header = ">BABX01044566.1"
		} else if i == 3 {
			header = ">CYP90C (one sequence)"
		} else if i == 4 {
			header = ">CYP90D (one sequence plus one pseudogene)"
		} else if i >= 5 && i <= 36 {
			header = ">CYP71A1 Ricinus communis helper"
		} else if i == 37 {
			header = ">Populus trichocarpa CYP75A13"
		}
		b.WriteString(header + "\n")
		b.WriteString(strings.Repeat("A", 400) + "*\n")
		if i == 517 {
			b.WriteString("20 False positive hits (3.7%)\n")
		}
	}
	records, excluded, err := parseJatrophaText(b.String())
	if err != nil {
		t.Fatal(err)
	}
	// 537 minus 32 Ricinus, one Populus, three structural/name-only
	// headings, and the 20 blocks below the false-positive heading.
	if len(records) != 481 || len(excluded) != 56 {
		t.Fatalf("records=%d excluded=%d", len(records), len(excluded))
	}
	if records[0].ID != "BABX01044566.1" || records[0].Symbol != "CYP735A22" || len(records[0].Sequence) != 400 {
		t.Fatalf("special CYP735A22 record=%#v", records[0])
	}
	if records[0].Status[0] != "source-terminal-stop" {
		t.Fatalf("terminal stop status=%v", records[0].Status)
	}
}
