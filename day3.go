package main

//Day3 solves the puzzles for day 3
func Day3() {

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
	side := 1

	for corner < num {
		side += sideStep

		if square == 1 {
			corner += step
			square++
			continue
		}

		corner = corner + square*step
		square++
	}

	return corner, square, side
}
