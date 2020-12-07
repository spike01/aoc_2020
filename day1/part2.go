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

	seen := make(map[int][]int)

	for scanner.Scan() {
		next := scanner.Text()
		i, err := strconv.Atoi(next)
		if err != nil {
			log.Fatalf("Could not convert number: %s", next)
		}

		// Check if there's an winning entry
		_, ok := seen[i]
		if ok && len(seen[i]) == 2 {
			first := i
			second := seen[i][0]
			third := seen[i][1]
			fmt.Printf("triple found: %d + %d + %d = %d\n", first, second, third, first+second+third)
			fmt.Printf("%d * %d * %d = %d\n", first, second, third, first*second*third)
			os.Exit(0)
		}

		remaining := 2020 - i
		seen[remaining] = make([]int, 0)
		seen[remaining] = append(seen[remaining], i)

		// Walk back over seen and update existing entries
		for k, v := range seen {
			// if we just stored this, skip
			if contains(v, i) {
				continue
			}
			// if list has 3 entries, delete it
			if len(v) == 3 {
				delete(seen, k)
			}
			seen[k-i] = append(v, i)
		}
	}
}

func contains(nums []int, i int) bool {
	for j := range nums {
		if nums[j] == i {
			return true
		}
	}
	return false
}
