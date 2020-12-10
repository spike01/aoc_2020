package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var adapters []int
	var max int

	adapters = append(adapters, 0)

	for sc.Scan() {
		next := sc.Text()
		i, err := strconv.Atoi(next)
		if err != nil {
			log.Fatalf("Could not convert: %s", err)
		}

		if i > max {
			max = i
		}
		adapters = append(adapters, i)
	}
	adapters = append(adapters, max+3)
	sort.Ints(adapters)

	var gaps []int

	for i, _ := range adapters {
		if i == len(adapters)-1 {
			break
		}
		diff := adapters[i+1] - adapters[i]
		gaps = append(gaps, diff)
	}

	var consolidated []int
	var count int

	for _, v := range gaps {
		if v == 3 {
			if count > 1 {
				consolidated = append(consolidated, count)
			}
			count = 0
			continue
		}
		count++
	}

	total := 1
	for _, v := range consolidated {
		switch v {
		case 4:
			total *= 7
		case 3:
			total *= 4
		case 2:
			total *= 2
		}
	}
	fmt.Println(total)
	os.Exit(0)
}
