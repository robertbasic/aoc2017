package main

import (
	"bufio"
	"strings"
	"testing"
)

var difftests = []struct {
	in  string
	out int
}{
	{"5 1 9 5", 8},
	{"7 5 3 ", 4},
	{"2 4 6 8", 6},
}

func TestCheckSum(t *testing.T) {
	sample := `5 1 9 5
	7 5 3
	2 4 6 8`

	scanner := bufio.NewScanner(strings.NewReader(sample))

	checksum := CheckSum(scanner)
	expected := 18
	if checksum != expected {
		t.Errorf("Got %d, expected %d", checksum, expected)
	}
}

func TestRowDiff(t *testing.T) {
	for _, tt := range difftests {
		s := RowDiff(tt.in)

		if s != int(tt.out) {
			t.Errorf("Got %d for %s, expected %d", s, tt.in, tt.out)
		}
	}
}
