package day19

import (
	"strings"
	"testing"
)

func TestWalk(t *testing.T) {
	tubes := `
     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ 

	`

	tr := strings.NewReader(tubes)
	rl, rs := walk(tr, 7)

	el := "ABCDEF"
	es := 38

	if rl != el {
		t.Errorf("Got %s, expected %s", rl, el)
	}

	if rs != es {
		t.Errorf("Got %d, expected %d", rs, es)
	}
}

func TestWalk2(t *testing.T) {
	tubes := `
     |
     |
     A
     |
     |
   +-+
   |
   B
	`

	tr := strings.NewReader(tubes)
	r, _ := walk(tr, 8)

	e := "AB"

	if r != e {
		t.Errorf("Got %s, expected %s", r, e)
	}
}
