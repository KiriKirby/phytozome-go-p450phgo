# Plant resource review: Prunus persica older version

- Source file: `plants-Prunus.persica.P450s.doc`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Prunus.persica.P450s.doc
- Source SHA-256: `7f8ef51bb05e76aab38c489c4c14042cc9d7ca7a0d4cb4ab1f870c0f7a40fe9f`
- Container inspected: legacy Word document, 3,066 paragraphs, no tables, 49 pages
- Normalized text lines: `3067`
- Header-delimited sequence blocks retained: `306`
- Source title claim: `305 sequences`
- Duplicate exact-header groups retained: `2`
- Duplicate literal-sequence groups retained: `2`
- Review status: `complete`

## Resource-specific interpretation

The document states that 305 sequences were retrieved by a Phytozome Biomart P450-keyword search and explicitly says they were not checked for accuracy or annotated manually. The body actually contains 306 complete `>accession_peptide|Ppersica|gene|transcript` blocks, each followed only by literal uppercase protein lines. Blocks 54/55 and 156/157 are two exact duplicate header-and-sequence pairs; all four blocks remain source-order records with distinct `RecordKey` values. Because this old resource provides no CYP assignment, `symbol` remains empty and is not inferred from the updated peach workbook. The transcript field is used as ID, while peptide accession and gene ID remain in provenance. Explicit terminal `*` markers are removed; all other source sequence content is retained. No sequence is repaired, annotated, merged, or deduplicated.

## Status counts

| Status | Records |
|---|---:|
| short-sequence | 51 |
| source-terminal-stop | 291 |
| source-unannotated | 306 |
| unusually-long-sequence | 2 |

## Representative blocks

| Block | Source line | Transcript ID | Sequence | Status | Header |
|---:|---:|---|---:|---|---|
| 1 | 12 | ppa004078m | 531 aa | source-unannotated;source-terminal-stop | >17641348_peptide\|Ppersica\|ppa004078m.g\|ppa004078m |
| 154 | 1549 | ppa015200m | 517 aa | source-unannotated;source-terminal-stop | >17652470_peptide\|Ppersica\|ppa015200m.g\|ppa015200m |
| 306 | 3058 | ppa005500m | 457 aa | source-unannotated;source-terminal-stop | >17650973_peptide\|Ppersica\|ppa005500m.g\|ppa005500m |
