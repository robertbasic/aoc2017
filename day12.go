package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Program holds the parsed message
// from which number to which numbers
type Program struct {
	from int
	to   []int
}

func (p *Program) parse(in string) {
	in = strings.Replace(in, "<->", ",", -1)

	parts := strings.Split(in, ",")

	from := StringNumberToInt(parts[0])
	to := []int{}

	for _, v := range parts[1:] {
		i := StringNumberToInt(v)
		to = append(to, i)
	}

	p.from = from
	p.to = to
}

// Day12 solves the puzzles for day 12
func Day12(logger *log.Logger) {
	file, _ := os.Open("./inputs/day12.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := []string{}

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	c := CountProgramsToZero(input)
	logger.Println("Pointing to zero:", c)

	g := CountProgramGroups(input)
	logger.Println("Program groups: ", g)
}

// CountProgramsToZero counts the number of messages
// that point to zero
func CountProgramsToZero(input []string) int {
	programs := BuildPrograms(input)
	pointers := []int{}
	seen := make(map[int]bool)

	pointers, seen = BuildPointers(programs[0], 0, pointers, seen, programs)

	return len(pointers)
}

// CountProgramGroups counts the unique program groups
func CountProgramGroups(input []string) int {
	programs := BuildPrograms(input)

	groups := make(map[int][]int)
	seen := make(map[int]bool)

	for from := range programs {
		groups, seen = BuildGroups(from, programs, groups, seen)
	}

	return len(groups)
}

// BuildGroups builds the distinct program groups
func BuildGroups(from int, programs map[int][]int, groups map[int][]int, seen map[int]bool) (map[int][]int, map[int]bool) {
	pointers := []int{}
	pointers, seen = BuildPointers(programs[from], from, pointers, seen, programs)

	if len(pointers) > 0 {
		groups[from] = pointers
	}

	return groups, seen
}

// BuildPointers builds a slice of numbers that point to zero
func BuildPointers(tos []int, parent int, pointers []int, seen map[int]bool, programs map[int][]int) ([]int, map[int]bool) {
	for _, to := range tos {
		if MapHasIntKeyBoolVal(to, seen) {
			continue
		}

		pointers = append(pointers, to)
		seen[to] = true
		pointers, seen = BuildPointers(programs[to], to, pointers, seen, programs)
	}

	return pointers, seen
}

// BuildPrograms builds the message routes
func BuildPrograms(input []string) map[int][]int {
	programs := make(map[int][]int)

	for _, s := range input {
		p := Program{}
		p.parse(s)

		programs[p.from] = p.to
	}

	return programs
}
