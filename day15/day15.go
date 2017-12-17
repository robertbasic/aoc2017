package day15

import (
	"encoding/hex"
	"fmt"
	"log"
	"math"
)

var af = 16807
var bf = 48271
var d = 2147483647

// Day15 solves the puzzles for day 15
func Day15(logger *log.Logger) {
	var a = 634
	var b = 301

	p := matchingPairs(40000000, a, b)
	logger.Println("Matching pairs: ", p)

	p = matchingPairsSpecRules(5000000, a, b)
	logger.Println("Matching pairs: ", p)
}

func matchingPairs(pairs int, a int, b int) int {
	var p int

	for i := 0; i < pairs; i++ {
		a = pair(a, af)
		b = pair(b, bf)
		ab := fmt.Sprintf("%016b", a)
		bb := fmt.Sprintf("%016b", b)

		ae := ab[len(ab)-16 : len(ab)]
		be := bb[len(bb)-16 : len(bb)]

		if ae == be {
			p++
		}
	}

	return p
}

func matchingPairsSpecRules(pairs int, a int, b int) int {
	var p int

	for i := 0; i < pairs; i++ {
		a = divPair(a, af, 4)
		b = divPair(b, bf, 8)
		ab := fmt.Sprintf("%016b", a)
		bb := fmt.Sprintf("%016b", b)

		ae := ab[len(ab)-16 : len(ab)]
		be := bb[len(bb)-16 : len(bb)]

		if ae == be {
			p++
		}
	}

	return p
}

func pair(i int, f int) int {
	p := math.Mod(float64(i*f), float64(d))
	return int(p)
}

func divPair(i int, f int, div int) int {
	var z bool
	for !z {
		i = int(math.Mod(float64(i*f), float64(d)))
		z = math.Mod(float64(i), float64(div)) == 0
	}
	return i
}

func dtob(d int) []byte {
	src := []byte(string(d))

	dst := make([]byte, hex.DecodedLen(len(src)))
	hex.Decode(dst, src)

	return dst
}
