package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var count int

type node struct {
	color    string
	contains map[string]*containment
}

type containment struct {
	node *node
	qty  int
}

func newNode(color string) *node {
	return &node{color: color, contains: make(map[string]*containment)}
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
		nodes[node.color] = node
	}
	nodes["o other"] = newNode("o other")

	f2, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f2.Close()
	sc = bufio.NewScanner(f2)
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
			qty, _ := strconv.Atoi(string(v[1]))
			node := nodes[color]
			node.contains[containsColor] = &containment{qty: qty, node: nodes[containsColor]}
		}
	}
	fmt.Println(nodes["shiny gold"])

	recursiveSearch(nodes["shiny gold"], []int{1})

	fmt.Println("Count:", count)
	os.Exit(0)
}

func recursiveSearch(n *node, prevLevel []int) {
	if len(n.contains) == 0 {
		return
	}

	for _, v := range n.contains {
		if v.qty != 0 {
			mult := 1
			for _, n := range prevLevel {
				mult = mult * n
			}
			count += v.qty * mult
		}
		recursiveSearch(v.node, append(prevLevel, v.qty))
	}
}
