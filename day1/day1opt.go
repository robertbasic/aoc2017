package day1

import (
	"io/ioutil"
	"log"
)

// Day1Opt solves the puzzles for day 1
func Day1Opt(logger *log.Logger, folder string) {
	buf, _ := ioutil.ReadFile(folder + "/day1.txt")

	sum := SumOpt(buf)
	half := HalfOpt(buf)

	logger.Println("The sum is: ", sum)
	logger.Println("The half is: ", half)
}

func SumOpt(input []byte) int64 {
	var sum int64
	var prev = input[len(input)-1 : len(input)][0]
	for _, b := range input {
		if b == prev {
			sum = sum + int64(b) - 48
		}
		prev = b
	}
	return sum
}

func HalfOpt(input []byte) int64 {
	var halfsum int64
	var step = len(input) / 2
	for pos := 0; pos < len(input)-step; pos++ {
		c := input[pos]
		p := input[pos+step]

		if c == p {
			halfsum = halfsum + (int64(c) - 48 + int64(p) - 48)
		}
	}

	return halfsum
}
