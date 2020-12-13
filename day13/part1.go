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
	var min int
	var id int

	for sc.Scan() {
		next := sc.Text()
		if start == 0 {
			start, err = strconv.Atoi(next)
			if err != nil {
				log.Fatalf("Could not parse: %s", err)
			}
			min = start
			continue
		}

		times := strings.Split(next, ",")
		for _, v := range times {
			if v != "x" {
				sched, err := strconv.Atoi(v)
				if err != nil {
					log.Fatalf("Could not parse: %s", err)
				}
				i := 1
				for i*sched < start {
					i++
				}
				d := i*sched - start
				if d < min {
					min = d
					id = sched
				}
			}
		}
	}

	fmt.Println(min * id)
	os.Exit(0)
}
