package day18

import (
	"log"
	"math"
	"strconv"
	"strings"
)

var instructions = []string{
	"set i 31",
	"set a 1",
	"mul p 17",
	"jgz p p",
	"mul a 2",
	"add i -1",
	"jgz i -2",
	"add a -1",
	"set i 127",
	"set p 622",
	"mul p 8505",
	"mod p a",
	"mul p 129749",
	"add p 12345",
	"mod p a",
	"set b p",
	"mod b 10000",
	"snd b",
	"add i -1",
	"jgz i -9",
	"jgz a 3",
	"rcv b",
	"jgz b -1",
	"set f 0",
	"set i 126",
	"rcv a",
	"rcv b",
	"set p a",
	"mul p -1",
	"add p b",
	"jgz p 4",
	"snd a",
	"set a b",
	"jgz 1 3",
	"snd b",
	"set f 1",
	"add i -1",
	"jgz i -11",
	"snd a",
	"jgz f -16",
	"jgz a -19",
}

// Day18 solves the puzzles for day 18
func Day18(logger *log.Logger) {
	s := RecoverSound(instructions)

	logger.Println("Sound recovered: ", s)
}

// RecoverSound recovers the last played sound
func RecoverSound(instructions []string) int {
	var regs = make(map[string]int)
	var snd int

loop:
	for i := 0; i < len(instructions); i++ {
		ip := strings.Fields(instructions[i])
		op := ip[0]
		reg := ip[1]

		if op != "jgz" && len(ip) == 3 {
			doValOp(op, reg, ip[2], regs)
		}

		switch op {
		case "jgz":
			t, j := jump(reg, ip[2], regs)
			if t > 0 {
				i += j - 1
			}
		case "snd":
			snd = regs[reg]
		case "rcv":
			t := regs[reg]
			if t != 0 {
				break loop
			}
		}
	}

	return snd
}

func doValOp(op string, reg string, val string, regs map[string]int) {
	v := 0
	if isReg(val) {
		v = regs[val]
	} else {
		v, _ = strconv.Atoi(val)
	}

	switch op {
	case "set":
		regs[reg] = v
	case "add":
		regs[reg] = regs[reg] + v
	case "mul":
		regs[reg] = regs[reg] * v
	case "mod":
		regs[reg] = int(math.Mod(float64(regs[reg]), float64(v)))
	}
}

func jump(reg string, val string, regs map[string]int) (int, int) {
	j := 0
	t := 0
	if isReg(reg) {
		t = regs[reg]
	} else {
		t, _ = strconv.Atoi(reg)
	}
	if isReg(val) {
		j = regs[val]
	} else {
		j, _ = strconv.Atoi(val)
	}
	return t, j
}

func isReg(r string) bool {
	return r >= "a" && r <= "z"
}
