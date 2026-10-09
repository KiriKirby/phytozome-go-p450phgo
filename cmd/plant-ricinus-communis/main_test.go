package main

import "testing"

func TestRicinusSequenceLineUsesExactReviewedForms(t *testing.T) {
	tests := []struct {
		line, want string
		ok         bool
	}{
		{"1622  PERFSAERQEDQLHKRNFLAFGAGAHQCL 1443", "PERFSAERQEDQLHKRNFLAFGAGAHQCL", true},
		{"VEDIIDVLLELEKSHREEFGAFQFSKDHIKAILM (0)", "VEDIIDVLLELEKSHREEFGAFQFSKDHIKAILM", true},
		{"xxxxxxxx", "xxxxxxxx", true},
		{"EXON 2", "", false},
		{"pseudogene", "", false},
		{"same seq as XP_002523235 (WHOLE SEQ)", "", false},
	}
	for _, tt := range tests {
		got, ok := ricinusSequenceLine(tt.line)
		if got != tt.want || ok != tt.ok {
			t.Errorf("ricinusSequenceLine(%q)=(%q,%v), want (%q,%v)", tt.line, got, ok, tt.want, tt.ok)
		}
	}
}
