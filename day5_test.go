package main

import "testing"

func TestJumps(t *testing.T) {
	i := []int{
		0: 0,
		1: 3,
		2: 0,
		3: 1,
		4: -3,
	}

	r := Jumps(i)

	e := 5

	if r != e {
		t.Errorf("Got %d, expected %d", r, e)
	}

	i = []int{
		0: 0,
		1: 3,
		2: 0,
		3: 1,
		4: -3,
	}

	r = CrazyJumps(i)

	e = 10

	if r != e {
		t.Errorf("Got %d, expected %d", r, e)
	}
}

func TestCrazyJumps(t *testing.T) {
	i := []int{
		0: 0,
		1: 3,
		2: 0,
		3: 1,
		4: -3,
	}

	r := CrazyJumps(i)

	e := 10

	if r != e {
		t.Errorf("Got %d, expected %d", r, e)
	}
}
