# Plant resource review: Thellungiella parvula

- Source file: `plants-Thellungiella.parvula.xlsx`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Thellungiella.parvula.xlsx
- Source SHA-256: `1557c72a9df6ad83f4d28f6e38b5e0f76f32013bf55bb0d8255c1bc5794024e1`
- Workbook sheet: `Sorted by CYP name`
- Used range: `A1:L213`
- Hidden rows: `[]`
- Hidden columns: `none`
- Merged cells: `none`
- Native table parts: `none`
- Named CYP rows accepted: `207`
- Accepted rows with literal sequence: `207`
- Unnamed candidate rows excluded: `5`
- Duplicate literal-sequence groups retained as distinct IDs: `6`
- Review status: `complete`

## Resource-specific interpretation

This workbook has one sheet. The header occupies row 1. For this workbook only, column H is `seq ID`, I is `best hit`, J is `%ID`, K is `CYP name`, and the blank-headed L column is the literal protein sequence. Rows 2–208 have a CYP name in K and are accepted. Rows 209–213 have empty I, J, and K values and remain excluded unnamed candidates even though L contains sequence-like text.

H is the Thellungiella gene/sequence ID and row-level key. I is retained as the displayed CYP symbol; K retains the source family classification. The source pseudo/P labels are status metadata, not grounds for deleting the row. Literal `O`, `X`, and `-` are preserved. Six pairs have identical literal sequences but different gene IDs, so every source row remains distinct. No character is repaired or removed.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X | 2 |
| nonstandard-O | 9 |
| short-sequence | 4 |
| source-gap | 2 |
| source-pseudogene-label | 5 |

## Duplicate literal-sequence groups

These rows are retained separately because their source IDs differ.

| Source rows |
|---|
| 46, 47 |
| 80, 81 |
| 117, 118 |
| 122, 123 |
| 128, 129 |
| 176, 177 |

## Boundary and representative rows

| Source row | seq ID | Best hit | CYP name | Sequence | Status |
|---:|---|---|---|---:|---|
| 2 | Thelparv1-1.3647 | CYP51G1 | CYP51G1 | 488 aa |  |
| 105 | Thelparv3-6.460 | CYP84A25 | CYP84A | 519 aa |  |
| 208 | Thelparv5-6.70 | CYP735A2 | CYP735A | 517 aa |  |
| 40 | Thelparv5-1.4278 | CYP71B4 | CYP71B pseudo | 333 aa | source-pseudogene-label;nonstandard-O;short-sequence |
| 80 | ThelparvUn1490.1 | CYP81D22 | CYP81D | 302 aa | ambiguous-X;short-sequence |
| 179 | Thelparv1-1.11852 | CYP705A32 | CYP705A | 463 aa | source-gap |
| 209 | Thelparv5-6.1082 | *(empty)* | *(empty)* | 406 aa (excluded) | excluded: CYP name column is empty |
| 213 | ThelparvUn619.1 | *(empty)* | *(empty)* | 318 aa (excluded) | excluded: CYP name column is empty |
