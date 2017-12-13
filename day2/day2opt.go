package day2

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day2Opt(logger *log.Logger, folder string) {
	file, _ := os.Open(folder + "/day2.txt")
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	scanner2 := bufio.NewScanner(file)

	cs := CheckSumOpt(scanner1)
	logger.Println("The checksum is: ", cs)

	csd := CheckSumDivs(scanner2)
	logger.Println("The divisible checksum is: ", csd)
}

func CheckSumOpt(scanner *bufio.Scanner) int {
	var cs int
	var d int

	for scanner.Scan() {
		d = RowDiffOpt(scanner.Text())
		cs = cs + d
	}

	return cs
}

func CheckSumDivsOpt(scanner *bufio.Scanner) int {
	var cs int
	var d int

	for scanner.Scan() {
		d = RowEvenDivOpt(scanner.Text())
		cs = cs + d
	}

	return cs
}

func RowDiffOpt(row string) int {
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

func RowEvenDivOpt(row string) int {
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
