package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not read file: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	slopes := []int{1, 3, 5, 7, 1}
	positions := []int{0, 0, 0, 0, 0}
	counts := []int{0, 0, 0, 0, 0}

	var line int

	for scanner.Scan() {
		next := scanner.Text()
		for i, n := range slopes {
			if !(i == 4 && line%2 != 0) { // special case, slope that skips every other line
				if string(next[positions[i]]) == "#" {
					counts[i]++
				}
				positions[i] += n
				positions[i] = positions[i] % len(next)
			}
		}
		line++
	}

	fmt.Println("Counts:", counts)
	trees := 1
	for _, n := range counts {
		trees *= n
	}
	fmt.Println("Trees:", trees)

	os.Exit(0)
}
