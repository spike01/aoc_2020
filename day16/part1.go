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

	var nearbyTicketSection, myTicketSection bool
	var errorRate int
	rules := make(map[int]struct{})

	for sc.Scan() {
		next := sc.Text()
		if next == "" {
			continue
		}
		if next == "your ticket:" {
			myTicketSection = true
			continue
		}
		if next == "nearby tickets:" {
			nearbyTicketSection = true
			myTicketSection = false
			continue
		}
		if myTicketSection {
			continue
		}
		if nearbyTicketSection {
			fields := strings.Split(next, ",")
			for _, f := range fields {
				field, err := strconv.Atoi(f)
				if err != nil {
					log.Fatalf("Could not convert: %s", err)
				}
				_, ok := rules[field]
				if !ok {
					errorRate += field
				}
			}
			continue
		}
		nameRules := strings.Split(next, ": ")
		rawRules := strings.Split(nameRules[1], " or ")
		leftRule := strings.Split(rawRules[0], "-")
		rightRule := strings.Split(rawRules[1], "-")
		addRule(rules, leftRule)
		addRule(rules, rightRule)
	}

	fmt.Println("Error rate:", errorRate)

	os.Exit(0)
}

func addRule(rules map[int]struct{}, rule []string) {
	start, err := strconv.Atoi(rule[0])
	if err != nil {
		log.Fatalf("Could not convert: %s", err)
	}
	end, err := strconv.Atoi(rule[1])
	if err != nil {
		log.Fatalf("Could not convert: %s", err)
	}
	for i := start; i <= end; i++ {
		rules[i] = struct{}{}
	}
}
