package day10

import (
	"fmt"
	"log"
	"math"
)

// Day10 solves the day 10 puzzles
func Day10(logger *log.Logger) {
	var lengths = []int{34, 88, 2, 222, 254, 93, 150, 0, 199, 255, 39, 32, 137, 136, 1, 167}
	checksum := elvenHashChecksum(256, lengths, 1)
	logger.Println("Checksum is: ", checksum)

	input := "34,88,2,222,254,93,150,0,199,255,39,32,137,136,1,167"
	hash := ElvenKnottedHash(input)

	logger.Println("Hash of the elven knot is: ", hash)
}

// ElvenKnottedHash hashes the input as per the elven rules
func ElvenKnottedHash(input string) string {
	var convertedlengths = make([]int, 0)
	for _, i := range input {
		convertedlengths = append(convertedlengths, int(i))
	}

	convertedlengths = append(convertedlengths, []int{17, 31, 73, 47, 23}...)

	ls := 256

	list := elvenHashList(ls, convertedlengths, 64)

	return denseHash(list, ls)
}

// elvenHashChecksum calculates the checksum
// for the elven hash
func elvenHashChecksum(listSize int, lengths []int, rounds int) int {
	list := elvenHashList(listSize, lengths, rounds)
	return list[0] * list[1]
}

func denseHash(list []int, ls int) string {
	var h string
	for i := 0; i < ls; i += 16 {
		sublist := list[i : i+16]
		x := 0
		for j := 0; j < len(sublist); j++ {
			x ^= sublist[j]
		}
		h += fmt.Sprintf("%02x", x)
	}

	return h
}

func elvenHashList(listSize int, lengths []int, rounds int) []int {
	var position int
	var skip int
	list := makeList(listSize)

	for rounds > 0 {
		for _, length := range lengths {
			list = elvenHash(list, listSize, length, position)
			position = nextPosition(position, listSize, skip+length)
			skip++
		}
		rounds--
	}

	return list
}

func elvenHash(list []int, listlength int, length int, position int) []int {
	var newlist = make([]int, listlength)
	copy(newlist, list)

	if length == 0 || length == 1 {
		return newlist
	}

	if position+length <= listlength {
		sublist := newlist[position : position+length]
		sublist = reverseList(sublist)

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

		sublist = reverseList(sublist)

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

func makeList(ls int) []int {
	var l = make([]int, ls)
	for i := 0; i < ls; i++ {
		l[i] = i
	}
	return l
}

func nextPosition(position int, listlength int, steps int) int {
	m := int(math.Mod(float64(steps), float64(listlength)))

	if position+m >= listlength {
		return position - listlength + m
	}

	return position + m
}

func reverseList(list []int) []int {
	var reversed = make([]int, len(list))
	var j = len(list) - 1

	for i := 0; i < len(list); i++ {
		reversed[i] = list[j]
		j--
	}

	return reversed
}
