package day2

import (
	"bufio"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

var difftestsopt = []struct {
	in  string
	out int
}{
	{"5 1 9 5", 8},
	{"7 5 3 ", 4},
	{"2 4 6 8", 6},
}

var divtestsopt = []struct {
	in  string
	out int
}{
	{"5 1 9 5", 21},
	{"7 5 3 ", 0},
	{"2 4 6 8", 11},
	{"5 9 2 8", 4},
	{"9 4 7 3", 3},
	{"3 8 6 5", 2},
}

func TestCheckSumOpt(t *testing.T) {
	sample := `5 1 9 5
	7 5 3
	2 4 6 8`

	scanner := bufio.NewScanner(strings.NewReader(sample))

	checksum := CheckSumOpt(scanner)
	expected := 18
	if checksum != expected {
		t.Errorf("Got %d, expected %d", checksum, expected)
	}
}

func TestCheckSumDivsOpt(t *testing.T) {
	sample := `5 9 2 8
	9 4 7 3
	3 8 6 5`

	scanner := bufio.NewScanner(strings.NewReader(sample))

	checksum := CheckSumDivs(scanner)
	expected := 9
	if checksum != expected {
		t.Errorf("Got %d, expected %d", checksum, expected)
	}
}

func TestRowDiffOpt(t *testing.T) {
	for _, tt := range difftestsopt {
		s := RowDiffOpt(tt.in)

		if s != int(tt.out) {
			t.Errorf("Got %d for %s, expected %d", s, tt.in, tt.out)
		}
	}
}

func TestRowEvenDivOpt(t *testing.T) {
	for _, tt := range divtestsopt {
		s := RowEvenDivOpt(tt.in)

		if s != int(tt.out) {
			t.Errorf("Got %d for %s, expected %d", s, tt.in, tt.out)
		}
	}
}

func BenchmarkDay2Opt(b *testing.B) {
	l := log.New(ioutil.Discard, "", 0)
	f := "../inputs"
	for n := 0; n < b.N; n++ {
		Day2(l, f)
	}
}
