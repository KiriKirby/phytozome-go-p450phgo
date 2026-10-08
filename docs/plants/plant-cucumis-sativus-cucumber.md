# Plant resource audit: Cucumis sativus (cucumber)

- Source file: `plants-Cucumis.sativus.xlsx`
- URL: https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/Cucumis.sativus.xlsx
- Detected container: `Excel workbook normalized to text`
- Resource-local records: `204`
- Records with accepted sequence: `201`
- Extraction method used for this resource: `explicit sequence column`
- Detected sequence-column header(s): `sequence`
- Rows with a non-empty sequence-column value: `229`

## Review rule

This resource is reviewed independently. FASTA extraction requires a CYP-bearing `>` header and sequence lines bounded by the next CYP header, a terminal `*`, or the first non-sequence annotation. Spreadsheet data is accepted only from an explicit `sequence` column. Alignment text, coordinates, descriptions, and unlabeled short fragments are rejected.

Missing sequences remain empty until this exact source file is manually verified; no other database is used as a substitute.
