# Plant resource review: Vitis vinifera

- Source file: `plants-vitis.doc`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/vitis.doc
- Source SHA-256: `37e8b56e9a9db299205b5ab80e0985697b55c4c21d14e2c5194fe8f255c4ec42`
- Container inspected: legacy Word document, 8,443 paragraphs, 138 pages
- Normalized split lines: `8,443`
- Source `>` blocks: `702`
- Accepted source-native Vitis blocks: `672`
- Accepted blocks with literal sequence: `672`
- Explicit foreign naming-reference blocks excluded: `30`
- Duplicate literal-sequence groups retained: `11`
- Review status: `complete`

## Resource-specific interpretation

The document is a working 2007 curation file, not plain FASTA. It interleaves Vitis WGS assemblies, GenPept entries, alleles, duplicates, fragments, pseudogenes, and explicit sequences from other plant species used as family-naming references. All 702 headers were reviewed in source order. The 30 blocks listed below are explicit foreign reference sequences and are excluded; Vitis headers that merely compare identity to another species remain included.

The title-page statement `591 sequences are present below` predates later revisions and does not match the current 702 literal blocks. The current file contains many alternate genome-project, allele, duplicate, and revised records. They are retained independently instead of being silently merged to force the historical count. Block 693 has a complete terminal-star-bounded CYP736A21 sequence followed by an unrelated unheaded protein; only its CYP sequence through normalized line 8067 is accepted.

Reviewed protein lines may carry leading/trailing genomic coordinates, phase markers, whitespace, and `&` joins. Those layout markers are removed; literal X/x, O, gap characters, and internal stops are preserved. Only one final `*` is removed. `$$$$`, ampersand separators, `(GAP)`, and lowercase prose are annotations. Nothing is translated, repaired, or externally completed.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X-or-x | 111 |
| internal-stop | 150 |
| short-sequence | 212 |
| source-fragment-or-missing-region | 128 |
| source-frameshift | 22 |
| source-gap-annotation | 23 |
| source-joined-piece | 18 |
| source-pseudogene | 255 |
| source-terminal-stop | 300 |

## Representative records

| Accepted index | Source block | Source line | ID | CYP symbol | Sequence | Status | Header |
|---:|---:|---:|---|---|---:|---|---|
| 1 | 1 | 91 | CAAP02000072.1 | CYP51G6 | 486 aa | source-terminal-stop | CYP51G6 CAAP02000072.1 81% to 51G1 Arab. |
| 337 | 348 | 4029 | CAAP02005229 | CYP82D10 | 526 aa | source-terminal-stop;source-fragment-or-missing-region | CYP82D10 CAAP02005229.1d 16585-14851 (-) strand, 89% to 89D22P |
| 672 | 702 | 8149 | CAAP02005006.1 | CYP736A27 | 492 aa | source-terminal-stop;source-pseudogene;source-fragment-or-missing-region;source-frameshift | CYP736A27 CAAP02005006.1 86% to CAAP02000243.1 CYP736 |

## Excluded foreign reference blocks

| Source block | Source line | Header | Reason |
|---:|---:|---|---|
| 5 | 129 | CYP71AH1 old 71A11 tobacco | Nicotiana tabacum naming reference |
| 6 | 138 | CYP71AH2 tobacco | Nicotiana tabacum naming reference |
| 7 | 149 | 71A9/CYP71AH3 Glycine max | Glycine max naming reference |
| 10 | 184 | CYP71AH6 Gossypium raimondii 58% to CAAP02005003.1a, 53% to 71A9/71AH3 | Gossypium raimondii naming reference |
| 71 | 831 | CYP71BG1 Solanum tuberosum | Solanum tuberosum naming reference |
| 72 | 844 | CYP71BG2 tomato breaker fruit Solanum lycopersicum | Solanum lycopersicum naming reference |
| 73 | 858 | CYP71BG3 CA993587.1 Gossypium hirsutum CO128388.1 Gossypium raimondii | Gossypium naming reference |
| 74 | 869 | CYP71BG4 DY280303.1, DY276238.1, Citrus clementina, | Citrus hybrid naming reference |
| 338 | 3904 | CYP82D1     Medicago sativa (alfalfa) | Medicago sativa naming reference |
| 339 | 3917 | CYP82D2 Populus | Populus naming reference |
| 340 | 3931 | CYP82D3 Coptis japonica AB374406 | Coptis japonica naming reference |
| 365 | 4235 | CYP82H1 Ammi majus AY532373.1 | Ammi majus naming reference |
| 401 | 4604 | CYP82J1 Populus | Populus naming reference |
| 402 | 4618 | CYP82K1 Populus | Populus naming reference |
| 403 | 4631 | CYP82L1 Populus | Populus naming reference |
| 404 | 4645 | CYP82L2 Populus | Populus naming reference |
| 593 | 6848 | CYP714A1 Arab | Arabidopsis naming reference |
| 594 | 6859 | CYP714A3 Populus | Populus naming reference |
| 595 | 6874 | CYP714B1 (japonica cultivar-group) chromosome 7 44% to 714A1 | rice naming reference |
| 596 | 6888 | CYP714C1 rice | rice naming reference |
| 597 | 6902 | CYP714C4 Lolium | Lolium naming reference |
| 598 | 6914 | CYP714D1 (japonica cultivar-group) chromosome 5 | rice naming reference |
| 599 | 6926 | CYP714E1 Medicago truncatula (barrel medic, Fabales) | Medicago truncatula naming reference |
| 600 | 6940 | CYP714E2 Populus | Populus naming reference |
| 601 | 6951 | CYP714E7 CR485106.1  CR485106 MTH2 Medicago truncatula genomic | Medicago truncatula naming reference |
| 602 | 6962 | CYP714F1 Populus | Populus naming reference |
| 603 | 6976 | CYP714E4X now CYP714G1  Populus | Populus naming reference |
| 604 | 6990 | CYP714E5X now CYP714G2  Populus | Populus naming reference |
| 605 | 7004 | 714G3/E8 PH2_Pehy Jixiang Han Petunia hybrida | Petunia hybrida naming reference |
| 606 | 7008 | CYP714E6X now CYP714G4 Populus | Populus naming reference |
