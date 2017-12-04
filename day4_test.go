package main

import "testing"

var passphrasetests = []struct {
	in  string
	out bool
}{
	{"a b c", true},
	{"a b a", false},
	{"aa bb ac", true},
	{"aa aa bb aa bb", false},
}

func TestValidPassphrase(t *testing.T) {
	for _, tt := range passphrasetests {
		v := ValidPassphrase(tt.in)

		if v != tt.out {
			t.Errorf("Got %t for %s, expected %t", v, tt.in, tt.out)
		}
	}
}
