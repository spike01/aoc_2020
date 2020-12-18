package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type coord struct {
	x int
	y int
	z int
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

func (s *space) String() string {
	var out strings.Builder
	var maxX, maxY, maxZ int
	var minX, minY, minZ int

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
		if coord.x < minX {
			minX = coord.x
		}
		if coord.y < minY {
			minY = coord.y
		}
		if coord.z < minZ {
			minZ = coord.z
		}
	}

	for z := minZ; z <= maxZ; z++ {
		fmt.Fprintf(&out, "z=%d\n", z)
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				cube, ok := s.cubes[*&coord{x: x, y: y, z: z}]
				if !ok {
					fmt.Fprintf(&out, ".")
					continue
				}
				fmt.Fprintf(&out, string(cube.state))
			}
			fmt.Fprintf(&out, "\n")
		}
		fmt.Fprintf(&out, "\n\n")
	}
	return out.String()
}

func (s *space) expand() {
	var maxX, maxY, maxZ int
	var minX, minY, minZ int

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
		if coord.x < minX {
			minX = coord.x
		}
		if coord.y < minY {
			minY = coord.y
		}
		if coord.z < minZ {
			minZ = coord.z
		}
	}

	for z := minZ - 1; z <= maxZ+1; z++ {
		for x := minX - 1; x <= maxX+1; x++ {
			for y := minY - 1; y <= maxY+1; y++ {
				coord := &coord{x: x, y: y, z: z}
				_, ok := s.cubes[*coord]
				if !ok {
					s.cubes[*coord] = &cube{state: '.'}
					continue
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

	for _, i := range x {
		for _, j := range y {
			for _, k := range z {
				coord := &coord{x: c.x + i, y: c.y + j, z: c.z + k}
				if !(i == 0 && j == 0 && k == 0) {
					neighbours = append(neighbours, coord)
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
			coord := &coord{x: x, y: y, z: 0}
			cube := &cube{state: v}
			s.cubes[*coord] = cube
		}
		y++
	}

	fmt.Println(s)

	var step int
	for step < 6 {
		fmt.Println("step", step+1)

		s.expand()
		s.setNextState()
		s.changeState()

		fmt.Println(s)
		step++
	}

	fmt.Println(s.active)

	os.Exit(0)
}
