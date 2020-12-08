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

	for pos <= len(lines) {
		_, ok := seen[pos]
		if ok {
			fmt.Println("Acc:", acc)
			os.Exit(1)
		}

		seen[pos] = struct{}{}
		next := lines[pos]

		line := strings.Split(next, " ")
		arg := line[1]
		sign := string(arg[0])
		val, err := strconv.Atoi(arg[1:])
		if err != nil {
			fmt.Println("Could not convert:", err)
		}

		switch instruction := line[0]; instruction {
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
