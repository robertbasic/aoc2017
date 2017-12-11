package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// Position holds the coordinates for the current position
type Position struct {
	x float64
	y float64
}

func (p Position) calc() int {
	fmt.Println(p.x, p.y)
	return int(math.Abs(p.x) + math.Abs(p.y))
}

func (p *Position) move(d string) {
	var x float64
	var y float64
	switch d {
	case "e":
		x, y = 1, 0
	case "w":
		x, y = -1, 0
	case "n":
		x, y = 0, 1
	case "s":
		x, y = 0, -1
	case "ne":
		x, y = 0.5, 0.5
	case "se":
		x, y = 0.5, -0.5
	case "sw":
		x, y = -0.5, -0.5
	case "nw":
		x, y = -0.5, 0.5
	}

	p.x = p.x + x
	p.y = p.y + y
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

	HexMove(&p, line)

	fmt.Println("Steps away: ", p.calc())
}

// HexMove moves the position around in a hex grid
func HexMove(p *Position, direction string) {
	for _, d := range strings.Split(direction, ",") {
		p.move(d)
	}
}
