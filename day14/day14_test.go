package day14

import "testing"

func TestCountGroups(t *testing.T) {
	binrows := []string{
		"01001011",
		"11011011",
		"01010001",
		"01110011",
		"00000110",
		"11100111",
		"00100000",
		"01111111",
	}
	r := CountGroups(binrows)
	e := 3

	if r != e {
		t.Errorf("Got %d, expected %d", r, e)
	}
}
