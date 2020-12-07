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
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var buf []string
	var count int

	for sc.Scan() {
		next := sc.Text()
		if next == "" {
			fmt.Println(buf)
			count += countUnique(buf)
			buf = []string{}
			continue
		}
		buf = append(buf, next)
	}

	// Final element, buf isn't flushed on last entry
	count += countUnique(buf)
	fmt.Println("Count:", count)
	os.Exit(0)
}

func countUnique(buf []string) int {
	seen := make(map[rune]struct{})
	for _, b := range buf {
		for _, c := range b {
			seen[c] = struct{}{}
		}
	}
	return len(seen)
}
