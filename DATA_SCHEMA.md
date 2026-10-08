# PGD schema

The database is bbolt. It contains a `records` bucket. Keys are stable zero-padded record numbers and values are JSON objects with these fields:

`ID`, `Category`, `Species`, `Symbol`, `Description`, `Sequence`, and `SourceURL`.

Empty source fields remain empty. `Category` is one of `animals`, `plants`, `fungi`, or `bacteria`. This format preserves the maximum information available from the source materials without requiring the application to access Dr Nelson online during search.

The optional `species` bucket contains JSON objects with `name`, `category`, `selectable`, and `description`. It lists every discovered species, including species without searchable records; those entries use `selectable: false` and remain visible but disabled in PHgo.
