# Plant resource review: Oryza brachyantha

- Source file: `plants-Oryza.brachyantha.xlsx`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Oryza.brachyantha.xlsx
- Source SHA-256: `cc1b3bc1a447e3e5d5e4961d019698503a4db80ff60ed960ba97098e7c72f59f`
- Workbook sheet: `sorted by CYP name`
- Used range: `A1:L321`
- Hidden rows: `[]`
- Hidden columns: `[]`
- Merged cells: `[]`
- Native table parts: `0`
- Accepted named rows: `316`
- Accepted rows with literal sequence: `316`
- Excluded rows: `4`
- Review status: `complete`

## Resource-specific interpretation

The workbook has exactly one worksheet, `sorted by CYP name`, with used range `A1:L321` and no hidden rows or columns, merged cells, or native tables. Row 1 names H=`Gotoh ID`, I=`Best hit`, J=`%ID`, and K=`CYP name`; L is the headerless literal protein column. Rows 2-317 form the continuous assigned region and A (after its source FASTA marker is removed) equals H. Rows 318-321 explicitly say `out of frame translation overlaps ...`, have J=`0` and an empty K assignment, and are excluded even though L contains out-of-frame letter strings. Literal O/X and source truncation remain unchanged; no row is repaired or supplemented.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X-or-x | 50 |
| nonstandard-O | 11 |
| short-sequence | 20 |
| source-best-hit-pseudogene-label | 7 |
| source-terminal-stop | 1 |

## Representative and excluded rows

| Row | Seq ID | Best hit | Assigned CYP | Sequence | Status |
|---:|---|---|---|---:|---|
| 2 | Oryzbrac006267462.1.2 | CYP51G1 | CYP51G1 | 489 aa |  |
| 160 | Oryzbrac006267384.1.12964 | CYP81P1 | CYP81P | 401 aa | ambiguous-X-or-x |
| 317 | Oryzbrac006267381.1.7194 | CYP735A4 | CYP735A | 477 aa |  |
| 318 | Oryzbrac006267373.1.18673 |  |  | 393 chars (excluded) | excluded: out-of-frame translation with no assigned CYP name |
| 319 | Oryzbrac006267375.1.15516 |  |  | 360 chars (excluded) | excluded: out-of-frame translation with no assigned CYP name |
| 320 | Oryzbrac006267377.1.8905 |  |  | 448 chars (excluded) | excluded: out-of-frame translation with no assigned CYP name |
| 321 | Oryzbrac006267383.1.9539 |  |  | 421 chars (excluded) | excluded: out-of-frame translation with no assigned CYP name |
