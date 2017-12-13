package day3

import "log"

//Day3 solves the puzzles for day 3
func Day3(logger *log.Logger) {
	s := Steps(368078)
	//s := Steps(1024)
	logger.Println("Steps required: ", s)
}

// Steps finds the number of steps required to
// take from <number> to 1
func Steps(number int) int {
	var steps int

	corner0, square, sidesize := Square(number)
	corner1 := corner0 - sidesize + 1
	corner2 := corner1 - sidesize + 1
	corner3 := corner2 - sidesize + 1

	mid0 := corner0 - square + 1
	mid1 := corner1 - square + 1
	mid2 := corner2 - square + 1
	mid3 := corner3 - square + 1

	if number == mid0 || number == mid1 || number == mid2 || number == mid3 {
		return square
	}

	if number <= corner0 && number > mid0 {
		steps = number - mid0
	} else if number <= corner1 && number > mid1 {
		steps = number - mid1
	} else if number <= corner2 && number > mid2 {
		steps = number - mid2
	} else if number <= corner3 && number > mid3 {
		steps = number - mid3
	} else if number >= corner1 && number < mid0 {
		steps = mid0 - number
	} else if number >= corner2 && number < mid1 {
		steps = mid1 - number
	} else if number >= corner3 && number < mid2 {
		steps = mid2 - number
	} else if number >= corner0 && number < mid3 {
		steps = mid3 - number
	}

	return steps + square - 1
}

// Square finds for a given number
// - the number at the bottom right corner,
// - the number of the square in which that number is
// - the number of digits in 1 side of the square
func Square(num int) (int, int, int) {
	step := 8
	sideStep := 2
	square := 1
	corner := 1
	sidesize := 1

	for corner < num {
		sidesize += sideStep

		if square == 1 {
			corner += step
			square++
			continue
		}

		corner = corner + square*step
		square++
	}

	return corner, square, sidesize
}
