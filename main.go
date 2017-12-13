package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	logger := log.New(os.Stdout, "", 0)

	Day1(logger)

	Day2(logger)

	Day3(logger)

	Day4(logger)

	Day5(logger)

	Day6(logger)

	Day7(logger)

	Day8(logger)

	Day9(logger)

	Day10(logger)

	Day11(logger)

	Day12(logger)

	Day13(logger)
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
