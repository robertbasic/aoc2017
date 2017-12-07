package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// Day7 solves the puzzles for day 7
func Day7() {
	file, _ := os.Open("./inputs/day7.txt")
	defer file.Close()

	// 	file := strings.NewReader(`pbga (66)
	// xhth (57)
	// ebii (61)
	// havc (66)
	// ktlj (57)
	// fwft (72) -> ktlj, cntj, xhth
	// qoyq (66)
	// padx (45) -> pbga, havc, qoyq
	// tknk (41) -> ugml, padx, fwft
	// jptl (61)
	// ugml (68) -> gyxo, ebii, jptl
	// gyxo (61)
	// cntj (57)
	// `)

	bottomTower := FindBottomTower(file)
	fmt.Println("Bottom tower is: ", bottomTower)
}

// FindBottomTower finds the bottom tower
func FindBottomTower(input io.Reader) string {
	holdingUp := make(map[string]bool)
	heldUp := make(map[string]bool)

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		t := scanner.Text()

		if t == "" {
			continue
		}

		if !strings.Contains(t, "->") {
			continue
		}

		re := regexp.MustCompile("^[a-z]+[^ ]*")

		holder := strings.TrimSpace(re.FindString(t))
		holdingUp[holder] = true

		re = regexp.MustCompile("( [a-z]+)")

		for _, heldee := range re.FindAllString(t, -1) {
			heldUp[strings.TrimSpace(heldee)] = true
		}
	}

	for k := range holdingUp {
		if _, ok := heldUp[k]; !ok {
			return k
		}
	}

	return ""
}
