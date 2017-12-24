package day18

import (
	"log"
	"math"
	"strconv"
	"strings"
	"sync"
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
	sch          chan int
	rch          chan int
	regs         map[string]int
	q            []int
	sends        int
}

func newProgram(id int, instructions []string) *Program {
	sch := make(chan int)
	regs := map[string]int{
		"p": id,
	}
	p := Program{
		id:           id,
		instructions: instructions,
		sch:          sch,
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
	var wg sync.WaitGroup
	wg.Add(2)

	p0 := newProgram(0, instructions)
	p1 := newProgram(1, instructions)

	p0.rch = p1.sch
	p1.rch = p0.sch

	log.Println("Running programs")

	go run(p0, &wg)
	go run(p1, &wg)

	wg.Wait()

	log.Println("Closing up shop")

	close(p0.sch)
	close(p1.sch)

	return p1.sends
}

func run(p *Program, wg *sync.WaitGroup) {
	log.Println("Running ID: ", p.id)

	var send = func(p *Program, reg string) {
		log.Println("Send by ", p.id, ": ", reg)
		p.sends++
		val := regVal(reg, p.regs)
		p.sch <- val
	}

	go func() {
		log.Println("Rec by ", p.id)
		val := <-p.rch
		p.q = append(p.q, val)
	}()

	var retries int
loop:
	for i := 0; i < len(p.instructions); i++ {
		ip := strings.Fields(p.instructions[i])
		op := ip[0]
		reg := ip[1]

		if op != "jgz" && len(ip) == 3 {
			doValOp(op, reg, ip[2], p.regs)
		}

		switch op {
		case "jgz":
			t, j := jump(reg, ip[2], p.regs)
			if t > 0 {
				i += j - 1
			}
		case "snd":
			log.Println("Send for ", p.id, ": ", reg)
			send(p, reg)
		case "rcv":
			//val := <-p.rch
			//p.regs[reg] = val
			//log.Println(p.id, ": ", p.q)
			if len(p.q) == 0 {
				i--
				retries++
				if retries > 5 {
					break loop
				}
			} else {
				val := p.q[0]
				p.regs[reg] = val
				p.q = p.q[1:]
			}
		}
	}
	log.Println(p.id, ": ", p.sends)
	wg.Done()
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
	v := regVal(reg, regs)

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
	t := regVal(reg, regs)
	j := regVal(val, regs)
	return t, j
}

func isReg(r string) bool {
	return r >= "a" && r <= "z"
}

func regVal(r string, regs map[string]int) int {
	var v int
	if isReg(r) {
		v = regs[r]
	} else {
		v, _ = strconv.Atoi(r)
	}
	return v
}
