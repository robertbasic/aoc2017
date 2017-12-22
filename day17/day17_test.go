package day17

import "testing"

func TestStep(t *testing.T) {
	r, p := step(3, 2017)

	e := 638

	if r[p+1] != e {
		t.Errorf("Got %d, expected %d", r, e)
	}
}
