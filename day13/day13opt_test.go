package day13

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestWalkOnFireOpt(t *testing.T) {
	input := `0: 3
1: 2
4: 4
6: 4`

	layers, md := FirewallLayersOpt(input)

	r := WalkOnFireOpt(layers, md)

	e := 24

	if r != e {
		t.Errorf("Got %d, expected %d", r, e)
	}
}

func TestEvadeDelayOpt(t *testing.T) {
	input := `0: 3
1: 2
4: 4
6: 4`

	layers, md := FirewallLayersOpt(input)

	r := EvadeDelayOpt(layers, md)

	e := 10

	if r != e {
		t.Errorf("Got %d, expected %d", r, e)
	}
}

func BenchmarkDay13Opt(b *testing.B) {
	l := log.New(ioutil.Discard, "", 0)
	f := "../inputs"
	for n := 0; n < b.N; n++ {
		Day13Opt(l, f)
	}
}
