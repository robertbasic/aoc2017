package day17

import (
	"math"
	"log"
)

// Day17 solves the puzzles for day 17
func Day17() {
	s, p := step(335, 2017)
	log.Println("Number after 2017: ", s[p+1])

	i := lotsofsteps(335, 50000000)
	log.Println(i)
}

func lotsofsteps(step int, max int) int {
	r := 0
	p := 0
	for i := 0; i < max + 1; i++ {
		p = int(math.Mod(float64(p+step), float64(i))) + 1

		if p == 1 {
			r = i
		}
	}
	return r
}

func step(step int, max int) ([]int, int) {
	s := []int{0}
	p := 0

	for i := 0; i < max; i++ {
		p = after(p, step, len(s))
		s = insertAt(s, p, i+1)
	}

	return s, p
}

func after(cp int, step int, l int) int {
	p := cp + step

	if p >= l {
		p = int(math.Mod(float64(p), float64(l)))
	}

	return p + 1
}

func insertAt(s []int, p int, v int) []int {
	s = append(s, 0)
	copy(s[p+1:], s[p:])
	s[p] = v
	return s
}
