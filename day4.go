package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Day4 solves the puzzles for day 4
func Day4() {
	file, _ := os.Open("./inputs/day4.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	n, h := NumberOfValidPassphrases(scanner)
	fmt.Println("Number of valid passphrases: ", n)
	fmt.Println("Number of valid hardened passphrases: ", h)
}

// NumberOfValidPassphrases finds the number of valid
// passphrases, for both regular and hardened security
func NumberOfValidPassphrases(scanner *bufio.Scanner) (int, int) {
	var n int
	var h int

	for scanner.Scan() {
		t := scanner.Text()
		if ValidPassphrase(t) {
			n++
		}

		if ValidHardenedPassphrase(t) {
			h++
		}

	}

	return n, h
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

// ValidHardenedPassphrase checks is the passphrase
// valid or not, based on the hardened rules
func ValidHardenedPassphrase(passphrase string) bool {
	words := strings.Fields(passphrase)

	uniques := make(map[string]bool)

	for _, v := range words {
		v = resort(v)
		uniques[v] = true
	}

	return len(words) == len(uniques)
}

func resort(str string) string {
	letters := strings.Split(str, "")

	sort.Strings(letters)

	return strings.Join(letters, "")
}
