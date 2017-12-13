package day4

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

var hardenedpassphrasetests = []struct {
	in  string
	out bool
}{
	{"a b c", true},
	{"a b a", false},
	{"aa bb ac", true},
	{"aa aa bb aa bb", false},
	{"ab ba", false},
	{"abcde fghij", true},
	{"abcde xyz ecdab", false},
	{"a ab abc abd abf abj", true},
	{"iiii oiii ooii oooi oooo", true},
	{"oiii ioii iioi iiio", false},
}

func TestValidPassphrase(t *testing.T) {
	for _, tt := range passphrasetests {
		v := ValidPassphrase(tt.in)

		if v != tt.out {
			t.Errorf("Got %t for %s, expected %t", v, tt.in, tt.out)
		}
	}
}

func TestValidHardenedPassphrase(t *testing.T) {
	for _, tt := range hardenedpassphrasetests {
		v := ValidHardenedPassphrase(tt.in)

		if v != tt.out {
			t.Errorf("Got %t for %s, expected %t", v, tt.in, tt.out)
		}
	}
}
