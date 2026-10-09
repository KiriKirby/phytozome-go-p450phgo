# Plant resource review: Volvox carteri

- Source file: `plants-volvox.doc`
- Source URL: `https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/volvox.doc`
- Source SHA-256: `594311939b4d10dd31680a84e12e7848f41658fc67603d344a04fbcec7167311`
- Physical `>` headers inspected: `23`
- Accepted Volvox CYP blocks: `19`
- Records with literal protein: `19`
- Excluded comparison records: `4` (two Chlamydomonas reconstructions, one bacterial protein, one human CYP7B1)
- Review status: `complete`

Every accepted Volvox block ends in an explicit terminal `*`. The resource-specific parser reads only pure residue lines and the document's observed coordinate/phase layouts until that stop, removes only the terminal marker, and does not bridge the explicit sequence-gap prose. Chlamydomonas CYP771A1/CYP772A1 reconstructions, their EST excerpts and alignments, the Plesiocystis bacterial protein, and human CYP7B1 remain comparison evidence only. Source order and literal lowercase `hq` in the Chlamydomonas comparison are inspected but the latter is not part of an accepted Volvox record.
