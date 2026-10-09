# Plant resource review: Ricinus communis

- Source file: `plants-Ricinus.communis.doc`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Ricinus.communis.doc
- Source SHA-256: `6778128f2a9c5fafdf97cc10c0d4c00ec384a1fb1e5bacda19e6b9d0867926c0`
- Container inspected: legacy Word document, 3,733 paragraphs, no tables, 80 pages
- Normalized split lines: `3,733`
- Source `>` blocks: `264`
- Accepted Ricinus CYP blocks with literal sequence: `263`
- Explicit bacterial contamination blocks excluded: `1`
- Duplicate literal-sequence groups retained: `1`
- Review status: `complete`

## Resource-specific interpretation

The title page says `261 sequences`, after combining many RefSeq duplicates. The literal document contains 264 `>` blocks: 263 CYP-labeled Ricinus blocks and one final accession-only block under the explicit `Contamination` heading whose own annotations say its top BLAST hits are bacteria. The contamination block is excluded. All 263 CYP blocks are retained in source order, including alternate entries and the two CYP71B64 blocks whose literal protein sequences are identical. The source's 261 claim is recorded as a discrepancy; records are not silently merged to force that count.

Protein lines in this document may carry leading/trailing coordinates, phase markers, and `&` joins. Those reviewed markers are removed while literal residue text remains. Lowercase `x`, uppercase `X`, and internal stops are retained. Only a final `*` is removed. The exact all-uppercase text `EXON 2` is an annotation, not four residues, and is explicitly rejected. Other lowercase prose cannot become sequence. No fragment is translated, repaired, extended, or filled from another database.

The preferred record ID is the source's XP accession, then its EEF accession, otherwise the CYP symbol. Duplicate names and sequences remain independent through the source-block number in `RecordKey`.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X-or-x | 11 |
| internal-stop | 29 |
| short-sequence | 35 |
| source-fragment-or-missing-region | 48 |
| source-frameshift | 6 |
| source-gap-annotation | 21 |
| source-joined-piece | 51 |
| source-pseudogene | 48 |
| source-terminal-stop | 27 |

## Representative records

| Accepted index | Source block | Source line | ID | CYP symbol | Sequence | Status | Header |
|---:|---:|---:|---|---|---:|---|---|
| 1 | 1 | 23 | CYP51G1 | CYP51G1 | 486 aa |  | CYP51G1 Ricinus communis |
| 132 | 132 | 1877 | XP_002524040.1 | CYP89A98 | 516 aa |  | CYP89A98 gi\|255566104\|ref\|XP_002524040.1  Ricinus communis |
| 263 | 263 | 3704 | XP_002525551.1 | CYP727B23 | 537 aa |  | CYP727B23 gi\|255569165\|ref\|XP_002525551.1  Ricinus communis |

## Excluded block

| Source block | Source line | Header | Reason |
|---:|---:|---|---|
| 264 | 3720 | gi\|255594353\|ref\|XP_002536077.1  Ricinus communis | source section and annotations explicitly identify this block as bacterial contamination |
