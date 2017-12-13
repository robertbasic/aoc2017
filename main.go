package main

import (
	"strconv"
	"strings"
)

func main() {
	//Day1()

	//Day2()

	//Day3()

	// Day4()

	// Day5()

	// Day6()

	// Day7()

	// Day8()

	// Day9()

	// Day10()

	// Day11()

	// Day12()

	Day13()
}

// StringNumberToInt converts a string number to an int
func StringNumberToInt(s string) int {
	i, _ := strconv.Atoi(strings.TrimSpace(s))
	return i
}

// MapHasIntKeyBoolVal checks if map[int]bool has key present
func MapHasIntKeyBoolVal(k int, m map[int]bool) bool {
	if _, ok := m[k]; ok {
		return true
	}
	return false
}
