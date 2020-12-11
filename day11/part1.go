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

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var row int
	var width int
	board := make(map[coord]*pos)

	for sc.Scan() {
		next := sc.Text()
		for column, v := range next {
			coord := &coord{row: row, column: column}
			pos := &pos{state: v}
			board[*coord] = pos
			width = len(next)
		}
		row++
	}

	setNeighbours(board, row, width)

	changed := 1
	for changed != 0 {
		changed = 0
		setNextState(board, row, width, &changed)
		changeState(board, row, width)
	}

	fmt.Println(countOccupied(board, row, width))

	os.Exit(0)
}

func setNeighbours(board map[coord]*pos, row, width int) {
	for r := 0; r < row; r++ {
		for c := 0; c < width; c++ {
			coord := &coord{row: r, column: c}
			pos := board[*coord]
			pos.neighbours = neighbours(coord, board)
		}
	}
}

func neighbours(c *coord, b map[coord]*pos) []*pos {
	var neighbours []*pos
	x := []int{-1, 0, 1}
	y := []int{-1, 0, 1}

	for _, i := range x {
		for _, j := range y {
			coord := &coord{row: c.row + i, column: c.column + j}
			n, ok := b[*coord]
			if ok && !(i == 0 && j == 0) {
				neighbours = append(neighbours, n)
			}
		}
	}
	return neighbours
}

func setNextState(board map[coord]*pos, row int, width int, changed *int) {
	for r := 0; r < row; r++ {
		for c := 0; c < width; c++ {
			coord := &coord{row: r, column: c}
			pos := board[*coord]
			if pos.state == '.' {
				pos.nextState = '.'
				continue
			}

			var occupied int
			for _, p := range pos.neighbours {
				if p.state == '#' {
					occupied++
				}
			}

			if pos.state == 'L' && occupied == 0 {
				pos.nextState = '#'
				*changed++
			}

			if pos.state == '#' && occupied >= 4 {
				pos.nextState = 'L'
				*changed++
			}
		}
	}
}

func changeState(b map[coord]*pos, row, width int) {
	for r := 0; r < row; r++ {
		for c := 0; c < width; c++ {
			coord := &coord{row: r, column: c}
			pos := b[*coord]
			pos.state = pos.nextState
		}
	}
}

func countOccupied(b map[coord]*pos, row, width int) int {
	var count int
	for r := 0; r < row; r++ {
		for c := 0; c < width; c++ {
			coord := &coord{row: r, column: c}
			pos := b[*coord]
			if pos.state == '#' {
				count++
			}
		}
	}
	return count
}

func printBoard(b map[coord]*pos, row, width int) {
	for r := 0; r < row; r++ {
		for c := 0; c < width; c++ {
			coord := &coord{row: r, column: c}
			fmt.Printf("%s", string(b[*coord].state))
		}
		fmt.Printf("\n")
	}
}
