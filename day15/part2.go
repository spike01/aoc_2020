package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var starters []int

	for sc.Scan() {
		next := sc.Text()
		split := strings.Split(next, ",")
		for _, v := range split {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("Could not convert: %s", err)
			}
			starters = append(starters, n)
		}
	}

	var next int

	target := 30000000
	turn := 1
	numbers := make(map[int]int)

	for turn < len(starters) {
		next = starters[turn-1]
		numbers[next] = turn
		turn++
	}

	next = starters[len(starters)-1]

	for turn < target {
		position, ok := numbers[next]

		if !ok {
			numbers[next] = turn
			next = 0
			turn++
			continue
		}

		numbers[next] = turn
		next = turn - position
		turn++
	}

	fmt.Println("Turn", turn, ":", next)

	os.Exit(0)
}
