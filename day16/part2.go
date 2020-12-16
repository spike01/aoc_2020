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
	var myTicket []int
	var names []string
	rules := make(map[int]map[string]struct{})      // i in range -> qualifying rules
	candidates := make(map[int]map[string]struct{}) // position in line -> qualifying rules

outer:
	for sc.Scan() {
		next := sc.Text()
		switch {
		case next == "":
		case next == "your ticket:":
			myTicketSection = true
		case next == "nearby tickets:":
			nearbyTicketSection = true
			myTicketSection = false
		case myTicketSection:
			fields := strings.Split(next, ",")
			for _, f := range fields {
				field, err := strconv.Atoi(f)
				if err != nil {
					log.Fatalf("Could not convert: %s", err)
				}
				myTicket = append(myTicket, field)
			}
		case nearbyTicketSection:
			fields := strings.Split(next, ",")
			for _, f := range fields {
				field, err := strconv.Atoi(f)
				if err != nil {
					log.Fatalf("Could not convert: %s", err)
				}
				_, ok := rules[field]
				if !ok {
					continue outer
				}
			}
			for i, f := range fields {
				field, err := strconv.Atoi(f)
				if err != nil {
					log.Fatalf("Could not convert: %s", err)
				}
				if len(candidates[i]) == 0 {
					candidates[i] = make(map[string]struct{})
					for _, n := range names {
						candidates[i][n] = struct{}{}
					}
				}
				if len(rules[field]) != 20 {
					notIn := difference(rules[field], names)
					delete(candidates[i], notIn)
				}
			}
		default:
			nameRules := strings.Split(next, ": ")
			rawRules := strings.Split(nameRules[1], " or ")
			leftRule := strings.Split(rawRules[0], "-")
			rightRule := strings.Split(rawRules[1], "-")
			name := nameRules[0]
			names = append(names, name)
			addRule(rules, leftRule, name)
			addRule(rules, rightRule, name)
		}
	}

	for uneliminated(candidates) {
		var key string
		for _, v := range candidates {
			if len(v) == 1 {
				for k, _ := range v {
					key = k
				}
			}
		}
		for _, v := range candidates {
			if len(v) != 1 {
				delete(v, key)
			}
		}
	}

	total := 1
	for i, v := range myTicket {
		for k, _ := range candidates[i] {
			if strings.HasPrefix(k, "departure") {
				total *= v
			}
		}
	}

	fmt.Println(total)

	os.Exit(0)
}

func addRule(rules map[int]map[string]struct{}, rule []string, name string) {
	start, err := strconv.Atoi(rule[0])
	if err != nil {
		log.Fatalf("Could not convert: %s", err)
	}
	end, err := strconv.Atoi(rule[1])
	if err != nil {
		log.Fatalf("Could not convert: %s", err)
	}
	for i := start; i <= end; i++ {
		if len(rules[i]) == 0 {
			rules[i] = make(map[string]struct{})
		}
		rules[i][name] = struct{}{}
	}
}

func difference(fieldRules map[string]struct{}, names []string) string {
	for _, n := range names {
		if _, ok := fieldRules[n]; !ok {
			return n
		}
	}
	return "notfound"
}

func uneliminated(c map[int]map[string]struct{}) bool {
	for _, v := range c {
		if len(v) == 1 {
			continue
		}
		return true
	}
	return false
}
