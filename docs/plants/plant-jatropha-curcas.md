# Plant resource review: Jatropha curcas

- Source file: `plants-Jatropha.P450s.2012.doc`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Jatropha.P450s.2012.doc
- Source SHA-256: `ab0a6347736fdd21093c00a5597fbdf52b9b040e487daee9e2f16c0f2b85c9d3`
- Container inspected: legacy Word document, 7,107 paragraphs, no tables, 129 pages
- Normalized split lines: `7,108`
- Source `>` blocks: `537`
- Accepted Jatropha records with literal sequence: `481`
- Excluded foreign, structural, or false-positive blocks: `56`
- Review status: `complete`

## Resource-specific interpretation

The title page says 484 Jatropha sequences after subtracting 31 Ricinus helpers and 20 false positives. Complete block inspection finds 537 `>` blocks: 32 headers explicitly identify Ricinus communis, one identifies Populus trichocarpa, and 20 occur below the document's exact `20 False positive hits (3.7%)` heading. Three additional `>` lines are not independent sequence records: `CYP90C (one sequence)` and `CYP90D (one sequence plus one pseudogene)` are family-count headings, while the name-only `CYP735A22` preface applies to the immediately following `BABX01044566.1` sequence block. Therefore the literal file contains 481 accepted Jatropha sequence records; the source's 484 claim is retained here as an explicit discrepancy rather than manufactured as empty or duplicate records.

For this document only, protein lines may contain separated leading/trailing genomic coordinates, `(0)`/`(1)`/`(2)`/`(?)`/`()` phase markers, and `&` joins. Those markers are removed while the residue text on the same line is retained. Lowercase `x`, uppercase `X`, `O`, and internal stops remain literal. Only an explicit final `*` on a record is removed; 175 records carry that marker. The words `COMPLETE` and `FINISHED` are annotations even though every letter is also an amino-acid code; they are explicitly rejected. Other lowercase prose cannot become sequence. No fragment is translated, repaired, extended, merged across `>` blocks, or filled from Ricinus or another database.

Accession-only Jatropha headers remain independent records with an empty CYP symbol, except `BABX01044566.1`, whose source name is supplied by the immediately preceding name-only `CYP735A22` preface. Duplicate accessions and adjacent-gene pieces remain block-distinct through unique `RecordKey` values.

## Status counts

| Status | Records |
|---|---:|
| ambiguous-X-or-x | 51 |
| internal-stop | 48 |
| nonstandard-O | 1 |
| short-sequence | 253 |
| source-fragment-or-missing-region | 275 |
| source-frameshift | 23 |
| source-gap-annotation | 27 |
| source-joined-piece | 147 |
| source-pseudogene | 104 |
| source-terminal-stop | 175 |

## Exclusion counts

| Reason | Blocks |
|---|---:|
| Ricinus comparison/assembly blocks | 32 |
| Populus comparison blocks | 1 |
| structural/name-only `>` headings | 3 |
| explicit false positives | 20 |

## Representative records

| Accepted index | Source block | Source line | ID | CYP symbol | Sequence | Status | Header |
|---:|---:|---:|---|---|---:|---|---|
| 1 | 1 | 68 | CYP51G1 | CYP51G1 | 486 aa | source-terminal-stop;ambiguous-X-or-x | CYP51G1 |
| 241 | 253 | 3364 | JcCA0270761.10 | CYP82C | 525 aa |  | CYP82C JcCA0270761.10 + |
| 481 | 517 | 6964 | JcCB0103981.10 | CYP727B23 | 559 aa | source-terminal-stop;source-fragment-or-missing-region | CYP727B23 JcCB0103981.10 - |

## Excluded blocks

| Source block | Source line | Header | Reason |
|---:|---:|---|---|
| 11 | 213 | CYP71B64 XM_002523189 Ricinus communis, XP_002523235.1 | explicit Ricinus communis assembly/comparison block |
| 12 | 217 | CYP71B65 EE255099 Ricinus communis EST, XP_002523238.1 | explicit Ricinus communis assembly/comparison block |
| 23 | 380 | CYP71B69 XM_002523196 Ricinus communis, XP_002523242.1 | explicit Ricinus communis assembly/comparison block |
| 25 | 395 | CYP71B78 XM_002522168 Ricinus communis, XP_002522214.1 | explicit Ricinus communis assembly/comparison block |
| 28 | 420 | CYP71B65 XM_002523192 Ricinus communis cytochrome P450, putative, mRNA | explicit Ricinus communis assembly/comparison block |
| 40 | 571 | CYP71D329 XM_002522857 Ricinus communis, XP_002522903.1 | explicit Ricinus communis assembly/comparison block |
| 119 | 1595 | CYP71AN23 Ricinus communis XM_002531819, XP_002531865.1 | explicit Ricinus communis assembly/comparison block |
| 135 | 1814 | CYP71BF5 XM_002511253.1 Ricinus communis | explicit Ricinus communis assembly/comparison block |
| 136 | 1822 | CYP71BF6 XM_002511252.1 Ricinus communis, EEF51900.1 | explicit Ricinus communis assembly/comparison block |
| 137 | 1827 | CYP71BF9 XM_002511247.1 Ricinus communis, EEF51895.1 | explicit Ricinus communis assembly/comparison block |
| 142 | 1901 | Populus trichocarpa CYP75A13 XM_002313967 | explicit Populus trichocarpa comparison block |
| 147 | 1959 | CYP75B63 Ricinus communis XM_002514619, XP_002514665.1 | explicit Ricinus communis assembly/comparison block |
| 307 | 4039 | CYP84A51 Ricinus communis XM_002512940 partial (exon 1) | explicit Ricinus communis assembly/comparison block |
| 354 | 4670 | CYP736A XM_002511922 Ricinus communis, EEF50637.1 | explicit Ricinus communis assembly/comparison block |
| 359 | 4705 | CYP736A XM_002531999 Ricinus communis, XP_002532045.1 | explicit Ricinus communis assembly/comparison block |
| 376 | 4927 | CYP714A20 Ricinus communis XM_002516439, XP_002516485.1 | explicit Ricinus communis assembly/comparison block |
| 380 | 4988 | CYP714M3P Ricinus communis XM_002531771 partial seq | explicit Ricinus communis assembly/comparison block |
| 390 | 5145 | CYP735A22 | name-only preface; the immediately following BABX01044566.1 block carries its literal assembled sequence and inherits this source name |
| 427 | 5715 | CYP90C (one sequence) | family count heading formatted with `>`; the following named blocks carry the records and sequences |
| 429 | 5736 | CYP90D (one sequence plus one pseudogene) | family count heading formatted with `>`; the following named blocks carry the records and sequences |
| 432 | 5779 | CYP90D Ricinus communis XM_002527368.1, XP_002527414.1 | explicit Ricinus communis assembly/comparison block |
| 438 | 5884 | CYP707A82 Ricinus communis XM_002510168, XP_002510214.1 | explicit Ricinus communis assembly/comparison block |
| 446 | 5987 | CYP716A56 Ricinus communis XM_002522891.1, XP_002522937.1 | explicit Ricinus communis assembly/comparison block |
| 450 | 6052 | CYP718 Ricinus communis XM_002520338, XP_002520384.1 | explicit Ricinus communis assembly/comparison block |
| 453 | 6096 | CYP722A1 Ricinus communis XM_002510377.1, XP_002510423.1 | explicit Ricinus communis assembly/comparison block |
| 455 | 6129 | CYP722C2 XM_002521491.1 Ricinus communis, EEF40808.1 | explicit Ricinus communis assembly/comparison block |
| 457 | 6158 | CYP722C3 Ricinus communis XM_002524287.1, XP_002524333.1 | explicit Ricinus communis assembly/comparison block |
| 460 | 6211 | CYP728D12 Ricinus communis XM_002522923, XP_002522969.1 | explicit Ricinus communis assembly/comparison block |
| 467 | 6318 | CYP86A74 Ricinus communis XM_002525562, XP_002525608.1 | explicit Ricinus communis assembly/comparison block |
| 469 | 6344 | CYP86A75 Ricinus communis  XM_002511829 framshift = & | explicit Ricinus communis assembly/comparison block |
| 471 | 6370 | CYP86A76 Ricinus communis XM_002509774.1, XP_002509820.1 | explicit Ricinus communis assembly/comparison block |
| 473 | 6386 | CYP86B14 XM_002523729.1 Ricinus communis, XP_002523775.1 | explicit Ricinus communis assembly/comparison block |
| 477 | 6418 | CYP86C13 Ricinus communis XM_002515007.1, XP_002515053.1 | explicit Ricinus communis assembly/comparison block |
| 505 | 6765 | CYP704A85P Ricinus communis XM_002529504 short, missing C-term | explicit Ricinus communis assembly/comparison block |
| 508 | 6808 | CYP704A83 Ricinus communis XM_002511699, EEF50414.1 | explicit Ricinus communis assembly/comparison block |
| 514 | 6919 | CYP97C23 XM_002519381 Ricinus communis XP_002519427.1 | explicit Ricinus communis assembly/comparison block |
| 518 | 6991 | JcCA0317921.20 -  not a P450 | block is below the document's explicit `20 False positive hits` heading |
| 519 | 7004 | JcCA0098461.10 +  /pseudo not P450 sequence | block is below the document's explicit `20 False positive hits` heading |
| 520 | 7013 | JcCA0282231.30 + /pseudo not P450 sequence | block is below the document's explicit `20 False positive hits` heading |
| 521 | 7017 | JcCA0126211.20 - /pseudo not P450 sequence | block is below the document's explicit `20 False positive hits` heading |
| 522 | 7022 | JcCA0291251.10 -  /pseudo not P450 sequence | block is below the document's explicit `20 False positive hits` heading |
| 523 | 7026 | JcCA0314781.30 -  not a P450 sequence | block is below the document's explicit `20 False positive hits` heading |
| 524 | 7034 | JcCB0233871.10 - /pseudo not a P450 sequence | block is below the document's explicit `20 False positive hits` heading |
| 525 | 7042 | JcCB0071431.10 -  not a P450 sequence | block is below the document's explicit `20 False positive hits` heading |
| 526 | 7049 | JcCB0298631.10 + /pseudo not a P450 sequence | block is below the document's explicit `20 False positive hits` heading |
| 527 | 7056 | JcCB0071491.10 -  not a P450 sequence | block is below the document's explicit `20 False positive hits` heading |
| 528 | 7064 | JcCB0352601.10 -  /pseudo not a P450 seq. | block is below the document's explicit `20 False positive hits` heading |
| 529 | 7069 | JcCB0292091.10 +  /pseudo not a P450 sequence | block is below the document's explicit `20 False positive hits` heading |
| 530 | 7073 | JcCB0025561.10 - /pseudo not a P450 seq. | block is below the document's explicit `20 False positive hits` heading |
| 531 | 7078 | JcCA0148621.20 - /pseudo not a P450 seq. | block is below the document's explicit `20 False positive hits` heading |
| 532 | 7081 | JcCB0476951.10 -  /pseudo not a P450 seq. | block is below the document's explicit `20 False positive hits` heading |
| 533 | 7088 | JcCB0675101.20 -  /pseudo not a P450 seq | block is below the document's explicit `20 False positive hits` heading |
| 534 | 7095 | JcCB1002561.10 -  /pseudo not a P450 seq. | block is below the document's explicit `20 False positive hits` heading |
| 535 | 7099 | JcPR03CKLC6.10 +  /pseudo not a P450 sequence | block is below the document's explicit `20 False positive hits` heading |
| 536 | 7102 | JcPR04GFG8C.10 +  /pseudo probable false positive hit | block is below the document's explicit `20 False positive hits` heading |
| 537 | 7105 | JcCD0116484.10 +  /pseudo probable false positive hit | block is below the document's explicit `20 False positive hits` heading |
