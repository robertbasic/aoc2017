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

type Program struct {
	id           int
	instructions []string
	regs         map[string]int64
	li           int
	q            []int64
	vs           int
}

func (p *Program) run(po *Program) {
loop:
	//log.Println(p.id, ": ", p.vs)
	for i := p.li; i < len(p.instructions); i++ {
		ip := strings.Fields(p.instructions[i])
		op := ip[0]
		reg := ip[1]

		if op != "jgz" && len(ip) == 3 {
			doValOp(op, reg, ip[2], p.regs)
			continue
		}

		switch op {
		case "jgz":
			t, j := jump(reg, ip[2], p.regs)
			if t > 0 {
				i += int(j) - 1
			}
		case "snd":
			p.vs++
			val := regVal(reg, p.regs)
			po.q = append(po.q, val)
		case "rcv":
			if len(p.q) != 0 {
				p.regs[reg] = p.q[0]
				p.q = p.q[1:]
			} else if p.li != 0 && len(p.q) == 0 && len(po.q) == 0 {
				break loop
			} else {
				p.li = i
				po.run(p)
			}
		}
	}
}

func newProgram(id int, instructions []string) *Program {
	regs := map[string]int64{
		"p": int64(id),
	}
	p := Program{
		id:           id,
		instructions: instructions,
		regs:         regs,
	}

	return &p
}

// Day18 solves the puzzles for day 18
func Day18() {
	//s := RecoverSound(instructions)

	//log.Println("Sound recovered: ", s)

	c := SendReceive(instructions)
	log.Println("Times sent a value: ", c)
}

func SendReceive(instructions []string) int {
	p0 := newProgram(0, instructions)
	p1 := newProgram(1, instructions)

	p0.run(p1)

	return p1.vs
}

// RecoverSound recovers the last played sound
func RecoverSound(instructions []string) int64 {
	var regs = make(map[string]int64)
	var snd int64

loop:
	for i := 0; i < len(instructions); i++ {
		ip := strings.Fields(instructions[i])
		op := ip[0]
		reg := ip[1]

		if op != "jgz" && len(ip) == 3 {
			doValOp(op, reg, ip[2], regs)
			continue
		}

		switch op {
		case "jgz":
			t, j := jump(reg, ip[2], regs)
			if t > 0 {
				i += int(j) - 1
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

func doValOp(op string, reg string, val string, regs map[string]int64) {
	v := regVal(val, regs)

	switch op {
	case "set":
		regs[reg] = v
	case "add":
		regs[reg] = regs[reg] + v
	case "mul":
		regs[reg] = regs[reg] * v
	case "mod":
		regs[reg] = int64(math.Mod(float64(regs[reg]), float64(v)))
	}
}

func jump(reg string, val string, regs map[string]int64) (int64, int64) {
	t := regVal(reg, regs)
	j := regVal(val, regs)
	return t, j
}

func isReg(r string) bool {
	return r >= "a" && r <= "z"
}

func regVal(r string, regs map[string]int64) int64 {
	var v int64
	if isReg(r) {
		v = regs[r]
	} else {
		vv, _ := strconv.Atoi(r)
		v = int64(vv)
	}
	return v
}
