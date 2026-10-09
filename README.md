# p450phgo.pgd data repository

This repository publishes the read-only `p450phgo.pgd` bbolt database consumed by phytozome GO's CYP keyword connector. The application never builds this database from the web pages.

The database is being rebuilt category by category, starting with every resource on the Dr. Nelson plants page. Each plant file is downloaded and interpreted independently; a generic CYP/FASTA scan is not accepted as release data. Completed resource parsers write reviewed CSV files under `sources/reviewed/plants/`, with source hashes and per-block provenance.

The application downloads the release asset on first CYP selection and stores it beside the executable as `p450phgo.pgd`. Set `PHGO_CYP_PGD_URL` during testing to point at a release asset or raw file.

The generator and source-analysis documentation belong here, alongside the generated PGD. Keep the generated database immutable between releases and publish a checksum/manifest with each update. Species whose resource has not completed its resource-specific review remain listed but disabled.

The current GitHub release asset is consumed through `manifest.json`; PHgo verifies both its byte length and SHA-256 before atomically replacing the installed database.
