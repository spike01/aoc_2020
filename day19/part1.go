package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type rule struct {
	left, right string
}

var ruleRegexp *regexp.Regexp

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	var part2 bool

	var count int
	rules := make(map[string]*rule)
	var zeroRule []string

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		next := sc.Text()
		if next == "" {
			part2 = true
			fmt.Println(zeroRule)

			for i, v := range zeroRule {
				zeroRule[i] = resolve(rules, v)
			}

			ruleRegexp = regexp.MustCompile(fmt.Sprintf("^%s$", strings.Join(zeroRule, "")))
			fmt.Println(ruleRegexp)

			continue
		}
		if part2 {
			if ruleRegexp.MatchString(next) {
				count++
			}
			continue
		}
		splitLine := strings.Split(next, ": ")
		key := splitLine[0]
		if key == "0" {
			zeroRule = strings.Split(splitLine[1], " ")
			continue
		}
		if splitLine[1][0] == '"' {
			fmt.Println(strings.Split(splitLine[1], "\"")[1])
			rules[key] = &rule{left: strings.Split(splitLine[1], "\"")[1]}
			continue
		}
		leftRight := strings.Split(splitLine[1], " | ")

		if len(leftRight) == 1 {
			rules[key] = &rule{left: leftRight[0]}
			continue
		}
		rules[key] = &rule{left: leftRight[0], right: leftRight[1]}
	}

	fmt.Println(count)

	os.Exit(0)
}

func resolve(rules map[string]*rule, v string) string {
	r, ok := rules[v]
	if !ok {
		log.Fatalf("Could not find rule: %s", v)
	}
	if len(r.left) == 1 { // Terminal rule, return the letter
		return r.left
	}
	var leftInts, rightInts []string
	if len(r.left) > 0 {
		leftInts = strings.Split(r.left, " ")
	}
	if len(r.right) > 0 {
		rightInts = strings.Split(r.right, " ")
	}
	var out strings.Builder
	for _, l := range leftInts {
		fmt.Fprintf(&out, resolve(rules, l))
	}
	if len(rightInts) > 0 {
		fmt.Fprintf(&out, "|")
	}
	for _, r := range rightInts {
		fmt.Fprintf(&out, resolve(rules, r))
	}
	return fmt.Sprintf("(%s)", out.String())
}
