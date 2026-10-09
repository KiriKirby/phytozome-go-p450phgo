# PGD schema

The database is bbolt. It contains a `records` bucket. Keys are stable zero-padded record numbers and values are JSON objects with these fields:

`ID`, optional `RecordKey`, `Category`, `Species`, `Symbol`, `Description`, `Sequence`, and `SourceURL`.

`ID` and `Symbol` preserve the biological/source identifier shown to users. `RecordKey` is a stable row-level key used when one resource contains several independently reviewed blocks with the same `ID`; it prevents FASTA lookup from resolving a duplicate name to the wrong block. Runtime readers fall back to `ID` for older records without `RecordKey`.

Empty source fields remain empty. `Category` is one of `animals`, `plants`, `fungi`, or `bacteria`. This format preserves the maximum information available from the source materials without requiring the application to access Dr Nelson online during search.

Plant release records are admitted only from reviewed resource-specific CSV inputs under `sources/reviewed/plants/`. The legacy normalized-text scanner is not a release input.

The optional `species` bucket contains JSON objects with `name`, `category`, `selectable`, and `description`. It lists every discovered species, including species without searchable records; those entries use `selectable: false` and remain visible but disabled in PHgo.
