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

var memRegexp = regexp.MustCompile(`\[(\d+)\]`)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var zeroMask int
	var oneMask int
	memory := make(map[int]int)

	for sc.Scan() {
		next := sc.Text()
		line := strings.Split(next, " = ")
		instruction := line[0]
		val := line[1]
		m := memRegexp.FindStringSubmatch(instruction)
		if len(m) > 0 {
			addr := m[1]
			a, err := strconv.Atoi(addr)
			if err != nil {
				log.Fatalf("Could not convert: %s", err)
			}
			v, err := strconv.Atoi(val)
			if err != nil {
				log.Fatalf("Could not convert: %s", err)
			}
			memory[a] = (v | oneMask) & zeroMask
			continue
		}

		// Set bitmasks
		zeroMask = 0xFFFFFFFFF // 2^36-1
		oneMask = 0

		for k := 0; k < len(val); k++ {
			if val[len(val)-1-k] == '1' {
				oneMask |= 1 << k
			}
			if val[len(val)-1-k] == '0' {
				zeroMask &= ^(1 << k)
			}
		}
	}

	var total int
	for _, v := range memory {
		total += v
	}

	fmt.Println(total)

	os.Exit(0)
}
