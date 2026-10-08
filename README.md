# p450phgo.pgd data repository

This repository publishes the read-only `p450phgo.pgd` bbolt database consumed by phytozome GO's CYP keyword connector. The application never builds this database from the web pages. A release must contain the complete four-category Dr Nelson dataset (animals, plants, fungi, bacteria), including records with missing fields as empty values.

The application downloads the release asset on first CYP selection and stores it beside the executable as `p450phgo.pgd`. Set `PHGO_CYP_PGD_URL` during testing to point at a release asset or raw file.

The generator and source-analysis documentation belong here, alongside the generated PGD. Keep the generated database immutable between releases and publish a checksum/manifest with each update.
