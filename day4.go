package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Day4 solves the puzzles for day 4
func Day4() {
	file, _ := os.Open("./inputs/day4.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	n := NumberOfValidPassphrases(scanner)
	fmt.Println("Number of valid passphrases: ", n)
}

// NumberOfValidPassphrases finds the number of valid
// passphrases
func NumberOfValidPassphrases(scanner *bufio.Scanner) int {
	var n int

	for scanner.Scan() {
		if ValidPassphrase(scanner.Text()) {
			n++
		}
	}

	return n
}

// ValidPassphrase checks is the passhprase
// valid or not
func ValidPassphrase(passphrase string) bool {
	words := strings.Fields(passphrase)

	uniques := make(map[string]bool)

	for _, v := range words {
		uniques[v] = true
	}

	return len(words) == len(uniques)
}
