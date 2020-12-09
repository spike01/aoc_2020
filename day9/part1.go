package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const preambleLength = 25

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var preamble []int
	var counter int

	for sc.Scan() {
		next := sc.Text()
		n, err := strconv.Atoi(next)
		if err != nil {
			log.Fatalf("Could not convert: %s", err)
		}

		if counter < 25 {
			preamble = append(preamble, n)
			counter++
			continue
		}

		if !findSum(preamble, n) {
			fmt.Println(n)
			os.Exit(0)
		}

		preamble = preamble[1:]
		preamble = append(preamble, n)
	}

	os.Exit(0)
}

func findSum(chunk []int, n int) bool {
	tmp := make([]int, len(chunk))
	copy(tmp, chunk)
	sort.Ints(tmp)
	l := 0
	r := len(tmp) - 1
	for l < r {
		sum := tmp[l] + tmp[r]
		switch {
		case sum == n:
			return true
		case sum < n:
			l++
		case sum > n:
			r--
		}
	}
	return false
}
