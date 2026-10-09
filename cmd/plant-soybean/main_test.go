package main

import (
	"strings"
	"testing"
)

func TestParseSoybeanSeparatesProteinBlocksFromESTDNA(t *testing.T) {
	text := strings.Join([]string{
		"Glycine max (soybean) P450s",
		"All soybean P450 sequences",
		">CYP51G1 Glycine max",
		"annotation",
		"MEIDSRFLNTGLLLVATILVAKLISAFIVPKSRKRV*",
		">CYP71D103P Glycine max",
		"pseudogene Frameshift = &",
		"100 IYLQLGETTTIIVSSPECVKEI & 200",
		">CYP83E22P Glycine max",
		"MKK??VQEEI*DEDDVQ",
		"955 CYTOCHROME P450 EST FOR Glycine max.",
		">gi|1|gb|DNA",
		"ACGTACGTACGTACGTACGT",
	}, "\n")
	records, ests, _, err := parseText(text, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(records) != 3 || ests != 1 {
		t.Fatalf("records=%d ESTs=%d", len(records), ests)
	}
	if records[0].Sequence != "MEIDSRFLNTGLLLVATILVAKLISAFIVPKSRKRV" || strings.Join(records[0].Status, ";") != "source-terminal-stop;short-sequence" {
		t.Fatalf("terminal block=%#v", records[0])
	}
	if records[1].Sequence != "IYLQLGETTTIIVSSPECVKEI" || strings.Join(records[1].Status, ";") != "source-pseudogene;source-frameshift-marker;short-sequence" {
		t.Fatalf("frameshift block=%#v", records[1])
	}
	if records[2].Sequence != "MKK??VQEEI*DEDDVQ" || strings.Join(records[2].Status, ";") != "source-pseudogene;internal-stop;source-question-mark-residue;short-sequence" {
		t.Fatalf("uncertain block=%#v", records[2])
	}
}

func TestSoybeanSequenceLineRequiresWholeReviewedParagraph(t *testing.T) {
	if got, ok := sequenceLine("1266237 DHLTYRSAIAVEWAMSELLR 1266656"); !ok || got != "DHLTYRSAIAVEWAMSELLR" {
		t.Fatalf("coordinate line=%q %t", got, ok)
	}
	if got, ok := sequenceLine("VLVNAWAIGRDP (0?) &"); !ok || got != "VLVNAWAIGRDP" {
		t.Fatalf("phase line=%q %t", got, ok)
	}
	if _, ok := sequenceLine("955 CYTOCHROME P450 EST FOR Glycine max"); ok {
		t.Fatal("accepted prose as protein")
	}
	if _, ok := sequenceLine("ACGT NUCLEOTIDE appendix"); ok {
		t.Fatal("accepted mixed annotation as protein")
	}
}
