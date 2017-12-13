package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// Day1 solves the puzzles for day 1
func Day1(logger *log.Logger) {
	var input string

	file, _ := os.Open("./inputs/day1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = scanner.Text()
	}

	sum := Sum(input)
	half := Half(input)

	logger.Println("The sum is: ", sum)
	logger.Println("The half is: ", half)
}

func Sum(input string) int64 {
	var sum int64
	var prev = input[len(input)-1:]
	for _, c := range input {
		char := string(c)
		if char == prev {
			i, _ := strconv.ParseInt(char, 10, 32)
			sum = sum + i
		}
		prev = char
	}
	return sum
}

func Half(input string) int64 {
	var halfsum int64
	var step = len(input) / 2
	var pos = 0
	for pos = 0; pos < len(input)-step; pos++ {
		char := string(input[pos])
		pair := string(input[pos+step])
		i, _ := strconv.ParseInt(char, 10, 32)
		ip, _ := strconv.ParseInt(pair, 10, 32)

		if i == ip {
			halfsum = halfsum + (i + ip)
		}
	}

	return halfsum
}
