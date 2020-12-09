package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const target = 14360655

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var set []int

	for sc.Scan() {
		next := sc.Text()
		n, err := strconv.Atoi(next)
		if err != nil {
			log.Fatalf("Could not convert: %s", err)
		}

		set = append(set, n)

		if sum(set) == target {
			fmt.Println(min(set) + max(set))
			os.Exit(0)
		}

		for sum(set) > target {
			set = set[1:]
		}
	}

	os.Exit(1)
}

func sum(chunk []int) int {
	var sum int
	for _, v := range chunk {
		sum += v
	}
	return sum
}

func min(chunk []int) int {
	min := math.MaxInt64
	for _, v := range chunk {
		if v < min {
			min = v
		}
	}
	return min
}

func max(chunk []int) int {
	max := 0
	for _, v := range chunk {
		if v > max {
			max = v
		}
	}
	return max
}
