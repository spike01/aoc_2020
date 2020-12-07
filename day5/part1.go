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

	var highest int64

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		next := sc.Text()
		row := next[0:7]
		row = strings.ReplaceAll(row, "F", "0")
		row = strings.ReplaceAll(row, "B", "1")

		column := next[7:10]
		column = strings.ReplaceAll(column, "L", "0")
		column = strings.ReplaceAll(column, "R", "1")

		rowNum, _ := strconv.ParseInt(row, 2, 32)
		columnNum, _ := strconv.ParseInt(column, 2, 32)
		id := rowNum*8 + columnNum

		if id > highest {
			highest = id
		}
	}

	fmt.Println("Highest:", highest)
	os.Exit(0)
}
