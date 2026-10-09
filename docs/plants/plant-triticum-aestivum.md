# Plant resource review: Triticum aestivum

- Source file: `plants-Triticum.aestivum.xlsx`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Triticum.aestivum.xlsx
- SHA-256: `ae064dab13fd31ea37999630a6704199a81838e481485749664a2de4b7c451e3`
- Sheet: `Sorted by CYP name`
- Used range: `A1:J1701`
- Accepted CYP-assigned rows: `1475`
- Excluded unassigned rows: `224`
- Excluded bacterial rows: `1`
- Duplicate sequence groups retained: `60`
- Review status: `complete`

This workbook has a distinct A:J layout: G is `Seq. ID`, H is `best hit`, I is `%ID`, and headerless J is the literal protein. It has no separate assigned-name column, so the source H value is retained as the CYP symbol without inventing a different family. Rows 2-1476 are the continuous region with CYP best hits and percentages. Rows 1477-1700 have neither, and are excluded even where J contains an out-of-frame translation; row 1701 explicitly labels `Bacillus cellulosilyticus` bacterial contamination and is excluded. X/O, gaps, short proteins and duplicate sequences remain literal and row-distinct.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X-or-x | 218 |
| nonstandard-O | 99 |
| short-sequence | 301 |
| source-best-hit-pseudogene-label | 28 |
| source-gap | 96 |

## Representative records

| Row | ID | CYP best hit | Length |
|---:|---|---|---:|
| 2 | Tritaesb4S4940376.9 | CYP51G1 | 488 |
| 739 | Tritaesa1S3314486.3 | CYP76V1 | 496 |
| 1476 | Tritaesd5L4564033.2 | CYP735A4 | 465 |
