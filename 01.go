package main

import (
	"strconv"
	"strings"
)

func solution_1_1(input string) int {
	lines := strings.Fields(input)
	count := 0
	prev, _ := strconv.Atoi(lines[0])
	for _, line := range lines[1:] {
		curr, _ := strconv.Atoi(line)
		if curr > prev {
			count++
		}
		prev = curr
	}
	return count
}

func solution_1_2(input string) int {
	lines := make([]int, 0)
	for _, line := range strings.Fields(input) {
		curr, _ := strconv.Atoi(line)
		lines = append(lines, curr)
	}

	count := 0
	prev := lines[0] + lines[1] + lines[2]
	for i := 3; i < len(lines); i++ {
		curr := prev + lines[i] - lines[i-3]
		if curr > prev {
			count++
		}
		prev = curr
	}
	return count
}
