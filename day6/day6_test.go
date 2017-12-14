package day6

import "testing"

var hightests = []struct {
	in [4]int
	op int
	ip int
}{
	{[4]int{0, 2, 7, 0}, 2, 7},
	{[4]int{4, 3, 2, 1}, 0, 4},
	{[4]int{0, 0, 0, 0}, 0, 0},
	{[4]int{0, 1, 1, 0}, 1, 1},
}

func TestCycle(t *testing.T) {
	i := [4]int{0, 2, 7, 0}

	r := Cycle(i[:])

	e := 5

	if r != e {
		t.Errorf("Got %d, expected %d", r, e)
	}
}

func TestHighest(t *testing.T) {
	for _, tt := range hightests {
		ph, pi := Highest(tt.in[:])

		if ph != tt.op && pi != tt.ip {
			t.Errorf("Got %d and %d for %#v, expected %d and %d", ph, pi, tt.in, tt.op, tt.ip)
		}
	}

	i := [16]int{4, 10, 4, 1, 8, 4, 9, 14, 5, 1, 14, 15, 0, 15, 3, 5}
	ph, pi := Highest(i[:])

	if ph != 11 && pi != 15 {
		t.Errorf("Got %d and %d for %#v, expected %d and %d", ph, pi, i, 11, 15)
	}
}
