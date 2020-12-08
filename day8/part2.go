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

	var lineCount int
	lines := make(map[int]string)

	for sc.Scan() {
		lines[lineCount] = sc.Text()
		lineCount++
	}

Outer:
	for i := 0; i < len(lines); i++ {

		var acc int
		var pos int
		seen := make(map[int]struct{})

		for pos <= len(lines) {
			if pos == len(lines) {
				fmt.Println("Acc:", acc)
				os.Exit(0)
			}
			_, ok := seen[pos]
			if ok {
				// Loop found, try modifying next line
				continue Outer
			}

			next, ok := lines[pos]
			if !ok {
				// Somehow we've jmped out of range
				continue Outer
			}

			seen[pos] = struct{}{}

			m := lineRegex.FindStringSubmatch(next)
			instruction := m[1]
			sign := m[2]
			val, err := strconv.Atoi(m[3])
			if err != nil {
				fmt.Println("Could not convert:", err)
			}

			if pos == i {
				switch instruction {
				case "jmp":
					instruction = "nop"
				case "nop":
					instruction = "jmp"
				}
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
	}

	os.Exit(1) // No solution found
}
