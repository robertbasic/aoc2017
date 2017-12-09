package main

import "testing"

var cancelledtests = []struct {
	in  string
	out string
}{
	{"{{<!>},{<!>},{<!>},{<a>}}", "{{<},{<},{<},{<a>}}"},
	{"{{<!!>},{<!!>},{<!!>},{<!!>}}", "{{<>},{<>},{<>},{<>}}"},
	{"{{<a!>},{<a!>},{<a!>},{<ab>}}", "{{<a},{<a},{<a},{<ab>}}"},
	{"{{<a!!!>},{<a!!!>},{<a!!!>},{<ab>}}", "{{<a},{<a},{<a},{<ab>}}"},
}

var garbagetests = []struct {
	in  string
	out string
}{
	{"{{<},{<},{<},{<a>test}}", "{{}}"},
	{"{{<>},{<>},{<>},{<>}}", "{{},{},{},{}}"},
	{"{test{<a},{<a},{<a},{<ab>}}", "{{}}"},
	{"<>", ""},
	{"<<<>", ""},
}

var scoretests = []struct {
	in  string
	out int
}{
	{"{}", 1},
	{"{{}}", 3},
	{"{{},{}}", 5},
	{"{{{},{},{{}}}}", 16},
	{"{{{}}}", 6},
}

var countgarbagetests = []struct {
	in  string
	out int
}{
	{"{{<},{<},{<},{<a>test}}", 13},
	{"{{<>},{<>},{<>},{<>}}", 0},
	{"{test{<a},{<a},{<a},{<ab>}}", 17},
	{"<>", 0},
	{"<<<>", 2},
	{"{{<1>},{<2>},{<3>},{<4>}}", 4},
}

func TestCountRemovedGarbage(t *testing.T) {
	for _, tt := range countgarbagetests {
		r := CountRemovedGarbage(tt.in)

		if r != tt.out {
			t.Errorf("Got %d for %s, expected %d", r, tt.in, tt.out)
		}
	}
}

func TestTotalScore(t *testing.T) {
	for _, tt := range scoretests {
		r := TotalScore(tt.in)

		if r != tt.out {
			t.Errorf("Got %d for %s, expected %d", r, tt.in, tt.out)
		}
	}
}

func TestRemoveGarbage(t *testing.T) {
	for _, tt := range garbagetests {
		r := RemoveGarbage(tt.in)

		if r != tt.out {
			t.Errorf("Got %s for %s, expected %s", r, tt.in, tt.out)
		}
	}
}

func TestCancelled(t *testing.T) {
	for _, tt := range cancelledtests {
		r := Cancelled(tt.in)

		if r != tt.out {
			t.Errorf("Got %s for %s, expected %s", r, tt.in, tt.out)
		}
	}
}
