package day13

import "testing"

func TestWalkOnFire(t *testing.T) {
	input := []string{
		"0: 3",
		"1: 2",
		"4: 4",
		"6: 4",
	}

	r := WalkOnFire(input)

	e := 24

	if r != e {
		t.Errorf("Got %d, expected %d", r, e)
	}
}

func TestEvadeDelay(t *testing.T) {
	input := []string{
		"0: 3",
		"1: 2",
		"4: 4",
		"6: 4",
	}

	r := EvadeDelay(input)

	e := 10

	if r != e {
		t.Errorf("Got %d, expected %d", r, e)
	}
}
