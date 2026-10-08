# PGD generator

Run from this directory:

```text
go run ./cmd/build -out p450phgo.pgd
```

The generator downloads the four Dr Nelson category indexes, resolves their published links, deduplicates records, and writes the runtime bbolt schema described in `DATA_SCHEMA.md`. Resource-specific extraction can be added to this repository without changing the application connector. Review generated records before publishing a release, especially legacy Word resources whose text may require manual conversion.
