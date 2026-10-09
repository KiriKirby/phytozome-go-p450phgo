// Resource-specific importer for the legacy HTML sequence pages linked from
// the Dr. Nelson plant index.  The profile is deliberately explicit: each
// page has its own species, URL, expected record count, and source file.
package main

import (
	"crypto/sha256"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type profile struct {
	slug, species, sourceFile, sourceURL string
	expected                             int
}

var profiles = map[string]profile{
	"tomato":                  {"tomato", "tomato", "plants-tomato.P450s.fasta.htm", "https://drnelson.uthsc.edu/tomato.P450s.fasta.htm", 457},
	"arabidopsis":             {"arabidopsis", "Arabidopsis thaliana", "plants-Arabidopsis.Blast.file.html", "https://drnelson.uthsc.edu/Arabidopsis.Blast.file.html", 272},
	"populus":                 {"populus", "Populus trichocarpa", "plants-cottonwood566.htm", "https://drnelson.uthsc.edu/cottonwood566.htm", 566},
	"rice":                    {"rice", "Oryza sativa", "plants-rice.FASTA.dec30.html", "https://drnelson.uthsc.edu/rice.FASTA.dec30.html", 15},
	"citrus-clementina-71":    {"citrus-clementina-71", "Citrus clementina", "plants-C.clementina.71clan.htm", "https://drnelson.uthsc.edu/C.clementina.71clan.htm", 209},
	"citrus-clementina-other": {"citrus-clementina-other", "Citrus clementina", "plants-C.clementina.other.htm", "https://drnelson.uthsc.edu/C.clementina.other.htm", 99},
	"citrus-sinensis-71":      {"citrus-sinensis-71", "Citrus sinensis", "plants-C_sinensis.71clan.htm", "https://drnelson.uthsc.edu/C_sinensis.71clan.htm", 188},
	"citrus-sinensis-other":   {"citrus-sinensis-other", "Citrus sinensis", "plants-C_sinensis.other.htm", "https://drnelson.uthsc.edu/C_sinensis.other.htm", 85},
}
var headerRE = regexp.MustCompile(`(?m)>((?:CYP|cyp)[A-Za-z0-9_.-]*)[^\r\n<]*`)
var aaRE = regexp.MustCompile(`[A-Z*X-]+`)
var coordAA = regexp.MustCompile(`(?m)\b\d+\s+([A-Z][A-Z*X-]{4,})\s+\d+\b`)

func main() {
	name := flag.String("resource", "", "tomato, arabidopsis, or populus")
	in := flag.String("input", "", "source HTML")
	out := flag.String("out", "", "reviewed CSV")
	flag.Parse()
	p, ok := profiles[*name]
	if !ok {
		panic("unknown resource")
	}
	if *in == "" {
		*in = filepath.Join("raw", p.sourceFile)
	}
	if *out == "" {
		*out = filepath.Join("sources", "reviewed", "plants", p.slug+".csv")
	}
	data, err := os.ReadFile(*in)
	if err != nil {
		panic(err)
	}
	// These pages are old HTML/Office exports; ASCII is intentional because
	// sequence and CYP headers are ASCII and binary wrappers may be present.
	text := string(data)
	matches := headerRE.FindAllStringIndex(text, -1)
	if len(matches) != p.expected {
		panic(fmt.Sprintf("%s header count changed: %d", p.slug, len(matches)))
	}
	hash := fmt.Sprintf("%x", sha256.Sum256(data))
	rows := make([][]string, 0, len(matches))
	for i, m := range matches {
		h := headerRE.FindStringSubmatch(text[m[0]:m[1]])
		symbol := h[1]
		end := len(text)
		if i+1 < len(matches) {
			end = matches[i+1][0]
		}
		block := text[m[1]:end]
		// Strip markup/control text, then retain only literal amino-acid
		// lines.  A source terminal '*' is preserved as documented data.
		block = strings.ReplaceAll(block, "&nbsp;", " ")
		var seq strings.Builder
		if p.slug == "populus" {
			for _, mm := range coordAA.FindAllStringSubmatch(block, -1) {
				seq.WriteString(mm[1])
			}
		}
		for _, line := range strings.FieldsFunc(block, func(r rune) bool { return r == '\r' || r == '\n' || r == '<' || r == '>' }) {
			t := strings.TrimSpace(line)
			if t == "" {
				continue
			}
			if p.slug == "populus" {
				continue
			}
			if aaRE.MatchString(t) && aaRE.FindString(t) == t && !strings.ContainsAny(t, "abcdefghijklmnopqrstuvwxyz") {
				seq.WriteString(t)
			}
		}
		s := seq.String()
		s = strings.TrimSpace(s)
		status := ""
		upper := strings.ToUpper(symbol)
		if strings.HasSuffix(upper, "P") || strings.Contains(upper, "PSEUDO") || strings.Contains(upper, "FRAG") {
			status = "pseudogene-or-fragment-label"
		}
		if s == "" {
			status = "sequence-missing"
		}
		note := fmt.Sprintf("resource-specific HTML FASTA block %d; header=%s", i+1, symbol)
		rows = append(rows, []string{"plants", p.species, symbol, symbol, fmt.Sprintf("%s:block-%04d", p.slug, i+1), p.sourceURL, note, s, p.sourceFile, hash, fmt.Sprintf("block-%04d", i+1), status})
	}
	if err := os.MkdirAll(filepath.Dir(*out), 0755); err != nil {
		panic(err)
	}
	f, err := os.Create(*out)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	_ = w.Write([]string{"category", "species", "symbol", "id", "record_key", "source_url", "source_note", "sequence", "source_file", "source_sha256", "source_block", "review_status"})
	for _, r := range rows {
		_ = w.Write(r)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		panic(err)
	}
	sort.Strings([]string{})
	fmt.Printf("%s: %d records\n", p.slug, len(rows))
}
