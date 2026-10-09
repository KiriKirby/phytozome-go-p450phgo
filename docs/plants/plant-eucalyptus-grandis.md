# Plant resource review: Eucalyptus grandis

- Source file: `plants-Eucalyptus.grandis.xlsx`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Eucalyptus.grandis.xlsx
- Source SHA-256: `3ca91184ea48b0a98e23112429e3935a537be372a68afa92479b6c32e641a557`
- Workbook sheet: `sorted by CYP name`
- Used range: `A1:L773`
- Hidden rows: `[]`
- Hidden columns: `none`
- Merged cells: `none`
- Native table parts: `none`
- Accepted CYP rows: `762`
- Accepted rows with literal sequence: `762`
- Non-CYP model/comparison rows excluded: `7`
- Missing worksheet row: `771` (blank)
- Legend rows excluded: `772-773`
- Review status: `complete`

## Resource-specific interpretation

This workbook uses A=`Gotoh's seq ID` (including a literal leading `>`), I=`blast ID`, J=`best hit`, K=`%ID`, and L=`sequence`. Rows 2-770 are data rows. A row is accepted only when J begins with `CYP`; seven rows whose J values are Eucalyptus model IDs are excluded. Worksheet row 771 is blank, and rows 772-773 are color-threshold legend text, not records. One or more explicit terminal `*` markers are removed and audited; internal `*`, O, uppercase X, lowercase x, and gaps are retained exactly and flagged. No sequence is otherwise repaired or deduplicated.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X-or-x | 21 |
| internal-stop | 12 |
| nonstandard-O | 52 |
| short-sequence | 53 |
| source-gap | 10 |
| source-pseudogene-label | 8 |

## Boundary and representative rows

| Row | Source ID | blast ID | Best hit | Sequence | Status |
|---:|---|---|---|---:|---|
| 2 | >Eucagran2.60235 | Eucagran2.60235 | CYP51G1 | 485 aa |  |
| 37 | >Eucagran6.28248 | Eucagran6.28248 | CYP71AT5P | 513 aa | source-pseudogene-label |
| 386 | >Eucagran4.20544 | Eucagran4.20544 | CYP88A50 | 491 aa |  |
| 763 | >Eucagran7.8672 | Eucagran7.8672rev | CYP74A16 | 138 aa | short-sequence |
| 764 | >Eucagran7.26860 | Eucagran7.26860 | CYP74A27 | 347 aa | short-sequence |
| 755 | >Eucagran3.75618 | Eucagran3.75618rev | Eucagran2.56063 | 339 aa (excluded) | excluded: non-CYP model/comparison row |
| 757 | >Eucagran1.32616 | Eucagran1.32616 | Eucagran1.32648 | 420 aa (excluded) | excluded: non-CYP model/comparison row |
| 758 | >Eucagran11.12675 | Eucagran11.12675 | Eucagran11.12879 | 485 aa (excluded) | excluded: non-CYP model/comparison row |
| 759 | >Eucagran3.30027 | Eucagran3.30027 | Eucagran2.2692 | 421 aa (excluded) | excluded: non-CYP model/comparison row |
| 768 | >Eucagran1.16165 | Eucagran1.16165 | Eucagran1.16038 | 313 aa (excluded) | excluded: non-CYP model/comparison row |
| 769 | >Eucagran1.26059 | Eucagran1.26059 | Eucagran4.33597 | 439 aa (excluded) | excluded: non-CYP model/comparison row |
| 770 | >Eucagran2.56543 | Eucagran2.56543 | Eucagran11.31130 | 407 aa (excluded) | excluded: non-CYP model/comparison row |
