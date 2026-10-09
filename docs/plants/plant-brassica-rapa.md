# Plant resource review: Brassica rapa

- Source file: `plants-Brassica.rapa.xlsx`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Brassica.rapa.xlsx
- Source SHA-256: `f1cde455d0748b832bf22c573e047441da8dee2d365d2d279030b3676837b1a1`
- Workbook sheet: `Sorted by CYP name`
- Used range: `A1:L383`
- Hidden rows: `[]`
- Hidden columns: `none`
- Merged cells: `none`
- Native table parts: `none`
- Named CYP rows accepted: `377`
- Accepted rows with literal sequence: `377`
- Unnamed candidate rows excluded: `5`
- Duplicate literal-sequence groups: `4`
- Review status: `complete`

## Resource-specific interpretation

This workbook has the sheet name `Sorted by CYP name`. Its exact data columns are H=`seq ID`, I=`best hit`, J=`%ID`, K=`CYP name`, and blank-headed L=literal protein sequence. Rows 2–378 are named and accepted; rows 379–383 have empty I/J/K classification cells and are excluded unnamed candidates. Source P suffixes, O, X, and gap characters remain literal and are flagged. Four groups of identical literal sequence remain separate because their IDs differ. No sequence is repaired or deduplicated.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X | 22 |
| nonstandard-O | 10 |
| short-sequence | 11 |
| source-gap | 3 |
| source-pseudogene-label | 18 |

## Boundary and representative rows

| Source row | seq ID | Best hit | CYP name | Sequence | Status |
|---:|---|---|---|---:|---|
| 2 | BrasrapaChr06.4366 | CYP51G1 | CYP51G | 488 aa |  |
| 30 | BrasrapaChr06.25739 | CYP71B18P | CYP71B | 326 aa | source-pseudogene-label;short-sequence |
| 190 | BrasrapaChr01.811 | CYP81H1 | CYP81H | 525 aa |  |
| 378 | BrasrapaChr02.9859 | CYP735A2 | CYP735A | 517 aa |  |
| 379 | Brasrapa24302.1 | *(empty)* | *(empty)* | 254 aa (excluded) | excluded: CYP name column is empty |
| 380 | BrasrapaChr04.3188 | *(empty)* | *(empty)* | 439 aa (excluded) | excluded: CYP name column is empty |
| 381 | BrasrapaChr06.22416 | *(empty)* | *(empty)* | 142 aa (excluded) | excluded: CYP name column is empty |
| 382 | BrasrapaChr06.25941 | *(empty)* | *(empty)* | 382 aa (excluded) | excluded: CYP name column is empty |
| 383 | BrasrapaChr07.25655 | *(empty)* | *(empty)* | 386 aa (excluded) | excluded: CYP name column is empty |
