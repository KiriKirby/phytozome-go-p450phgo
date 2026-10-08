# PGD schema

The database is bbolt. It contains a `records` bucket. Keys are stable zero-padded record numbers and values are JSON objects with these fields:

`ID`, `Category`, `Species`, `Symbol`, `Description`, `Sequence`, and `SourceURL`.

Empty source fields remain empty. `Category` is one of `animals`, `plants`, `fungi`, or `bacteria`. This format preserves the maximum information available from the source materials without requiring the application to access Dr Nelson online during search.
