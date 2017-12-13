package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day2(logger *log.Logger) {
	file, _ := os.Open("./inputs/day2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cs := CheckSum(scanner)
	logger.Println("The checksum is: ", cs)

	file, _ = os.Open("./inputs/day2.txt")
	defer file.Close()
	scanner = bufio.NewScanner(file)
	cs = CheckSumDivs(scanner)
	logger.Println("The divisible checksum is: ", cs)
}

func CheckSum(scanner *bufio.Scanner) int {
	var cs int
	var d int

	for scanner.Scan() {
		d = RowDiff(scanner.Text())
		cs = cs + d
	}

	return cs
}

func CheckSumDivs(scanner *bufio.Scanner) int {
	var cs int
	var d int

	for scanner.Scan() {
		d = RowEvenDiv(scanner.Text())
		cs = cs + d
	}

	return cs
}

func RowDiff(row string) int {
	var diff int
	var h int
	var l int
	ints := strings.Fields(row)

	for _, s := range ints {
		i, _ := strconv.Atoi(s)
		if i < l || l == 0 {
			l = i
		}
		if i > h {
			h = i
		}
	}

	diff = h - l

	return diff
}

func RowEvenDiv(row string) int {
	var div int

	ints := strings.Fields(row)

	for i := 0; i < len(ints); i++ {
		dividend, _ := strconv.ParseFloat(ints[i], 64)

		for j := 0; j < len(ints); j++ {
			if i == j {
				continue
			}

			divisor, _ := strconv.ParseFloat(ints[j], 64)

			if divisor > dividend {
				continue
			}

			if math.Mod(dividend, divisor) == 0.0 {
				d := dividend / divisor
				div = div + int(d)
			}
		}
	}

	return div
}
