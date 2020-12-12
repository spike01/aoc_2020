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
			w_x, w_y = rotate(value, w_x, w_y)
		case 'R':
			w_x, w_y = rotate(flip(value), w_x, w_y)
		case 'F':
			x += (value * w_x)
			y += (value * w_y)
		}
	}

	fmt.Println(abs(x) + abs(y))

	os.Exit(0)
}

func flip(n int) int {
	if n == 90 {
		return 270
	}
	if n == 270 {
		return 90
	}
	return n
}

func rotate(value, w_x, w_y int) (int, int) {
	switch value {
	case 90:
		return -w_y, w_x
	case 180:
		return -w_x, -w_y
	case 270:
		return w_y, -w_x
	}
	return w_x, w_y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
