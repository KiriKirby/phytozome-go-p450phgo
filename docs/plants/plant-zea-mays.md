# Plant resource review: Zea mays

- Source file: `plants-zea.doc`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/zea.doc
- SHA-256: `96f8a5f71ed55f34ce1dba75949c71c892d3115ec3a140b8c2ab42ec7302a9bd`
- Container inspected: legacy Word, 155 paragraphs, 3 pages, no tables
- Source family-sorted `>` groups: `26`
- Accepted source groups: `26`
- Groups with literal protein fragments: `18`
- Groups without literal protein: `8`
- Distinct accessions in alphabetical appendix: `42`
- Review status: `complete`

The title says `39 Zea mays P450 ESTs`, but the document's own alphabetical appendix contains 42 distinct accession rows. The family-sorted body represents those accessions in 26 annotation groups (and mentions the same 42 distinct accessions, including T70647 in the AI622303 comparison). The release therefore preserves the 26 actual source groups instead of inventing one protein per accession. Eighteen groups contain literal coordinate-delimited protein fragments; eight contain annotation only and retain an empty sequence. The alphabetical appendix is an index and creates no duplicate records. X and internal stops remain literal; only a terminal `*` is removed.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X | 7 |
| est-protein-fragment | 18 |
| grouped-accessions | 7 |
| internal-stop | 1 |
| source-sequence-not-present | 8 |
| source-terminal-stop | 1 |

## Representative groups

| Block | Line | ID | CYP assignment | Length | Source header |
|---:|---:|---|---|---:|---|
| 1 | 8 | AI770623 | CYP51 | 219 | AI770623, AI621427, AI649583, AI621417, T12664 = CYP51 98% to sorghum CYP51 |
| 4 | 21 | AI714669 | CYP71C2 | 0 | AI714669, AI943993, AI943991, AI649716, AI395920, AI932184, AI891364 97% identical to |
| 12 | 43 | AI855377 | CYP72A5 | 268 | AI855377, AI637224 = 72A5 |
| 23 | 89 | AI734373 | CYP98A1 | 282 | AI734373 T18826 86% to 98A1 AI881302 80% to 98A1 |
| 26 | 104 | AI947887 | CYP714A2 | 127 | AI947887 55% to 714A2 |
