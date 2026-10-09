# Plant resource review: Setaria italica

- Source file: `plants-Setaria.italica.xlsx`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Setaria.italica.xlsx
- Source SHA-256: `2fae1256d5cfcbbe9bdc9b1fa41c6ba10b7632447dd40b51d637ee7d6c3673c6`
- Workbook sheet: `Sorted by CYP name`
- Used range: `A1:K419`
- Hidden rows: `[]`
- Hidden columns: `[]`
- Merged cells: `[]`
- Native table parts: `0`
- Accepted named rows: `413`
- Accepted rows with literal sequence: `413`
- Excluded rows: `5`
- Review status: `complete`

## Resource-specific interpretation

The sole `Sorted by CYP name` sheet has exact used range `A1:K419`, no hidden rows or columns, merged cells, or native tables. G/J/K are the literal sequence ID, assigned CYP, and headerless protein columns. Rows 2-414 are the continuous assigned region. Rows 415-419 carry model text in A/G and out-of-frame translations containing O/X in K, but H/I/J are empty; they are excluded. The extra locus/strand text in A is checked to start with G rather than being mistaken for the ID itself. Literal O/X/gaps, short sequences, pseudogene hits and three duplicate-sequence groups remain row-distinct.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X-or-x | 3 |
| nonstandard-O | 12 |
| short-sequence | 5 |
| source-best-hit-pseudogene-label | 9 |
| source-gap | 1 |

## Representative and excluded rows

| Row | Seq ID | Best hit | Assigned CYP | Sequence | Status |
|---:|---|---|---|---:|---|
| 2 | Setaital8.26299 | CYP51G1 | CYP51G1 | 488 aa |  |
| 208 | Setaital9.4984 | CYP81A27P | CYP81A | 523 aa | source-best-hit-pseudogene-label |
| 414 | Setaital6.27518 | CYP735A30 | CYP735A | 530 aa |  |
| 415 | Setaital2.21776 Setaital2 - |  |  | 384 chars (excluded) | excluded: no best hit, percent identity, or assigned CYP name; K is an out-of-frame translation |
| 416 | Setaital424.1 Setaital424 - |  |  | 355 chars (excluded) | excluded: no best hit, percent identity, or assigned CYP name; K is an out-of-frame translation |
| 417 | Setaital5.46687 Setaital5 + |  |  | 365 chars (excluded) | excluded: no best hit, percent identity, or assigned CYP name; K is an out-of-frame translation |
| 418 | Setaital8.26320 Setaital8 + |  |  | 424 chars (excluded) | excluded: no best hit, percent identity, or assigned CYP name; K is an out-of-frame translation |
| 419 | Setaital8.38960 Setaital8 - |  |  | 468 chars (excluded) | excluded: no best hit, percent identity, or assigned CYP name; K is an out-of-frame translation |
