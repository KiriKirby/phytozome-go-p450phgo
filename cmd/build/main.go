package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	bolt "go.etcd.io/bbolt"
)

type record struct {
	ID, Category, Species, Symbol, Description, Sequence, SourceURL string
}
type speciesRecord struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Selectable  bool   `json:"selectable"`
	Description string `json:"description"`
}

var pages = map[string]string{
	"animals":  "https://drnelson.uthsc.edu/animals/",
	"plants":   "https://drnelson.uthsc.edu/plants/",
	"fungi":    "https://drnelson.uthsc.edu/fungal-genomes/",
	"bacteria": "https://drnelson.uthsc.edu/bacteria/",
}

var bracketSpeciesRE = regexp.MustCompile(`\[([^\]]{3,120})\]`)

func main() {
	out := flag.String("out", "p450phgo.pgd", "output PGD path")
	sources := flag.String("sources", "sources", "reviewed structured source directory")
	docs := flag.String("species-docs", "docs/species", "generated per-species audit documentation")
	extracted := flag.String("extracted", "extracted-v4", "Office-normalized resource text directory")
	resources := flag.String("resource-index", "sources/resources.csv", "resource manifest CSV")
	flag.Parse()
	_, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	// Resource text is retained for audit and parser development. Formal
	// records are emitted only by resource-specific parsers plus reviewed CSV.
	records, err := parseExtracted(*resources, *extracted)
	if err != nil {
		panic(err)
	}
	reviewed, err := readReviewedSources(*sources)
	if err != nil {
		panic(err)
	}
	records = mergeRecords(records, reviewed)
	allSpecies, err := readManifestSpecies(*resources)
	if err != nil {
		panic(err)
	}
	if err := write(*out, records, allSpecies); err != nil {
		panic(err)
	}
	if err := writeSpeciesDocs(*docs, records); err != nil {
		panic(err)
	}
	if err := writeResourceAuditDocs(*resources, filepath.Join(filepath.Dir(*docs), "resources")); err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d records to %s\n", len(records), *out)
}

func writeResourceAuditDocs(manifestPath, root string) error {
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return err
	}
	data = []byte(strings.TrimPrefix(string(data), "\ufeff"))
	rows, err := csv.NewReader(strings.NewReader(string(data))).ReadAll()
	if err != nil {
		return err
	}
	if len(rows) < 2 {
		return nil
	}
	h := map[string]int{}
	for i, v := range rows[0] {
		h[strings.ToLower(strings.TrimSpace(v))] = i
	}
	get := func(row []string, k string) string {
		i := h[k]
		if i < 0 || i >= len(row) {
			return ""
		}
		return strings.TrimSpace(row[i])
	}
	if err := os.MkdirAll(root, 0o755); err != nil {
		return err
	}
	for _, row := range rows[1:] {
		label, cat, local, source := get(row, "species_label"), get(row, "category"), get(row, "local_file"), get(row, "source_url")
		if label == "" {
			continue
		}
		slug := regexp.MustCompile(`[^A-Za-z0-9._-]+`).ReplaceAllString(strings.ToLower(cat+"-"+label), "-")
		slug = strings.Trim(slug, "-")
		body := fmt.Sprintf("# Resource audit: %s\n\n- Category: `%s`\n- Resource: `%s`\n- URL: %s\n- Parser status: `pending resource-specific review`\n- Formal PGD records: `not published`\n\nThis resource was downloaded and normalized with Office COM. It is intentionally excluded from the searchable PGD until its species relationship, identifier fields, sequence handling, and parser validation are documented. Aggregate resources require separate per-species extraction.\n", label, cat, local, source)
		if err := os.WriteFile(filepath.Join(root, slug+".md"), []byte(body), 0o644); err != nil {
			return err
		}
	}
	return nil
}

func readManifestSpecies(path string) ([]speciesRecord, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data = []byte(strings.TrimPrefix(string(data), "\ufeff"))
	rows, err := csv.NewReader(strings.NewReader(string(data))).ReadAll()
	if err != nil {
		return nil, err
	}
	if len(rows) < 2 {
		return nil, nil
	}
	h := map[string]int{}
	for i, v := range rows[0] {
		h[strings.ToLower(strings.TrimSpace(v))] = i
	}
	get := func(row []string, k string) string {
		i := h[k]
		if i < 0 || i >= len(row) {
			return ""
		}
		return strings.TrimSpace(row[i])
	}
	seen := map[string]bool{}
	out := []speciesRecord{}
	for _, row := range rows[1:] {
		name, cat := get(row, "species_label"), get(row, "category")
		if name == "" {
			continue
		}
		key := cat + "|" + name
		if seen[key] {
			continue
		}
		seen[key] = true
		out = append(out, speciesRecord{Name: name, Category: cat, Selectable: false, Description: "Listed by Dr Nelson resource; record requires resource-specific review."})
	}
	return out, nil
}

// parseExtracted parses only normalized Office text produced by
// scripts/extract-office.ps1. It deliberately never reads the original DOC/XLSX
// bytes or scans printable binary strings. Each record keeps the resource URL
// and the manifest label so the resulting PGD is auditable.
func parseExtracted(manifestPath, extractedDir string) ([]record, error) {
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, err
	}
	data = []byte(strings.TrimPrefix(string(data), "\ufeff"))
	r := csv.NewReader(strings.NewReader(string(data)))
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(rows) < 2 {
		return nil, nil
	}
	h := map[string]int{}
	for i, v := range rows[0] {
		h[strings.ToLower(strings.TrimSpace(v))] = i
	}
	get := func(row []string, key string) string {
		i := h[key]
		if i < 0 || i >= len(row) {
			return ""
		}
		return strings.TrimSpace(row[i])
	}
	var out []record
	for _, row := range rows[1:] {
		label, category, source, local := get(row, "species_label"), get(row, "category"), get(row, "source_url"), get(row, "local_file")
		if local == "" {
			continue
		}
		data, e := os.ReadFile(filepath.Join(extractedDir, strings.TrimSuffix(local, filepath.Ext(local))+".txt"))
		if e != nil {
			continue
		}
		seen := map[string]bool{}
		for _, line := range strings.Split(string(data), "\n") {
			lineSpecies := label
			if m := bracketSpeciesRE.FindStringSubmatch(line); len(m) == 2 {
				candidate := strings.TrimSpace(strings.ReplaceAll(m[1], "_", " "))
				if !strings.Contains(strings.ToLower(candidate), "predicted") && !strings.Contains(strings.ToLower(candidate), "cytochrome") {
					lineSpecies = candidate
				}
			}
			for _, name := range cypRE.FindAllString(line, -1) {
				name = strings.ToUpper(name)
				if seen[name] {
					continue
				}
				seen[name] = true
				seq := ""
				for _, candidate := range strings.Fields(line) {
					candidate = strings.Trim(candidate, ",;()[]")
					if len(candidate) >= 50 && strings.Trim(candidate, "ACDEFGHIKLMNPQRSTVWY") == "" {
						seq = candidate
						break
					}
				}
				out = append(out, record{ID: name, Category: category, Species: lineSpecies, Symbol: name, Sequence: seq, Description: "Parsed from normalized resource text; source file: " + local, SourceURL: source})
			}
		}
	}
	return out, nil
}

func aggregateResource(category, label string) bool {
	l := strings.ToLower(label)
	if category == "bacteria" {
		return true
	}
	for _, token := range []string{"public", "all named", "lepidoptera", "environmental", "taxonomic group", "partial collection"} {
		if strings.Contains(l, token) {
			return true
		}
	}
	return false
}

func readReviewedSources(root string) ([]record, error) {
	matches, err := filepath.Glob(filepath.Join(root, "*.csv"))
	if err != nil {
		return nil, err
	}
	var out []record
	for _, path := range matches {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		data = []byte(strings.TrimPrefix(string(data), "\ufeff"))
		rows, err := csv.NewReader(strings.NewReader(string(data))).ReadAll()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		if len(rows) < 2 {
			continue
		}
		header := map[string]int{}
		for i, v := range rows[0] {
			header[strings.ToLower(strings.TrimSpace(v))] = i
		}
		required := []string{"category", "species", "symbol", "source_url", "source_note"}
		valid := true
		for _, name := range required {
			if _, ok := header[name]; !ok {
				valid = false
				break
			}
		}
		if !valid {
			continue // non-record audit CSVs share this directory
		}
		for _, row := range rows[1:] {
			get := func(name string) string {
				i := header[name]
				if i >= len(row) {
					return ""
				}
				return strings.TrimSpace(row[i])
			}
			symbol := get("symbol")
			species := get("species")
			if symbol == "" || species == "" {
				continue
			}
			out = append(out, record{ID: first(get("id"), symbol), Category: get("category"), Species: species, Symbol: symbol, Description: get("source_note"), Sequence: get("sequence"), SourceURL: get("source_url")})
		}
	}
	return out, nil
}
func first(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}
func mergeRecords(groups ...[]record) []record {
	seen := map[string]int{}
	var out []record
	for _, group := range groups {
		for _, r := range group {
			k := strings.ToLower(strings.Join([]string{r.Category, r.Species, r.Symbol, r.ID, r.SourceURL}, "|"))
			if i, ok := seen[k]; ok {
				if out[i].Description == "" {
					out[i].Description = r.Description
				}
				if out[i].Sequence == "" {
					out[i].Sequence = r.Sequence
				}
				continue
			}
			seen[k] = len(out)
			out = append(out, r)
		}
	}
	sort.Slice(out, func(i, j int) bool {
		return strings.ToLower(out[i].Species+"|"+out[i].Symbol) < strings.ToLower(out[j].Species+"|"+out[j].Symbol)
	})
	return out
}
func writeSpeciesDocs(root string, records []record) error {
	if err := os.MkdirAll(root, 0o755); err != nil {
		return err
	}
	old, err := filepath.Glob(filepath.Join(root, "*.md"))
	if err != nil {
		return err
	}
	for _, path := range old {
		if err := os.Remove(path); err != nil {
			return err
		}
	}
	bySpecies := map[string][]record{}
	for _, r := range records {
		if strings.TrimSpace(r.Species) != "" {
			bySpecies[r.Species] = append(bySpecies[r.Species], r)
		}
	}
	for species, rows := range bySpecies {
		name := regexp.MustCompile(`[^A-Za-z0-9._-]+`).ReplaceAllString(strings.ToLower(species), "-")
		name = strings.Trim(name, "-")
		if name == "" {
			continue
		}
		var b strings.Builder
		fmt.Fprintf(&b, "# %s\n\nCategory: `%s`\n\nRecords: %d\n\n| CYP / ID | Description | Source |\n|---|---|---|\n", species, rows[0].Category, len(rows))
		for _, r := range rows {
			fmt.Fprintf(&b, "| %s | %s | %s |\n", md(r.Symbol), md(r.Description), md(r.SourceURL))
		}
		if err := os.WriteFile(filepath.Join(root, name+".md"), []byte(b.String()), 0o644); err != nil {
			return err
		}
	}
	return nil
}
func md(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(s), "|", "\\|"), "\n", " ")
}

func collect(ctx context.Context) ([]record, error) {
	client := &http.Client{Timeout: 60 * time.Second}
	linkRE := regexp.MustCompile(`(?is)<a[^>]+href=["']([^"']+)["'][^>]*>(.*?)</a>`)
	stripRE := regexp.MustCompile(`(?is)<[^>]+>`)
	seen := map[string]bool{}
	var links []record
	for category, page := range pages {
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, page, nil)
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(io.LimitReader(resp.Body, 16<<20))
		resp.Body.Close()
		if err != nil {
			return nil, err
		}
		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("%s: %s", page, resp.Status)
		}
		base, _ := url.Parse(page)
		for _, match := range linkRE.FindAllSubmatch(body, -1) {
			href := strings.TrimSpace(string(match[1]))
			label := strings.Join(strings.Fields(stripRE.ReplaceAllString(string(match[2]), " ")), " ")
			if href == "" || label == "" {
				continue
			}
			ref, err := base.Parse(href)
			if err != nil || ref.Host == "" {
				continue
			}
			key := category + "|" + ref.String() + "|" + label
			if seen[key] {
				continue
			}
			seen[key] = true
			id := filepath.Base(strings.TrimSuffix(ref.Path, "/"))
			if id == "." || id == "" {
				id = label
			}
			if ext := strings.ToLower(filepath.Ext(ref.Path)); ext == ".doc" || ext == ".xlsx" || ext == ".txt" || ext == ".fasta" || ext == ".fa" {
				links = append(links, record{ID: id, Category: category, Species: label, Symbol: label, Description: label, SourceURL: ref.String()})
			}
		}
	}
	records := fetchResources(ctx, client, links)
	sort.Slice(records, func(i, j int) bool {
		if records[i].Category != records[j].Category {
			return records[i].Category < records[j].Category
		}
		return records[i].ID < records[j].ID
	})
	return records, nil
}

var cypRE = regexp.MustCompile(`(?i)\bCYP[0-9]{1,4}[A-Z]{1,4}[0-9]{1,4}(?:\.[0-9]+)?\b`)

func fetchResources(ctx context.Context, client *http.Client, links []record) []record {
	workers := 12
	if len(links) < workers {
		workers = len(links)
	}
	jobs := make(chan record)
	out := make(chan []record, len(links))
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for link := range jobs {
				out <- fetchResource(ctx, client, link)
			}
		}()
	}
	go func() {
		for _, l := range links {
			jobs <- l
		}
		close(jobs)
		wg.Wait()
		close(out)
	}()
	var records []record
	for batch := range out {
		records = append(records, batch...)
	}
	return records
}
func fetchResource(ctx context.Context, client *http.Client, link record) []record {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, link.SourceURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		return []record{link}
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return []record{link}
	}
	data, err := io.ReadAll(io.LimitReader(resp.Body, 64<<20))
	if err != nil {
		return []record{link}
	}
	text := extractPrintable(data)
	names := cypRE.FindAllString(text, -1)
	if len(names) == 0 {
		return []record{link}
	}
	seen := map[string]bool{}
	out := make([]record, 0, len(names))
	for _, name := range names {
		key := strings.ToUpper(name)
		if seen[key] {
			continue
		}
		seen[key] = true
		out = append(out, record{ID: key, Category: link.Category, Species: link.Species, Symbol: key, Description: link.Description, SourceURL: link.SourceURL})
	}
	return out
}
func extractPrintable(data []byte) string {
	var b strings.Builder
	b.Grow(len(data))
	space := false
	for _, c := range data {
		if c >= 32 && c <= 126 {
			b.WriteByte(c)
			space = false
		} else if !space {
			b.WriteByte(' ')
			space = true
		}
	}
	return b.String()
}

func write(path string, records []record, manifestSpecies []speciesRecord) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil && filepath.Dir(path) != "." {
		return err
	}
	tmp := path + ".tmp"
	db, err := bolt.Open(tmp, 0o644, nil)
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b, e := tx.CreateBucket([]byte("records"))
		if e != nil {
			return e
		}
		for i, r := range records {
			v, e := json.Marshal(r)
			if e != nil {
				return e
			}
			if e = b.Put([]byte(fmt.Sprintf("%08d", i)), v); e != nil {
				return e
			}
		}
		speciesBucket, e := tx.CreateBucket([]byte("species"))
		if e != nil {
			return e
		}
		seen := map[string]bool{}
		index := 0
		for _, r := range records {
			key := r.Category + "|" + r.Species
			if strings.TrimSpace(r.Species) == "" || seen[key] {
				continue
			}
			seen[key] = true
			s := speciesRecord{Name: r.Species, Category: r.Category, Selectable: true, Description: r.Description}
			v, e := json.Marshal(s)
			if e != nil {
				return e
			}
			if e = speciesBucket.Put([]byte(fmt.Sprintf("%08d", index)), v); e != nil {
				return e
			}
			index++
		}
		for i := range manifestSpecies {
			for _, r := range records {
				if strings.EqualFold(r.Category, manifestSpecies[i].Category) && strings.EqualFold(r.Species, manifestSpecies[i].Name) {
					manifestSpecies[i].Selectable = true
					manifestSpecies[i].Description = "Parsed records available from the normalized Dr Nelson resource."
					break
				}
			}
		}
		for _, s := range manifestSpecies {
			key := s.Category + "|" + s.Name
			if seen[key] {
				continue
			}
			seen[key] = true
			v, e := json.Marshal(s)
			if e != nil {
				return e
			}
			if e = speciesBucket.Put([]byte(fmt.Sprintf("%08d", index)), v); e != nil {
				return e
			}
			index++
		}
		return nil
	})
	closeErr := db.Close()
	if err == nil {
		err = closeErr
	}
	if err != nil {
		_ = os.Remove(tmp)
		return err
	}
	return os.Rename(tmp, path)
}
