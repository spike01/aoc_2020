package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

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

		// Graph construction - don't nuke nodes if created as part of containedBy
		_, ok := nodes[color]
		if !ok {
			node := newNode(color)
			nodes[node.color] = node
		}

		// Matches are in form:     0       1  2
		//                      [n col col, n, col col]
		containedBags := bagsRegexp.FindAllStringSubmatch(next, -1)

		for _, v := range containedBags {
			containingColor := v[2]

			// If node doesn't already exist, create it in the list
			node, ok := nodes[containingColor]
			if !ok {
				node = newNode(containingColor)
				nodes[node.color] = node
			}
			node.containedBy[color] = nodes[color]
		}
	}

	set := make(map[string]struct{})
	recursiveSearch(nodes["shiny gold"], set)

	fmt.Println("Set len:", len(set))
	os.Exit(0)
}

// Simple DFS search, whatever
func recursiveSearch(n *node, set map[string]struct{}) {
	if len(n.containedBy) == 0 {
		return
	}
	for k, v := range n.containedBy {
		set[k] = struct{}{}
		recursiveSearch(v, set)
	}
}
