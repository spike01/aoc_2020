package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	var seats []int

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

		seats = append(seats, int(id))
	}

  sort.Ints(seats)
  var mySeat int
  for i, n := range seats {
     if n == len(seats) {
        break
     }
     if seats[i] + 1 != seats[i + 1] {
       mySeat = seats[i] + 1
     }
  }
  fmt.Println("Seat ID:", mySeat)
	os.Exit(0)
}
