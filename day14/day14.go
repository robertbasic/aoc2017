package day14

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/robertbasic/aoc2017/day10"
)

// Day14 solves the puzzles for day 14
func Day14() {
	var bits int
	for i := 0; i < 128; i++ {
		h := knotHash("amgozmfv-" + strconv.Itoa(i))
		bs := htob(h)

		for _, b := range bs {
			bits += countOnes(fmt.Sprintf("%08b\n", b))
		}
	}
	fmt.Println(bits)
}

func countOnes(input string) int {
	var ones int

	for _, i := range input {
		if i == 49 {
			ones++
		}
	}

	return ones
}

func htob(h string) []byte {
	src := []byte(h)

	dst := make([]byte, hex.DecodedLen(len(src)))
	hex.Decode(dst, src)

	return dst
}

func knotHash(input string) string {
	var list = make([]int, 256)
	for l := 0; l < 256; l++ {
		list[l] = l
	}

	var lengths = make([]int, 0)
	for _, c := range input {
		lengths = append(lengths, int(c))
	}

	lengths = append(lengths, []int{17, 31, 73, 47, 23}...)

	_, hashedlist := day10.ElvenHashChecksum(list, lengths, 64)

	return day10.DenseHash(hashedlist)
}
