# Plant resource review: Panicum virgatum

- Source file: `plants-Panicum.virgatum.xlsx`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Panicum.virgatum.xlsx
- Source SHA-256: `138182c4d9a9f920d3830d34cc35c1b11bee035d49da7f8966492807113f91be`
- Workbook sheet: `sorted by CYP name`
- Used range: `A1:G810`
- Hidden rows: `[]`
- Hidden columns: `[]`
- Merged cells: `[]`
- Native table parts: `0`
- Accepted named rows: `806`
- Accepted rows with literal sequence: `806`
- Excluded rows: `0`
- Review status: `complete`

## Resource-specific interpretation

The workbook contains two full data sheets: `Sorted by length` (`A1:F807`) and `sorted by CYP name` (`A1:G810`). They are alternate orderings, not two releases, so only the CYP-name sheet contributes records. Rows 2-807 contain 806 assigned literal G-column sequences; rows 808-809 are blank and row 810 is the legend `green <55% identical to a named P450`. Every D length equals the literal G length. Cross-sheet multiset comparison matches 802 rows exactly and isolates four source revisions/differences: Panivirg326850.2 rev, Panivirg327867.1, Panivirg356643.1rev, and Panivirg23744.8 (the name sheet assigns CYP76N1P while the length sheet says `not a P450`). The CYP-name sheet is the reviewed release authority. Ninety X-bearing and 189 short sequences remain literal.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X-or-x | 90 |
| short-sequence | 189 |
| source-best-hit-pseudogene-label | 17 |

## Representative and excluded rows

| Row | Seq ID | Best hit | Assigned CYP | Sequence | Status |
|---:|---|---|---|---:|---|
| 2 | Panivirg24881.7 | CYP51G1 | CYP51G | 289 aa | short-sequence |
| 404 | Panivirg17853.7 | CYP87A15 | CYP87A | 497 aa |  |
| 807 | Panivirg146243.1 | CYP76F35 | CYP | 215 aa | ambiguous-X-or-x;short-sequence |
