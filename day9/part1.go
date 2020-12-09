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
	for i := 0; i < len(tmp); i++ {
		for j := len(tmp) - 1; 0 < j; j-- {
			if tmp[j] > n {
				continue
			}
			if tmp[i] > n {
				continue
			}
			if tmp[i]+tmp[j] == n {
				return true
			}
		}
	}
	return false
}
