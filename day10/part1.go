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

	var ones int
	var threes int

	for i, _ := range adapters {
		if i == len(adapters)-1 {
			break
		}
		switch diff := adapters[i+1] - adapters[i]; diff {
		case 3:
			threes++
		case 1:
			ones++
		}
	}

	fmt.Println("ones:", ones)
	fmt.Println("threes:", threes)
	fmt.Println("ones * threes:", ones*threes)
	os.Exit(0)
}
