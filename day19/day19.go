package day19

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

var directions = map[int][]int{
	1: {0, 1},
	2: {1, 0},
	3: {0, -1},
	4: {-1, 0},
}

// Day19 solves the puzzles for day 19
func Day19(folder string) {
	file, _ := os.Open(folder + "/day19.txt")

	l, s := walk(file, 200)
	log.Println("The letters are: ", l)
	log.Println("Number of steps: ", s)
}

func walk(tubes io.Reader, size int) (string, int) {
	var letters string
	var steps int
	var end bool
	direction := 2
	tubemap := maptubes(tubes, size)

	sign := '|'
	ri := 0
	ci := strings.IndexRune(tubemap[0], sign)
	for !end {
		steps++
		sign, ri, ci = move(ri, ci, direction, tubemap)

		if sign == '|' || sign == '-' {
			continue
		}

		if sign >= 'A' && sign <= 'Z' {
			letters += string(sign)
			continue
		}

		if sign == '+' {
			direction = turn(ri, ci, direction, tubemap)
		}

		if sign == ' ' {
			end = true
		}
	}

	return letters, steps
}

func move(r int, c int, dir int, tubemap []string) (rune, int, int) {
	r += directions[dir][0]
	c += directions[dir][1]

	if r >= len(tubemap) {
		return ' ', r, c
	}

	if c >= len(tubemap[r]) {
		return ' ', r, c
	}

	return rune(tubemap[r][c]), r, c
}

func turn(r int, c int, dir int, tubemap []string) int {
	var nd int

	vd := func(d int) bool {
		nr := r + directions[d][0]
		nc := c + directions[d][1]

		if nr >= len(tubemap) {
			return false
		}

		if nc >= len(tubemap[nr]) {
			return false
		}

		if tubemap[nr][nc] == ' ' {
			return false
		}
		return true
	}

	nd = dir - 1

	if nd == 0 {
		nd = 4
	}

	if vd(nd) {
		return nd
	}

	nd = dir + 1

	if nd == 5 {
		nd = 1
	}

	if vd(nd) {
		return nd
	}

	panic("Can't turn")
}

func maptubes(tubes io.Reader, size int) []string {
	scanner := bufio.NewScanner(tubes)

	tubemap := make([]string, size)

	i := 0
	for scanner.Scan() {
		tube := scanner.Text()

		if strings.TrimSpace(tube) == "" {
			continue
		}

		tubemap[i] = tube
		i++
	}

	return tubemap
}
