package day18

import "testing"

func TestRecoverSound(t *testing.T) {
	instructions := []string{
		"set a 1",
		"add a 2",
		"mul a a",
		"mod a 5",
		"snd a",
		"set a 0",
		"rcv a",
		"jgz a -1",
		"set a 1",
		"jgz a -2",
	}

	r := RecoverSound(instructions)

	e := 4

	if r != e {
		t.Errorf("Got %d, expected %d", r, e)
	}
}

func TestSendReceive(t *testing.T) {
	instructions := []string{
		"snd 1",
		"snd 2",
		"snd p",
		"rcv a",
		"rcv b",
		"rcv c",
		"rcv d",
	}

	SendReceive(instructions)
}
