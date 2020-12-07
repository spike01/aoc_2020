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

	var pos int
	var count int

	for scanner.Scan() {
		next := scanner.Text()
		if string(next[pos]) == "#" {
			count++
		}
		pos += 3
		pos = pos % len(next)
	}

	fmt.Println("Trees:", count)

	os.Exit(0)
}
