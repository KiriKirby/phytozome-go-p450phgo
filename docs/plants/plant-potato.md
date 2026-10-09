# Plant resource review: potato

- Source file: `plants-potato.P450s.doc`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/potato.P450s.doc
- Source SHA-256: `057367d8adcd0018119c1d7458de4205a747b8f40e8e91854db20ee144d21933`
- Container inspected: legacy Word document, 13,474 paragraphs, no tables
- Accepted source blocks: `920`
- Blocks with literal protein sequence: `899`
- Blocks without accepted sequence: `21`
- Explicit foreign/comparison blocks excluded: `12`
- Review status: `complete`

## Resource-specific interpretation

This file is not ordinary FASTA. Each `>` block may contain a potato gene model, genomic coordinates, exon-boundary markers such as `(0)`/`(1)`/`(?)`, joined fragments marked with `&`, pseudogenes, EST fragments, alternate versions, or an explicitly foreign comparison sequence. The parser is exclusive to `potato.P450s.doc`. It preserves every accepted potato block in source order, does not merge equal CYP names, strips only a terminal `*`, and retains internal stops and literal `X` residues. `Query`/`Sbjct` alignments and the explicitly named tobacco, tomato, eggplant, Nicotiana, Solanum phureja, and Capsicum comparison blocks are excluded.

Known source uncertainty remains visible in `review_status`; no residue, splice boundary, frameshift, gap, or missing terminus is repaired.

## Status counts

| Status | Records |
|---|---:|
| pseudogene | 26 |
| partial | 201 |
| frameshift | 70 |
| sequence-gap | 6 |
| uncertain-boundary | 15 |
| internal-stop | 135 |
| ambiguous-residue | 9 |
| short-sequence | 366 |
| unusually-long-sequence | 0 |
| sequence-missing | 21 |

## Excluded comparison blocks

| Source line | Header | Reason |
|---:|---|---|
| 2209 | FI033502.1 tobacco FH434438.1, FI006728.1, FH605148.1, FH346097.1, ET755410.1\| | tobacco comparison sequence |
| 7307 | CYP80F5 tobacco GSS FH610574.1, FH441154.1, FH610657.1, FH441233.1 | tobacco comparison sequence |
| 7362 | CYP80N1 ortholog from eggplant | eggplant BLAST comparison |
| 7395 | CYP80N1 ortholog from tobacco | tobacco BLAST comparison |
| 7418 | CYP80N1 ortholog from Nicotiana benthamiana | Nicotiana BLAST comparison |
| 7435 | CYP80N1 ortholog from Nicotiana benthamiana | Nicotiana BLAST comparison |
| 7448 | CYP80N1 ortholog Solanum phureja | Solanum phureja comparison |
| 9785 | FT251464.1 tomato GSS sequence 96% to CYP94C.3 potato, not found in ver 1.00. | tomato comparison sequence |
| 11037 | Tobacco FI058991.1 FH882184.1, FH882707.1, FH220054.1, FH882708.1, FH421254.1 | tobacco comparison sequence |
| 11078 | scaffold02618 tomato | tomato comparison scaffold |
| 11085 | scaffold02618 tomato | tomato comparison scaffold |
| 11102 | CYP704B Capsicum annuum EU620581 | Capsicum annuum comparison sequence |

## Accepted records

| Block | Source line | ID / symbol | Sequence | Status | Header |
|---:|---:|---|---:|---|---|
| 1 | 13 | CYP51G1 | 487 aa |  | CYP51G1 |
| 2 | 27 | CYP51G1 | missing | sequence-missing | CYP51G1 |
| 3 | 34 | CYP51-pseudogene-1 | 103 aa | pseudogene; short-sequence | CYP51 pseudogene 1 |
| 4 | 42 | CYP51-pseudogene-1 | missing | sequence-missing | CYP51 pseudogene 1 |
| 5 | 47 | CYP51-pseudogene-2 | 99 aa | pseudogene; internal-stop; short-sequence | CYP51 pseudogene 2 |
| 6 | 56 | CYP51-pseudogene-2 | missing | sequence-missing | CYP51 pseudogene 2 |
| 7 | 62 | CYP71D241 | 499 aa |  | CYP71D241 CYP71B.1 |
| 8 | 81 | CYP71D240P | 494 aa | internal-stop | CYP71D240P CYP71B.2Pv1 |
| 9 | 99 | CYP71B.2Pv2 | 45 aa | short-sequence | CYP71B.2Pv2 |
| 10 | 107 | CYP71D239P | 407 aa | sequence-gap; internal-stop | CYP71D239P CYP71B.3P |
| 11 | 124 | CYP71D238P | 406 aa | partial; uncertain-boundary; internal-stop | CYP71D238P CYP71B.4P |
| 12 | 139 | CYP71B.4Pv2 | 44 aa | short-sequence | CYP71B.4Pv2 |
| 13 | 147 | CYP71D237P | 30 aa | short-sequence | CYP71D237P CYP71B.5P |
| 14 | 155 | CYP71D236P | 442 aa | uncertain-boundary | CYP71D236P CYP71B.6P |
| 15 | 172 | CYP71D235P | 44 aa | short-sequence | CYP71D235P CYP71B.7P |
| 16 | 180 | CYP71D186P | 579 aa |  | CYP71D186P CYP71B.8P |
| 17 | 201 | CYP71D234P | 54 aa | short-sequence | CYP71D234P CYP71B.9P |
| 18 | 209 | CYP71D233P | 289 aa | short-sequence | CYP71D233P CYP71B.10P |
| 19 | 222 | CYP71D232 | 495 aa | frameshift | CYP71D232 CYP71B.11 |
| 20 | 240 | CYP71D185bP | 165 aa | short-sequence | CYP71D185bP CYP71B.12P |
| 21 | 252 | CYP71D185a | 504 aa | frameshift | CYP71D185a CYP71B.13 |
| 22 | 271 | CYP71D243P | 378 aa | partial; frameshift | CYP71D243P CYP71B.14P |
| 23 | 288 | CYP71D242 | 202 aa | partial; short-sequence | CYP71D242 CYP71B.15 |
| 24 | 302 | CYP71D184 | 206 aa | short-sequence | CYP71D184 CYP71B.16v1 |
| 25 | 313 | CYP71B.16v2 | 62 aa | short-sequence | CYP71B.16v2 |
| 26 | 321 | CYP71D244P | 495 aa | internal-stop | CYP71D244P CYP71B.17P |
| 27 | 340 | CYP71D245P | 286 aa | short-sequence | CYP71D245P CYP71B.18P |
| 28 | 354 | CYP71D246P | 53 aa | short-sequence | CYP71D246P CYP71B.19P |
| 29 | 362 | CYP71D6v1 | 501 aa |  | CYP71D6v1 Solanum chacoense (Chaco potato) |
| 30 | 378 | CYP71D7 | 500 aa |  | CYP71D7   Solanum chacoense (Chaco potato) |
| 31 | 396 | CYP71D249 | 501 aa | pseudogene; partial; frameshift; internal-stop | CYP71D249 CYP71D.1 |
| 32 | 414 | CYP71D250 | 500 aa | partial | CYP71D250 CYP71D.2 |
| 33 | 431 | CYP71D7v2 | 500 aa | partial | CYP71D7v2 = old CYP71D.3 |
| 34 | 448 | CYP71D6v2 | 501 aa | partial | CYP71D6v2 CYP71D.4 |
| 35 | 465 | CYP71D4bP | 502 aa |  | CYP71D4bP CYP71D.5P |
| 36 | 485 | CYP71D4a | 517 aa | partial | CYP71D4a    Solanum tuberosum cv. Datura (potato) |
| 37 | 502 | CYP71D4a | 502 aa |  | CYP71D4a CYP71D.6 |
| 38 | 518 | CYP71D251P | 77 aa | short-sequence | CYP71D251P CYP71D.7P pseudogene C-term |
| 39 | 526 | CYP71D252P | 490 aa |  | CYP71D252P CYP71D.8P |
| 40 | 546 | CYP71D252P-de2b | 43 aa | short-sequence | CYP71D252P-de2b CYP71D.8P-de2b |
| 41 | 553 | CYP71D253 | 502 aa | frameshift; ambiguous-residue | CYP71D253 CYP71D.9 |
| 42 | 570 | CYP71D254P | 59 aa | internal-stop; short-sequence | CYP71D254P CYP71D.10P pseudogene |
| 43 | 577 | CYP71D255P | 359 aa |  | CYP71D255P CYP71D.11P pseudogene joined with CYP71D.12P pseudogene |
| 44 | 594 | CYP71D198 | 509 aa |  | CYP71D198 ortholog CYP71D.13v1 |
| 45 | 609 | CYP71D.13v2 | 205 aa | short-sequence | CYP71D.13v2 |
| 46 | 620 | CYP71D262P | 524 aa | partial; frameshift | CYP71D262P CYP71D.14P |
| 47 | 640 | CYP71D261P | 173 aa | short-sequence | CYP71D261P CYP71D.15P |
| 48 | 650 | CYP71D260P | 509 aa |  | CYP71D260P CYP71D.16P |
| 49 | 667 | CYP71D259P | 430 aa | internal-stop | CYP71D259P CYP71D.17P |
| 50 | 684 | CYP71D258 | 505 aa | partial | CYP71D258 CYP71D.18 |
| 51 | 702 | CYP71D257P | 56 aa | short-sequence | CYP71D257P CYP71D.19P |
| 52 | 710 | CYP71D256P | 415 aa | partial; frameshift | CYP71D256P CYP71D.20P |
| 53 | 726 | CYP71D189b | 440 aa |  | CYP71D189b ortholog CYP71D.21 |
| 54 | 743 | CYP71D264 | 502 aa | partial | CYP71D264 CYP71D.22 |
| 55 | 760 | CYP71D263P | 344 aa | partial; internal-stop; short-sequence | CYP71D263P CYP71D.23P |
| 56 | 776 | CYP71D189a | 499 aa | pseudogene; partial; internal-stop | CYP71D189a CYP71D.24v1 |
| 57 | 792 | CYP71D.24v2 | 287 aa | short-sequence | CYP71D.24v2 |
| 58 | 804 | CYP71D208 | 503 aa |  | CYP71D208 ortholog CYP71D.25 |
| 59 | 822 | CYP71D187 | 503 aa | partial | CYP71D187 ortholog CYP71D.26 |
| 60 | 840 | CYP71D265 | 501 aa | partial; frameshift | CYP71D265 CYP71D.27P |
| 61 | 859 | CYP71D266 | 514 aa |  | CYP71D266 CYP71D.28 |
| 62 | 875 | CYP71D.28v2 | 40 aa | short-sequence | CYP71D.28v2 |
| 63 | 882 | CYP71D266-de2b | 23 aa | short-sequence | CYP71D266-de2b CYP71D.28-de2b |
| 64 | 889 | CYP71D267P | 490 aa |  | CYP71D267P CYP71D.29P |
| 65 | 909 | CYP71D268P | 514 aa | partial; frameshift; internal-stop | CYP71D268P CYP71D.30P |
| 66 | 926 | CYP71D.30Pv2 | 39 aa | short-sequence | CYP71D.30Pv2 |
| 67 | 934 | CYP71D206 | 504 aa |  | CYP71D206 possible ortholog CYP71D.31v1 |
| 68 | 950 | CYP71D.31v2 | 36 aa | short-sequence | CYP71D.31v2 |
| 69 | 958 | CYP71D205 | 521 aa |  | CYP71D205 ortholog CYP71D.32 |
| 70 | 975 | CYP71D209 | 504 aa |  | CYP71D209 ortholog CYP71D.33 |
| 71 | 993 | CYP71D273P | 34 aa | short-sequence | CYP71D273P CYP71D.34P |
| 72 | 1001 | CYP71D272P | 496 aa | partial; frameshift; internal-stop | CYP71D272P CYP71D.35P |
| 73 | 1019 | CYP71D271P | 316 aa | partial; internal-stop; short-sequence | CYP71D271P CYP71D.36P |
| 74 | 1033 | CYP71D270P | 151 aa | short-sequence | CYP71D270P CYP71D.37P |
| 75 | 1042 | CYP71D269P | 553 aa | partial; internal-stop | CYP71D269P CYP71D.38P |
| 76 | 1063 | CYP71D225 | 508 aa | partial; internal-stop | CYP71D225 CYP71D.39 |
| 77 | 1082 | CYP71D213 | 506 aa | partial | CYP71D213 CYP71D.40 |
| 78 | 1100 | CYP71D224 | 495 aa | partial; internal-stop | CYP71D224 CYP71D.41 |
| 79 | 1116 | CYP71D.41v2 | 57 aa | short-sequence | CYP71D.41v2 |
| 80 | 1123 | CYP71D224-de1b | 49 aa | internal-stop; short-sequence | CYP71D224-de1b CYP71D.41-de1b |
| 81 | 1130 | CYP71D212 | 502 aa | frameshift | CYP71D212 CYP71D.42 |
| 82 | 1147 | CYP71D212-de1b | 37 aa | internal-stop; short-sequence | CYP71D212-de1b CYP71D.42-de1b |
| 83 | 1154 | CYP71D211 | 494 aa | partial; frameshift | CYP71D211 CYP71D.43v1 |
| 84 | 1172 | CYP71D.43v2 | 33 aa | short-sequence | CYP71D.43v2 |
| 85 | 1178 | CYP71D.43v3 | 40 aa | short-sequence | CYP71D.43v3 |
| 86 | 1186 | CYP71D210 | 499 aa | pseudogene; partial; frameshift | CYP71D210 CYP71D.44 |
| 87 | 1204 | CYP71D274 | 505 aa | partial | CYP71D274 CYP71D.45 |
| 88 | 1222 | CYP71D204a | 507 aa |  | CYP71D204a ortholog CYP71D.46 |
| 89 | 1240 | CYP71D220bP | 543 aa | partial; frameshift | CYP71D220bP CYP71D.47Pv1 |
| 90 | 1258 | CYP71D.47Pv2 | 72 aa | short-sequence | CYP71D.47Pv2 |
| 91 | 1267 | CYP71D220a | 509 aa |  | CYP71D220a = OLD CYP71D231 possible ortholog CYP71D.48v1 |
| 92 | 1284 | CYP71D.48v2 | 84 aa | short-sequence | CYP71D.48v2 |
| 93 | 1291 | CYP71D.48v3 | 57 aa | short-sequence | CYP71D.48v3 |
| 94 | 1300 | CYP71D230P | 202 aa | short-sequence | CYP71D230P CYP71D.49P |
| 95 | 1312 | CYP71D219 | 510 aa | partial; frameshift | CYP71D219 CYP71D.50 |
| 96 | 1330 | CYP71D229P | 365 aa | partial; frameshift; internal-stop | CYP71D229P CYP71D.51P |
| 97 | 1348 | CYP71D228P | 221 aa | internal-stop; short-sequence | CYP71D228P CYP71D.52P |
| 98 | 1360 | CYP71D218 | 510 aa | partial | CYP71D218 CYP71D.53 |
| 99 | 1378 | CYP71D227P | 520 aa | partial; internal-stop | CYP71D227P CYP71D.54P |
| 100 | 1397 | CYP71D217P | 355 aa | partial; frameshift; internal-stop | CYP71D217P CYP71D.55P |
| 101 | 1414 | CYP71D226 | 510 aa | partial | CYP71D226 CYP71D.56v1 |
| 102 | 1430 | CYP71D.56v2 | 303 aa | short-sequence | CYP71D.56v2 |
| 103 | 1443 | CYP71D216 | 510 aa | partial | CYP71D216 POSSIBLE ortholog CYP71D.57 |
| 104 | 1461 | CYP71D.58P | 203 aa | short-sequence | CYP71D.58P |
| 105 | 1472 | CYP71D276 | 505 aa |  | CYP71D276 CYP71D.59v1 |
| 106 | 1488 | CYP71D.59v2 | 82 aa | short-sequence | CYP71D.59v2 |
| 107 | 1497 | CYP71D282P | 309 aa | internal-stop; short-sequence | CYP71D282P CYP71D.60P |
| 108 | 1514 | CYP71D281P | 62 aa | internal-stop; short-sequence | CYP71D281P CYP71D.61P |
| 109 | 1522 | CYP71D280P | 45 aa | short-sequence | CYP71D280P CYP71D.62P |
| 110 | 1531 | CYP71D279P | 487 aa |  | CYP71D279P CYP71D.63P |
| 111 | 1549 | CYP71D278P | 95 aa | short-sequence | CYP71D278P CYP71D.64P |
| 112 | 1559 | CYP71D277P | 87 aa | short-sequence | CYP71D277P CYP71D.65P |
| 113 | 1568 | CYP71D221 | 417 aa | pseudogene; partial | CYP71D221 CYP71D.66 |
| 114 | 1588 | CYP71D248P | 446 aa | partial; internal-stop | CYP71D248P CYP71D.67P |
| 115 | 1606 | CYP71D283P | 181 aa | internal-stop; short-sequence | CYP71D283P CYP71D.68P |
| 116 | 1619 | CYP71D284P | 173 aa | internal-stop; short-sequence | CYP71D284P CYP71D.69P |
| 117 | 1630 | CYP71D247 | 410 aa | partial | CYP71D247 CYP71D.70 |
| 118 | 1647 | CYP71D204b | 201 aa | sequence-gap; short-sequence | CYP71D204b CYP71D.71v1 |
| 119 | 1658 | CYP71D.71v2 | 37 aa | short-sequence | CYP71D.71v2 |
| 120 | 1664 | CYP71D.71v3 | 37 aa | short-sequence | CYP71D.71v3 |
| 121 | 1670 | CYP71D.71v4 | 34 aa | short-sequence | CYP71D.71v4 |
| 122 | 1678 | CYP71D200 | 498 aa |  | CYP71D200 ortholog CYP71D.72 |
| 123 | 1697 | CYP71D207 | 503 aa |  | CYP71D207 ortholog CYP71D.73 |
| 124 | 1715 | CYP71D284 | 505 aa | frameshift; ambiguous-residue | CYP71D284 CYP71D.74 |
| 125 | 1734 | CYP71BP1 | 502 aa |  | CYP71BP1 CYP71D.75 |
| 126 | 1752 | CYP71D201 | 496 aa | partial; frameshift | CYP71D201 CYP71D.76 |
| 127 | 1771 | CYP71D285P | 161 aa | short-sequence | CYP71D285P CYP71D.77P |
| 128 | 1781 | CYP71D286P | 159 aa | short-sequence | CYP71D286P CYP71D.78P |
| 129 | 1791 | CYP71D287 | 229 aa | short-sequence | CYP71D287 CYP71D.79 |
| 130 | 1803 | CYP71D288P | 485 aa | internal-stop | CYP71D288P CYP71D.80P |
| 131 | 1822 | CYP71D289P | 183 aa | short-sequence | CYP71D289P CYP71D.81P |
| 132 | 1834 | CYP71D290P | 52 aa | pseudogene; short-sequence | CYP71D290P CYP71D.82P |
| 133 | 1844 | CYP71D291P | 52 aa | short-sequence | CYP71D291P CYP71D.83P |
| 134 | 1852 | CYP71D292P | 52 aa | short-sequence | CYP71D292P CYP71D.84P |
| 135 | 1860 | CYP71D293P | 52 aa | short-sequence | CYP71D293P CYP71D.85P |
| 136 | 1868 | CYP71D294P | 52 aa | short-sequence | CYP71D294P CYP71D.86P |
| 137 | 1876 | CYP71D295P | 52 aa | short-sequence | CYP71D295P CYP71D.87P |
| 138 | 1884 | CYP71D296P | 52 aa | short-sequence | CYP71D296P CYP71D.88P |
| 139 | 1892 | CYP71D297P | 52 aa | short-sequence | CYP71D297P CYP71D.89P |
| 140 | 1900 | CYP71D298P | 52 aa | short-sequence | CYP71D298P CYP71D.90P |
| 141 | 1908 | CYP71D299P | 52 aa | short-sequence | CYP71D299P CYP71D.91P |
| 142 | 1916 | CYP71D300P | 52 aa | short-sequence | CYP71D300P CYP71D.92P |
| 143 | 1924 | CYP71D301P | 52 aa | short-sequence | CYP71D301P CYP71D.93P |
| 144 | 1932 | CYP71D302P | 52 aa | short-sequence | CYP71D302P CYP71D.94P |
| 145 | 1940 | CYP71D303P | 52 aa | short-sequence | CYP71D303P CYP71D.95P |
| 146 | 1948 | CYP71D304P | 52 aa | short-sequence | CYP71D304P CYP71D.96P |
| 147 | 1956 | CYP71D305P | 52 aa | short-sequence | CYP71D305P CYP71D.97P |
| 148 | 1964 | CYP71D306P | 70 aa | short-sequence | CYP71D306P CYP71D.98P |
| 149 | 1973 | CYP71D307P | 83 aa | short-sequence | CYP71D307P CYP71D.99P |
| 150 | 1982 | CYP71D202P | 213 aa | short-sequence | CYP71D202P CYP71D.100P |
| 151 | 1994 | CYP71D308P | 73 aa | short-sequence | CYP71D308P CYP71D.101P |
| 152 | 2003 | CYP71D309P | 126 aa | internal-stop; short-sequence | CYP71D309P CYP71D.102P |
| 153 | 2013 | CYP71D310P | 58 aa | short-sequence | CYP71D310P CYP71D.103 |
| 154 | 2022 | CYP71AH9 | 495 aa |  | CYP71AH9 ortholog CYP71AH.1 |
| 155 | 2041 | CYP71AT30 | 498 aa | partial; frameshift | CYP71AT30 CYP71AT.1 |
| 156 | 2058 | CYP71AT31 | 498 aa | partial | CYP71AT31 CYP71AT.2 |
| 157 | 2075 | CYP71AT32P | 432 aa | internal-stop | CYP71AT32P CYP71AT.3Pv1 |
| 158 | 2091 | CYP71AT.3Pv2 | 36 aa | short-sequence | CYP71AT.3Pv2 |
| 159 | 2099 | CYP71AT33 | 495 aa | partial; frameshift | CYP71AT33 CYP71AT.4v1 |
| 160 | 2115 | CYP71AT.4v2 | 40 aa | short-sequence | CYP71AT.4v2 |
| 161 | 2123 | CYP71AT13a | 497 aa |  | CYP71AT13a CYP71AT.5v1 |
| 162 | 2138 | CYP71AT.5v2 | 33 aa | short-sequence | CYP71AT.5v2 |
| 163 | 2144 | CYP71AT.5v3 | 34 aa | short-sequence | CYP71AT.5v3 |
| 164 | 2152 | CYP71AT34 | 499 aa |  | CYP71AT34 CYP71AT.6v1 |
| 165 | 2167 | CYP71AT.6v2 | 82 aa | short-sequence | CYP71AT.6v2 |
| 166 | 2174 | CYP71AT.6v3 | 58 aa | short-sequence | CYP71AT.6v3 |
| 167 | 2180 | CYP71AT.6v4 | 41 aa | short-sequence | CYP71AT.6v4 |
| 168 | 2186 | CYP71AT.6v5 | 56 aa | short-sequence | CYP71AT.6v5 |
| 169 | 2194 | CYP71AT36 | 495 aa | frameshift | CYP71AT36 CYP71AT.7 |
| 171 | 2225 | CYP71AT37Pa | 296 aa | internal-stop; short-sequence | CYP71AT37Pa CYP71AT.8P |
| 172 | 2238 | CYP71AT38Pa | 137 aa | internal-stop; short-sequence | CYP71AT38Pa CYP71AT.9Pa |
| 173 | 2246 | CYP71AT39P | 236 aa | short-sequence | CYP71AT39P CYP71AT.9Pb |
| 174 | 2258 | CYP71AT13b | 498 aa |  | CYP71AT13b CYP71AT.10 |
| 175 | 2276 | CYP71AT17a | 495 aa | partial | CYP71AT17a CYP71AT.11v1 |
| 176 | 2291 | CYP71AT.11v2 | 72 aa | short-sequence | CYP71AT.11v2 |
| 177 | 2298 | CYP71AT.11v3 | 150 aa | short-sequence | CYP71AT.11v3 |
| 178 | 2306 | CYP71AT.11v4 | 49 aa | short-sequence | CYP71AT.11v4 |
| 179 | 2312 | CYP71AT.11v5 | 48 aa | short-sequence | CYP71AT.11v5 |
| 180 | 2318 | CYP71AT.1v6 | 38 aa | short-sequence | CYP71AT.1v6 |
| 181 | 2326 | CYP71AT17a-de2b | 204 aa | short-sequence | CYP71AT17a-de2b CYP71AT.12Pv1 |
| 182 | 2335 | CYP71AT.12Pv2 | 44 aa | short-sequence | CYP71AT.12Pv2 |
| 183 | 2341 | CYP71AT.12Pv3 | 63 aa | internal-stop; short-sequence | CYP71AT.12Pv3 |
| 184 | 2350 | CYP71AT41Pv1 | 266 aa | internal-stop; short-sequence | CYP71AT41Pv1 CYP71AT.13Pv1 |
| 185 | 2361 | CYP71AT.13Pv2 | 40 aa | short-sequence | CYP71AT.13Pv2 |
| 186 | 2367 | CYP71AT.13Pv3 | 33 aa | short-sequence | CYP71AT.13Pv3 |
| 187 | 2375 | CYP71AT15v1 | 491 aa | partial; internal-stop | CYP71AT15v1 CYP71AT.14 |
| 188 | 2393 | CYP71AT42Pv1 | 463 aa | internal-stop | CYP71AT42Pv1 CYP71AT.15Pv1 |
| 189 | 2409 | CYP71AT.15Pv2 | 80 aa | short-sequence | CYP71AT.15Pv2 |
| 190 | 2415 | CYP71AT.15Pv3 | 34 aa | short-sequence | CYP71AT.15Pv3 |
| 191 | 2421 | CYP71AT.15Pv4 | 42 aa | short-sequence | CYP71AT.15Pv4 |
| 192 | 2429 | CYP71AT16v1 | 321 aa | partial; short-sequence | CYP71AT16v1 ortholog CYP71AT.16v1 |
| 193 | 2442 | CYP71AT.16v2 | 39 aa | short-sequence | CYP71AT.16v2 |
| 194 | 2450 | CYP71AT37Pb | 79 aa | internal-stop; short-sequence | CYP71AT37Pb CYP71AT.17P |
| 195 | 2458 | CYP71AT38Pb | 117 aa | short-sequence | CYP71AT38Pb CYP71AT.19P |
| 196 | 2464 | CYP71AT.18P | 28 aa | internal-stop; short-sequence | CYP71AT.18P |
| 197 | 2473 | CYP71AT17b | 495 aa |  | CYP71AT17b CYP71AT.20 |
| 198 | 2490 | CYP71AT1a | 495 aa | partial | CYP71AT1a CYP71AT.21 |
| 199 | 2508 | CYP71AT20a | 496 aa | partial; frameshift | CYP71AT20a CYP71AT.22 |
| 200 | 2526 | CYP71AT21 | 491 aa |  | CYP71AT21 possible ortholog CYP71AT.23 |
| 201 | 2543 | CYP71AT22 | 497 aa |  | CYP71AT22 CYP71AT.24 |
| 202 | 2560 | CYP71AT43 | 496 aa |  | CYP71AT43 CYP71AT.25 |
| 203 | 2577 | CYP71AT23 | 496 aa |  | CYP71AT23 possible ortholog CYP71AT.26 |
| 204 | 2594 | CYP71AT1b | 496 aa |  | CYP71AT1b CYP71AT.27 |
| 205 | 2614 | CYP71AT20bP | 293 aa | short-sequence | CYP71AT20bP CYP71AT.28P |
| 206 | 2629 | CYP71AT25P | 151 aa | short-sequence | CYP71AT25P CYP71AT.29P |
| 207 | 2640 | CYP71AT16v2 | 488 aa |  | CYP71AT16v2 CYP71AT.30P |
| 208 | 2659 | CYP71AT42v2 | 498 aa |  | CYP71AT42v2 CYP71AT.31v1 |
| 209 | 2676 | CYP71AT.31v2 | 40 aa | short-sequence | CYP71AT.31v2 |
| 210 | 2684 | CYP71AT15v2 | 494 aa |  | CYP71AT15v2 CYP71AT.32v1 |
| 211 | 2700 | CYP71AT.32v2 | 41 aa | short-sequence | CYP71AT.32v2 |
| 212 | 2708 | CYP71AT41v2 | 497 aa |  | CYP71AT41v2 CYP71AT.33v1 |
| 213 | 2724 | CYP71AT.33v2 | 31 aa | short-sequence | CYP71AT.33v2 |
| 214 | 2730 | CYP71AT.33 | 82 aa | partial; short-sequence | CYP71AT.33 |
| 215 | 2742 | CYP71AT24 | 503 aa |  | CYP71AT24 ortholog CYP71AT.34 |
| 216 | 2760 | CYP71AT49P | 488 aa | internal-stop | CYP71AT49P CYP71AT.35P |
| 217 | 2778 | CYP71AT49P-de1b | 98 aa | internal-stop; short-sequence | CYP71AT49P-de1b CYP71AT.36P |
| 218 | 2788 | CYP71AT50P | 183 aa | internal-stop; short-sequence | CYP71AT50P CYP71AT.37P |
| 219 | 2798 | CYP71AT27a | 458 aa | partial | CYP71AT27a CYP71AT.38 |
| 220 | 2817 | CYP71AT27b | 497 aa | frameshift | CYP71AT27b CYP71AT.39 |
| 221 | 2835 | CYP71AT27b-de2b | 131 aa | short-sequence | CYP71AT27b-de2b CYP71AT.40P |
| 222 | 2846 | CYP71AT27b-de2c | 202 aa | short-sequence | CYP71AT27b-de2c CYP71AT.41P |
| 223 | 2857 | CYP71AT46 | 477 aa | partial | CYP71AT46 CYP71AT.42 |
| 224 | 2876 | CYP71AT45P | 154 aa | internal-stop; short-sequence | CYP71AT45P CYP71AT.43Pv1 |
| 225 | 2885 | CYP71AT.43Pv2 | 38 aa | short-sequence | CYP71AT.43Pv2 |
| 226 | 2893 | CYP71AT44P | 144 aa | short-sequence | CYP71AT44P CYP71AT.44P |
| 227 | 2903 | CYP71AT51P | 499 aa |  | CYP71AT51P CYP71AT.45P |
| 228 | 2922 | CYP71AT47 | 498 aa | partial | CYP71AT47 CYP71AT.46 |
| 229 | 2941 | CYP71AT48 | 454 aa | partial; frameshift | CYP71AT48 CYP71AT.47 |
| 230 | 2960 | CYP71AT52P | 45 aa | short-sequence | CYP71AT52P CYP71AT.48P |
| 231 | 2968 | CYP71AT53P | 91 aa | short-sequence | CYP71AT53P CYP71AT.49P |
| 232 | 2978 | CYP71AT28P | 77 aa | internal-stop; short-sequence | CYP71AT28P CYP71AT.50P |
| 233 | 2987 | CYP71AT29P | 23 aa | short-sequence | CYP71AT29P CYP71AT.51P |
| 234 | 2995 | CYP71AT35P | 50 aa | short-sequence | CYP71AT35P CYP71AT.52P |
| 235 | 3003 | CYP71AT40P | 21 aa | short-sequence | CYP71AT40P CYP71AT.53P |
| 236 | 3011 | CYP71AT54P | 48 aa | short-sequence | CYP71AT54P CYP71AT.54P |
| 237 | 3020 | CYP71AT55P | 62 aa | short-sequence | CYP71AT55P CYP71AT.55P |
| 238 | 3028 | CYP71AU32 | 502 aa | partial; internal-stop | CYP71AU32 ortholog CYP71AU.1 syntenic |
| 239 | 3045 | CYP71AU2 | 500 aa |  | CYP71AU2 ortholog CYP71AU.2 syntenic |
| 240 | 3062 | CYP71AU33 | 518 aa |  | CYP71AU33 ortholog CYP71AU.3 |
| 241 | 3080 | CYP71AX14P | 118 aa | partial; short-sequence | CYP71AX14P CYP71AX.1P not in tomato |
| 242 | 3089 | CYP71AX15 | 502 aa | partial; frameshift; ambiguous-residue | CYP71AX15 CYP71AX.2 not in tomato |
| 243 | 3106 | CYP71AX.2v2 | 47 aa | short-sequence | CYP71AX.2v2 |
| 244 | 3112 | CYP71AX.2v3 | 53 aa | short-sequence | CYP71AX.2v3 |
| 245 | 3118 | CYP71AX.2v4 | 42 aa | partial; short-sequence | CYP71AX.2v4 |
| 246 | 3126 | CYP71AX16 | 496 aa |  | CYP71AX16 CYP71AX.3v1 not in tomato |
| 247 | 3142 | CYP71AX.3v2 | 48 aa | short-sequence | CYP71AX.3v2 |
| 248 | 3150 | CYP71AX17P | 381 aa |  | CYP71AX17P CYP71AX.4P not in tomato |
| 249 | 3166 | CYP71AX18P | 417 aa | partial; frameshift | CYP71AX18P CYP71AX.5Pv1 not in tomato |
| 250 | 3181 | CYP71AX.5Pv2 | 45 aa | short-sequence | CYP71AX.5Pv2 |
| 251 | 3189 | CYP71AX19P | 133 aa | short-sequence | CYP71AX19P CYP71AX.6Pv1 not in tomato |
| 252 | 3197 | CYP71AX.6Pv2 | 65 aa | short-sequence | CYP71AX.6Pv2 |
| 253 | 3205 | CYP71AX20 | 500 aa | partial | CYP71AX20 CYP71AX.7 not in tomato |
| 254 | 3222 | CYP71AX2 | 503 aa | internal-stop | CYP71AX2 ortholog CYP71AX.8 syntenic |
| 255 | 3239 | CYP71AX3 | 498 aa | partial | CYP71AX3 possible ortholog CYP71AX.9 |
| 256 | 3257 | CYP71AX4 | 499 aa | partial | CYP71AX4 CYP71AX.10 |
| 257 | 3275 | CYP71AX7 | 502 aa |  | CYP71AX7 ortholog CYP71AX.11 syntenic |
| 258 | 3293 | CYP71AX8P | 512 aa | pseudogene; partial | CYP71AX8P ortholog CYP71AX.12 |
| 259 | 3313 | CYP71AX21 | 501 aa |  | CYP71AX21 CYP71AX.13 not in tomato |
| 260 | 3331 | CYP71AX9 | 495 aa |  | CYP71AX9 possible ortholog CYP71AX.14 |
| 261 | 3349 | CYP71AX22P | 493 aa | pseudogene; partial; frameshift | CYP71AX22P CYP71AX.15P not in tomato |
| 262 | 3367 | CYP71AX23 | 503 aa |  | CYP71AX23 CYP71AX.16 not in tomato |
| 263 | 3385 | CYP71AX24P | 491 aa | internal-stop | CYP71AX24P CYP71AX.17P not in tomato |
| 264 | 3403 | CYP71AX10 | 499 aa |  | CYP71AX10 ortholog CYP71AX.18 syntenic |
| 265 | 3421 | CYP71AX11 | 496 aa |  | CYP71AX11 possible ortholog CYP71AX.19 |
| 266 | 3438 | CYP71AX25 | 497 aa | partial | CYP71AX25 CYP71AX.20 not in tomato |
| 267 | 3456 | CYP71AX12 | 497 aa | partial | CYP71AX12 possible ortholog CYP71AX.21 |
| 268 | 3473 | CYP71AX13 | 506 aa |  | CYP71AX13 possible ortholog CYP71AX.22 |
| 269 | 3490 | CYP71AX26 | 481 aa | partial | CYP71AX26 CYP71AX.23 |
| 270 | 3507 | CYP71AX26-de1b | 68 aa | short-sequence | CYP71AX26-de1b CYP71AX.23-de1b |
| 271 | 3515 | CYP71AX27P | 398 aa | internal-stop | CYP71AX27P CYP71AX.24P |
| 272 | 3532 | CYP71AX28P | 85 aa | internal-stop; short-sequence | CYP71AX28P CYP71AX.25P |
| 273 | 3542 | CYP71AX29P | 97 aa | internal-stop; short-sequence | CYP71AX29P CYP71AX.26P |
| 274 | 3552 | CYP71BE17 | 501 aa |  | CYP71BE17 ortholog CYP71BE.1 |
| 275 | 3570 | CYP71BE.2P | 108 aa | internal-stop; short-sequence | CYP71BE.2P |
| 276 | 3583 | CYP71BL1 | 503 aa |  | CYP71BL1 ortholog CYP71BE.3 |
| 277 | 3603 | CYP71BE.4P | 506 aa | partial; uncertain-boundary | CYP71BE.4P |
| 278 | 3623 | CYP71BE.5P | 213 aa | internal-stop; short-sequence | CYP71BE.5P |
| 279 | 3635 | CYP71BE18 | 500 aa |  | CYP71BE18 ortholog CYP71BE.6 |
| 280 | 3654 | CYP71BM1 | 489 aa | partial | CYP71BM1 CYP71BE.7 |
| 281 | 3673 | CYP71BG1 | 510 aa |  | CYP71BG1 ortholog of CYP71BG2 |
| 282 | 3692 | CYP71BG.2P | 80 aa | internal-stop; short-sequence | CYP71BG.2P |
| 283 | 3701 | CYP71BP3 | 497 aa |  | CYP71BP3 ortholog CYP71x.1v1 probable CYP71D member |
| 284 | 3720 | CYP71x.1v2 | 207 aa | short-sequence | CYP71x.1v2 |
| 285 | 3732 | CYP71BN1P | 480 aa | partial; internal-stop | CYP71BN1P CYP71y.1P |
| 286 | 3753 | CYP71z.1P | 137 aa | partial; internal-stop; short-sequence | CYP71z.1P |
| 287 | 3766 | CYP72A182 | 495 aa |  | CYP72A182 POSSIBLE ortholog CYP72A.1 |
| 288 | 3784 | CYP72A182-ie1b | 21 aa | short-sequence | CYP72A182-ie1b CYP72A.1-ie1b |
| 289 | 3791 | CYP72A218P | 489 aa | partial; uncertain-boundary; internal-stop | CYP72A218P CYP72A.2P |
| 290 | 3809 | CYP72A218P-de1b | 56 aa | internal-stop; short-sequence | CYP72A218P-de1b CYP72A.2P-de1b |
| 291 | 3816 | CYP72A181 | 492 aa |  | CYP72A181 POSSIBLE ortholog CYP72A.3 |
| 292 | 3836 | CYP72A180 | 492 aa |  | CYP72A180 POSSIBLE ortholog CYP72A.4 |
| 293 | 3856 | CYP72A217 | 489 aa |  | CYP72A217 CYP72A.5 |
| 294 | 3877 | CYP72A216P | 492 aa | frameshift | CYP72A216P CYP72A.6 |
| 295 | 3898 | CYP72A179 | 492 aa | partial | CYP72A179 POSSIBLE ortholog CYP72A.7v1 |
| 296 | 3918 | CYP72A.7v2 | 34 aa | short-sequence | CYP72A.7v2 |
| 297 | 3926 | CYP72A178 | 509 aa | partial; frameshift | CYP72A178 POSSIBLE ortholog CYP72A.8 |
| 298 | 3947 | CYP72A177 | 505 aa |  | CYP72A177 POSSIBLE ortholog CYP72A.9 |
| 299 | 3966 | CYP72A176cP | 461 aa |  | CYP72A176cP CYP72A.10Pv1 |
| 300 | 3985 | CYP72A.10Pv2 | 40 aa | short-sequence | CYP72A.10Pv2 |
| 301 | 3993 | CYP72A176bP | 396 aa | partial | CYP72A176bP CYP72A.11P |
| 302 | 4012 | CYP72A176aP | 379 aa | partial | CYP72A176aP CYP72A.12Pv1 |
| 303 | 4031 | CYP72A.12Pv2 | 40 aa | short-sequence | CYP72A.12Pv2 |
| 304 | 4039 | CYP72A175P | 489 aa | partial; internal-stop | CYP72A175P CYP72A.13P |
| 305 | 4060 | CYP72A175P-de1b | 56 aa | internal-stop; short-sequence | CYP72A175P-de1b CYP72A.13-de1b |
| 306 | 4067 | CYP72A215 | 487 aa | frameshift | CYP72A215 CYP72A.14v1 |
| 307 | 4089 | CYP72A215-de4b | 24 aa | short-sequence | CYP72A215-de4b CYP72A.14-de4b |
| 308 | 4096 | CYP72A.14v2 | 47 aa | short-sequence | CYP72A.14v2 |
| 309 | 4104 | CYP72A174P | 473 aa |  | CYP72A174P ortholog CYP72A.15P |
| 310 | 4125 | CYP72A173 | 454 aa | partial; frameshift | CYP72A173 POSSIBLE ortholog CYP72A.16v1 |
| 311 | 4145 | CYP72A.16v2 | 72 aa | short-sequence | CYP72A.16v2 |
| 312 | 4154 | CYP72A172 | 496 aa |  | CYP72A172 ortholog CYP72A.17v1 |
| 313 | 4174 | CYP72A172-de1b2b | 167 aa | short-sequence | CYP72A172-de1b2b CYP72A.17-de1b2b |
| 314 | 4182 | CYP72A172-de1c2c | 167 aa | short-sequence | CYP72A172-de1c2c CYP72A.17-de1c2c |
| 315 | 4192 | CYP72A.17v2 | 60 aa | partial; short-sequence | CYP72A.17v2 fragment |
| 316 | 4200 | CYP72A.17v3 | 29 aa | short-sequence | CYP72A.17v3 |
| 317 | 4208 | CYP72A171 | 495 aa |  | CYP72A171 POSSIBLE ortholog CYP72A.18 |
| 318 | 4228 | CYP72A214P | 501 aa | partial; uncertain-boundary | CYP72A214P CYP72A.19P |
| 319 | 4249 | CYP72A.20P | 453 aa | uncertain-boundary | CYP72A.20P |
| 320 | 4269 | CYP72A.21P | 517 aa | frameshift | CYP72A.21P |
| 321 | 4292 | CYP72A.22Pv1 | 491 aa | frameshift | CYP72A.22Pv1 |
| 322 | 4314 | CYP72A.22Pv2 | 166 aa | short-sequence | CYP72A.22Pv2 |
| 323 | 4326 | CYP72A.23Pv1 | 408 aa | partial; uncertain-boundary | CYP72A.23Pv1 |
| 324 | 4347 | CYP72A.23Pv2 | 46 aa | short-sequence | CYP72A.23Pv2 |
| 325 | 4354 | CYP72A.23Pv3 | 41 aa | short-sequence | CYP72A.23Pv3 |
| 326 | 4362 | CYP72A.23Pv4 | 30 aa | short-sequence | CYP72A.23Pv4 |
| 327 | 4370 | CYP72A204 | 505 aa |  | CYP72A204 CYP72A.24 |
| 328 | 4392 | CYP72A186 | 494 aa |  | CYP72A186 POSSIBLE ortholog CYP72A.25 |
| 329 | 4412 | CYP72A185 | 494 aa |  | CYP72A185 POSSIBLE ortholog CYP72A.26 |
| 330 | 4433 | CYP72A213 | 491 aa | frameshift | CYP72A213 CYP72A.27 |
| 331 | 4453 | CYP72A212P | 478 aa | uncertain-boundary | CYP72A212P CYP72A.28Pv1 |
| 332 | 4474 | CYP72A.28Pv2 | 30 aa | short-sequence | CYP72A.28Pv2 |
| 333 | 4480 | CYP72A.28Pv3 | 80 aa | short-sequence | CYP72A.28Pv3 |
| 334 | 4489 | CYP72A211P | 235 aa | internal-stop; short-sequence | CYP72A211P CYP72A.29P |
| 335 | 4503 | CYP72A210P | 361 aa | internal-stop | CYP72A210P CYP72A.30P |
| 336 | 4519 | CYP72A209 | 494 aa |  | CYP72A209 CYP72A.31v1 |
| 337 | 4539 | CYP72A209-de1b2b | 148 aa | short-sequence | CYP72A209-de1b2b CYP72A.31-de1b2b |
| 338 | 4550 | CYP72A209-de1c | 50 aa | short-sequence | CYP72A209-de1c CYP72A.31-de1c |
| 339 | 4558 | CYP72A.31v2 | 49 aa | short-sequence | CYP72A.31v2 |
| 340 | 4565 | CYP72A184b | 495 aa |  | CYP72A184b CYP72A.32v1 |
| 341 | 4585 | CYP72A.32v2 | 60 aa | short-sequence | CYP72A.32v2 |
| 342 | 4593 | CYP72A.32v3 | 35 aa | short-sequence | CYP72A.32v3 |
| 343 | 4601 | CYP72A184a | 461 aa |  | CYP72A184a CYP72A.33 |
| 344 | 4622 | CYP72A183 | 475 aa |  | CYP72A183 ortholog CYP72A.34 |
| 345 | 4642 | CYP72A.x | missing | sequence-missing | CYP72A.x scaffold |
| 346 | 4645 | CYP72A.x | 477 aa |  | CYP72A.x = 100% to CYP72A.35a, 35b, 35c |
| 347 | 4661 | CYP72A208a | 474 aa |  | CYP72A208a ortholog CYP72A.35av1 |
| 348 | 4682 | CYP72A.35av2 | 465 aa | partial | CYP72A.35av2 |
| 349 | 4702 | CYP72A208b | 469 aa |  | CYP72A208b CYP72A.35b |
| 350 | 4724 | CYP72A193 | 485 aa |  | CYP72A193 ortholog CYP72A.36 |
| 351 | 4742 | CYP72A193-de1b | 23 aa | short-sequence | CYP72A193-de1b CYP72A.36-de1b |
| 352 | 4749 | CYP72A194a | 495 aa | pseudogene; frameshift | CYP72A194a POSSIBLE ortholog CYP72A.37 |
| 353 | 4770 | CYP72A194a-de4b | 40 aa | short-sequence | CYP72A194a-de4b CYP72A.37-de4b |
| 354 | 4777 | CYP72A195Pa | 492 aa | partial | CYP72A195Pa CYP72A.38P |
| 355 | 4797 | CYP72A.38Pv2 | 79 aa | short-sequence | CYP72A.38Pv2 |
| 356 | 4806 | CYP72A194Pb | 270 aa | short-sequence | CYP72A194Pb CYP72A.39P |
| 357 | 4819 | CYP72A195b | 482 aa |  | CYP72A195b CYP72A.40v1 |
| 358 | 4838 | CYP72A.40v2 | 55 aa | short-sequence | CYP72A.40v2 |
| 359 | 4846 | CYP72A188 | 496 aa |  | CYP72A188 ortholog CYP72A.41 |
| 360 | 4867 | CYP72A189 | 497 aa |  | CYP72A189 CYP72A.42 |
| 361 | 4887 | CYP72A190 | 495 aa |  | CYP72A190 POSSIBLE ortholog CYP72A.43 |
| 362 | 4907 | CYP72A191P | 520 aa | pseudogene; frameshift | CYP72A191P CYP72A.44P |
| 363 | 4927 | CYP72A192P | 492 aa | frameshift | CYP72A192P CYP72A.45P |
| 364 | 4947 | CYP72A187 | 489 aa |  | CYP72A187 POSSIBLE ortholog CYP72A.46 |
| 365 | 4965 | CYP72A51 | 487 aa |  | CYP72A51 POSSIBLE ortholog CYP72A.47 |
| 366 | 4985 | CYP72D8 | 489 aa |  | CYP72D8 ortholog CYP72A.48 |
| 367 | 5006 | CYP72A.49P | 57 aa | short-sequence | CYP72A.49P |
| 368 | 5015 | CYP72A.50P | 53 aa | internal-stop; short-sequence | CYP72A.50P |
| 369 | 5023 | CYP72A.51P | 31 aa | short-sequence | CYP72A.51P |
| 370 | 5031 | CYP72A.52P | 42 aa | short-sequence | CYP72A.52P |
| 371 | 5039 | CYP72A.53P | 41 aa | short-sequence | CYP72A.53P |
| 372 | 5047 | CYP73A96 | 471 aa |  | CYP73A96 ortholog CYP73A.3 |
| 373 | 5065 | CYP73A24c | 505 aa |  | CYP73A24c CYP73A66v1   Solanum tuberosum  (potato, Solanales) |
| 374 | 5081 | CYP73A24c | 483 aa |  | CYP73A24c CYP73A.1 = CYP73A66a(v1) (1 aa diff) |
| 375 | 5097 | CYP73A24d | 476 aa | pseudogene; frameshift | CYP73A24d CYP73A66b(v2)P |
| 376 | 5115 | CYP74A1 | 530 aa |  | CYP74A1     Solanum tuberosum (potato) |
| 377 | 5133 | CYP74A1 | 530 aa |  | CYP74A1     Solanum tuberosum (potato) |
| 378 | 5150 | CYP74A6v1 | 508 aa |  | CYP74A6v1   Solanum tuberosum (potato) |
| 379 | 5175 | CYP74A6v1 | 510 aa |  | CYP74A6v1   Solanum tuberosum (potato) |
| 380 | 5192 | CYP74A6v2 | 507 aa |  | CYP74A6v2   Solanum tuberosum (potato) |
| 381 | 5206 | CYP74B3 | 476 aa |  | CYP74B3 ortholog CYP74B.1 |
| 382 | 5223 | CYP74C10v1 | 491 aa |  | CYP74C10v1  Solanum tuberosum (potato) CYP74C3 ortholog |
| 383 | 5245 | CYP74C10v2 | 491 aa |  | CYP74C10v2   Solanum tuberosum (potato) |
| 384 | 5266 | CYP74C10v3 | 267 aa | short-sequence | CYP74C10v3   Solanum tuberosum (potato) |
| 385 | 5277 | CYP74C10v4 | 491 aa | partial | CYP74C10v4 |
| 386 | 5295 | CYP74C.1P | 317 aa | internal-stop; short-sequence | CYP74C.1P |
| 387 | 5311 | CYP74C.2P | 49 aa | short-sequence | CYP74C.2P |
| 388 | 5319 | CYP74C4 | 383 aa |  | CYP74C4 CYP74C.3 |
| 389 | 5337 | CYP74C.4P | 178 aa | internal-stop; short-sequence | CYP74C.4P |
| 390 | 5349 | CYP74D2 | 478 aa |  | CYP74D2      Solanum tuberosum (potato) |
| 391 | 5369 | CYP74D2 | 417 aa |  | CYP74D2      Solanum tuberosum (potato) |
| 392 | 5386 | CYP75A18v1 | 509 aa |  | CYP75A18v1  Solanum tuberosum (potato) |
| 393 | 5404 | CYP75A18v2 | 508 aa |  | CYP75A18v2  Solanum tuberosum (potato) |
| 394 | 5426 | CYP75A18v3 | 308 aa | short-sequence | CYP75A18v3    Solanum tuberosum |
| 395 | 5438 | CYP75A18P | 492 aa | pseudogene; partial | CYP75A18P |
| 396 | 5459 | CYP75B49 | 514 aa |  | CYP75B49 ortholog |
| 397 | 5479 | CYP76A22 | 507 aa | partial; internal-stop | CYP76A22 ortholog CYP76A.1 |
| 398 | 5496 | CYP76A23 | 506 aa |  | CYP76A23 CYP76A.2 |
| 399 | 5513 | CYP76A24 | 515 aa |  | CYP76A24 CYP76A.3 |
| 400 | 5531 | CYP76A21 | 500 aa |  | CYP76A21 possible ortholog CYP76A.4 |
| 401 | 5549 | CYP76A25 | 501 aa |  | CYP76A25 CYP76A.5 |
| 402 | 5567 | CYP76A6 | 501 aa |  | CYP76A6 possible ortholog CYP76A.6 |
| 403 | 5585 | CYP76A.7P | 140 aa | short-sequence | CYP76A.7P |
| 404 | 5595 | CYP76A19 | 500 aa |  | CYP76A19 possible ortholog CYP76A.8 |
| 405 | 5613 | CYP76A20 | 513 aa | partial | CYP76A20 ortholog CYP76A.9 |
| 406 | 5631 | CYP76A.10P | 103 aa | short-sequence | CYP76A.10P |
| 407 | 5641 | CYP76B17 | 491 aa |  | CYP76B17 ortholog CYP76B.1 |
| 408 | 5659 | CYP76B39P | 487 aa |  | CYP76B39P ortholog CYP76B.2P |
| 409 | 5677 | CYP76B38 | 497 aa |  | CYP76B38 CYP76B.3 |
| 410 | 5694 | CYP76B18 | 493 aa |  | CYP76B18 CYP76B.4 |
| 411 | 5712 | CYP76B19 | 492 aa |  | CYP76B19 CYP76B.5 |
| 412 | 5730 | CYP76B37P | 302 aa | short-sequence | CYP76B37P CYP76B.6P |
| 413 | 5743 | CYP76B36P | 77 aa | internal-stop; short-sequence | CYP76B36P CYP76B.7P |
| 414 | 5752 | CYP76B35P | 201 aa | short-sequence | CYP76B35P CYP76B.8Pv1 |
| 415 | 5762 | CYP76B.8Pv2 | 47 aa | short-sequence | CYP76B.8Pv2 |
| 416 | 5770 | CYP76B34 | 493 aa |  | CYP76B34 CYP76B.9 |
| 417 | 5788 | CYP76B34-de1b | 23 aa | short-sequence | CYP76B34-de1b CYP76B.9-de1b |
| 418 | 5795 | CYP76B21a | 495 aa |  | CYP76B21a paralog CYP76B.10 |
| 419 | 5812 | CYP76B21b | 492 aa |  | CYP76B21b paralog CYP76B.11v1 |
| 420 | 5827 | CYP76B.11v2 | 43 aa | short-sequence | CYP76B.11v2 |
| 421 | 5835 | CYP76B.12P | 361 aa | internal-stop | CYP76B.12P |
| 422 | 5852 | CYP76B32P | 486 aa | frameshift | CYP76B32P ortholog CYP76B.13 |
| 423 | 5871 | CYP76B40Pa | 55 aa | short-sequence | CYP76B40Pa CYP76B.14P |
| 424 | 5881 | CYP76B41Pa | 98 aa | short-sequence | CYP76B41Pa CYP76B.15Pv1 |
| 425 | 5889 | CYP76B.15Pv2 | 58 aa | short-sequence | CYP76B.15Pv2 |
| 426 | 5897 | CYP76B42Pa | 376 aa | partial | CYP76B42Pa CYP76B.16P |
| 427 | 5915 | CYP76B43 | 492 aa |  | CYP76B43 CYP76B.17 |
| 428 | 5935 | CYP76B40Pb | 58 aa | short-sequence | CYP76B40Pb CYP76B.18P |
| 429 | 5944 | CYP76B41Pb | 104 aa | internal-stop; short-sequence | CYP76B41Pb CYP76B.19P |
| 430 | 5953 | CYP76B42b | 492 aa |  | CYP76B42b CYP76B.20 |
| 431 | 5969 | CYP76B.20v2 | 35 aa | short-sequence | CYP76B.20v2 |
| 432 | 5980 | CYP76B11P | 492 aa | internal-stop | CYP76B11P ortholog CYP76B.21P |
| 433 | 6000 | CYP76B44P | 104 aa | internal-stop; short-sequence | CYP76B44P CYP76B.22P |
| 434 | 6010 | CYP76B45 | 492 aa |  | CYP76B45 CYP76B.23 |
| 435 | 6029 | CYP76B46P | 144 aa | internal-stop; short-sequence | CYP76B46P CYP76B.24P |
| 436 | 6039 | CYP76B47P | 95 aa | internal-stop; short-sequence | CYP76B47P CYP76B.25P |
| 437 | 6048 | CYP76B48 | 489 aa |  | CYP76B48 CYP76B.26 |
| 438 | 6067 | CYP76B12P | 489 aa | partial; frameshift; internal-stop | CYP76B12P ortholog CYP76B.27P |
| 439 | 6086 | CYP76B13a | 496 aa |  | CYP76B13a paralog CYP76B.28 |
| 440 | 6103 | CYP76B13b | 497 aa | partial; frameshift | CYP76B13b paralog CYP76B.29 |
| 441 | 6120 | CYP76B49P | 495 aa | internal-stop | CYP76B49P CYP76B.30P |
| 442 | 6138 | CYP76B50 | 495 aa | pseudogene; partial; internal-stop | CYP76B50 CYP76B.31 |
| 443 | 6155 | CYP76B50-de3b | 30 aa | partial; short-sequence | CYP76B50-de3b CYP76B.31-de3b |
| 444 | 6163 | CYP76B54 | 495 aa | partial | CYP76B54 CYP76B.32 |
| 445 | 6180 | CYP76B30 | 496 aa |  | CYP76B30 ortholog CYP76B.33 |
| 446 | 6198 | CYP76B26 | 477 aa |  | CYP76B26 possible ortholog CYP76B.34 |
| 447 | 6216 | CYP76B25 | 477 aa |  | CYP76B25 possible ortholog CYP76B.35 |
| 448 | 6235 | CYP76B51 | 469 aa | internal-stop | CYP76B51 CYP76B.36 |
| 449 | 6253 | CYP76B24 | 465 aa |  | CYP76B24 ortholog CYP76B.37 |
| 450 | 6271 | CYP76B23 | 469 aa |  | CYP76B23 ortholog CYP76B.38 |
| 451 | 6289 | CYP76B53P-de3b | 167 aa | short-sequence | CYP76B53P-de3b CYP76B.39P-de3b |
| 452 | 6299 | CYP76B53P | 480 aa | partial; frameshift; uncertain-boundary | CYP76B53P CYP76B.39P |
| 453 | 6319 | CYP76B52 | 468 aa | partial | CYP76B52 CYP76B.40 |
| 454 | 6337 | CYP76B55 | 515 aa | partial | CYP76B55 CYP76B.41 |
| 455 | 6355 | CYP76B56P | 393 aa | partial | CYP76B56P CYP76B.42P |
| 456 | 6371 | CYP76B28P | 496 aa | internal-stop | CYP76B28P ortholog CYP76B.43P |
| 457 | 6389 | CYP76B.44P | 362 aa | internal-stop | CYP76B.44P |
| 458 | 6405 | CYP76B27P | 447 aa | internal-stop | CYP76B27P CYP76B.45P |
| 459 | 6424 | CYP76G10 | 509 aa |  | CYP76G10 ortholog CYP76G.1 |
| 460 | 6444 | CYP76Y8 | 492 aa | internal-stop | CYP76Y8 CYP76Y.1 |
| 461 | 6461 | CYP76Y7 | 501 aa |  | CYP76Y7 CYP76Y.2 |
| 462 | 6478 | CYP76Y5-de1b2b | 164 aa | partial; short-sequence | CYP76Y5-de1b2b CYP76Y.3-de1b2b |
| 463 | 6489 | CYP76Y5 | 502 aa |  | CYP76Y5 possible ortholog CYP76Y.3 |
| 464 | 6506 | CYP76Y9P | 494 aa | partial | CYP76Y9P CYP76Y.4P |
| 465 | 6525 | CYP76Y10P | 178 aa | short-sequence | CYP76Y10P CYP76Y.5P |
| 466 | 6535 | CYP76Y11P | 32 aa | short-sequence | CYP76Y11P CYP76Y.6P |
| 467 | 6543 | CYP76Y12P | 191 aa | short-sequence | CYP76Y12P CYP76Y.7P |
| 468 | 6556 | CYP77A20 | 503 aa |  | CYP77A20 ortholog CYP77A.1 |
| 469 | 6573 | CYP77A19 | 511 aa |  | CYP77A19 ortholog CYP77A.2 |
| 470 | 6591 | CYP77B11 | 508 aa |  | CYP77B11 ortholog CYP77B |
| 471 | 6610 | CYP78A77 | 531 aa |  | CYP78A77 ortholog CYP78A.1 |
| 472 | 6628 | CYP78A74 | 496 aa |  | CYP78A74 ortholog CYP78A.2 |
| 473 | 6645 | CYP78A79 | 519 aa |  | CYP78A79 ortholog CYP78A.3 |
| 474 | 6664 | CYP78A78 | 487 aa |  | CYP78A78 ortholog CYP78A.4 |
| 475 | 6681 | CYP78A76 | 512 aa | partial | CYP78A76 ortholog CYP78A.5 |
| 476 | 6699 | CYP78A75 | 512 aa |  | CYP78A75 ortholog CYP78A.6 |
| 477 | 6718 | CYP78A80 | 484 aa |  | CYP78A80 CYP78A.7 lost in tomato |
| 478 | 6738 | CYP79A.1P | 485 aa |  | CYP79A.1P |
| 479 | 6757 | CYP79A.1P-de1b | 39 aa | short-sequence | CYP79A.1P-de1b |
| 480 | 6765 | CYP79A.1P-de1c | 45 aa | short-sequence | CYP79A.1P-de1c |
| 481 | 6772 | CYP79A.1P-de1d | 45 aa | short-sequence | CYP79A.1P-de1d |
| 482 | 6778 | CYP79A52 | 521 aa |  | CYP79A52 ortholog CYP79A.2 |
| 483 | 6796 | CYP79A59 | 518 aa |  | CYP79A59 CYP79A.3 |
| 484 | 6814 | CYP79A59-de2b | 41 aa | short-sequence | CYP79A59-de2b CYP79A.3-de2b |
| 485 | 6822 | CYP79A58P | 204 aa | short-sequence | CYP79A58P CYP79A.4P |
| 486 | 6834 | CYP79A57P | 89 aa | internal-stop; short-sequence | CYP79A57P CYP79A.5P |
| 487 | 6842 | CYP79A56 | 510 aa |  | CYP79A56 CYP79A.6 one frameshift, possible pseudogene |
| 488 | 6860 | CYP79A55P | 538 aa | frameshift; ambiguous-residue | CYP79A55P CYP79A.7P |
| 489 | 6880 | CYP79A54P | 367 aa | internal-stop | CYP79A54P CYP79A.8P |
| 490 | 6898 | CYP79A53P | 90 aa | short-sequence | CYP79A53P CYP79A.9P |
| 491 | 6907 | CYP79A.10P | 516 aa |  | CYP79A.10P |
| 492 | 6924 | CYP79A.10-de1b | 103 aa | short-sequence | CYP79A.10-de1b |
| 493 | 6932 | CYP79A.10-de1c | 59 aa | short-sequence | CYP79A.10-de1c |
| 494 | 6938 | CYP79A.11P | 426 aa | sequence-gap | CYP79A.11P |
| 495 | 6954 | CYP79A32 | 519 aa |  | CYP79A32 ortholog CYP79A.12 |
| 496 | 6972 | CYP79A.13P | 207 aa | frameshift; short-sequence | CYP79A.13P |
| 497 | 6984 | CYP79A.14P | 246 aa | short-sequence | CYP79A.14P |
| 498 | 6996 | CYP79A.15P | 181 aa | internal-stop; short-sequence | CYP79A.15P |
| 499 | 7005 | CYP79A.16P | 454 aa | sequence-gap | CYP79A.16P |
| 500 | 7024 | CYP79A.17P | 134 aa | short-sequence | CYP79A.17P |
| 501 | 7034 | CYP79A.18P | 516 aa |  | CYP79A.18P |
| 502 | 7052 | CYP79A.19P | 74 aa | short-sequence | CYP79A.19P |
| 503 | 7060 | CYP80M11 | 487 aa |  | CYP80M11 CYP80E.1 |
| 504 | 7077 | CYP80M10 | 486 aa |  | CYP80M10 CYP80E.2 |
| 505 | 7094 | CYP80M9P | 221 aa | short-sequence | CYP80M9P CYP80E.3P |
| 506 | 7106 | CYP80M8P | 156 aa | short-sequence | CYP80M8P CYP80E.4P |
| 507 | 7117 | CYP80M7P | 106 aa | internal-stop; short-sequence | CYP80M7P CYP80E.5P |
| 508 | 7127 | CYP80M6P | 486 aa |  | CYP80M6P CYP80E.6P |
| 509 | 7145 | CYP80M5P | 486 aa | internal-stop | CYP80M5P CYP80E.7P |
| 510 | 7161 | CYP80M5P-de2b | 24 aa | short-sequence | CYP80M5P-de2b CYP80E.7P-de2b |
| 511 | 7169 | CYP80M4 | 486 aa | partial; frameshift; ambiguous-residue | CYP80M4 CYP80E.8 |
| 512 | 7187 | CYP80M3P | 396 aa | partial | CYP80M3P CYP80E.9P |
| 513 | 7208 | CYP80M1 | 486 aa |  | CYP80M1 ortholog CYP80E.10 |
| 514 | 7225 | CYP80E6 | 495 aa |  | CYP80E6 possible ortholog CYP80E.11 |
| 515 | 7243 | CYP80E9 | 496 aa |  | CYP80E9 CYP80E.12 |
| 516 | 7261 | CYP80E7 | 495 aa | partial | CYP80E7 possible ortholog CYP80E.13 |
| 517 | 7279 | CYP80E.14P | 68 aa | short-sequence | CYP80E.14P |
| 518 | 7289 | CYP80F4 | 499 aa | partial | CYP80F4 CYP80F.1 |
| 520 | 7323 | CYP80F3 | 494 aa | partial | CYP80F3 ortholog CYP80F.2 |
| 521 | 7341 | CYP80N1 | 505 aa | partial | CYP80N1 CYP80x.1 |
| 527 | 7462 | CYP81B38a | 512 aa |  | CYP81B38a CYP81B.1v1 |
| 528 | 7477 | CYP81B.1v2 | 74 aa | short-sequence | CYP81B.1v2 |
| 529 | 7486 | CYP81B41 | 509 aa | partial; frameshift | CYP81B41 possible ortholog CYP81B.2 |
| 530 | 7504 | CYP81B40 | 507 aa |  | CYP81B40 possible ortholog CYP81B.3 |
| 531 | 7522 | CYP81B46P | 83 aa | internal-stop; short-sequence | CYP81B46P CYP81B.4P |
| 532 | 7531 | CYP81B45P | 104 aa | short-sequence | CYP81B45P CYP81B.5P |
| 533 | 7540 | CYP81B39 | 455 aa |  | CYP81B39 ortholog CYP81B.6 |
| 534 | 7559 | CYP81B44P | 179 aa | internal-stop; short-sequence | CYP81B44P CYP81B.7P |
| 535 | 7570 | CYP81B38b | 512 aa | partial; frameshift | CYP81B38b CYP81B.8 |
| 536 | 7587 | CYP81B38b-de1b | 40 aa | short-sequence | CYP81B38b-de1b CYP81B.8-de1b |
| 537 | 7594 | CYP81B38a | 512 aa |  | CYP81B38a CYP81B.9 |
| 538 | 7613 | CYP81B43P | 507 aa | partial; frameshift | CYP81B43P CYP81B.10P |
| 539 | 7633 | CYP81B37 | 509 aa |  | CYP81B37 possible ortholog CYP81B.11 |
| 540 | 7650 | CYP81C8 | 516 aa | partial | CYP81C8 ortholog CYP81C.1 |
| 541 | 7668 | CYP81C9P | 517 aa | partial; internal-stop | CYP81C9P possible ortholog CYP81C.2P |
| 542 | 7687 | CYP81C.3P | 521 aa | internal-stop; ambiguous-residue | CYP81C.3P |
| 543 | 7705 | CYP81C10 | 510 aa |  | CYP81C10 ortholog CYP81C.4 |
| 544 | 7723 | CYP81C11 | 513 aa | partial | CYP81C11 ortholog CYP81C.5 |
| 545 | 7741 | CYP81C.6P | 130 aa | short-sequence | CYP81C.6P |
| 546 | 7751 | CYP81C.7P | 68 aa | short-sequence | CYP81C.7P |
| 547 | 7761 | CYP81Y1 | 484 aa |  | CYP81Y1 ortholog CYP81x.1 (divergent gene in the CYP81C cluster) |
| 548 | 7779 | CYP81V.1P | 199 aa | partial; internal-stop; short-sequence | CYP81V.1P |
| 549 | 7789 | CYP81V.2P | 489 aa |  | CYP81V.2P |
| 550 | 7806 | CYP81Q31 | 500 aa |  | CYP81Q31 ortholog CYP81V.3 |
| 551 | 7823 | CYP81Q30 | 494 aa |  | CYP81Q30 ortholog CYP81V.4 |
| 552 | 7840 | CYP82D41 | 517 aa |  | CYP82D41 CYP82D.1 |
| 553 | 7857 | CYP82D41-DE1B | 63 aa | internal-stop; short-sequence | CYP82D41-DE1B CYP82D.1-de1b |
| 554 | 7864 | CYP82D42 | 517 aa |  | CYP82D42 CYP82D.2 |
| 555 | 7880 | CYP82D34 | 511 aa | partial | CYP82D34 CYP82D.3 |
| 556 | 7898 | CYP82D34-de1b | 22 aa | short-sequence | CYP82D34-de1b CYP82D.3-de1b |
| 557 | 7905 | CYP82D39 | 517 aa | partial; internal-stop | CYP82D39 CYP82D.4 |
| 558 | 7924 | CYP82D37P | 323 aa | internal-stop; short-sequence | CYP82D37P CYP82D.5P |
| 559 | 7939 | CYP82D40 | 520 aa |  | CYP82D40 possible ortholog CYP82D.6 |
| 560 | 7958 | CYP82D43 | 511 aa |  | CYP82D43 CYP82D.7 |
| 561 | 7977 | CYP82D44 | 462 aa |  | CYP82D44 CYP82D.8 |
| 562 | 7995 | CYP82D45 | 521 aa |  | CYP82D45 CYP82D.9 |
| 563 | 8013 | CYP82C22 | 479 aa |  | CYP82C22 ortholog CYP82D.10 |
| 564 | 8031 | CYP82E14P | 36 aa | short-sequence | CYP82E14P CYP82E.1P |
| 565 | 8036 | CYP82E13v1 | 530 aa | partial | CYP82E13v1 CYP82E.2v1 |
| 566 | 8055 | CYP82E13v2 | 145 aa | short-sequence | CYP82E13v2 CYP82E.2v2 |
| 567 | 8063 | CYP82E.2 | missing | sequence-missing | CYP82E.2 |
| 568 | 8066 | CYP82E.2 | 145 aa | partial; short-sequence | CYP82E.2 |
| 569 | 8076 | CYP82E12 | 530 aa | partial; frameshift | CYP82E12 CYP82E.3 |
| 570 | 8096 | CYP82E11 | 527 aa |  | CYP82E11 ortholog CYP82E.4 |
| 571 | 8113 | CYP82E7 | 221 aa | partial; short-sequence | CYP82E7   Solanum tuberosum   (potato) |
| 572 | 8124 | CYP82M3 | 524 aa |  | CYP82M3 ortholog CYP82E.5 CYP82M3 |
| 573 | 8143 | CYP82M2P | 520 aa | uncertain-boundary; internal-stop | CYP82M2P ortholog CYP82E.6P CYP82M2P |
| 574 | 8160 | CYP82U2 | 518 aa | partial | CYP82U2 CYP82L.1 |
| 575 | 8178 | CYP82U1 | 511 aa | partial | CYP82U1 ortholog CYP82L.2 |
| 576 | 8196 | CYP82V1 | 520 aa | partial | CYP82V1 CYP82S.1 |
| 577 | 8215 | CYP82W1 | 545 aa |  | CYP82W1 ortholog CYP82x.1 |
| 578 | 8235 | CYP82y.1P | 220 aa | internal-stop; short-sequence | CYP82y.1P |
| 579 | 8248 | CYP83 | 285 aa | partial; short-sequence | CYP83       Solanum tuberosum (potato) |
| 580 | 8260 | CYP83C | 199 aa | partial; short-sequence | CYP83C      Solanum tuberosum (potato) |
| 581 | 8270 | CYP83 | 248 aa | short-sequence | CYP83     Solanum tuberosum (potato) |
| 582 | 8281 | SGN-U270131 | 258 aa | ambiguous-residue; short-sequence | SGN-U270131 Solanum tuberosum (potato) [3 ESTs aligned] |
| 583 | 8291 | CYP84A2 | 519 aa | partial | CYP84A2 ortholog CYP84A |
| 584 | 8306 | CYP84A | missing | sequence-missing | CYP84A |
| 585 | 8314 | CYP85A1 | 358 aa | partial | CYP85A1 |
| 586 | 8332 | CYP85A1 | missing | sequence-missing | CYP85A1 |
| 587 | 8338 | CYP85A3 | 426 aa |  | CYP85A3 |
| 588 | 8355 | CYP85A3 | missing | sequence-missing | CYP85A3 |
| 589 | 8360 | CYP86A33 | 521 aa |  | CYP86A33    Solanum tuberosum (potato) |
| 590 | 8376 | CYP86A33 | 487 aa |  | CYP86A33 (2 aa diffs) |
| 591 | 8394 | CYP86A69 | 548 aa |  | CYP86A69 ortholog CYP86A.1 |
| 592 | 8413 | CYP86A68 | 535 aa |  | CYP86A68 ortholog CYP86A.2 |
| 593 | 8430 | CYP86A.3P | 123 aa | short-sequence | CYP86A.3P |
| 594 | 8440 | CYP86A.4P | 404 aa | uncertain-boundary; internal-stop | CYP86A.4P |
| 595 | 8456 | CYP86B12 | 540 aa |  | CYP86B12 ortholog CYP86B.1 |
| 596 | 8474 | CYP86G1 | 521 aa | partial | CYP86G1 ortholog CYP86x.1 |
| 597 | 8494 | CYP86 | 83 aa | internal-stop; short-sequence | CYP86 pseudogene |
| 598 | 8500 | CYP87A21 | 366 aa |  | CYP87A21 ortholog CYP87A.1 |
| 599 | 8519 | CYP87A.2P | 358 aa | partial | CYP87A.2P pseudogene |
| 600 | 8541 | CYP87A20 | 363 aa |  | CYP87A20 possible ortholog CYP87A.3 |
| 601 | 8560 | CYP87A19 | 362 aa |  | CYP87A19 ortholog CYP87A.4 |
| 602 | 8578 | CYP87A.5P | 320 aa | internal-stop; short-sequence | CYP87A.5P pseudogene |
| 603 | 8593 | CYP87A.6P | 362 aa |  | CYP87A.6P pseudogene (three frameshifts) |
| 604 | 8614 | CYP87A.7P | 199 aa | short-sequence | CYP87A.7P pseudogene |
| 605 | 8627 | CYP87E3 | 421 aa |  | CYP87E3 ortholog CYP87B |
| 606 | 8646 | CYP88G1 | 479 aa |  | CYP88G1 ortholog CYP88A.1 |
| 607 | 8666 | CYP88A35 | 435 aa |  | CYP88A35 ortholog CYP88A.2 |
| 608 | 8686 | CYP88A.3P | 239 aa | internal-stop; short-sequence | CYP88A.3P |
| 609 | 8702 | CYP88A.4P | 60 aa | internal-stop; short-sequence | CYP88A.4P |
| 610 | 8712 | CYP88B1 | 428 aa |  | CYP88B1 ortholog CYP88B |
| 611 | 8732 | CYP88C3 | 410 aa | partial | CYP88C3 CYP88C.1 |
| 612 | 8752 | CYP88C2 | 424 aa |  | CYP88C2 CYP88C.2 |
| 613 | 8772 | CYP88C4 | 428 aa |  | CYP88C4 CYP88C.3 one frameshift |
| 614 | 8794 | CYP88C5 | 427 aa |  | CYP88C5 CYP88C.4 |
| 615 | 8812 | CYP88C6P | 242 aa | internal-stop; short-sequence | CYP88C6P CYP88C.5P |
| 616 | 8825 | CYP89A69 | 508 aa |  | CYP89A69 ortholog CYP89A.1 |
| 617 | 8843 | CYP89A75 | 514 aa |  | CYP89A75 CYP89A.2 found pseudogene CYP89A69-de1b on scaffold06019 95% |
| 618 | 8860 | CYP89A74 | 513 aa |  | CYP89A74 CYP89A.3 |
| 619 | 8877 | CYP89A72 | 517 aa | partial; frameshift | CYP89A72 CYP89A.4 |
| 620 | 8895 | CYP89A70 | 520 aa |  | CYP89A70 CYP89A.5v1 |
| 621 | 8910 | CYP89A.5v2 | 57 aa | pseudogene; short-sequence | CYP89A.5v2 |
| 622 | 8921 | CYP89A73 | 508 aa |  | CYP89A73 ortholog CYP89A.6 |
| 623 | 8938 | CYP89A.6-de1b | 81 aa | short-sequence | CYP89A.6-de1b pseudogene |
| 624 | 8946 | CYP89A.6-de1c | 24 aa | short-sequence | CYP89A.6-de1c pseudogene |
| 625 | 8952 | CYP89A.7P | 511 aa | partial; frameshift | CYP89A.7P pseudogene |
| 626 | 8971 | CYP89A.8P | 433 aa | partial; frameshift | CYP89A.8P pseudogene |
| 627 | 8988 | CYP89A.9P | 406 aa | partial; frameshift | CYP89A.9P pseudogene |
| 628 | 9005 | CYP89A.10P | 365 aa | internal-stop | CYP89A.10P pseudogene |
| 629 | 9020 | CYP89A.11P | 508 aa | internal-stop | CYP89A.11P pseudogene |
| 630 | 9038 | CYP89A.12P | 256 aa | short-sequence | CYP89A.12P pseudogene |
| 631 | 9051 | CYP89A.13P | 85 aa | partial; short-sequence | CYP89A.13P pseudogene |
| 632 | 9059 | CYP89A.14P | 107 aa | short-sequence | CYP89A.14P pseudogene |
| 633 | 9067 | CYP89A.15Pv1 | 429 aa | partial; internal-stop | CYP89A.15Pv1 pseudogene |
| 634 | 9082 | CYP89A.15Pv2 | 82 aa | short-sequence | CYP89A.15Pv2 |
| 635 | 9091 | CYP89A.16P | 156 aa | short-sequence | CYP89A.16P pseudogene |
| 636 | 9100 | CYP89A.17P | 33 aa | short-sequence | CYP89A.17P pseudogene |
| 637 | 9107 | CYP90A5 | 446 aa |  | CYP90A5 ortholog CYP90A |
| 638 | 9124 | CYP90A | missing | sequence-missing | CYP90A |
| 639 | 9129 | CYP90B3 | 373 aa |  | CYP90B3 ortholog CYP90B |
| 640 | 9146 | CYP90B | missing | sequence-missing | CYP90B |
| 641 | 9152 | CYP90C2 | 379 aa |  | CYP90C2 ortholog CYP90C |
| 642 | 9169 | CYP90C | missing | sequence-missing | CYP90C |
| 643 | 9175 | CYP90D19 | 374 aa | partial | CYP90D19 ortholog CYP90D |
| 644 | 9192 | CYP90D | missing | sequence-missing | CYP90D |
| 645 | 9198 | CYP92B5 | 506 aa |  | CYP92B5 ortholog CYP92B.1 |
| 646 | 9215 | CYP92B16 | 506 aa |  | CYP92B16 CYP92B.2 |
| 647 | 9232 | CYP92B17P | 76 aa | internal-stop; short-sequence | CYP92B17P CYP92B.3P |
| 648 | 9241 | CYP92B18P | 43 aa | short-sequence | CYP92B18P CYP92B.4P |
| 649 | 9248 | CYP92B19P | 108 aa | short-sequence | CYP92B19P CYP92B.5P |
| 650 | 9257 | CYP92B20 | 506 aa |  | CYP92B20 CYP92B.6 |
| 651 | 9275 | CYP92B21v1 | 299 aa | short-sequence | CYP92B21v1 CYP92B.7v1 |
| 652 | 9288 | CYP92B21v2 | 129 aa | short-sequence | CYP92B21v2 CYP92B.7v2 |
| 653 | 9298 | CYP92B22 | 509 aa |  | CYP92B22 CYP92B.8 |
| 654 | 9315 | CYP92B23P | 270 aa | short-sequence | CYP92B23P CYP92B.9P |
| 655 | 9328 | CYP92B9 | 508 aa |  | CYP92B9 ortholog CYP92B.10 |
| 656 | 9345 | CYP92B7P | 498 aa | partial; frameshift | CYP92B7P ortholog CYP92B.11P |
| 657 | 9365 | CYP92B24P | 599 aa | pseudogene; partial; internal-stop | CYP92B24P CYP92B.12P |
| 658 | 9390 | CYP92B25 | 502 aa |  | CYP92B25 CYP92B.13 |
| 659 | 9408 | CYP92B26 | 503 aa |  | CYP92B26 CYP92B.14 |
| 660 | 9426 | CYP92B27P | 486 aa | pseudogene; internal-stop | CYP92B27P CYP92B.15P |
| 661 | 9444 | CYP92B14P | 499 aa | pseudogene; partial; internal-stop | CYP92B14P ortholog CYP92B.16P |
| 662 | 9462 | CYP92B12P | 452 aa | partial | CYP92B12P CYP92B.17P |
| 663 | 9480 | CYP92A51P | 507 aa | partial; frameshift | CYP92A51P ortholog CYP92A.1P |
| 664 | 9497 | CYP92A50 | 509 aa |  | CYP92A50 ortholog CYP92A.2 |
| 665 | 9515 | CYP93A42 | 520 aa | partial | CYP93A42 ortholog CYP93A |
| 666 | 9533 | CYP93A | missing | sequence-missing | CYP93A |
| 667 | 9539 | CYP94B18 | 495 aa |  | CYP94B18 ortholog CYP94B.1v1 |
| 668 | 9554 | CYP94B.1v2 | 31 aa | short-sequence | CYP94B.1v2 |
| 669 | 9562 | CYP94B20 | 489 aa |  | CYP94B20 ortholog CYP94B.2 |
| 670 | 9579 | CYP94B19b | 492 aa |  | CYP94B19b CYP94B.3v1 |
| 671 | 9594 | CYP94B19frag1 | 47 aa | short-sequence | CYP94B19frag1 CYP94B.3v2 |
| 672 | 9602 | CYP94B19frag2 | 43 aa | short-sequence | CYP94B19frag2 CYP94B.3v3 |
| 673 | 9610 | CYP94B19Pa | 488 aa | frameshift | CYP94B19Pa CYP94B.4P pseudogene, one frameshift |
| 674 | 9628 | CYP94B17 | 520 aa |  | CYP94B17 ortholog CYP94B.5 |
| 675 | 9645 | CYP94B.6P | 230 aa | internal-stop; short-sequence | CYP94B.6P |
| 676 | 9655 | CYP94B.7P | 48 aa | internal-stop; short-sequence | CYP94B.7P |
| 677 | 9662 | CYP94A26 | 504 aa |  | CYP94A26 ortholog CYP94A.1 |
| 678 | 9680 | CYP94A25 | 525 aa |  | CYP94A25 ortholog CYP94A.2 |
| 679 | 9698 | CYP94A24 | 501 aa |  | CYP94A24 ortholog CYP94A.3 |
| 680 | 9715 | CYP94A32P | 261 aa | internal-stop; short-sequence | CYP94A32P CYP94A.4P |
| 681 | 9728 | CYP94C29 | 489 aa |  | CYP94C29 ortholog CYP94C.1 |
| 682 | 9745 | CYP94C30b | 496 aa |  | CYP94C30b CYP94C.2v1 paralog to CYP94C30 |
| 683 | 9760 | CYP94C30frag | 119 aa | short-sequence | CYP94C30frag CYP94C.2v2 |
| 684 | 9769 | CYP94C30a | 496 aa | partial | CYP94C30a CYP94C.3 paralog to CYP94C30 |
| 686 | 9794 | CYP94C31 | 509 aa | partial | CYP94C31 CYP94C.4 not in tomato |
| 687 | 9813 | CYP94D34 | 495 aa | partial | CYP94D34 ortholog CYP94D.1 |
| 688 | 9832 | CYP94D34-de1b | 148 aa | short-sequence | CYP94D34-de1b CYP94D.2P |
| 689 | 9842 | CYP94K1 | 485 aa | partial | CYP94K1 CYP94x.1 (not a close match to tomato) |
| 690 | 9862 | CYP96A52 | 501 aa |  | CYP96A52 possible ortholog CYP96A.1 |
| 691 | 9878 | CYP96A53 | 506 aa |  | CYP96A53 ortholog CYP96A.2v1 |
| 692 | 9895 | CYP96A51 | 517 aa | partial; frameshift | CYP96A51 possible ortholog CYP96A.3 |
| 693 | 9915 | CYP96A.4P | 507 aa | partial; internal-stop | CYP96A.4P |
| 694 | 9932 | CYP96A49 | 505 aa |  | CYP96A49 possible ortholog CYP96A.5 |
| 695 | 9950 | CYP96A55 | 524 aa | partial | CYP96A55 CYP96A.6 |
| 696 | 9969 | CYP96A44 | 517 aa |  | CYP96A44 possible ortholog CYP96A.7v1 |
| 697 | 9986 | CYP96A45 | 511 aa | partial | CYP96A45 possible ortholog CYP96A.8 |
| 698 | 10004 | CYP96A.7/8v2 | 48 aa | short-sequence | CYP96A.7/8v2 |
| 699 | 10014 | CYP96A.7/8v3 | 43 aa | short-sequence | CYP96A.7/8v3 |
| 700 | 10024 | CYP9A46v2 | 508 aa |  | CYP9A46v2 CYP96A.9v1 |
| 701 | 10039 | CYP96A.9v2 | 27 aa | short-sequence | CYP96A.9v2 |
| 702 | 10047 | CYP96A.9v3 | 34 aa | short-sequence | CYP96A.9v3 |
| 703 | 10055 | CYP96A.10P | 190 aa | short-sequence | CYP96A.10P |
| 704 | 10065 | CYP96A.11P | 508 aa | partial; frameshift | CYP96A.11P |
| 705 | 10082 | CYP96A48 | 510 aa |  | CYP96A48 ortholog CYP96A.12 |
| 706 | 10099 | CYP96A.13P | 509 aa | partial; frameshift; internal-stop | CYP96A.13P |
| 707 | 10115 | CYP9A46v1 | 508 aa |  | CYP9A46v1 CYP96A.14 |
| 708 | 10132 | CYP96A56 | 513 aa |  | CYP96A56 CYP96A.15 |
| 709 | 10149 | CYP96A.16P | 214 aa | internal-stop; short-sequence | CYP96A.16P |
| 710 | 10158 | CYP96A.17P | 462 aa | sequence-gap; internal-stop | CYP96A.17P |
| 711 | 10175 | CYP96A.18P | 143 aa | internal-stop; short-sequence | CYP96A.18P |
| 712 | 10186 | CYP97A29 | 511 aa |  | CYP97A29 ortholog |
| 713 | 10210 | CYP97B22 | 499 aa | partial | CYP97B22 ortholog CYP97B |
| 714 | 10235 | CYP97C11 | 517 aa |  | CYP97C11 ortholog |
| 715 | 10257 | CYP98A57b | 499 aa |  | CYP98A57b CYP98A.1 |
| 716 | 10275 | CYP98A53bP | 295 aa | partial; short-sequence | CYP98A53bP CYP98A.2P |
| 717 | 10291 | CYP98A57aP | 382 aa | partial | CYP98A57aP CYP98A.3Pv1 |
| 718 | 10304 | CYP98A.3Pv2 | 37 aa | short-sequence | CYP98A.3Pv2 |
| 719 | 10312 | CYP98A.4P | missing | sequence-missing | CYP98A.4P |
| 720 | 10315 | CYP98A.4P | 78 aa | short-sequence | CYP98A.4P |
| 721 | 10321 | CYP98A53aP | 516 aa | partial | CYP98A53aP possible ortholog CYP98A.4Pv1 |
| 722 | 10339 | CYP98A4Pv2 | 78 aa | short-sequence | CYP98A4Pv2 |
| 723 | 10346 | CYP98A.4Pv3 | 46 aa | short-sequence | CYP98A.4Pv3 |
| 724 | 10352 | CYP98A.4Pv4 | 54 aa | short-sequence | CYP98A.4Pv4 |
| 725 | 10358 | CYP98A53c | 162 aa | partial; short-sequence | CYP98A53c CYP98A.10 |
| 726 | 10368 | CYP98A53d | 162 aa | partial; short-sequence | CYP98A53d CYP98A.11 |
| 727 | 10380 | CYP98A52v1 | 509 aa | partial | CYP98A52v1 ortholog CYP98A.5v1 |
| 728 | 10397 | CYP98A.5v2 | 50 aa | short-sequence | CYP98A.5v2 |
| 729 | 10403 | CYP98A.5v3 | 56 aa | short-sequence | CYP98A.5v3 |
| 730 | 10409 | CYP98A.5v4 | 91 aa | short-sequence | CYP98A.5v4 |
| 731 | 10416 | CYP98A.5 | missing | sequence-missing | CYP98A.5 |
| 732 | 10419 | CYP98A.5 | 115 aa | short-sequence | CYP98A.5 |
| 733 | 10427 | CYP98A52v2 | 509 aa | sequence-gap | CYP98A52v2 ortholog CYP98A.8 |
| 734 | 10446 | CYP98A52v2-ie1b | 109 aa | partial; short-sequence | CYP98A52v2-ie1b CYP98A.8-ie1b |
| 735 | 10456 | CYP98A51 | 508 aa |  | CYP98A51 ortholog CYP98A.6 |
| 736 | 10474 | CYP98A55c | 214 aa | short-sequence | CYP98A55c CYP98A.7, 100% to CYP98A.13, 100% to CYP98A.14 |
| 737 | 10485 | CYP98A.7 | missing | sequence-missing | CYP98A.7 |
| 738 | 10488 | CYP98A.7 | 214 aa | partial; short-sequence | CYP98A.7 |
| 739 | 10496 | CYP98A55b | 507 aa |  | CYP98A55b CYP98A.13 paralog of CYP98A55 tomato |
| 740 | 10512 | CYP98A55a | 507 aa |  | CYP98A55a CYP98A.14 paralog of CYP98A55 tomato |
| 741 | 10530 | CYP98A53e | 311 aa | short-sequence | CYP98A53e CYP98A.9 |
| 742 | 10546 | CYP98A56 | 511 aa |  | CYP98A56 ortholog CYP98A.12 |
| 743 | 10564 | CYP701A30b | 514 aa | partial | CYP701A30b CYP701A.1 |
| 744 | 10581 | CYP701A30a | 514 aa |  | CYP701A30a CYP701A.2 |
| 745 | 10597 | CYP701A1.3 | 52 aa | short-sequence | CYP701A1.3 |
| 746 | 10604 | CYP703A13 | 521 aa |  | CYP703A13 ortholog CYP703A |
| 747 | 10623 | CYP704A75 | 526 aa |  | CYP704A75 CYP704A.1 |
| 748 | 10642 | CYP704A75-de4b | 25 aa | short-sequence | CYP704A75-de4b CYP704A.1-de4b |
| 749 | 10649 | CYP704A74b | 505 aa |  | CYP704A74b CYP704A.2 |
| 750 | 10670 | CYP704A73b | 499 aa |  | CYP704A73b CYP704A.3a |
| 751 | 10689 | CYP704A73a | 499 aa |  | CYP704A73a CYP704A.3b = old CYP704A.7 |
| 752 | 10708 | CYP704A.3av2 | 245 aa | short-sequence | CYP704A.3av2 = old CYP704A.14 |
| 753 | 10720 | CYP704A.3av3 | 36 aa | short-sequence | CYP704A.3av3 |
| 754 | 10728 | CYP704A.3av4 | 22 aa | short-sequence | CYP704A.3av4 |
| 755 | 10736 | CYP704A66a | 506 aa |  | CYP704A66a paralog CYP704A.4av1 |
| 756 | 10757 | CYP704A66a-de1b | 32 aa | short-sequence | CYP704A66a-de1b CYP704A.4-de1b |
| 757 | 10763 | CYP704A66b | 504 aa |  | CYP704A66b CYP704A.4b = old CYP704A.8 |
| 758 | 10783 | CYP704A66b-de1b | 32 aa | short-sequence | CYP704A66b-de1b CYP704A.8-de1b |
| 759 | 10790 | CYP704A.4av2 | 182 aa | short-sequence | CYP704A.4av2 = old CYP704A.18 |
| 760 | 10800 | CYP704A.a4v3 | 94 aa | short-sequence | CYP704A.a4v3 |
| 761 | 10807 | CYP704A.a4v4 | 36 aa | short-sequence | CYP704A.a4v4 |
| 762 | 10816 | CYP704A73c | 216 aa | short-sequence | CYP704A73c CYP704A.5P |
| 763 | 10828 | CYP704A73c-de1b | 50 aa | short-sequence | CYP704A73c-de1b CYP704A.5P-de1b |
| 764 | 10837 | CYP704A74a | 506 aa |  | CYP704A74a CYP704A.6v1 |
| 765 | 10854 | CYP704A.6v2 | 39 aa | short-sequence | CYP704A.6v2 |
| 766 | 10862 | CYP704A.6v3 | 37 aa | short-sequence | CYP704A.6v3 |
| 767 | 10870 | CYP704A72P | 328 aa | uncertain-boundary; short-sequence | CYP704A72P CYP704A.9P |
| 768 | 10885 | CYP704A72P-de1b | 50 aa | short-sequence | CYP704A72P-de1b CYP704A.9P-de1b |
| 769 | 10894 | CYP704A65P | 500 aa | partial; frameshift; internal-stop | CYP704A65P possible ortholog CYP704A.10P |
| 770 | 10913 | CYP704A64 | 515 aa | partial | CYP704A64 ortholog CYP704A.11 |
| 771 | 10932 | CYP704A63a | 509 aa |  | CYP704A63a CYP704A.12 adjacent paralogs |
| 772 | 10950 | CYP704A63b | 515 aa |  | CYP704A63b CYP704A.13 adjacent paralogs |
| 773 | 10969 | CYP704A71Pb | 46 aa | internal-stop; short-sequence | CYP704A71Pb CYP704A.15P |
| 774 | 10977 | CYP704A71Pa | 396 aa | internal-stop; ambiguous-residue | CYP704A71Pa CYP704A.16P |
| 775 | 10996 | CYP704A70P | 506 aa | partial; uncertain-boundary | CYP704A70P ortholog CYP704A.17P |
| 776 | 11018 | CYP704A76 | 511 aa |  | CYP704A76 CYP704A.18 |
| 778 | 11060 | CYP704B30 | 513 aa |  | CYP704B30 ortholog CYP704B.1 |
| 782 | 11116 | CYP706C12P | 525 aa | partial; frameshift | CYP706C12P ortholog CYP706C.1P |
| 783 | 11136 | CYP706C14 | 532 aa |  | CYP706C14 possible ortholog CYP706C.2 |
| 784 | 11154 | CYP706C15P | 490 aa | partial | CYP706C15P ortholog CYP706C.3P syntenic |
| 785 | 11173 | CYP706C16 | 532 aa |  | CYP706C16 ortholog CYP706C.4 |
| 786 | 11191 | CYP706C16-de1b | 30 aa | short-sequence | CYP706C16-de1b CYP706C.4-de1b |
| 787 | 11198 | CYP706C17 | 521 aa |  | CYP706C17 possible ortholog CYP706C.5 |
| 788 | 11216 | CYP706C18 | 521 aa |  | CYP706C18 ortholog CYP706C.6 |
| 789 | 11232 | CYP706C.18 | 180 aa | short-sequence | CYP706C.18 identical to CYP706C.6 |
| 790 | 11243 | CYP706C18-de1b | 154 aa | short-sequence | CYP706C18-de1b ortholog CYP706C.7P |
| 791 | 11255 | CYP706C19 | 523 aa |  | CYP706C19 possible ortholog CYP706C.8 |
| 792 | 11273 | CYP706C20 | 523 aa |  | CYP706C20 ortholog CYP706C.9 |
| 793 | 11291 | CYP706C21 | 521 aa |  | CYP706C21 ortholog CYP706C.10 |
| 794 | 11309 | CYP706C21-de1b | 215 aa | short-sequence | CYP706C21-de1b CYP706C.11P no tomato ortholog |
| 795 | 11322 | CYP706C22 | 524 aa |  | CYP706C22 possible ortholog CYP706C.12 |
| 796 | 11340 | CYP706C23 | 521 aa |  | CYP706C23 possible ortholog CYP706C.13 |
| 797 | 11358 | CYP706C24 | 524 aa |  | CYP706C24 possible ortholog CYP706C.14 |
| 798 | 11376 | CYP706C25Pa | 187 aa | short-sequence | CYP706C25Pa ORTHOLOG CYP706C.15P N-term syntenic |
| 799 | 11388 | CYP706C25Pb | 328 aa | partial; short-sequence | CYP706C25Pb CYP706C.17P C-term |
| 800 | 11404 | CYP706C26 | 517 aa |  | CYP706C26 CYP706C.16 no tomato ortholog |
| 801 | 11422 | CYP706G4 | 515 aa |  | CYP706G4 ortholog CYP706G.1v1 |
| 802 | 11439 | CYP706G.1v2 | 54 aa | short-sequence | CYP706G.1v2 |
| 803 | 11447 | CYP706G5v1 | 517 aa |  | CYP706G5v1 possible ortholog CYP706G.3 |
| 804 | 11465 | CYP706G5v2 | 307 aa | short-sequence | CYP706G5v2 CYP706G.2 |
| 805 | 11480 | CYP706G6P | 288 aa | internal-stop; short-sequence | CYP706G6P CYP706G.4P |
| 806 | 11493 | CYP707A22 | 469 aa |  | CYP707A22     Solanum tuberosum  (potato) |
| 807 | 11506 | CYP707A22 | 458 aa |  | CYP707A22 ortholog of CYP707A7 tomato |
| 808 | 11528 | CYP707A69 | 442 aa |  | CYP707A69 ortholog CYP707A.1 |
| 809 | 11550 | CYP707A23 | 412 aa |  | CYP707A23 (98% to DQ206631) |
| 810 | 11569 | CYP707A23 | 475 aa | partial | CYP707A23    Solanum tuberosum (potato) |
| 811 | 11590 | CYP707A24 | 415 aa |  | CYP707A24 ortholog of CYP707A8 tomato |
| 812 | 11613 | CYP710A11 | 501 aa |  | CYP710A11 ortholog |
| 813 | 11628 | CYP710A11 | missing | sequence-missing | CYP710A11 ortholog |
| 814 | 11634 | CYP711A22 | 455 aa | partial | CYP711A22 potato ortholog of CYP711A21 tomato |
| 815 | 11651 | CYP711A22 | missing | sequence-missing | CYP711A22 potato |
| 816 | 11657 | CYP711.1P | 242 aa | pseudogene; internal-stop; short-sequence | CYP711.1P pseudogene 1 potato |
| 817 | 11670 | CYP711.1P | 269 aa | pseudogene; internal-stop; short-sequence | CYP711.1P |
| 818 | 11683 | CYP711.2P | 240 aa | pseudogene; internal-stop; short-sequence | CYP711.2P pseudogene 2 potato |
| 819 | 11694 | CYP711.2P | 267 aa | pseudogene; internal-stop; short-sequence | CYP711.2P |
| 820 | 11707 | CYP711.3P | 180 aa | pseudogene; internal-stop; short-sequence | CYP711.3P pseudogene 3 potato |
| 821 | 11716 | CYP711.3P | 208 aa | internal-stop; short-sequence | CYP711.3P |
| 822 | 11727 | CYP711.4P | 241 aa | pseudogene; internal-stop; short-sequence | CYP711.4P pseudogene 4 potato |
| 823 | 11737 | CYP711.4P | 268 aa | internal-stop; short-sequence | CYP711.4P |
| 824 | 11751 | CYP712G1 | 524 aa | internal-stop | CYP712G1 ortholog CYP712 |
| 825 | 11769 | CYP714A17 | 500 aa | partial | CYP714A17 ortholog CYP714A.1 |
| 826 | 11789 | CYP714E15 | 458 aa | partial | CYP714E15 ortholog CYP714E |
| 827 | 11808 | CYP714G9 | 487 aa |  | CYP714G9 ortholog CYP714G.1 |
| 828 | 11827 | CYP715A15a | 496 aa |  | CYP715A15a CYP715A.1 |
| 829 | 11844 | CYP715A15b | 501 aa |  | CYP715A15b CYP715A.2 |
| 830 | 11863 | CYP716A13 | 472 aa |  | CYP716A13 Solanum tuberosum (potato, Solanales) |
| 831 | 11881 | CYP716A13 | 476 aa |  | CYP716A13    Solanum tuberosum (potato, Solanales) |
| 832 | 11896 | CYP716A46 | 475 aa |  | CYP716A46 ortholog CYP716A.1 |
| 833 | 11914 | CYP716A45 | 486 aa |  | CYP716A45 CYP716A.2 |
| 834 | 11934 | CYP716A42 | 471 aa |  | CYP716A42 ortholog CYP716A.3 |
| 835 | 11952 | CYP716A43 | 478 aa |  | CYP716A43 ortholog CYP716A.4 |
| 836 | 11971 | CYP716C6 | 479 aa |  | CYP716C6 ortholog CYP716C.1 |
| 837 | 11989 | CYP716H1v1 | 474 aa |  | CYP716H1v1     Solanum tuberosum (potato) |
| 838 | 12004 | CYP716H1v2 | 472 aa |  | CYP716H1v2 Solanum tuberosum (potato) |
| 839 | 12022 | CYP716H3P | 438 aa | partial; frameshift | CYP716H3P |
| 840 | 12046 | CYP716H4b | 479 aa |  | CYP716H4b CYP716H.1 |
| 841 | 12064 | CYP716H4a | 479 aa |  | CYP716H4a CYP716H.2 |
| 842 | 12081 | CYP716H5P | 465 aa |  | CYP716H5P CYP716H.3P |
| 843 | 12100 | CYP716Q3 | 480 aa | partial; uncertain-boundary | CYP716Q3 CYP716x.1 |
| 844 | 12119 | CYP716Q1 | 474 aa |  | CYP716Q1 ortholog CYP716x.2 |
| 845 | 12135 | CYP716 | 142 aa | short-sequence | CYP716      Solanum tuberosum (potato, Solanales) |
| 846 | 12148 | CYP716x.3P | 450 aa | uncertain-boundary; internal-stop | CYP716x.3P |
| 847 | 12166 | CYP718A7 | 493 aa | partial | CYP718A7   Solanum tuberosum (potato) |
| 848 | 12186 | CYP718A7 | missing | sequence-missing | CYP718A7 Solanum tuberosum (potato) |
| 849 | 12192 | CYP720A1 | 424 aa | partial | CYP720A1 ortholog CYP720A |
| 850 | 12214 | CYP721A.1 | missing | sequence-missing | CYP721A.1 |
| 851 | 12217 | CYP721A.1 | 104 aa | short-sequence | CYP721A.1 |
| 852 | 12225 | CYP721A27v2 | 350 aa |  | CYP721A27v2       Solanum tuberosum (potato, Solanales) |
| 853 | 12239 | CYP721A27 | 481 aa |  | CYP721A27 ortholog CYP721A.1 |
| 854 | 12258 | CYP721A26 | 414 aa |  | CYP721A26 ortholog CYP721A.2 |
| 855 | 12270 | CYP722A1 | 432 aa | partial | CYP722A1 ortholog CYP722A.1 |
| 856 | 12293 | CYP722C1 | 390 aa | partial | CYP722C1 ortholog CYP722C.1 |
| 857 | 12314 | CYP724B2 | 474 aa | partial | CYP724B2 |
| 858 | 12336 | CYP724B16 | 430 aa |  | CYP724B16 ortholog CYP724B.2 |
| 859 | 12355 | CYP724B17 | 466 aa | partial | CYP724B17 CYP724C.1 subfamily is new in potato tomato |
| 860 | 12377 | CYP728B22b | 469 aa |  | CYP728B22b CYP728B.1 paralog of CYP728B22 |
| 861 | 12393 | CYP728B22a | 475 aa | pseudogene | CYP728B22a CYP728B.2 paralog of CYP728B22 |
| 862 | 12411 | CYP733A1 | 424 aa | pseudogene; partial; frameshift | CYP733A1 ortholog CYP733A.1v1 |
| 863 | 12431 | CYP733A.1v2 | 478 aa |  | CYP733A.1v2 |
| 864 | 12451 | CYP734A22 | 484 aa | partial | CYP734A22 ortholog CYP734A.1 |
| 865 | 12472 | CYP734A7 | 468 aa |  | CYP734A7 ortholog CYP734A.2 |
| 866 | 12492 | CYP734A8 | 525 aa |  | CYP734A8 ortholog CYP734A.3 |
| 867 | 12512 | CYP734A.4P | 110 aa | short-sequence | CYP734A.4P |
| 868 | 12521 | CYP735A8 | 514 aa |  | CYP735A8   Solanum tuberosum (potato) |
| 869 | 12534 | CYP735A8v1 | 490 aa | partial | CYP735A8v1 |
| 870 | 12551 | CYP735A8v2 | 104 aa | short-sequence | CYP735A8v2 fragment |
| 871 | 12560 | CYP735A20P | 492 aa | partial | CYP735A20P CYP735A.1P pseudogene |
| 872 | 12582 | CYP736A66bP | 503 aa |  | CYP736A66bP CYP736A.1Pv1 |
| 873 | 12601 | CYP736A.1Pv2 | 46 aa | short-sequence | CYP736A.1Pv2 |
| 874 | 12609 | CYP736A67b | 494 aa |  | CYP736A67b ortholog CYP736A.2v1 |
| 875 | 12624 | CYP736A.2v2 | 85 aa | short-sequence | CYP736A.2v2 |
| 876 | 12632 | CYP736A68 | 496 aa | partial | CYP736A68 ortholog CYP736A.3 |
| 877 | 12649 | CYP736A79eP | 204 aa | short-sequence | CYP736A79eP CYP736A.4P |
| 878 | 12661 | CYP736A79fP | 204 aa | short-sequence | CYP736A79fP CYP736A.5P |
| 879 | 12672 | CYP736A79a | 497 aa |  | CYP736A79a CYP736A.6v1 |
| 880 | 12687 | CYP736A.6v2 | 88 aa | short-sequence | CYP736A.6v2 |
| 881 | 12696 | CYP736A79bP | 293 aa | short-sequence | CYP736A79bP CYP736A.7P |
| 882 | 12708 | CYP736A79c | 428 aa | partial | CYP736A79c CYP736A.8v1 |
| 883 | 12725 | CYP736A.8v2 | 35 aa | short-sequence | CYP736A.8v2 |
| 884 | 12733 | CYP736A79d | 504 aa |  | CYP736A79d CYP736A.9 |
| 885 | 12750 | CYP736A80P | 342 aa | partial; short-sequence | CYP736A80P CYP736A.10P |
| 886 | 12767 | CYP736A67a | 318 aa | short-sequence | CYP736A67a CYP736A.11 |
| 887 | 12781 | CYP736A66a | 494 aa | partial; frameshift | CYP736A66a CYP736A.12 |
| 888 | 12798 | CYP736A81bP | 303 aa | short-sequence | CYP736A81bP CYP736A.13P |
| 889 | 12813 | CYP736A64b | 492 aa | partial | CYP736A64b CYP736A.14v1 |
| 890 | 12829 | CYP736A.14v2 | 492 aa | partial | CYP736A.14v2 old CYP736A.39v1 |
| 891 | 12845 | CYP736A.14v3 | 319 aa | partial; short-sequence | CYP736A.14v3 |
| 892 | 12859 | CYP736.14 | missing | sequence-missing | CYP736.14 |
| 893 | 12862 | CYP736.14 | 313 aa | partial; short-sequence | CYP736.14 |
| 894 | 12875 | CYP736A.14v4 | 40 aa | short-sequence | CYP736A.14v4 CYP736A.39v2 |
| 895 | 12883 | CYP736A65bP | 132 aa | internal-stop; short-sequence | CYP736A65bP CYP736A.15P |
| 896 | 12893 | CYP736A81aP | 37 aa | short-sequence | CYP736A81aP CYP736A.16P |
| 897 | 12903 | CYP736A61b | 494 aa | partial; frameshift | CYP736A61b CYP736A.17v1 |
| 898 | 12919 | CYP736A.17v2 | 48 aa | short-sequence | CYP736A.17v2 |
| 899 | 12927 | CYP736A64a | 494 aa | partial | CYP736A64a CYP736A.18 |
| 900 | 12944 | CYP736A61aP | 427 aa |  | CYP736A61aP CYP736A.19P |
| 901 | 12961 | CYP736A65a | 493 aa |  | CYP736A65a CYP736A.20 |
| 902 | 12979 | CYP736A89 | 497 aa | partial | CYP736A89 CYP736A.21 |
| 903 | 12997 | CYP736A76a | 497 aa | partial; frameshift | CYP736A76a CYP736A.22 |
| 904 | 13015 | CYP736A76b | 498 aa | partial; frameshift | CYP736A76b CYP736A.23P |
| 905 | 13034 | CYP736A82 | 498 aa |  | CYP736A82 CYP736A.24 |
| 906 | 13051 | CYP736A82-de1b | 264 aa | internal-stop; short-sequence | CYP736A82-de1b CYP736A.25P |
| 907 | 13064 | CYP736A88 | 494 aa | partial; frameshift | CYP736A88 CYP736A.26 |
| 908 | 13082 | CYP736A77 | 498 aa |  | CYP736A77 CYP736A.27 |
| 909 | 13099 | CYP736A78 | 492 aa |  | CYP736A78 ortholog CYP736A.28 |
| 910 | 13116 | CYP736A74 | 493 aa |  | CYP736A74 ortholog CYP736A.29 |
| 911 | 13133 | CYP736A90P | 435 aa | partial | CYP736A90P CYP736A.30P |
| 912 | 13151 | CYP736A91P | 204 aa | short-sequence | CYP736A91P CYP736A.31P |
| 913 | 13160 | CYP736A | 225 aa | short-sequence | CYP736A    Solanum tuberosum (potato, Solanales) |
| 914 | 13171 | CYP736A92 | 494 aa |  | CYP736A92 CYP736A.32 |
| 915 | 13189 | CYP736A93 | 494 aa |  | CYP736A93 CYP736A.33v1 |
| 916 | 13205 | CYP736A.33v2 | 74 aa | short-sequence | CYP736A.33v2 |
| 917 | 13214 | CYP736A59 | 493 aa |  | CYP736A59 ortholog CYP736A.34 |
| 918 | 13232 | CYP736A60 | 504 aa | partial; frameshift | CYP736A60 ortholog CYP736A.35 |
| 919 | 13249 | CYP736A94P | 496 aa | partial; internal-stop | CYP736A94P CYP736A.36P |
| 920 | 13267 | CYP736A72 | 507 aa |  | CYP736A72 ortholog CYP736A.37 |
| 921 | 13285 | CYP736A95P | 474 aa | internal-stop | CYP736A95P CYP736A.38P |
| 922 | 13303 | CYP736A96P | 170 aa | short-sequence | CYP736A96P CYP736A.40P |
| 923 | 13314 | CYP749A20 | 506 aa |  | CYP749A20 ortholog CYP749.1 |
| 924 | 13335 | CYP749A19 | 482 aa |  | CYP749A19 ortholog CYP749.2 |
| 925 | 13356 | CYP749.3 | 290 aa | short-sequence | CYP749.3 pseudogene |
| 926 | 13374 | CYP749.4 | 209 aa | short-sequence | CYP749.4 pseudogene |
| 927 | 13384 | CYP749.5P | 454 aa | partial; internal-stop | CYP749.5P pseudogene |
| 928 | 13409 | CYP749.6P | 146 aa | short-sequence | CYP749.6P pseudogene |
| 929 | 13418 | CYP749A.7P | 134 aa | short-sequence | CYP749A.7P |
| 930 | 13429 | CYP749A.8P | 561 aa | internal-stop | CYP749A.8P |
| 931 | 13460 | CYP85A | 226 aa | partial; ambiguous-residue; short-sequence | CYP85A     Solanum chacoense (Chaco potato) |
| 932 | 13468 | CYP721 | 125 aa | partial; short-sequence | CYP721       Solanum chacoense (chaco potato, Solanales) |
