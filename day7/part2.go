package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type node struct {
	color    string
	contains map[string]*contents
}

type contents struct {
	qty  int
	node *node
}

func newNode(color string) *node {
	return &node{
		color:    color,
		contains: make(map[string]*contents),
	}
}

var colorRegexp = regexp.MustCompile(`(\w+ \w+) bags contain`)
var bagsRegexp = regexp.MustCompile(`(\d) (\w+ \w+)`)

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
		color := colorRegexp.FindStringSubmatch(next)[1]

		// Graph construction - don't nuke nodes if created as part of contents
		node, ok := nodes[color]
		if !ok {
			node = newNode(color)
			nodes[node.color] = node
		}

		// regexp.MustCompile( `(\d) (\w+ \w+)`)
		// Matches are in form:  0 (left-most)        -> 1    2 (capture groups)
		//                     ((\d) (\w+ \w+  )  (\d) (\w+   w+)   )
		//                      [(n color color), (n), (color color)]
		containedBags := bagsRegexp.FindAllStringSubmatch(next, -1)

		for _, v := range containedBags {
			containedColor := v[2]

			qty, err := strconv.Atoi(v[1])
			if err != nil {
				fmt.Println("Could not convert:", v[1], err)
			}

			// If node doesn't already exist, create it in the list
			containedNode, ok := nodes[containedColor]
			if !ok {
				containedNode = newNode(containedColor)
				nodes[containedNode.color] = containedNode
			}

			node.contains[containedColor] = &contents{qty: qty, node: nodes[containedColor]}
		}
	}

	var count int
	recursiveSearch(nodes["shiny gold"], []int{1}, &count)

	fmt.Println("Count:", count)
	os.Exit(0)
}

// Simple DFS search, whatever
func recursiveSearch(n *node, prevLevel []int, count *int) {
	if len(n.contains) == 0 {
		return
	}

	for _, v := range n.contains {
		if v.qty != 0 {
			mult := 1
			for _, n := range prevLevel {
				mult = mult * n
			}
			*count += v.qty * mult
		}
		recursiveSearch(v.node, append(prevLevel, v.qty), count)
	}
}
