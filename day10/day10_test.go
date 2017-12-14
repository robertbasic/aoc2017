package day10

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestElvenHashChecksum(t *testing.T) {
	var list = []int{0, 1, 2, 3, 4}
	var lengths = []int{3, 4, 1, 5}

	r, _ := ElvenHashChecksum(list, lengths, 1)

	e := 12

	if r != e {
		t.Errorf("Got %d, expected %d", r, e)
	}
}

func BenchmarkDay10(b *testing.B) {
	l := log.New(ioutil.Discard, "", 0)
	for n := 0; n < b.N; n++ {
		Day10(l)
	}
}
