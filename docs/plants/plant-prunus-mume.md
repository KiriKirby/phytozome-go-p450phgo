# Plant resource review: Prunus mume

- Source file: `plants-Prunus.mume.xlsx`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Prunus.mume.xlsx
- Source SHA-256: `6a9d91dfa1ef369dbe6b7c399126a7a7568b431f3a9e477bb3fc7a34001d96e9`
- Workbook sheet: `Sorted by CYP name`
- Used range: `A1:L288`
- Hidden rows: `[]`
- Hidden columns: `[]`
- Merged cells: `[]`
- Native table parts: `0`
- Accepted named rows: `282`
- Accepted rows with literal sequence: `282`
- Excluded rows: `5`
- Review status: `complete`

## Resource-specific interpretation

The reviewed workbook has one worksheet and no hidden rows or columns, merged cells, or native tables. Row 1 is the exact header, including the trailing space in the I-column heading `best hit `. Rows 2-283 are the complete assigned region: A is the Gotoh source ID, H is the sequence ID, I is the best hit, J is percent identity, K is the source-assigned CYP label, and L is the literal protein sequence. Rows 281-283 deliberately retain the broad source label `CYP` because the workbook supplies that assignment together with a specific best hit and percent identity. Rows 284-288 have IDs and sequence-like text but empty I/J/K assignments and extensive non-protein O/X content, so they are excluded as unnamed models. One gap, literal O/X residues, eleven short sequences, two digit-P best-hit labels, and four duplicate-sequence pairs remain exactly as supplied. No sequence is repaired or deduplicated.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X-or-x | 14 |
| nonstandard-O | 9 |
| short-sequence | 11 |
| source-best-hit-pseudogene-label | 2 |
| source-gap | 1 |

## Representative and excluded rows

| Row | Seq ID | Best hit | Assigned CYP | Sequence | Status |
|---:|---|---|---|---:|---|
| 2 | Prunmume007362126.1.236 | CYP51G1 | CYP51G1 | 486 aa |  |
| 140 | Prunmume007362363.1.141 | CYP89A18 | CYP89A | 515 aa | ambiguous-X-or-x |
| 283 | Prunmume007362430.1.83 | CYP706B1 | CYP | 341 aa | short-sequence |
| 284 | Prunmume007362030.1.1 |  |  | 429 chars (excluded) | excluded: no best hit, percent identity, or assigned CYP name |
| 285 | Prunmume007362084.1.353 |  |  | 395 chars (excluded) | excluded: no best hit, percent identity, or assigned CYP name |
| 286 | Prunmume007362123.1.547 |  |  | 413 chars (excluded) | excluded: no best hit, percent identity, or assigned CYP name |
| 287 | Prunmume007362123.1.749 |  |  | 417 chars (excluded) | excluded: no best hit, percent identity, or assigned CYP name |
| 288 | Prunmume007362434.1.111 |  |  | 337 chars (excluded) | excluded: no best hit, percent identity, or assigned CYP name |
