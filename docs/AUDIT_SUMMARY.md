# Complete Dr. Nelson resource audit

This summary is generated from the current normalized resource set. Every listed species/resource is represented in its own Markdown file under `species/` or `resources/`. No external sequence database is consulted.

Species/resource entries: 167

| Category | PGD records | Verified sequences |
|---|---:|---:|
| animals | 3542 | 1933 |
| plants | 9085 | 2523 |
| fungi | 3233 | 2078 |
| bacteria | 2850 | 2137 |

Sequence acceptance rules:

- Real FASTA blocks: `>` header containing a CYP identifier, followed by wrapped amino-acid lines until the next header.
- Structured tables: sequence is accepted only from an explicitly labeled `sequence` or `protein sequence` column.
- Scores, alignment snippets, comments, accession text, and ambiguous unlabeled fields are not sequences.
- Missing or ambiguous sequences remain empty in PGD and are documented as `missing`; they are never filled from another database.
