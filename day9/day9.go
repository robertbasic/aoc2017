package day9

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Day9 solves the puzzles for day 9
func Day9(logger *log.Logger) {
	file, _ := os.Open("./inputs/day9.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var in string

	for scanner.Scan() {
		in = scanner.Text()
	}

	cin := Cancelled(in)

	ts := TotalScore(RemoveGarbage(cin))
	logger.Println("Total score is: ", ts)

	rg := CountRemovedGarbage(cin)
	logger.Println("Total removed garbage is: ", rg)
}

// CountRemovedGarbage counts the number of chars removed
// while removing the garbage
func CountRemovedGarbage(in string) int {
	g := false
	var c int

	for i := 0; i < len(in); i++ {
		if in[i] == '<' && g == false {
			c--
			g = true
		}

		if g == true {
			c++
		}

		if in[i] == '>' && g == true {
			c--
			g = false
		}
	}

	return c
}

// TotalScore counts the total score
// for all the groups
func TotalScore(in string) int {
	var t int
	var l int

	for _, r := range in {
		if r == '{' {
			l++
			t += l
		}
		if r == '}' {
			l--
		}
	}

	return t
}

// RemoveGarbage removes all that is between
// < and >, the < > characters as well
// and also removes anything that is not a {, }, or a , (comma)
func RemoveGarbage(in string) string {
	if !strings.ContainsRune(in, '<') && !strings.ContainsRune(in, '>') {
		return in
	}

	out := make([]string, len(in))
	g := false

	for i := 0; i < len(in); i++ {
		if in[i] == '<' && g == false {
			g = true
		}

		if g == false && (in[i] == '{' || in[i] == '}' || in[i] == ',') {
			out[i] = string(in[i])
		}

		if in[i] == '>' && g == true {
			g = false
		}
	}

	return strings.Join(out, "")
}

// Cancelled cancels everything there is to be
// cancelled in the in string
func Cancelled(in string) string {
	if !strings.ContainsRune(in, '!') {
		return in
	}

	out := make([]string, len(in))

	for i := 0; i < len(in); i++ {
		j := getj(i)

		if out[j] != "!" {
			out[i] = string(in[i])
		}

		if out[j] == "!" {
			out[j] = ""
		}
	}

	return strings.Join(out, "")
}

func getj(i int) int {
	if i != 0 {
		return i - 1
	}
	return 0
}
