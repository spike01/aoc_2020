package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var directions = []rune{'N', 'E', 'S', 'W'}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	x := 0
	y := 0
	facing := 1

	for sc.Scan() {
		next := sc.Text()
		action := next[0]
		value, err := strconv.Atoi(next[1:])
		if err != nil {
			log.Fatalf("Could not parse: %s", err)
		}

		switch action {
		case 'N':
			y += value
		case 'S':
			y -= value
		case 'E':
			x += value
		case 'W':
			x -= value
		case 'L':
			facing = posMod(facing-(value/90), 4)
		case 'R':
			facing = posMod(facing+(value/90), 4)
		case 'F':
			switch directions[facing] {
			case 'N':
				y += value
			case 'S':
				y -= value
			case 'E':
				x += value
			case 'W':
				x -= value
			}
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
