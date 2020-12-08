package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var lineRegex = regexp.MustCompile(`(acc|jmp|nop) (\+|-)(\d+)`)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var acc int
	var pos int
	var lines []string

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	seen := make(map[int]struct{})

	for pos < len(lines) {
		_, ok := seen[pos]
		if ok {
			fmt.Println("Acc:", acc)
			os.Exit(1)
		}

		seen[pos] = struct{}{}
		next := lines[pos]

		m := lineRegex.FindStringSubmatch(next)
		instruction := m[1]
		sign := m[2]
		val, err := strconv.Atoi(m[3])
		if err != nil {
			fmt.Println("Could not convert:", err)
		}

		switch instruction {
		case "acc":
			if sign == "+" {
				acc += val
			}
			if sign == "-" {
				acc -= val
			}
		case "jmp":
			if sign == "+" {
				pos += val
				continue
			}
			if sign == "-" {
				pos -= val
				continue
			}
		}
		pos++
	}

	os.Exit(0)
}
