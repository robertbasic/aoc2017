package main

import (
	"fmt"
	"math"
)

// Day10 solves the puzzles for day 10
func Day10() {
	var list = make([]int, 256)
	var lengths = []int{34, 88, 2, 222, 254, 93, 150, 0, 199, 255, 39, 32, 137, 136, 1, 167}
	for i := 0; i < 256; i++ {
		list[i] = i
	}

	cs := ElvenHashChecksum(list, lengths, 1)
	fmt.Println("Checksum is: ", cs)
}

// ElvenHashChecksum calculates the checksum
// for the elven hash
func ElvenHashChecksum(list []int, lengths []int, rounds int) int {
	var position int
	var skip int
	var listlength = len(list)

	for rounds > 0 {
		for _, length := range lengths {
			list = ElvenHash(list, listlength, length, position)
			position = NextPosition(position, listlength, skip+length)
			skip++
		}
		rounds--
	}

	return list[0] * list[1]
}

// ElvenHash hashes the list as per elven rules
func ElvenHash(list []int, listlength int, length int, position int) []int {
	if length == 1 {
		return list
	}

	if position+length <= listlength {
		sublist := list[position : position+length]
		sublist = ReverseList(sublist)

		var j int
		for i := position; i < length; i++ {
			list[i] = sublist[j]
			j++
		}
	} else {
		uppersublist := list[position:listlength]
		uppersublistlength := len(uppersublist)
		lowersublist := list[0 : length-uppersublistlength]

		sublist := make([]int, 0)
		sublist = append(sublist, uppersublist...)
		sublist = append(sublist, lowersublist...)

		sublist = ReverseList(sublist)

		var k int
		for i, j := position, 0; i < listlength; i, j = i+1, j+1 {
			list[i] = sublist[j]
			k = j
		}

		for i := 0; i < len(lowersublist); i++ {
			k++
			list[i] = sublist[k]
		}
	}

	return list
}

// NextPosition calculates the next position
func NextPosition(position int, listlength int, steps int) int {
	m := int(math.Mod(float64(steps), float64(listlength)))

	if position+m >= listlength {
		return position - listlength + m
	}

	return position + m
}

// ReverseList reverses a list
func ReverseList(list []int) []int {
	var reversed = make([]int, len(list))
	var j = len(list) - 1

	for i := 0; i < len(list); i++ {
		reversed[i] = list[j]
		j--
	}

	return reversed
}
