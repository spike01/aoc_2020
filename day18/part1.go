package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var total uint64

	for sc.Scan() {
		next := sc.Text()
		res := process(next)
		total += res
	}

	fmt.Println(total)

	os.Exit(0)
}

func process(expr string) uint64 {
	var total uint64

	var op rune
	var subExpr []rune
	var inSub bool
	var stack int

	for i, r := range expr {
		if i == 0 && r != '(' {
			v, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatalf("Could not convert: %s", err)
			}
			total = uint64(v)
			continue
		}

		if inSub {
			if r == ')' && stack == 0 {
				inSub = false
				sub := process(string(subExpr))
				subExpr = []rune{}
				if total == 0 {
					total = sub
					continue
				}
				if op == '+' {
					total += sub
					continue
				}
				if op == '*' {
					total *= sub
					continue
				}
				continue
			}
			if r == '(' {
				stack++
			}
			if r == ')' {
				stack--
			}
			subExpr = append(subExpr, r)
			continue
		}

		switch r {
		case ' ':
			continue
		case '(':
			inSub = true
			continue
		case ')':
			log.Fatalf("should never reach")
		case '+':
			op = r
		case '*':
			op = r
		default:
			c, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatalf("Could not convert: %s", err)
			}
			if op == '+' {
				total += uint64(c)
				continue
			}
			if op == '*' {
				total *= uint64(c)
				continue
			}
		}
	}

	return total
}
