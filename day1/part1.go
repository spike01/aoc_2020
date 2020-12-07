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
		log.Fatalf("Could not read file: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	seen := make(map[int]struct{})

	for scanner.Scan() {
		next := scanner.Text()
		i, err := strconv.Atoi(next)
		if err != nil {
			log.Fatalf("Could not convert number: %s", next)
		}

		pair := 2020 - i
		seen[pair] = struct{}{}

		_, ok := seen[i]
		if ok {
			fmt.Printf("Pair found: %d * %d = %d\n", i, pair, i*pair)
			break
		}
	}
	os.Exit(0)
}
