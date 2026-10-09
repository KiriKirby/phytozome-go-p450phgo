# Plant resource review: Fragaria ananassa

- Source file: `plants-Fragaria.ananassa.xlsx`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Fragaria.ananassa.xlsx
- Source SHA-256: `0d10575eb3ab9dffd8ea35717cdbab720a9f95a70e4b4110b11bfed9ff697665`
- Workbook sheet: `Sorted by CYP name`
- Used range: `A1:L225`
- Hidden rows: `[]`
- Hidden columns: `[]`
- Merged cells: `[]`
- Native table parts: `0`
- Accepted named rows: `206`
- Accepted rows with literal sequence: `206`
- Excluded rows: `18`
- Review status: `complete`

## Resource-specific interpretation

This workbook is not parsed with the ordinary H-column ID rule. Row 1 explicitly names A as `Gotoh ID`, and every A value in rows 2-225 begins with `>`. For accepted rows 2-206, H repeats A without `>`; row 207 is a real assigned CYP88A50/CYP88A record whose H cell instead says `opposite strand`. The reviewed output therefore uses A with exactly one leading `>` removed as the stable ID for every accepted row. Rows 208-225 also say `opposite strand`, but I and K are both `no hit` and J is empty, so those unassigned models are excluded. All 206 assigned rows keep their literal L-column sequences, including 16 gaps, literal X/O, 62 short fragments, and two digit-P best-hit labels. No sequence is inferred, repaired, or deduplicated.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X-or-x | 16 |
| nonstandard-O | 20 |
| short-sequence | 62 |
| source-best-hit-pseudogene-label | 2 |
| source-gap | 16 |

## Representative and excluded rows

| Row | Seq ID | Best hit | Assigned CYP | Sequence | Status |
|---:|---|---|---|---:|---|
| 2 | Fraganan_rscf00000033.1.21 | CYP51G1 | CYP51G1 | 487 aa |  |
| 104 | Fraganan_rscf00001030.1.4 | CYP92A84 | CYP92A | 513 aa |  |
| 207 | Fraganan_icon00003587_a.1.1rev | CYP88A50 | CYP88A | 251 aa | short-sequence |
| 208 | Fraganan_icon00004271_a.1.1 |  |  | 152 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 209 | Fraganan_icon00008119_a.1.1 |  |  | 15 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 210 | Fraganan_icon00009633_a.1.1 |  |  | 201 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 211 | Fraganan_icon00009753_a.1.0 |  |  | 213 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 212 | Fraganan_icon00011181_a.1.1 |  |  | 206 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 213 | Fraganan_icon00012471_a.1.1 |  |  | 195 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 214 | Fraganan_icon00014753_a.1.1 |  |  | 144 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 215 | Fraganan_icon00022837_a.1.1 |  |  | 122 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 216 | Fraganan_icon00044082_a.1.0 |  |  | 118 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 217 | Fraganan_icon19548641_s.1.0 |  |  | 82 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 218 | Fraganan_icon20387067_s.1.1 |  |  | 160 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 219 | Fraganan_icon20544974_s.1.1 |  |  | 156 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 220 | Fraganan_rscf00000073.1.127 |  |  | 434 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 221 | Fraganan_rscf00000151.1.20 |  |  | 307 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 222 | Fraganan_rscf00000911.1.20 |  |  | 125 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 223 | Fraganan_rscf00007124.1.1 |  |  | 392 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 224 | Fraganan_rscf00007255.1.1 |  |  | 333 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
| 225 | Fraganan_rscf00007376.1.1 |  |  | 267 chars (excluded) | excluded: source marks the model as opposite strand with no hit and no CYP assignment |
