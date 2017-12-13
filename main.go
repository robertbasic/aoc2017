package main

import (
	"log"
	"os"

	"github.com/robertbasic/aoc2017/day1"
	"github.com/robertbasic/aoc2017/day10"
	"github.com/robertbasic/aoc2017/day11"
	"github.com/robertbasic/aoc2017/day12"
	"github.com/robertbasic/aoc2017/day13"
	"github.com/robertbasic/aoc2017/day2"
	"github.com/robertbasic/aoc2017/day3"
	"github.com/robertbasic/aoc2017/day4"
	"github.com/robertbasic/aoc2017/day5"
	"github.com/robertbasic/aoc2017/day6"
	"github.com/robertbasic/aoc2017/day7"
	"github.com/robertbasic/aoc2017/day8"
	"github.com/robertbasic/aoc2017/day9"
)

func main() {
	logger := log.New(os.Stdout, "", 0)

	f := "./inputs"

	day1.Day1(logger, f)

	day2.Day2(logger)

	day3.Day3(logger)

	day4.Day4(logger)

	day5.Day5(logger)

	day6.Day6(logger)

	day7.Day7(logger)

	day8.Day8(logger)

	day9.Day9(logger)

	day10.Day10(logger)

	day11.Day11(logger)

	day12.Day12(logger)

	day13.Day13(logger)
}
