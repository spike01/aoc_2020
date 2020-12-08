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

	// Ok, let's try moving through a file without storing every line in
	// memory...  Rules:
	// 1. if we actually read a line, store it
	// 2. if we have to jump back to a line we haven't seen yet, we can go from
	//    the last read line
	// 3. why do i do this to myself...
	//
	//
	// tiny 5 line example (input_smol_smol.txt) - expect a loop
	//
	// 0 nop +1
	// 1 acc +1
	// 2 jmp +3
	// 3 acc +2
	// 4 acc +1
	// 5 jmp -2
	//
	// seen: [0,1,2,5,3,4] (exit on return to 5)
	// acc: 4

	var acc int
	var pos int
	var line int

	lines := make(map[int]string)
	seen := make(map[int]struct{})

	for sc.Scan() {
		if pos > line {
			lines[line] = sc.Text()
			line++
			continue
		}

		// Termination condition - have we been here before?
		_, ok := seen[pos]
		if ok {
			fmt.Println("Acc:", acc)
			os.Exit(0)
		}

		// Read the line, even if it's a jmp we need to read and store it
		next := sc.Text()
		fmt.Printf("%s pos=%d line=%d acc=%d seen=%v lines=%v\n", next, pos, line, acc, seen, lines)
		lines[line] = next

		// Decipher the instruction
		m := lineRegex.FindStringSubmatch(next)
		instruction := m[1]
		sign := m[2]
		val, err := strconv.Atoi(m[3])
		if err != nil {
			fmt.Println("Could not convert:", err)
		}

		// Store the fact that we have seen the instruction
		seen[pos] = struct{}{}

		// Do the thing
		switch instruction {
		case "acc":
			if sign == "+" {
				acc += val
			}
			if sign == "-" {
				acc -= val
			}
			pos++
		case "jmp":
			if sign == "+" {
				pos += val
			}
			if sign == "-" {
				pos -= val
			}
		case "nop":
			pos++
		}

		for pos <= line {
			// Termination condition - have we been here before?
			_, ok := seen[pos]
			if ok {
				fmt.Println("Acc:", acc)
				os.Exit(0)
			}

			// Read the line, even if it's a jmp we need to read and store it
			next := lines[pos]
			fmt.Printf("%s pos=%d line=%d acc=%d seen=%v lines=%v\n", next, pos, line, acc, seen, lines)
			lines[line] = next

			// Decipher the instruction
			m := lineRegex.FindStringSubmatch(next)
			instruction := m[1]
			sign := m[2]
			val, err := strconv.Atoi(m[3])
			if err != nil {
				fmt.Println("Could not convert:", err)
			}

			// Store the fact that we have seen the instruction
			seen[pos] = struct{}{}

			// Do the thing
			switch instruction {
			case "acc":
				if sign == "+" {
					acc += val
				}
				if sign == "-" {
					acc -= val
				}
				pos++
			case "jmp":
				if sign == "+" {
					pos += val
				}
				if sign == "-" {
					pos -= val
				}
			case "nop":
				pos++
			}
		}
		line++
	}

	os.Exit(1) // No solution found? Or return acc?
}
