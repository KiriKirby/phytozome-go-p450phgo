package main

import (
	"context"
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

func main() {
	out := flag.String("out", "p450phgo.pgd", "output PGD path")
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	records, err := collect(ctx)
	if err != nil {
		panic(err)
	}
	if err := write(*out, records); err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d records to %s\n", len(records), *out)
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

func write(path string, records []record) error {
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
