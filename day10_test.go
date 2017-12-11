package main

import "testing"

func TestElvenHashChecksum(t *testing.T) {
	var list = []int{0, 1, 2, 3, 4}
	var lengths = []int{3, 4, 1, 5}

	r := ElvenHashChecksum(list, lengths, 1)

	e := 12

	if r != e {
		t.Errorf("Got %d, expected %d", r, e)
	}
}
