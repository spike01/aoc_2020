package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coord struct {
	x, y, z, w int
}

type cube struct {
	state     rune
	nextState rune
}

type space struct {
	cubes  map[coord]*cube
	active int
}

func newSpace() *space {
	return &space{
		cubes: make(map[coord]*cube),
	}
}

func (s *space) expand() {
	var maxX, maxY, maxZ, maxW int
	var minX, minY, minZ, minW int

	for coord, _ := range s.cubes {
		if coord.x > maxX {
			maxX = coord.x
		}
		if coord.y > maxY {
			maxY = coord.y
		}
		if coord.z > maxZ {
			maxZ = coord.z
		}
		if coord.w > maxW {
			maxW = coord.w
		}
		if coord.x < minX {
			minX = coord.x
		}
		if coord.y < minY {
			minY = coord.y
		}
		if coord.z < minZ {
			minZ = coord.z
		}
		if coord.w < minW {
			minW = coord.w
		}
	}

	for w := minW - 1; w <= maxW+1; w++ {
		for z := minZ - 1; z <= maxZ+1; z++ {
			for x := minX - 1; x <= maxX+1; x++ {
				for y := minY - 1; y <= maxY+1; y++ {
					coord := &coord{x: x, y: y, z: z, w: w}
					_, ok := s.cubes[*coord]
					if !ok {
						s.cubes[*coord] = &cube{state: '.'}
						continue
					}
				}
			}
		}
	}
}

func (s *space) neighbours(c *coord) []*coord {
	var neighbours []*coord
	x := []int{-1, 0, 1}
	y := []int{-1, 0, 1}
	z := []int{-1, 0, 1}
	w := []int{-1, 0, 1}

	for _, i := range x {
		for _, j := range y {
			for _, k := range z {
				for _, l := range w {
					coord := &coord{x: c.x + i, y: c.y + j, z: c.z + k, w: c.w + l}
					if !(i == 0 && j == 0 && k == 0 && l == 0) {
						neighbours = append(neighbours, coord)
					}
				}
			}
		}
	}
	return neighbours
}

func (s *space) setNextState() {
	for co, cu := range s.cubes {
		var active int
		for _, coor := range s.neighbours(&co) {
			cub, ok := s.cubes[*coor]
			if ok && cub.state == '#' {
				active++
			}
		}
		if cu.state == '#' && active == 2 || active == 3 {
			cu.nextState = '#'
			continue
		}
		if cu.state == '.' && active == 3 {
			cu.nextState = '#'
			continue
		}
		cu.nextState = '.'
	}
}

func (s *space) changeState() {
	s.active = 0
	for _, c := range s.cubes {
		if c.nextState == '#' {
			s.active++
		}
		c.state = c.nextState
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var y int
	s := newSpace()

	for sc.Scan() {
		next := sc.Text()
		for x, v := range next {
			coord := &coord{x: x, y: y, z: 0, w: 0}
			cube := &cube{state: v}
			s.cubes[*coord] = cube
		}
		y++
	}

	var step int
	for step < 6 {
		s.expand()
		s.setNextState()
		s.changeState()
		step++
	}

	fmt.Println(s.active)

	os.Exit(0)
}
