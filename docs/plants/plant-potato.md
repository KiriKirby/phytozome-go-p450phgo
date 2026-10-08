# Plant resource audit: potato

- Source file: `plants-potato.P450s.doc`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/potato.P450s.doc
- Detected container: `Word document normalized to text`
- Resource-local records: `432`
- Records with accepted sequence: `77`
- Extraction method used for this resource: `FASTA block; explicit sequence column`
- Detected sequence-column header(s): ``
- Rows with a non-empty sequence-column value: `3`

## Review rule

This resource is reviewed independently. FASTA extraction requires a CYP-bearing `>` header and sequence lines bounded by the next CYP header, a terminal `*`, or the first non-sequence annotation. Spreadsheet data is accepted only from an explicit `sequence` column. Alignment text, coordinates, descriptions, and unlabeled short fragments are rejected.

Missing sequences remain empty until this exact source file is manually verified; no other database is used as a substitute.
