package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// Day6 solves the day 6 puzzles
func Day6() {
	i := [16]int{4, 10, 4, 1, 8, 4, 9, 14, 5, 1, 14, 15, 0, 15, 3, 5}
	c := Cycle(i[:])
	fmt.Println("Cycles: ", c)
}

// Cycle finds the number of cycles required
// after which the banks have a previosly
// seen blocks configuration
func Cycle(input []int) int {
	var cycles int
	var l = len(input)

	blocks := make(map[string]bool)

	for {
		cycles++
		p, i := Highest(input)
		input[p] = 0

		for i > 0 {
			p = nextp(p, l)
			input[p]++
			i--
		}

		k := getk(input)
		fmt.Println(k)
		if _, ok := blocks[k]; ok {
			break
		}

		blocks[k] = true
	}

	return cycles
}

// Highest finds the highest number in input
// and it's position in the slice
func Highest(input []int) (int, int) {
	var ph, ih int

	for p, i := range input {
		if i > ih {
			ph = p
			ih = i
		}
	}

	return ph, ih
}

func nextp(p int, l int) int {
	p++

	if p >= l {
		p = 0
	}

	return p
}

func getk(input []int) string {
	var b bytes.Buffer

	for p, i := range input {
		b.WriteString(fmt.Sprintf("%s:%s", strconv.Itoa(p), strconv.Itoa(i)))
	}

	return b.String()
}
