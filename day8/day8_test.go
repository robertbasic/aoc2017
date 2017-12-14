package day8

import (
	"strings"
	"testing"
)

func TestFindLargestAtEnd(t *testing.T) {
	input := `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`

	r, r1 := FindLargestAtEnd(strings.NewReader(input))

	e := 1
	e1 := 10

	if r != e {
		t.Errorf("Got %d, expeced %d", r, e)
	}

	if r1 != e1 {
		t.Errorf("Got %d, expeced %d", r1, e1)
	}
}
