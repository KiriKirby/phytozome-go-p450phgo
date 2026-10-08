# PGD generator

Run from this directory:

```text
go run ./cmd/build -out p450phgo.pgd
```

The generator downloads the four Dr Nelson category indexes, resolves their published links, deduplicates records, and writes the runtime bbolt schema described in `DATA_SCHEMA.md`. Resource-specific extraction can be added to this repository without changing the application connector. Review generated records before publishing a release, especially legacy Word resources whose text may require manual conversion.

FASTA extraction is strictly local to the downloaded Dr. Nelson resources. The
normalized Office text may contain bare CR line endings and wrapped FASTA
records, so the parser normalizes CR/CRLF, recognizes a `>` header containing a
CYP name, and consumes amino-acid lines until the next header. A sequence is
published only when it is actually present in that resource (at least 30 amino
acids). No UniProt, NCBI, or other external sequence source is used. Generated
species pages report `present (N aa)` or `missing` for every record, making
extraction gaps auditable against the original Dr. Nelson file.
