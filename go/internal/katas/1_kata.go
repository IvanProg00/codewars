package katas

import (
	"strings"
)

var dnaReplacer *strings.Replacer = strings.NewReplacer(
	"T", "A",
	"A", "T",
	"C", "G",
	"G", "C",
)

func DNAStrand(dna string) string {
	return dnaReplacer.Replace(dna)
}
