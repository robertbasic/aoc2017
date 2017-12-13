package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type register struct {
	name         string
	value        int
	largestValue int
}

func (r *register) doop(op string, val int) {
	switch op {
	case "inc":
		r.value = r.value + val
	case "dec":
		r.value = r.value - val
	}

	if r.value > r.largestValue {
		r.largestValue = r.value
	}
}

func (r register) satisfies(condition string, value int) bool {
	switch condition {
	case "==":
		return r.value == value
	case "!=":
		return r.value != value
	case ">":
		return r.value > value
	case ">=":
		return r.value >= value
	case "<":
		return r.value < value
	case "<=":
		return r.value <= value
	}

	return false
}

type instruction struct {
	valreg  register
	op      string
	val     int
	condreg register
	cond    string
	condval int
}

// Day8 solves the puzzles for day 8
func Day8(logger *log.Logger) {
	file, _ := os.Open("./inputs/day8.txt")
	defer file.Close()

	i, ie := FindLargestAtEnd(file)

	logger.Println("Largest value in registries at the end: ", i)
	logger.Println("Largest value in registries ever: ", ie)
}

// FindLargestAtEnd finds the largest value that
// shows up in the registers over time
func FindLargestAtEnd(input io.Reader) (int, int) {
	var l int
	var le int

	scanner := bufio.NewScanner(input)

	regs := make(map[string]register)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		ins := makeInstruction(scanner.Text(), regs)

		if ins.condreg.satisfies(ins.cond, ins.condval) {
			ins.valreg.doop(ins.op, ins.val)
		}

		if ins.valreg.value > le {
			le = ins.valreg.value
		}

		regs[ins.valreg.name] = ins.valreg
	}

	for _, reg := range regs {
		if reg.value > l {
			l = reg.value
		}
	}

	return l, le
}

func makeInstruction(line string, regs map[string]register) instruction {
	parts := strings.Fields(line)

	valregname := parts[0]
	condregname := parts[4]

	if _, ok := regs[valregname]; !ok {
		valreg := register{name: valregname}
		regs[valregname] = valreg
	}
	if _, ok := regs[condregname]; !ok {
		condreg := register{name: condregname}
		regs[condregname] = condreg
	}

	valreg := regs[valregname]
	condreg := regs[condregname]
	val, _ := strconv.Atoi(parts[2])
	condval, _ := strconv.Atoi(parts[6])

	return instruction{
		valreg:  valreg,
		op:      parts[1],
		val:     val,
		condreg: condreg,
		cond:    parts[5],
		condval: condval,
	}
}
