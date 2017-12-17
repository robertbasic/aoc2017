package day14

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"

	"github.com/robertbasic/aoc2017/day10"
)

// Day14 solves the puzzles for day 14
func Day14(logger *log.Logger) {
	var bits int
	var binrows = make([]string, 128)
	for i := 0; i < 128; i++ {
		h := day10.ElvenKnottedHash("amgozmfv-" + strconv.Itoa(i))
		bs := htob(h)

		for _, b := range bs {
			binrow := fmt.Sprintf("%08b", b)

			bits += countOnes(binrow)

			binrows[i] += binrow
		}
	}

	g := CountGroups(binrows)

	fmt.Println("Number of used squares: ", bits)
	fmt.Println("Number of groups: ", g)
}

func CountGroups(rows []string) int {
	groups := buildGroups(rows)

	var mg int
	for _, g := range groups {
		if g > mg {
			mg = g
		}
	}

	return mg
}

func buildGroups(rows []string) map[string]int {
	var cg int
	var group = make(map[string]int)

	for i, row := range rows {
		for j, b := range row {
			if b != 49 {
				continue
			}

			if _, ok := group[k(i, j)]; ok {
				continue
			}

			cg++
			cg, group = visit(i, j, rows, cg, group)
		}
	}

	return group
}

func visit(i int, j int, rows []string, cg int, group map[string]int) (int, map[string]int) {
	if i < 0 || j < 0 || i > len(rows)-1 || j > len(rows)-1 {
		return cg, group
	}

	if rows[i][j] != 49 {
		return cg, group
	}

	if g, ok := group[k(i, j)]; ok {
		return g, group
	}

	group[k(i, j)] = cg

	cg, group = visit(i, j+1, rows, cg, group)
	cg, group = visit(i+1, j, rows, cg, group)
	cg, group = visit(i, j-1, rows, cg, group)
	cg, group = visit(i-1, j, rows, cg, group)

	return cg, group
}

func k(i int, j int) string {
	return fmt.Sprintf("%d:%d", i, j)
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
