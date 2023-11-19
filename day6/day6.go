package day6

import (
	"aoc-2019/dag"
	"aoc-2019/set"
	"fmt"
	"math"
	"strings"
)

func Solution(input []string) {
	edges := make([][]string, len(input))

	for i, line := range input {
		edges[i] = strings.Split(line, ")")
	}

	orbitDag := dag.FromEdgeList(edges)

	orbits := dayOne(orbitDag)
	skips := dayTwo(orbitDag)

	fmt.Printf("Day One: %d\n", orbits)
	fmt.Printf("Day Two: %d\n", skips)
}

func dayTwo(orbitDag *dag.DAG[string]) int {
	commonAncestors := set.NewSet[*dag.Node[string]]()

	orbitDag.BFS(func(node *dag.Node[string]) {
		var visited []*dag.Node[string]
		var santa bool
		var you bool

		node.BFS(func(n *dag.Node[string]) {
			if n.Value == "YOU" {
				you = true
			}

			if n.Value == "SAN" {
				santa = true
			}

			if you && santa {
				commonAncestors.Add(node)
			}
		}, visited)
	})

	minSkips := math.MaxInt
	// for every common ancestor, figure out how long it takes to get to either node
	for ca := range commonAncestors {
		var visited []*dag.Node[string]
		var distanceYou int
		var distanceSanta int

		ca.BFS(func(n *dag.Node[string]) {
			if n.Value == "YOU" {
				curr := n
				for curr.Parent.Value != ca.Value {
					curr = curr.Parent
					distanceYou++
				}
			}

			if n.Value == "SAN" {
				curr := n
				for curr.Parent.Value != ca.Value {
					curr = curr.Parent
					distanceSanta++
				}
			}
		}, visited)

		skips := distanceSanta + distanceYou
		if skips < minSkips {
			minSkips = skips
		}
	}

	return minSkips
}

func dayOne(orbitDag *dag.DAG[string]) int {
	var orbits int

	orbitDag.BFS(func(n *dag.Node[string]) {
		current := n
		for current.Parent != nil {
			current = current.Parent
			orbits++
		}
	})

	return orbits
}
