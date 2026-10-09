# Plant resource review: Carica papaya

- Source file: `plants-papaya.doc`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/papaya.doc
- Source SHA-256: `9b1c28c43941a5614155129fe47c2cad9387586778df90963b1fee78894a24cd`
- Container inspected: legacy Word document, 3,112 paragraphs, no tables
- Header-delimited blocks: `225`
- Papaya blocks accepted with literal sequence: `182`
- Explicit Vitis vinifera comparison blocks excluded: `43`
- Source-labeled pseudogenes: `39`
- Duplicate CYP-name groups retained: `1`
- Duplicate literal-sequence groups: `0`
- Review status: `complete`

## Resource-specific interpretation

The document declares 182 sequences: 143 genes and 39 pseudogenes. It contains 225 `>` blocks in total. Exactly 43 headers explicitly identify `Vitis vinifera`; these are the comparison sequences mentioned on the title page and are excluded. The remaining 182 blocks are retained in source order and all contain literal protein sequence.

For this document only, sequence lines are uppercase amino-acid text (with lowercase `x` normalized to literal `X`), optionally bounded by separated genomic coordinates and `(0)`/`(1)`/`(2)` phase markers. One literal `--` gap, every `X`, every internal `*`, fragments, frameshifts, and pseudogene blocks are retained. Only terminal `*` markers are removed. Three malformed coordinate/phase lines and the two annotation lines separating the strands of CYP712A10P are handled by exact source line and record ID. Any other non-empty line after a sequence begins is an error, so the parser cannot silently jump over a new annotation.

CYP729A17 occurs twice as source versions v1 and v2 and remains two records with distinct `RecordKey` values. No sequence is repaired, translated, filled, or deduplicated.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X | 13 |
| internal-stop | 39 |
| short-sequence | 30 |
| source-frameshift | 13 |
| source-gap | 1 |
| source-partial-or-gap | 43 |
| source-pseudogene-label | 39 |

## Representative and boundary blocks

| Block | Source line | CYP name | Sequence | Status | Header |
|---:|---:|---|---:|---|---|
| 1 | 14 | CYP51G1 | 484 aa |  | CYP51G1 |
| 106 | 1440 | CYP87A11 | 387 aa | source-partial-or-gap;ambiguous-X | CYP87A11 old b |
| 224 | 3087 | CYP736A15 | 507 aa |  | CYP736A15 old a |
| 225 | 3100 | CYP736A16 | 498 aa |  | CYP736A16 old c |

## Excluded comparison blocks

- Block 2, line 28: `CYP51G1 AM475390.2 Vitis vinifera`
- Block 23, line 300: `CYP71BE1 AM445470.2  Vitis vinifera 52% to 71D16`
- Block 42, line 579: `CYP73A78 AM455281.2 Vitis vinifera`
- Block 46, line 629: `CYP74B13 AM441513 PLN 18-MAY-2007 Vitis vinifera`
- Block 48, line 655: `CYP75A28 AJ880356.1 Vitis vinifera`
- Block 50, line 682: `CYP75B32   Vitis vinifera (core eudicotyledons; Vitales)`
- Block 52, line 714: `CYP76A10 AM451133.2  Vitis vinifera`
- Block 56, line 761: `CYP76G6 AM452176.2  Vitis vinifera`
- Block 58, line 791: `CYP77A14 AM463740.2 Vitis vinifera complement(2718..4265)`
- Block 60, line 816: `CYP77B6 AM482217.2 Vitis vinifera`
- Block 94, line 1275: `CYP85A6 AM431608 Vitis vinifera 71% to CYP85Aa, 70% to 85Ab`
- Block 95, line 1288: `CYP85A7 DQ235273.1 Vitis vinifera brassinosteroid-6-oxidase (BR6OX1)`
- Block 101, line 1367: `CYP86B7 AM486428.2 Vitis vinifera`
- Block 103, line 1396: `CYP86C8 AM462286.2 Vitis vinifera`
- Block 114, line 1542: `CYP89A38 old a AM423953.2  Vitis vinifera 64% to papaya`
- Block 115, line 1554: `CYP89A39 old b AM423953.2  Vitis vinifera 66% to papaya (best blast hit)`
- Block 116, line 1565: `CYP89A40 old c AM437626.2  Vitis vinifera 64% to papaya`
- Block 117, line 1577: `CYP89A41 old d AM433215.2  Vitis vinifera 63% to papaya`
- Block 119, line 1603: `CYP90A16 AM463180 Vitis vinifera`
- Block 121, line 1633: `CYP90B12 AM441474.1 Vitis vinifera`
- Block 123, line 1666: `CYP90C5 AM429218 Vitis vinifera complement(join(25517..25610,25704..25831,25899..26005,`
- Block 125, line 1697: `CYP90D8 AM482768.1 Vitis vinifera missing last exon`
- Block 127, line 1730: `CYP92A32 AM446822.2  Vitis vinifera`
- Block 129, line 1756: `CYP93A9 AM429328.2 Vitis vinifera`
- Block 146, line 1980: `CYP97A11 AM476239.2 Vitis vinifera missing N-term and C-term`
- Block 148, line 2019: `CYP97B15 AM472981.1 Vitis vinifera partial`
- Block 150, line 2049: `CYP97C12 AM463116.2 Vitis vinifera`
- Block 152, line 2083: `CYP98A43 AM435080.1 Vitis vinifera`
- Block 154, line 2114: `CYP701A19 AM434546.2 Vitis vinifera`
- Block 158, line 2176: `CYP703A9 AM475919.1  Vitis vinifera`
- Block 178, line 2449: `CYP710A18 AM446467.2 Vitis vinifera`
- Block 180, line 2475: `CYP711A14 Vitis vinifera AM474585.2 73% to CYP711A7v1 Populus trichocarpa`
- Block 186, line 2567: `CYP715A6 AM468263 Vitis vinifera`
- Block 188, line 2596: `CYP716A15 AM471070.1  Vitis vinifera`
- Block 190, line 2623: `CYP718A4 AM448959.2 Vitis vinifera`
- Block 194, line 2680: `CYP722A1    Vitis vinifera`
- Block 199, line 2742: `CYP727A7 AM439016.2 Vitis vinifera and AM454319.2 N-term`
- Block 201, line 2776: `CYP728B5 AM457336.1 Vitis vinifera with 3 frameshifts`
- Block 202, line 2790: `CYP728B6 AM457336.1 Vitis vinifera`
- Block 203, line 2803: `CYP728B7 AM442348.1 Vitis vinifera 68% to papaya`
- Block 204, line 2816: `CYP728B8 AM426541.1 Vitis vinifera 56% to papaya`
- Block 211, line 2885: `CYP728G1   Vitis vinifera (grapes)`
- Block 220, line 3026: `CYP734A13 AM443674.1 Vitis vinifera 75% to Pisum`
