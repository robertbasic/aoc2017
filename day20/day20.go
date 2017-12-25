package day20

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type coord struct {
	x int
	y int
	z int
}

func (c *coord) add(oc coord) {
	c.x += oc.x
	c.y += oc.y
	c.z += oc.z
}

func (c coord) dist() int {
	return int(math.Abs(float64(c.x))) + int(math.Abs(float64(c.y))) + int(math.Abs(float64(c.z)))
}

type Particle struct {
	id int
	p  coord
	v  coord
	a  coord
	d  int
}

func (p *Particle) move() {
	p.v.add(p.a)
	p.p.add(p.v)
	p.d = p.p.dist()
}

func newParticle(i int, p coord, v coord, a coord) *Particle {
	prt := Particle{
		id: i,
		p:  p,
		v:  v,
		a:  a,
	}
	return &prt
}

// Day20 solves the puzzles for day 20
func Day20(folder string) {
	file, _ := os.Open(folder + "/day20.txt")

	scanner := bufio.NewScanner(file)

	var particles []Particle

	var i int
	for scanner.Scan() {
		p, v, a := parse(scanner.Text())
		prt := newParticle(i, p, v, a)
		particles = append(particles, *prt)
		i++
	}

	c := closest(particles)
	log.Println("Closest particle is: ", c.id)
}

func closest(particles []Particle) Particle {
	var wg sync.WaitGroup
	wg.Add(len(particles))

	var c Particle

	pch := make(chan Particle, len(particles))

	for _, particle := range particles {
		go move(particle, 285, pch, &wg)
	}

	wg.Wait()

	close(pch)

	for particle := range pch {
		if c.d == 0 || particle.d < c.d {
			c = particle
		}
	}

	return c
}

func move(particle Particle, iterations int, pch chan Particle, wg *sync.WaitGroup) {
	for i := 0; i < iterations; i++ {
		particle.move()
	}
	pch <- particle
	wg.Done()
}

func parse(pva string) (coord, coord, coord) {
	parts := strings.Fields(pva)
	p := parsecoord(parts[0][3 : len(parts[0])-2])
	v := parsecoord(parts[1][3 : len(parts[1])-2])
	a := parsecoord(parts[2][3 : len(parts[2])-1])
	return p, v, a
}

func parsecoord(s string) coord {
	cs := strings.Split(s, ",")
	x, _ := strconv.Atoi(cs[0])
	y, _ := strconv.Atoi(cs[1])
	z, _ := strconv.Atoi(cs[2])
	c := coord{
		x: x,
		y: y,
		z: z,
	}
	return c
}
