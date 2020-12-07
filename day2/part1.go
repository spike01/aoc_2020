package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type line struct {
	min  int
	max  int
	char string
	pw   string
}

func newLineFromMatches(matches []string) *line {
	m := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			m[name] = matches[i]
		}
	}
	min, err := strconv.Atoi(m["min"])
	if err != nil {
		log.Fatalf("Could not convert: %s", err)
	}

	max, err := strconv.Atoi(m["max"])
	if err != nil {
		log.Fatalf("Could not convert: %s", err)
	}

	return &line{
		min:  min,
		max:  max,
		char: m["char"],
		pw:   m["pw"],
	}
}

func (l *line) isValid() bool {
	c := strings.Count(l.pw, l.char)
	return l.min <= c && c <= l.max
}

var r = regexp.MustCompile(`(?P<min>\d+)-(?P<max>\d+) (?P<char>\w): (?P<pw>\w+)`)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not read file: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var count int

	for scanner.Scan() {
		next := scanner.Text()
		m := r.FindStringSubmatch(next)
		l := newLineFromMatches(m)
		if l.isValid() {
			count++
		}
	}
	fmt.Println("Valid passwords:", count)
	os.Exit(0)
}
