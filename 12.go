package main

import (
	"strings"
	"unicode"
)

type Node string
type Graph map[Node][]Node

const START_NODE Node = "start"
const END_NODE Node = "end"

func solution_12_1(input string) int {
	graph := parse_graph(strings.Split(input, "\n"))
	return count_paths(START_NODE, graph, makeSet[Node](0), true)
}

func parse_graph(edges []string) Graph {
	graph := make(Graph, 0)
	for _, edge := range edges {
		nodes := strings.Split(edge, "-")
		n1, n2 := Node(nodes[0]), Node(nodes[1])
		graph[n1] = append(graph[n1], n2)
		graph[n2] = append(graph[n2], n1)
	}

	return graph
}

func count_paths(curr Node, graph Graph, visited Set[Node], visited_extra bool) int {
	if unicode.IsLower([]rune(curr)[0]) {
		visited.add(curr)
	}

	count := 0
	for _, n := range graph[curr] {
		if n == END_NODE {
			count++
			continue
		}
		if visited.contains(n) {
			if !visited_extra && n != START_NODE {
				count += count_paths(n, graph, visited.copy(), true)
			}
			continue
		}

		count += count_paths(n, graph, visited.copy(), visited_extra)
	}

	return count
}

func solution_12_2(input string) int {
	graph := parse_graph(strings.Split(input, "\n"))
	return count_paths(START_NODE, graph, makeSet[Node](0), false)
}
