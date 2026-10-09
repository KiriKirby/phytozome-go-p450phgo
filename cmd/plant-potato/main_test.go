package main

import (
	"strings"
	"testing"
)

func TestPotatoSequenceLine(t *testing.T) {
	tests := []struct {
		line string
		want string
	}{
		{"348444 MELGDNKILNVGLLLVATLLVAKLISALIMPR 348333", "MELGDNKILNVGLLLVATLLVAKLISALIMPR"},
		{"VQDNIYSNRPKTVAISYLTYDRADMAFADYGPFWRQMRKLCVMKLFSRKRAESWDSVRD (0)", "VQDNIYSNRPKTVAISYLTYDRADMAFADYGPFWRQMRKLCVMKLFSRKRAESWDSVRD"},
		{"KSVVNEEDIQNLPYFKAVIKETFRLYPPV & 159946", "KSVVNEEDIQNLPYFKAVIKETFRLYPPV"},
		{"QIAGIESTSLXQQFMPEFFNLXLX TLSLPINLPNTNYYRGFQARKIL", "QIAGIESTSLXQQFMPEFFNLXLXTLSLPINLPNTNYYRGFQARKIL"},
		{"Query  295  EIFPAGTGTLTSTIEWAMAELVRNKEVMKKLNSELQN  353", ""},
		{"L++NIDYKGQDFEFLPFGAGRRMCPGLPFATKQ+HLILAYLVYHFEWS", ""},
	}
	for _, tt := range tests {
		got, _ := potatoSequenceLine(tt.line)
		if got != tt.want {
			t.Errorf("potatoSequenceLine(%q)=%q want %q", tt.line, got, tt.want)
		}
	}
}

func TestParsePotatoBlocksPreservesDuplicatesAndExcludesComparisons(t *testing.T) {
	text := strings.Join([]string{
		">CYP51G1", "note", "1 MPEPTIDEACDEFGHIKLMNPQRSTVWY* 99",
		">CYP51G1", "coordinates only",
		">CYP80N1 ortholog from eggplant", "MPEPTIDEACDEFGHIKLMNPQRSTVWY*",
		">SGN-U270131 Solanum tuberosum (potato) [3 ESTs aligned]", "1 VAFPIFLIFLLSGKRNLPPGPIGLPFIGNLHQY 174",
	}, "\n")
	records, excluded := parsePotatoText(text)
	if len(records) != 3 || len(excluded) != 1 {
		t.Fatalf("records=%d excluded=%d", len(records), len(excluded))
	}
	if records[0].RecordKey == records[1].RecordKey || records[0].ID != records[1].ID {
		t.Fatalf("duplicate identity handling: %#v %#v", records[0], records[1])
	}
	if records[2].ID != "SGN-U270131" || records[2].Sequence == "" {
		t.Fatalf("EST record=%#v", records[2])
	}
}
