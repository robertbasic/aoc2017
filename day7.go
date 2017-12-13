package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type tower struct {
	name          string
	weight        int
	carriedWeight int
	totalWeight   int
	towers        towers
}

func (t *tower) adjustWeightBy(w int) {
	t.weight = t.weight - w
	t.totalWeight = t.carriedWeight + t.weight
}

type towers []tower

// Day7 solves the puzzles for day 7
func Day7(logger *log.Logger) {
	file, _ := os.Open("./inputs/day7.txt")
	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	bottomTower := FindBottomTower(lines)
	logger.Println("Bottom tower is: ", bottomTower)

	w := BalanceTower(bottomTower, lines)
	logger.Println("Weight required to balance tower", w)
}

// BalanceTower finds the weight required to balance the tower
func BalanceTower(root string, subs []string) int {
	t := buildTower(root, subs)

	t = findOOBTower(t)

	return t.weight
}

func findOOBTower(t tower) tower {
	ws := make(map[int]towers)
	for _, st := range t.towers {
		ws[st.totalWeight] = append(ws[st.totalWeight], st)
	}

	var wd int
	var oobt tower

	for w, t := range ws {
		if len(ws[w]) == 1 {
			wd += w
			oobt = t[0]
		} else {
			wd -= w
		}
	}

	if wd > 0 {
		oobt.adjustWeightBy(wd)

		return findOOBTower(oobt)
	}

	return t
}

func buildTower(root string, subs []string) tower {
	t := tower{name: root}

	for _, sub := range subs {
		if strings.HasPrefix(sub, t.name) {
			t.weight = getweight(sub)

			re := regexp.MustCompile("( [a-z]+)")

			for _, subtower := range re.FindAllString(sub, -1) {
				t.towers = append(t.towers, buildTower(strings.TrimSpace(subtower), subs))
			}
		}
	}

	cw := 0

	for _, st := range t.towers {
		cw += st.weight + st.carriedWeight
	}

	t.carriedWeight = cw
	t.totalWeight = t.weight + cw

	return t
}

func getweight(l string) int {
	re := regexp.MustCompile("[0-9]+")
	w, _ := strconv.Atoi(re.FindString(l))
	return w
}

// FindBottomTower finds the bottom tower
func FindBottomTower(lines []string) string {
	holdingUp := make(map[string]bool)
	heldUp := make(map[string]bool)

	for _, l := range lines {
		if l == "" {
			continue
		}

		if !strings.Contains(l, "->") {
			continue
		}

		re := regexp.MustCompile("^[a-z]+[^ ]*")

		holder := strings.TrimSpace(re.FindString(l))
		holdingUp[holder] = true

		re = regexp.MustCompile("( [a-z]+)")

		for _, heldee := range re.FindAllString(l, -1) {
			heldUp[strings.TrimSpace(heldee)] = true
		}
	}

	for k := range holdingUp {
		if _, ok := heldUp[k]; !ok {
			return k
		}
	}

	return ""
}
