package main

import "testing"

func TestCountToZero(t *testing.T) {
	input := []string{
		"0 <-> 2",
		"1 <-> 7",
		"2 <-> 0, 3, 4",
		"3 <-> 2, 4",
		"4 <-> 2, 3, 6",
		"5 <-> 6",
		"6 <-> 4, 5",
		"7 <-> 1",
	}

	r := CountProgramsToZero(input)

	e := 6

	if r != e {
		t.Errorf("Got %d, expected %d", r, e)
	}
}
