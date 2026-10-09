# Plant resource review: Ostreococcus

- Source file: `plants-ostreococcus.doc`
- Source URL: `https://drnelson.uthsc.edu/wp-content/uploads/sites/130/resources/ostreococcus.doc`
- Source SHA-256: `12c744823cab1fd85471b10f024100b2438f4465e61e314f5f7e8367b06ef622`
- Physical `>CYP` headers: `30`
- Accepted records: `30`
- Actual species split: `10` Ostreococcus tauri, `11` Ostreococcus lucimarinus, `9` Ostreococcus RCC809
- Review status: `complete`

The preamble says each of the three species has ten P450s, but the physical headers do not have a 10/10/10 split: O. lucimarinus includes two explicit CYP97A15 model blocks while the RCC809 section contains nine headers. All 30 source blocks remain distinct rather than being merged or padded to force the prose count. Species are assigned from the source header identifiers and exact RCC809 section boundary, not inferred from similarity text. Each block is bounded by the next `>CYP` header; coordinate numbers in the two `CYP800A1` layouts are removed only as line-edge coordinates, while residue letters, X/x, internal stops and duplicate/alternate models remain literal. Only an explicit terminal `*` is removed.
