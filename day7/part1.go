package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var count int

type node struct {
	color       string
	containedBy map[string]*node
}

func newNode(color string) *node {
	return &node{
		color:       color,
		containedBy: make(map[string]*node),
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	nodes := make(map[string]*node)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		next := sc.Text()
		split := strings.Split(next, "bags")
		split = strings.Split(strings.Join(split, " "), "contain")
		color := split[0]
		color = strings.TrimSpace(color)
		node := newNode(color)
		nodes[color] = node
	}
	nodes["o other"] = newNode("o other")

	f.Seek(0, io.SeekStart)

	sc = bufio.NewScanner(f)
	for sc.Scan() {
		next := sc.Text()
		split := strings.Split(next, "bags")
		split = strings.Split(strings.Join(split, " "), "contain")
		color := split[0]
		color = strings.TrimSpace(color)
		contains := strings.ReplaceAll(split[1], ".", "")
		containsSplit := strings.Split(contains, ",")

		for _, v := range containsSplit {
			containsColor := v[2:]
			containsColor = strings.ReplaceAll(containsColor, "bags", "")
			containsColor = strings.ReplaceAll(containsColor, "bag", "")
			containsColor = strings.TrimSpace(containsColor)
			node := nodes[containsColor]
			node.containedBy[color] = nodes[color]
		}
	}

	set := make(map[string]struct{})
	recursiveSearch(nodes["shiny gold"], set)

	fmt.Println("Set len:", len(set))
	os.Exit(0)
}

func recursiveSearch(n *node, set map[string]struct{}) {
	if len(n.containedBy) == 0 {
		return
	}
	for k, v := range n.containedBy {
		set[k] = struct{}{}
		recursiveSearch(v, set)
	}
}
