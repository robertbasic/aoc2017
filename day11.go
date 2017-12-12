package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

// Position holds the coordinates for the current position
type Position struct {
	x []int
	y []int
}

func (p Position) calc() int {
	var x, y int

	for _, i := range p.x {
		x += i
	}

	for _, i := range p.y {
		y += i
	}

	x, y = int(math.Abs(float64(x))), int(math.Abs(float64(y)))

	c := []int{x, y}
	sort.Ints(c)

	return c[len(c)-1]
}

func (p *Position) move(d string) {
	var x int
	var y int

	switch d {
	case "n":
		x, y = 0, -1
	case "s":
		x, y = 0, 1
	case "ne":
		x, y = 1, -1
	case "se":
		x, y = 1, 0
	case "sw":
		x, y = -1, 1
	case "nw":
		x, y = -1, 0
	}

	p.x = append(p.x, x)
	p.y = append(p.y, y)
}

// Day11 solves the puzzles for day 11
func Day11() {
	file, _ := os.Open("./inputs/day11.txt")
	defer file.Close()

	var line string

	scanner := bufio.NewScanner(file)

	p := Position{}

	for scanner.Scan() {
		line = scanner.Text()
	}

	m := HexMove(&p, line)

	fmt.Println("Steps away: ", p.calc())
	fmt.Println("Got to furthest: ", m)
}

// HexMove moves the position around in a hex grid
func HexMove(p *Position, direction string) int {
	var m int
	for _, d := range strings.Split(direction, ",") {
		p.move(d)

		c := p.calc()

		if c > m {
			m = c
		}
	}
	return m
}
