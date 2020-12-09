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

type computer struct {
	sc    *bufio.Scanner
	acc   int
	pos   int
	line  int
	lines map[int]string
	seen  map[int]struct{}
}

func newComputer(sc *bufio.Scanner) *computer {
	return &computer{
		sc:    sc,
		lines: make(map[int]string),
		seen:  make(map[int]struct{}),
	}
}

func (c *computer) execute(next string) {
	m := lineRegex.FindStringSubmatch(next)
	op, sign := m[1], m[2]
	val, err := strconv.Atoi(m[3])
	if err != nil {
		fmt.Println("Could not convert:", err)
	}

	switch op {
	case "acc":
		if sign == "+" {
			c.acc += val
		}
		if sign == "-" {
			c.acc -= val
		}
		c.pos++
	case "jmp":
		if sign == "+" {
			c.pos += val
		}
		if sign == "-" {
			c.pos -= val
		}
	case "nop":
		c.pos++
	}
}

func (c *computer) printState(next string) {
	fmt.Printf("%s pos=%d line=%d acc=%d seen=%v lines=%v\n", next, c.pos, c.line, c.acc, len(c.seen), len(c.lines))
}

func (c *computer) process(next string) {
	_, ok := c.seen[c.pos]
	if ok {
		fmt.Println("Acc:", c.acc)
		os.Exit(0)
	}

	c.printState(next)
	c.lines[c.pos] = next
	c.seen[c.pos] = struct{}{}
	c.execute(next)
}

// Amusingly, trying to not read the entire file only saves reading 29 lines.
// Probably the only improvement from here is to actually move through the
// file based on newlines...
func (c *computer) run() {
	for c.sc.Scan() {
		// Catch up lines
		if c.pos > c.line {
			c.lines[c.line] = c.sc.Text() // Need to store these in case we need to jump back
			c.line++
			continue
		}

		// Line and pos are equal
		next := c.sc.Text()
		c.process(next)

		// Read from history until caught up
		for c.pos <= c.line {
			next := c.lines[c.pos]
			c.process(next)
		}
		c.line++
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	c := newComputer(sc)
	c.run()

	os.Exit(1) // No solution found? Or return acc?
}
