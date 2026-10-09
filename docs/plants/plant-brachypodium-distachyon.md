# Plant resource review: Brachypodium distachyon

- Source file: `plants-Brachypodium.FASTA.doc`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Brachypodium.FASTA.doc
- SHA-256: `d10a7ca25a08dc4f98e8a86a925b401b6f31337f59d37f89107f78e6f0894d9a`
- Container inspected: legacy Word, 2,652 paragraphs, 48 pages, no tables
- Source `>` blocks: `279`
- Accepted Brachypodium literal sequence blocks: `278`
- Explicit foreign blocks excluded: `1`
- Duplicate sequence groups retained: `0`
- Review status: `complete`

All header-delimited blocks were inspected. Block 74 is explicitly `Sorghum bicolor` and is excluded; the other 278 CYP blocks belong to the Brachypodium resource, including no-model fragments and pseudogenes. This working annotation document embeds genomic coordinates, phase markers, `&` joins and deletion prose among literal residue lines. Only those reviewed layout tokens are removed. X/x, gaps and internal stops remain literal, and only a final `*` is removed. Nothing is repaired, translated, merged, or externally completed.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X-or-x | 2 |
| internal-stop | 18 |
| short-sequence | 22 |
| source-fragment-or-missing-region | 14 |
| source-gap | 1 |
| source-joined-piece | 22 |
| source-pseudogene | 54 |
| source-terminal-stop | 255 |

## Representative records

| Accepted index | Source block | Line | ID | CYP | Length |
|---:|---:|---:|---|---|---:|
| 1 | 1 | 1 | Bradi4g25930 | CYP51G1 | 490 |
| 140 | 141 | 1355 | Bradi3g47750 | CYP86E1 | 532 |
| 278 | 279 | 2643 | Bradi4g29860.1 | CYP735A4 | 529 |

Excluded block 74 at line 698: `CYP71AM1 Sorghum bicolor XM_002451987` — explicit Sorghum bicolor comparison sequence.
