package day10

import (
	"testing"
)

func TestElvenKnottedHash(t *testing.T) {
	input := "34,88,2,222,254,93,150,0,199,255,39,32,137,136,1,167"
	r := ElvenKnottedHash(input)

	e := "a7af2706aa9a09cf5d848c1e6605dd2a"

	if r != e {
		t.Errorf("Got %s, expected %s", r, e)
	}
}

func TestElvenHashChecksum(t *testing.T) {
	var listSize = 5
	var lengths = []int{3, 4, 1, 5}

	r := elvenHashChecksum(listSize, lengths, 1)

	e := 12

	if r != e {
		t.Errorf("Got %d, expected %d", r, e)
	}
}

func BenchmarkElvenKnottedHash(b *testing.B) {
	i := "34,88,2,222,254,93,150,0,199,255,39,32,137,136,1,167"
	for n := 0; n < b.N; n++ {
		ElvenKnottedHash(i)
	}
}
