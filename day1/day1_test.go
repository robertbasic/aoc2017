package day1

import (
	"io/ioutil"
	"log"
	"testing"
)

var sumtests = []struct {
	in  string
	out int
}{
	{"1122", 3},
	{"1111", 4},
	{"1212", 0},
	{"1234", 0},
	{"2222", 8},
	{"91212129", 9},
}

var halftests = []struct {
	in  string
	out int
}{
	{"1111", 4},
	{"1212", 6},
	{"1221", 0},
	{"2222", 8},
	{"123425", 4},
	{"123123", 12},
}

func TestSum(t *testing.T) {
	for _, tt := range sumtests {
		s := Sum(tt.in)

		if s != int64(tt.out) {
			t.Errorf("Got %d for %s, expected %d", s, tt.in, tt.out)
		}
	}
}

func TestHalf(t *testing.T) {
	for _, tt := range halftests {
		h := Half(tt.in)

		if h != int64(tt.out) {
			t.Errorf("Got %d for %s, expected %d", h, tt.in, tt.out)
		}
	}
}

func BenchmarkDay1(b *testing.B) {
	l := log.New(ioutil.Discard, "", 0)
	f := "../inputs"
	for n := 0; n < b.N; n++ {
		Day1(l, f)
	}
}
