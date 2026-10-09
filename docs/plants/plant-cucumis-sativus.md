# Plant resource review: Cucumis sativus

- Source file: `plants-Cucumis.sativus.xlsx`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Cucumis.sativus.xlsx
- Source SHA-256: `a7e4a4182923d3b350b9ae75bd1322aa7ab899619d05885c5989b9233e07e7e3`
- Workbook sheet: `sorted by CYP name`
- Used range: `A1:J230`
- Hidden rows: `[]`
- Hidden columns: `[]`
- Merged cells: `[]`
- Native table parts: `0`
- Accepted rows: `229`
- Accepted rows with literal sequence: `229`
- Review status: `complete`

## Resource-specific interpretation

This workbook contains one sheet named `sorted by CYP name` with `A1:J230`. The reviewed fields are A=`Gotoh's seq ID`, H=`seq ID`, I=`CYP name`, and J=`sequence`; rows 2-230 are all accepted because every row has a CYP symbol and a non-empty literal sequence. A is retained as source provenance while H is the output ID. The one leading cell-space in row 147 is removed as surrounding cell whitespace. Explicit terminal `*` markers are removed and audited; internal `*`, gaps, O, and X are retained exactly. Rows are not deduplicated and no sequence is repaired.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X | 4 |
| internal-stop | 5 |
| nonstandard-O | 4 |
| short-sequence | 16 |
| source-gap | 3 |
| source-pseudogene-label | 27 |
| source-terminal-stop | 7 |

## Representative rows

| Row | Gotoh ID | Seq ID | CYP name | Sequence | Status |
|---:|---|---|---|---:|---|
| 2 | Cucusati00919.1591 | Cucusati00919.1591 | CYP51G1 | 487 aa |  |
| 116 | Cucusati03487.421 | Cucusati03487.421 | CYP87D19 | 475 aa |  |
| 230 | Cucusati02653.1189 | Cucusati02653.1189 | CYP749A47 | 520 aa |  |
