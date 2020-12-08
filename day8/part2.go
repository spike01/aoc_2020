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

	var lineCount int
	var acc int
	var pos int

	lines := make(map[int]string)

	for sc.Scan() {
		lines[lineCount] = sc.Text()
		lineCount++
	}

	seen := make(map[int]struct{})

	fmt.Println("Lines:", len(lines))

Outer:
	for i := 0; i < len(lines); i++ {

		// Reset for new attempt
		acc = 0
		pos = 0
		seen = make(map[int]struct{})

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

			line := strings.Split(next, " ")
			arg := line[1]
			sign := string(arg[0])
			val, err := strconv.Atoi(arg[1:])
			if err != nil {
				fmt.Println("Could not convert:", err)
			}

			instruction := line[0]

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

	os.Exit(0)
}
