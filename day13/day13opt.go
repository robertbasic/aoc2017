package day13

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

// ScannerOpt is a scanner in a layer
type ScannerOpt struct {
	rang      int
	position  int
	direction int
}

func (s ScannerOpt) positionAtZero(steps int) bool {
	r := s.rang - 1
	d := steps - r
	div := int(math.Trunc(float64(d / r)))
	m := int(math.Mod(float64(d), float64(r)))

	if math.Mod(float64(div), 2) == 0 {
		return r-m == 0
	}

	return m == 0
}

// Day13Opt solves the puzzles for day 13
func Day13Opt(logger *log.Logger, folder string) {
	buf, _ := ioutil.ReadFile(folder + "/day13.txt")
	input := string(buf)

	layers, md := FirewallLayersOpt(input)

	sev := WalkOnFireOpt(layers, md)
	logger.Println("Severity is: ", sev)

	delay := EvadeDelayOpt(layers, md)
	logger.Println("Delay for: ", delay)
}

// EvadeDelayOpt calculates the delay required to evade all scanners
func EvadeDelayOpt(layers map[int]ScannerOpt, md int) int {
	var delay = 0

	for {
		var cd = -1
		var caught = false

		for cd <= md {
			cd++

			if _, ok := layers[cd]; !ok {
				continue
			}

			cs := layers[cd]

			if cs.positionAtZero(cd + delay) {
				caught = true
				break
			}
		}

		if !caught {
			break
		}

		delay += 2
	}

	return delay
}

// WalkOnFireOpt walks us through the firewall
func WalkOnFireOpt(layers map[int]ScannerOpt, md int) int {
	var cd = -1
	var sev = 0

	for cd <= md {
		cd++

		if _, ok := layers[cd]; !ok {
			continue
		}

		cs := layers[cd]

		if cs.positionAtZero(cd) {
			sev += cd * cs.rang
		}
	}

	return sev
}

// FirewallLayersOpt reads the firewall layers
func FirewallLayersOpt(input string) (map[int]ScannerOpt, int) {
	lines := strings.Split(input, "\n")
	ls := make(map[int]ScannerOpt, len(lines))
	var md int

	for _, i := range lines {
		sci := strings.Index(i, ":")
		d := StringNumberToIntOpt(i[:sci])
		r := StringNumberToIntOpt(i[sci+1:])
		s := ScannerOpt{rang: r, direction: 1}

		ls[d] = s

		md = d
	}

	return ls, md
}

// StringNumberToIntOpt converts a string number to an int
func StringNumberToIntOpt(s string) int {
	i, _ := strconv.Atoi(strings.TrimSpace(s))
	return i
}
