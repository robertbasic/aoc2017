package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// Scanner is a scanner in a layer
type Scanner struct {
	rang      int
	position  int
	direction int
}

func (s *Scanner) move(cp int) {
	r := s.rang - 1
	d := cp - r
	div := int(math.Trunc(float64(d / r)))
	m := int(math.Mod(float64(d), float64(r)))

	if math.Mod(float64(div), 2) == 0 {
		s.position = r - m
	} else {
		s.position = m
	}
}

func (s Scanner) caught() bool {
	if s.position == 0 {
		return true
	}

	return false
}

// Layer is a layer in the firewall
type Layer struct {
	depth   int
	scanner Scanner
}

func (l Layer) sev() int {
	return l.depth * l.scanner.rang
}

// Layers holds all the layers
type Layers map[int]Layer

// Day13 solves the puzzles for day 13
func Day13() {
	file, _ := os.Open("./inputs/day13.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	sev := WalkOnFire(input)
	fmt.Println("Severity is: ", sev)

	delay := EvadeDelay(input)
	fmt.Println("Delay for: ", delay)
}

// EvadeDelay calculates the delay required to evade all scanners
func EvadeDelay(input []string) int {
	var delay = 0
	layers, md := FirewallLayers(input)

	for {
		var cd = -1
		var caught = false

		for cd <= md {
			cd++

			if _, ok := layers[cd]; !ok {
				continue
			}

			cl := layers[cd]
			cl.scanner.move(cd + delay)

			if cl.scanner.caught() {
				caught = true
			}
		}

		if !caught {
			break
		}

		delay += 2
	}

	return delay
}

// WalkOnFire walks us through the firewall
func WalkOnFire(input []string) int {
	var cd = -1
	var sev = 0
	layers, md := FirewallLayers(input)

	for cd <= md {
		cd++

		if _, ok := layers[cd]; !ok {
			continue
		}

		cl := layers[cd]
		cl.scanner.move(cd)

		if cl.scanner.caught() {
			sev += cl.sev()
		}
	}

	return sev
}

// FirewallLayers reads the firewall layers
func FirewallLayers(input []string) (Layers, int) {
	ls := make(Layers)
	var md int

	for _, i := range input {
		parts := strings.Split(i, ": ")
		d := StringNumberToInt(parts[0])
		r := StringNumberToInt(parts[1])
		s := Scanner{rang: r, direction: 1}
		l := Layer{depth: d, scanner: s}

		ls[d] = l

		md = d
	}

	return ls, md
}
