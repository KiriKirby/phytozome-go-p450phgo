# Plant resource review: Nelumbo nucifera

- Annotation workbook: `plants-Lotus.P450s.Oct31.2012.xlsx`
- Workbook URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Lotus.P450s.Oct31.2012.xlsx
- Workbook SHA-256: `7a79841a8c66473596bf5ea1c2d7e3719304eda72e5408ebdafa3b51794434ba`
- Sequence Word file: `plants-Lotus.P450.set.doc`
- Sequence URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Lotus.P450.set.doc
- Sequence SHA-256: `99fbd1983129e64c2dd6e21b04197fe6c9dc22a43cb6cbdb95631394324994c0`
- Workbook sheets: `CYPs in name order` (`A1:M384`), `CYPs in scaffold order` (`A1:M376`), `Gene Pairs by WGD` (`A1:I41`)
- Workbook annotation/model rows: `372` (172 gene rows, 179 primary pseudogene rows, 17 merged-model continuation rows)
- Workbook curated-sequence summary: `355 sequences, 175 genes, 180 pseudogenes`
- Word container inspected: legacy Word, 2,948 paragraphs, 77 pages, no tables
- Word `>` blocks: `366`
- Accepted Nelumbo literal sequence blocks: `364`
- Temporary Aquilegia blocks excluded: `2`
- Duplicate literal-sequence groups retained: `3`
- Review status: `complete`

## Resource-specific interpretation

The workbook and Word file are two parts of one Nelumbo release. The workbook contains curated names, coordinates, gene/pseudogene calls, merged-model notes, scaffold ordering, and WGD pairs, but no amino-acid sequence column. It therefore acts as the annotation ledger and does not independently create empty-sequence PGD records. Its 372 model rows resolve, per its own footer, to 355 curated sequences because 17 rows are continuation models in merged pseudogenes.

The Word file supplies literal protein sequences. Its preface says that, beyond the 355 sequence set, nine redundant contig sequences and two temporary `Lotus japonicus` sequences are present. The only two explicit foreign blocks are instead labelled `Aquilegia` in their headers (blocks 276-277); this source inconsistency is preserved in the audit, and those two non-Nelumbo blocks are excluded. The remaining 364 Nelumbo blocks are retained in source order, including the nine redundant contig representations.

Protein lines use several reviewed layouts: whole proteins with spaces, scaffold fragments with one or two leading coordinates, trailing coordinates, phase markers, source `&` joins, X/x, gaps and internal stops. Layout metadata is removed while literal uncertainty remains; only one terminal `*` is removed. No sequence is translated, repaired or externally completed.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X-or-x | 23 |
| internal-stop | 94 |
| short-sequence | 161 |
| source-fragment-or-missing-region | 44 |
| source-gap | 1 |
| source-gap-annotation | 30 |
| source-joined-piece | 15 |
| source-pseudogene | 204 |
| source-terminal-stop | 70 |

## Representative records

| Accepted index | Source block | Source line | ID | CYP symbol | Sequence | Status |
|---:|---:|---:|---|---|---:|---|
| 1 | 1 | 26 | maker-scaffold_5-snap-gene-84.25-mRNA-1 | CYP51G1a | 487 aa | source-terminal-stop |
| 183 | 183 | 1402 | CYP89A84P | CYP89A84P | 101 aa | source-pseudogene;internal-stop;short-sequence |
| 364 | 366 | 2940 | augustus_masked-scaffold_1-processed-gene-119.5-mRNA-1 | CYP736A100P | 501 aa | source-terminal-stop;source-pseudogene;internal-stop |

## Excluded blocks

| Source block | Source line | Header | Reason |
|---:|---:|---|---|
| 276 | 2154 | CYP96P9 DR944473.1 Aquilegia | source preface explicitly says two temporary Aquilegia sequences are included |
| 277 | 2165 | CYP96P10 DR932492 Aquilegia | source preface explicitly says two temporary Aquilegia sequences are included |
