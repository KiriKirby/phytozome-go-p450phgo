# Plant resource review: Capsicum annuum

The Dr. Nelson page/legacy manifest label was `Capsicum annum`; the workbook filename and accepted scientific spelling are `Capsicum annuum`, which is the canonical PGD species name.

- Source file: `plants-Capsicum.annuum.xlsx`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Capsicum.annuum.xlsx
- Source SHA-256: `918847325edeb98b388220f1c2517aa1177853a3562b1d0868fbd11ae3fbc294`
- Workbook sheet: `Sorted by CYP name`
- Used range: `A1:L650`
- Hidden rows: `[]`
- Merged cells: `none`
- Named CYP rows accepted: `617`
- Accepted rows with literal sequence: `617`
- Unnamed candidate rows excluded: `32`
- Review status: `complete`

## Resource-specific interpretation

This workbook has one sheet and no hidden rows, hidden columns, merged cells, or native table part. The header occupies row 1. For this workbook only, column H is `seq. ID`, column I is `best hit`, column J is `%ID`, column K is `CYP name`, and column L is the full source protein sequence even though its header cell is blank. Rows 2–618 have a named CYP family in K and are accepted. Rows 619–650 have an empty K value and remain excluded unnamed candidates.

The displayed CYP symbol is the exact I-column best hit; H remains the gene/sequence ID and row-level key. Literal `X`, `O`, and `-` characters are retained and flagged. No sequence character is repaired or removed.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X | 25 |
| nonstandard-O | 62 |
| short-sequence | 40 |
| source-gap | 10 |

## Boundary and representative rows

| Source row | seq. ID | Best hit | CYP name | Sequence | Status |
|---:|---|---|---|---:|---|
| 2 | Capsannu08.1090 | CYP51G1 | CYP51G1 | 487 aa |  |
| 310 | Capsannu04.10165 | CYP81B40 | CYP81B | 506 aa |  |
| 618 | Capsannu00.281985 | CYP94A48 | CYP | 372 aa |  |
| 619 | Capsannu00.144276 |  | *(empty)* | 386 aa (excluded) | excluded: CYP name column is empty |
| 650 | Capsannu12.216154 |  | *(empty)* | 361 aa (excluded) | excluded: CYP name column is empty |
