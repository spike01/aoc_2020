package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type line struct {
	pos1 int
	pos2 int
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
	pos1, err := strconv.Atoi(m["pos1"])
	if err != nil {
		log.Fatalf("Could not convert: %s", err)
	}

	pos2, err := strconv.Atoi(m["pos2"])
	if err != nil {
		log.Fatalf("Could not convert: %s", err)
	}

	return &line{
		pos1: pos1 - 1, // positions are 1-indexed
		pos2: pos2 - 1, // positions are 1-indexed
		char: m["char"],
		pw:   m["pw"],
	}
}

func (l *line) isValid() bool {
	if l.pos1 > len(l.pw) || l.pos2 > len(l.pw) {
		return false
	}
	firstMatch := string(l.pw[l.pos1]) == l.char
	secondMatch := string(l.pw[l.pos2]) == l.char
	return firstMatch != secondMatch
}

var r = regexp.MustCompile(`(?P<pos1>\d+)-(?P<pos2>\d+) (?P<char>\w): (?P<pw>\w+)`)

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
