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

	var oneMask int
	var xMask int
	var baseAddr int
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

			var bits []int
			// Find set bits
			bit := 1
			for {
				if bit > xMask {
					break
				}
				if xMask&bit != 0 {
					bits = append(bits, bit)
				}
				bit = bit << 1
			}

			var masks []int
			var bitMask []int

			for range bits {
				bitMask = append(bitMask, 0)
			}
			masks = append(masks, 0)

			for sum(bitMask) < len(bitMask) {
				incrementOne(bitMask)
				var sum int
				for i, v := range bitMask {
					if v == 1 {
						sum += bits[i]
					}
				}
				masks = append(masks, sum)
			}

			baseAddr = a | oneMask

			for _, i := range masks {
				memory[baseAddr&^xMask|i] = v
			}

			continue
		}

		// Set bitmasks
		oneMask = 0
		xMask = 0

		for k := 0; k < len(val); k++ {
			if val[len(val)-1-k] == 'X' {
				xMask |= 1 << k
			}
			if val[len(val)-1-k] == '1' {
				oneMask |= 1 << k
			}
		}
	}

	fmt.Println(mapSum(memory))

	os.Exit(0)
}

func sum(ints []int) int {
	var total int
	for _, v := range ints {
		total += v
	}
	return total
}

func mapSum(intMap map[int]int) int {
	var total int
	for _, v := range intMap {
		total += v
	}
	return total
}

// This "increments" left->right, because it doesn't actually matter which
// order as long as we do all permutations of the reduced bitfield
func incrementOne(bits []int) {
	carry := true
	for i, _ := range bits {
		if carry {
			if bits[i] == 0 {
				bits[i] = 1
				carry = false
				continue
			}
			bits[i] = 0
			carry = true
		}
	}
}
