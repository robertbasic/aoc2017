package day11

import (
	"testing"
)

var movetests = []struct {
	in   string
	outp int
	outm int
}{
	{"ne,ne,ne", 3, 3},
	{"ne,ne,sw,sw", 0, 2},
	{"ne,ne,s,s", 2, 2},
	{"se,sw,se,sw,sw", 3, 3},
}

func TestHexMove(t *testing.T) {
	for _, tt := range movetests {
		p := Position{}
		m := HexMove(&p, tt.in)

		r := p.calc()

		if r != tt.outp {
			t.Errorf("Got %d for %s, expected %d", r, tt.in, tt.outp)
		}

		if m != tt.outm {
			t.Errorf("Got %d for %s, expected %d", m, tt.in, tt.outm)
		}
	}
}
