package day16

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"fmt"
)

type instruction struct {
	op string
	spinarg int
	swapargs []int
	partnerargs []string
}

// Day16 solves the puzzles for day 16
func Day16(folder string) {
	file, _ := os.Open(folder + "/day16.txt")

	scanner := bufio.NewScanner(file)

	programs := "abcdefghijklmnop"

	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	instructions := processInput(line)

	programs = dance(programs, instructions)
	log.Println("The new lineup is: ", programs)

	programs = "abcdefghijklmnop"
	for i := 0; i < 40; i++ {
		programs = dance(programs, instructions)
		fmt.Println(i, ": ", programs)
	}
	log.Println("The new lineup is: ", programs)
}

func processInput(input string) []instruction {
	var instructions []instruction

	r := bufio.NewReader(strings.NewReader(input))

	for {
		ins, err := r.ReadBytes(',')

		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}

		instruction := instruction{}

		switch ins[0] {
		case 's':
			size, _ := strconv.Atoi(string(ins[1 : len(ins)-1]))
			instruction.op = string(ins[0])
			instruction.spinarg = size
		case 'x':
			f, t := getft(ins)
			if f > t {
				f, t = t, f
			}
			instruction.op = string(ins[0])
			instruction.swapargs = []int{f, t}
		case 'p':
			x, y := getxy(ins)
			instruction.op = string(ins[0])
			instruction.partnerargs = []string{x, y}
		}

		instructions = append(instructions, instruction)

		if err == io.EOF {
			break
		}
	}

	return instructions
}

func dance(programs string, instructions []instruction) string {
	for _, instruction := range instructions {
		switch instruction.op {
		case "s":
			programs = spin(programs, instruction.spinarg)
		case "x":
			programs = swap(programs, instruction.swapargs[0], instruction.swapargs[1])
		case "p":
			programs = partner(programs, instruction.partnerargs[0], instruction.partnerargs[1])
		}
	}

	return programs
}

func spin(programs string, size int) string {
	return programs[len(programs)-size:] + programs[0:len(programs)-size]
}

func swap(programs string, from int, to int) string {
	return programs[0:from] + programs[to:to+1] + programs[from+1:to] + programs[from:from+1] + programs[to+1:]
}

func partner(programs string, x string, y string) string {
	f := strings.Index(programs, x)
	t := strings.Index(programs, y)

	if f > t {
		f, t = t, f
	}

	return swap(programs, f, t)
}

func getft(ins []byte) (int, int) {
	fb, tb := getftb(ins)
	f, _ := strconv.Atoi(string(fb))
	t, _ := strconv.Atoi(string(tb))

	if f > t {
		f, t = t, f
	}

	return f, t
}

func getxy(ins []byte) (string, string) {
	xb, yb := getftb(ins)

	return string(xb), string(yb)
}

func getftb(ins []byte) ([]byte, []byte) {
	var fb []byte
	var tb []byte
	var wf = true

	for i := 1; i < len(ins); i++ {
		if ins[i] == ',' {
			continue
		}
		if ins[i] == '/' {
			wf = false
			continue
		}

		if wf {
			fb = append(fb, ins[i])
		} else {
			tb = append(tb, ins[i])
		}
	}

	return fb, tb
}
