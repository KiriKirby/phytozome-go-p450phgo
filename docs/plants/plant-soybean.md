# Plant resource review: soybean

- Source file: `plants-soybean.doc`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/soybean.doc
- Source SHA-256: `1e06e73024df91a66d73774bce123f17de19fd8159af9be59c279fcca3508342`
- Container inspected: legacy Word document, 10,775 paragraphs, no tables, 199 pages
- Normalized text lines: `10783`
- Accepted named CYP blocks: `171`
- Accepted blocks with literal protein sequence: `115`
- Accepted blocks without literal protein sequence: `56`
- Excluded EST nucleotide headers after the explicit appendix marker: `953`
- Duplicate exact-header groups retained: `8`
- Duplicate literal-sequence groups retained: `2`
- Review status: `complete`

## Resource-specific interpretation

The document is divided at its explicit `955 CYTOCHROME P450 EST FOR Glycine max` heading. Before that heading, 171 `>CYP...` annotation blocks are accepted in source order; 115 contain literal protein paragraphs and 56 provide names and annotations without a protein sequence. After the heading, all `>gi|...` records are nucleotide EST data and are excluded rather than translated. Protein paragraphs are accepted only when the entire paragraph consists of source amino-acid characters plus reviewed coordinate, phase, ampersand, or question-mark notation. Coordinate numbers, parenthesized phase markers, whitespace, and ampersand frameshift separators are layout annotations, not residues. The one `??` residue segment and the one internal stop remain literal. Only terminal stop markers are removed. Pseudogene, fragment, missing-region, and frameshift evidence remains in review status. Duplicate names and sequences remain block-distinct. No sequence is translated, repaired, completed, or obtained externally.

## Status counts

| Status | Records |
|---|---:|
| internal-stop | 1 |
| sequence-missing | 56 |
| short-sequence | 16 |
| source-fragment-or-missing-region | 12 |
| source-frameshift-marker | 7 |
| source-pseudogene | 9 |
| source-question-mark-residue | 1 |
| source-terminal-stop | 52 |

## Representative blocks

| Block | Source line | Symbol | Sequence | Status | Header |
|---:|---:|---|---:|---|---|
| 1 | 84 | CYP51G1 | 475 aa |  | >CYP51G1     Glycine max (soybeans, Fabales) |
| 86 | 1139 | CYP85A13 | 464 aa | source-fragment-or-missing-region;source-terminal-stop | >CYP85A13 fragments Glycine max (soybean) |
| 171 | 2259 | CYP736A | 124 aa | source-pseudogene;short-sequence | >CYP736A pseudogene  Glycine max (soybean) |
