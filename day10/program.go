package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type node struct {
	index    int
	value    int
	children []*node
}

func nodeChildren(n *node, nodes []int) []*node {
	children := []*node{}

	for i := n.index + 1; ; i++ {
		if i >= len(nodes) {
			return children
		}
		if nodes[i]-nodes[n.index] <= 3 {
			children = append(children, &node{index: i, value: nodes[i], children: []*node{}})
		} else {
			break
		}
	}

	return children
}

func tree(adapters []int) *node {
	root := &node{index: 0, value: adapters[0], children: []*node{}}
	current := []*node{root}
	next := []*node{}
	allNodes := map[int]*node{}

	for {
		for _, n := range current {
			children := nodeChildren(n, adapters)

			for _, c := range children {
				if val, exists := allNodes[c.value]; exists {
					n.children = append(n.children, val)
				} else {
					n.children = append(n.children, c)
					next = append(next, c)
					allNodes[c.value] = c
				}
			}
		}

		if len(next) == 0 {
			break
		}
		current = next
		next = nil
	}

	return root
}

func countLeaf(n *node, leaves *map[int]int) int {
	sum := 0

	if v, exists := (*leaves)[n.value]; exists {
		return v
	}
	if len(n.children) == 0 {
		return 1
	}

	for _, c := range n.children {
		sum += countLeaf(c, leaves)
	}

	(*leaves)[n.value] = sum
	return sum
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to open input file.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	adapters := []int{0}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		adapters = append(adapters, n)
	}

	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	diffs := map[int]int{}
	for i := 1; i < len(adapters); i++ {
		diffs[adapters[i]-adapters[i-1]]++
	}

	// part 1
	fmt.Println(diffs[1] * diffs[3])

	// part 2
	m := &map[int]int{}
	t := tree(adapters)

	fmt.Println(countLeaf(t, m))
	file.Close()
}
