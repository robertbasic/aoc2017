package day16

import "testing"

func TestDance(t *testing.T) {
	p := "abcde"
	l := "s1,x3/4,pe/b"
	i := processInput(l)

	r := dance(p, i)
	e := "baedc"

	if r != e {
		t.Errorf("Got %s, expected %s", r, e)
	}
}
