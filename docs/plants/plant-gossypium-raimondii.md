# Plant resource review: Gossypium raimondii

- Source file: `plants-Gossypium.raimondii.xlsx`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Gossypium.raimondii.xlsx
- Source SHA-256: `ccd3b612830c93df4410ae5a07cef81c799d43044cacb19328e1b89027c6a977`
- Workbook sheet: `Sorted by CYP name`
- Used range: `A1:L462`
- Hidden rows: `[]`
- Hidden columns: `none`
- Merged cells: `none`
- Native table parts: `none`
- Named CYP rows accepted: `449`
- Accepted rows with literal sequence: `449`
- Unnamed candidate rows excluded: `12`
- Duplicate literal-sequence groups: `2`
- Review status: `complete`

## Resource-specific interpretation

This workbook has one sheet named `Sorted by CYP name`. For this file, H=`seq. ID`, I=`best hit`, J=`%ID`, K=`CYP name`, and blank-headed L is the literal protein sequence. Rows 2-450 are named and accepted. Rows 451-462 have empty I/J/K classification cells and are excluded unnamed candidates. A pseudogene flag requires a CYP token ending in digit-plus-P. Literal O, X, and gap characters are retained. Row 161 contains the source's 1,520-aa literal value and is retained intact with an unusually-long flag; it is not truncated. Two groups of identical literal sequence remain separate because the source IDs differ. No sequence is repaired or deduplicated.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X | 4 |
| nonstandard-O | 37 |
| short-sequence | 18 |
| source-gap | 2 |
| source-pseudogene-label | 3 |
| unusually-long-sequence | 1 |

## Boundary and representative rows

| Source row | seq. ID | Best hit | CYP name | Sequence | Status |
|---:|---|---|---|---:|---|
| 2 | GossraimChr05.57164 | CYP51G1 | CYP51G | 486 aa |  |
| 55 | GossraimChr13.47505 | CYP71BE20 | CYP71 | 201 aa | source-gap;short-sequence |
| 161 | GossraimChr12.33569 | CYP82C49 | CYP82C | 1520 aa | nonstandard-O;unusually-long-sequence |
| 226 | GossraimChr08.52938 | CYP87A37 | CYP87A | 474 aa |  |
| 450 | GossraimChr13.18698 | CYP749A9 | CYP749A | 519 aa |  |
| 451 | GossraimChr02.8412 | *(empty)* | *(empty)* | 409 aa (excluded) | excluded: CYP name and best hit columns are empty |
| 452 | GossraimChr05.2972 | *(empty)* | *(empty)* | 399 aa (excluded) | excluded: CYP name and best hit columns are empty |
| 453 | GossraimChr05.3004 | *(empty)* | *(empty)* | 368 aa (excluded) | excluded: CYP name and best hit columns are empty |
| 454 | GossraimChr08.56016 | *(empty)* | *(empty)* | 375 aa (excluded) | excluded: CYP name and best hit columns are empty |
| 455 | GossraimChr09.15117 | *(empty)* | *(empty)* | 349 aa (excluded) | excluded: CYP name and best hit columns are empty |
| 456 | GossraimChr09.17807 | *(empty)* | *(empty)* | 432 aa (excluded) | excluded: CYP name and best hit columns are empty |
| 457 | GossraimChr09.24343 | *(empty)* | *(empty)* | 328 aa (excluded) | excluded: CYP name and best hit columns are empty |
| 458 | GossraimChr10.6396 | *(empty)* | *(empty)* | 380 aa (excluded) | excluded: CYP name and best hit columns are empty |
| 459 | GossraimChr13.16735 | *(empty)* | *(empty)* | 327 aa (excluded) | excluded: CYP name and best hit columns are empty |
| 460 | GossraimChr13.3544 | *(empty)* | *(empty)* | 404 aa (excluded) | excluded: CYP name and best hit columns are empty |
| 461 | GossraimChr13.38434 | *(empty)* | *(empty)* | 313 aa (excluded) | excluded: CYP name and best hit columns are empty |
| 462 | GossraimChr13.4567 | *(empty)* | *(empty)* | 314 aa (excluded) | excluded: CYP name and best hit columns are empty |
