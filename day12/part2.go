package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	x := 0
	y := 0

	w_x := 10
	w_y := 1

	for sc.Scan() {
		next := sc.Text()
		action := next[0]
		value, err := strconv.Atoi(next[1:])
		if err != nil {
			log.Fatalf("Could not parse: %s", err)
		}

		switch action {
		case 'N':
			w_y += value
		case 'S':
			w_y -= value
		case 'E':
			w_x += value
		case 'W':
			w_x -= value
		case 'L':
			rotation := value / 90
			switch rotation {
			case 1:
				tmp_x := w_x
				tmp_y := w_y
				w_x = -tmp_y
				w_y = tmp_x
			case 2:
				tmp_x := w_x
				tmp_y := w_y
				w_x = -tmp_x
				w_y = -tmp_y
			case 3:
				tmp_x := w_x
				tmp_y := w_y
				w_x = tmp_y
				w_y = -tmp_x
			}
		case 'R':
			rotation := value / 90
			switch rotation {
			case 1:
				tmp_x := w_x
				tmp_y := w_y
				w_x = tmp_y
				w_y = -tmp_x
			case 2:
				tmp_x := w_x
				tmp_y := w_y
				w_x = -tmp_x
				w_y = -tmp_y
			case 3:
				tmp_x := w_x
				tmp_y := w_y
				w_x = -tmp_y
				w_y = tmp_x
			}
		case 'F':
			x += (value * w_x)
			y += (value * w_y)
		}
	}

	fmt.Println(abs(x) + abs(y))

	os.Exit(0)
}

func posMod(n, modulus int) int {
	mod := n % modulus
	if mod < 0 {
		return mod + modulus
	}
	return mod
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
