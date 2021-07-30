package katas

import (
	"fmt"
	"strings"
)

type Symbol struct {
	symbol1 rune
	symbol2 rune
}

var (
	symbols = []Symbol{
		{symbol1: 'A', symbol2: 'T'},
		{symbol1: 'C', symbol2: 'G'},
	}
	dnaReplacer *strings.Replacer = strings.NewReplacer(
		"A", "T",
		"T", "A",
		"C", "G",
		"G", "C",
	)
)

func RunKata1() {
	PrintWithDifferentsMethod("ATTGC") // TAACG
	PrintWithDifferentsMethod("GTAT")  // CATA
}

func PrintWithDifferentsMethod(dna string) {
	fmt.Printf("1: %s\n", DNAStrand1(dna))
	fmt.Printf("2: %s\n", DNAStrand2(dna))
}

func DNAStrand1(dna string) string {
	res := ""
dnaFor:
	for _, c := range dna {
		for _, symbol := range symbols {
			if c == symbol.symbol1 {
				res += string(symbol.symbol2)
				continue dnaFor
			}
			if c == symbol.symbol2 {
				res += string(symbol.symbol1)
				continue dnaFor
			}
		}
		res += string(c)
	}
	return res
}

func DNAStrand2(dna string) string {
	return dnaReplacer.Replace(dna)
}
