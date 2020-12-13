package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var start int
	var scheds []int

	for sc.Scan() {
		next := sc.Text()
		if start == 0 {
			start, err = strconv.Atoi(next)
			if err != nil {
				log.Fatalf("Could not parse: %s", err)
			}
			continue
		}

		times := strings.Split(next, ",")
		for _, v := range times {
			if v == "x" {
				scheds = append(scheds, 0)
				continue
			}
			sched, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("Could not parse: %s", err)
			}
			scheds = append(scheds, sched)
		}
	}

	step := scheds[0]
	var time int
	for j, v := range scheds[1:] {
		if v == 0 {
			continue
		}
		for (time+j+1)%v != 0 {
			time += step
		}
		step *= v
	}
	fmt.Println(time)
	os.Exit(0)
}
