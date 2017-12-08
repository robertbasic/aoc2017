package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestFindBottomTower(t *testing.T) {
	in := `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)
`

	scanner := bufio.NewScanner(strings.NewReader(in))

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	r := FindBottomTower(lines)
	e := "tknk"

	if r != e {
		t.Errorf("Got %s, expected %s", r, e)
	}
}

func TestBalanceTower(t *testing.T) {
	lines := []string{
		"pbga (66)",
		"xhth (57)",
		"ebii (61)",
		"havc (66)",
		"ktlj (57)",
		"fwft (72) -> ktlj, cntj, xhth",
		"qoyq (66)",
		"padx (45) -> pbga, havc, qoyq",
		"tknk (41) -> ugml, padx, fwft",
		"jptl (61)",
		"ugml (68) -> gyxo, ebii, jptl",
		"gyxo (61)",
		"cntj (57)",
	}

	root := "tknk"

	r := BalanceTower(root, lines)

	e := 60

	if r != e {
		t.Errorf("Got %d, expected %d", r, e)
	}
}
