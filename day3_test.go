package main

import "testing"

var squaretests = []struct {
	in  int
	out []int
}{
	{8, []int{9, 2, 3}},
	{54, []int{81, 5, 9}},
	{1024, []int{1089, 17, 33}},
}

func TestSquare(t *testing.T) {
	for _, tt := range squaretests {
		corner, square, side := Square(tt.in)

		if corner != tt.out[0] {
			t.Errorf("Got %d for %d at corner, expected %d", corner, tt.in, tt.out[0])
		}
		if square != tt.out[1] {
			t.Errorf("Got %d for %d at square, expected %d", square, tt.in, tt.out[1])
		}
		if side != tt.out[2] {
			t.Errorf("Got %d for %d at side, expected %d", side, tt.in, tt.out[2])
		}

	}
}
