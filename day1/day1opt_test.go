package day1

import (
	"io/ioutil"
	"log"
	"testing"
)

var sumtestsopt = []struct {
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

var halftestsopt = []struct {
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

func TestSumOpt(t *testing.T) {
	for _, tt := range sumtestsopt {
		s := SumOpt([]byte(tt.in))

		if s != int64(tt.out) {
			t.Errorf("Got %d for %s, expected %d", s, tt.in, tt.out)
		}
	}
}

func TestHalfOpt(t *testing.T) {
	for _, tt := range halftestsopt {
		h := HalfOpt([]byte(tt.in))

		if h != int64(tt.out) {
			t.Errorf("Got %d for %s, expected %d", h, tt.in, tt.out)
		}
	}
}

func BenchmarkDay1Opt(b *testing.B) {
	l := log.New(ioutil.Discard, "", 0)

	for n := 0; n < b.N; n++ {
		Day1Opt(l)
	}
}
