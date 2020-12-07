package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
  defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		next := sc.Text()
		fmt.Println(next)
	}

	os.Exit(0)
}
