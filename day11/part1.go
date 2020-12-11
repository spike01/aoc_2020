package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coord struct {
	row    int
	column int
}

type pos struct {
	state      rune
	nextState  rune
	neighbours []*pos
}

type board struct {
	positions map[coord]*pos
	rows      int
	columns   int
	changed   bool
}

func newBoard() *board {
	return &board{
		positions: make(map[coord]*pos),
		changed:   true,
	}
}

func (b *board) iterate(f func(*pos, *coord)) {
	for r := 0; r < b.rows; r++ {
		for c := 0; c < b.columns; c++ {
			coord := &coord{row: r, column: c}
			pos := b.positions[*coord]
			f(pos, coord)
		}
	}
}

func (b *board) setNeighbours() {
	b.iterate(func(p *pos, c *coord) {
		p.neighbours = b.neighbours(c)
	})
}

func (b *board) neighbours(c *coord) []*pos {
	var neighbours []*pos
	x := []int{-1, 0, 1}
	y := []int{-1, 0, 1}

	for _, i := range x {
		for _, j := range y {
			coord := &coord{row: c.row + i, column: c.column + j}
			n, ok := b.positions[*coord]
			if ok && !(i == 0 && j == 0) {
				neighbours = append(neighbours, n)
			}
		}
	}
	return neighbours
}

func (b *board) setNextState() {
	b.changed = false
	b.iterate(func(p *pos, c *coord) {
		if p.state == '.' {
			p.nextState = '.'
			return
		}

		var occupied int
		for _, n := range p.neighbours {
			if n.state == '#' {
				occupied++
			}
		}

		if p.state == 'L' && occupied == 0 {
			p.nextState = '#'
			b.changed = true
		}

		if p.state == '#' && occupied >= 4 {
			p.nextState = 'L'
			b.changed = true
		}
	})
}

func (b *board) changeState() {
	b.iterate(func(p *pos, c *coord) {
		p.state = p.nextState
	})
}

func (b *board) countOccupied() int {
	var count int
	b.iterate(func(p *pos, c *coord) {
		if p.state == '#' {
			count++
		}
	})
	return count
}

func (b *board) print() {
	b.iterate(func(p *pos, c *coord) {
		fmt.Printf("%s", string(p.state))
		if c.column == b.columns-1 {
			fmt.Printf("\n")
		}
	})
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var row int
	b := newBoard()

	for sc.Scan() {
		next := sc.Text()
		for column, v := range next {
			coord := &coord{row: row, column: column}
			pos := &pos{state: v}
			b.positions[*coord] = pos
			b.columns = len(next)
		}
		row++
	}
	b.rows = row

	b.setNeighbours()

	for b.changed {
		b.setNextState()
		b.changeState()
	}

	fmt.Println(b.countOccupied())

	os.Exit(0)
}
