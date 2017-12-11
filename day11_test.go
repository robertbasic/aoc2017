package main

import "testing"

var movetests = []struct {
	in  string
	out int
}{
	{"ne,ne,ne", 3},
	{"ne,ne,sw,sw", 0},
	{"ne,ne,s,s", 2},
	{"se,sw,se,sw,sw", 3},
}

func TestHexMove(t *testing.T) {
	for _, tt := range movetests {
		p := Position{}
		r := HexMove(p, tt.in)

		if r != tt.out {
			t.Errorf("Got %d for %s, expected %d", r, tt.in, tt.out)
		}
	}
}
