package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2() {
	file, _ := os.Open("./inputs/day2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cs := CheckSum(scanner)

	fmt.Println("The checksum is: ", cs)
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
