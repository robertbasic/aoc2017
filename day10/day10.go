package day10

import (
	"fmt"
	"log"
	"math"
)

// Day10 solves the puzzles for day 10
func Day10(logger *log.Logger) {
	var list = make([]int, 256)
	var lengths = []int{34, 88, 2, 222, 254, 93, 150, 0, 199, 255, 39, 32, 137, 136, 1, 167}
	for i := 0; i < 256; i++ {
		list[i] = i
	}

	cs, _ := ElvenHashChecksum(list, lengths, 1)
	logger.Println("Checksum is: ", cs)

	var stringlengths = "34,88,2,222,254,93,150,0,199,255,39,32,137,136,1,167"
	var convertedlengths = make([]int, 0)
	for _, i := range stringlengths {
		convertedlengths = append(convertedlengths, int(i))
	}

	convertedlengths = append(convertedlengths, []int{17, 31, 73, 47, 23}...)

	_, hashedlist := ElvenHashChecksum(list, convertedlengths, 64)
	dense := DenseHash(hashedlist)
	logger.Println("Hashed string is: ", dense)
}

// DenseHash denses the hashed list
func DenseHash(list []int) string {
	var h string
	for i := 0; i < 256; i += 16 {
		sublist := list[i : i+16]
		x := 0
		for j := 0; j < len(sublist); j++ {
			x ^= sublist[j]
		}
		h += fmt.Sprintf("%02x", x)
	}

	return h
}

// ElvenHashChecksum calculates the checksum
// for the elven hash
func ElvenHashChecksum(list []int, lengths []int, rounds int) (int, []int) {
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

	return list[0] * list[1], list
}

// ElvenHash hashes the list as per elven rules
func ElvenHash(list []int, listlength int, length int, position int) []int {
	var newlist = make([]int, listlength)
	copy(newlist, list)

	if length == 0 || length == 1 {
		return newlist
	}

	if position+length <= listlength {
		sublist := newlist[position : position+length]
		sublist = ReverseList(sublist)

		var j int
		for i := position; i < position+length; i++ {
			newlist[i] = sublist[j]
			j++
		}
	} else {
		uppersublist := newlist[position:listlength]
		uppersublistlength := len(uppersublist)
		lowersublist := list[0 : length-uppersublistlength]

		sublist := make([]int, 0)
		sublist = append(sublist, uppersublist...)
		sublist = append(sublist, lowersublist...)

		sublist = ReverseList(sublist)

		var k int
		for i, j := position, 0; i < listlength; i, j = i+1, j+1 {
			newlist[i] = sublist[j]
			k = j
		}

		for i := 0; i < len(lowersublist); i++ {
			k++
			newlist[i] = sublist[k]
		}
	}

	return newlist
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
