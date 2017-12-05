package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Day5 solves the puzzles for day 5
func Day5() {
	instructions := make([]int, 0)
	ci := make([]int, 0)

	file, _ := os.Open("./inputs/day5.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input, _ := strconv.Atoi(scanner.Text())
		instructions = append(instructions, input)
		ci = append(ci, input)
	}

	j := Jumps(instructions)

	fmt.Println("Number of jumps: ", j)

	cj := CrazyJumps(ci)

	fmt.Println("Number of crazy jumps: ", cj)
}

// Jumps finds the number of jumps
// to jump out of the maze
func Jumps(instructions []int) int {
	var jumps int
	var pos int
	var ins int
	var out bool
	var l = len(instructions)

	for !out {
		if pos >= l {
			break
		}

		ins = instructions[pos]
		instructions[pos]++
		pos = pos + ins

		jumps++
	}

	return jumps
}

// CrazyJumps finds the number of jumps
// to jump out of the maze
func CrazyJumps(cis []int) int {
	var jumps int
	var pos int
	var ins int
	var out bool
	var l = len(cis)

	for !out {
		if pos >= l {
			break
		}

		ins = cis[pos]
		if ins < 3 {
			cis[pos]++
		} else {
			cis[pos]--
		}
		pos = pos + ins

		jumps++
	}

	return jumps
}
