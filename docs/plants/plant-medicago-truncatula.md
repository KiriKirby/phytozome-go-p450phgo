# Plant resource review: Medicago truncatula

- Source file: `plants-medicago.seqs.doc`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/medicago.seqs.doc
- Source SHA-256: `0d4967190bb35bf639018a90473e402d5ae25b09ea5d8fa3d5944aec904e0199`
- Container inspected: legacy Word document, 3,950 paragraphs, no tables, 64 pages
- Normalized split lines: `3951`
- Source `>` sequence pieces: 376
- Accepted Medicago P450/candidate-P450 pieces: 349
- Accepted pieces with literal protein sequence: 349
- Explicit other-species assembly/comparison pieces excluded: 24
- Explicit false-positive pieces excluded: 3
- Duplicate exact-header groups retained: 4
- Duplicate literal-sequence groups retained: 9
- Review status: `complete`

## Resource-specific interpretation

The preamble calls the entries `376 sequence pieces`, warns that some are duplicates, and says some other-species pieces are present for gene assembly. Each Medicago header-delimited piece is therefore retained independently in source order; accession-only genomic/cDNA pieces are not merged into adjacent named CYPs, and duplicate or alternate pieces are not collapsed. Twenty-four blocks explicitly belonging to Medicago sativa/alfalfa, Glycine max/soybean, Lotus japonicus, Stevia rebaudiana, Vitis vinifera, or Zinnia elegans are comparison/assembly helpers and are excluded from the Medicago species. The three blocks below the literal `False positives` heading are excluded. All remaining 349 blocks contain literal protein. Coordinates, standalone phase markers, whitespace, and wrapping quotes are document layout, not residues. The ABE79358 block has its protein after the exact same-line `[Medicago truncatula]` marker and is handled only by that reviewed layout. Internal stops, X, question marks, short pieces, pseudogenes, and fragments remain literal; only explicit terminal stops are removed. Nothing is translated, repaired, completed, or fetched elsewhere.

## Status counts

| Status | Records |
|---|---:|
| internal-stop | 37 |
| short-sequence | 180 |
| source-accession-piece | 135 |
| source-fragment-or-missing-region | 29 |
| source-pseudogene | 41 |
| source-terminal-stop | 72 |

## Excluded source pieces

| Block | Header prefix | Reason |
|---:|---|---|
| 64 | >CYP73A3 Medicago sativa | foreign Medicago sativa comparison |
| 69 | >CYP76E1 From Chris Steele | foreign Medicago sativa/alfalfa comparison |
| 72 | >CYP76E2 TC107627 | foreign Medicago sativa assembly helper |
| 73 | >76F5 C Steele | foreign Medicago sativa/alfalfa comparison |
| 76 | >76F6 Length | foreign Medicago sativa/alfalfa comparison |
| 128 | >CYP82D1 C Steele | foreign Medicago sativa/alfalfa comparison |
| 132 | >CYP83E1 C Steele | foreign Medicago sativa/alfalfa comparison |
| 151 | >CYP84A19 CYP84Ms1 | foreign Medicago sativa comparison |
| 152 | >CYP84A20 CYP84Ms2 | foreign Medicago sativa comparison |
| 179 | >CYP93C7v1 AF195801 | foreign Medicago sativa comparison |
| 180 | >CYP93C8 AF195800 | foreign Medicago sativa comparison |
| 181 | >CYP93C AF195802 | foreign Medicago sativa comparison |
| 182 | >CYP93C5 Glycine max | foreign Glycine max comparison |
| 192 | >CYP706A11 AP006082.1 | foreign Lotus japonicus comparison |
| 232 | >AJ410089.1 Medicago sativa | foreign Medicago sativa assembly helper |
| 255 | >CYP74B4v1  Medicago sativa | foreign Medicago sativa comparison |
| 278 | >CYP707A16 Glycine max | foreign Glycine max comparison |
| 287 | >CYP716D4 Stevia rebaudiana | foreign Stevia rebaudiana comparison |
| 294 | >CYP724 DT014285.1 Vitis vinifera | foreign Vitis vinifera comparison |
| 297 | >CX704924.1 Glycine max | foreign Glycine max assembly helper |
| 300 | >Zinnia elegans CYP733 | foreign Zinnia elegans comparison |
| 301 | >CG816619.1 Glycine max | foreign Glycine max assembly helper |
| 372 | >Soybean CYP727 | foreign soybean comparison after no-member CYP727 note |
| 373 | >Lotus japonicus CYP727 | foreign Lotus japonicus comparison after no-member CYP727 note |
| 374 | >CR331796.1 | explicit false positive: ubiquitin ligase |
| 375 | >CG954156.1 | explicit false positive |
| 376 | >CG969307.1 | explicit false positive |

## Representative accepted pieces

| Accepted index | Source block | Source line | ID | Symbol | Sequence | Status | Header |
|---:|---:|---:|---|---|---:|---|---|
| 1 | 1 | 14 | CYP51G1 | CYP51G1 | 489 aa |  | >CYP51G1 DQ335779                1712 bp    mRNA    linear   PLN 04AUG2006 |
| 175 | 188 | 2118 | CR321277.1 |  | 195 aa | source-accession-piece;short-sequence | >CR321277.1 Medicago truncatula genomic |
| 349 | 371 | 3899 | CYP711A12 | CYP711A12 | 541 aa |  | >CYP711A12 CR956434.13  M.truncatula on chromosome 3 |
