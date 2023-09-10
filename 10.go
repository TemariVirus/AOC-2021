package main

import (
	"sort"
	"strings"
)

var paren_map = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}
var paren_opens = makeSetFrom([]rune{'(', '[', '{', '<'})
var paren_closes = makeSetFrom([]rune{')', ']', '}', '>'})

func solution_10_1(input string) int {
	scores := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	sum := 0
	for _, line := range strings.Split(input, "\n") {
		stack := make([]rune, 0)
		for _, char := range line {
			if paren_opens.contains(char) {
				stack = append(stack, char)
				continue
			}

			open := stack[len(stack)-1]
			if char == paren_map[open] {
				stack = stack[:len(stack)-1]
				continue
			}

			sum += scores[char]
			break
		}
	}

	return sum
}

func solution_10_2(input string) int {
	scores := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	line_scores := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		stack := make([]rune, 0)
		incomplete := true
		for _, char := range line {
			if paren_opens.contains(char) {
				stack = append(stack, char)
				continue
			}

			open := stack[len(stack)-1]
			if char == paren_map[open] {
				stack = stack[:len(stack)-1]
			} else {
				incomplete = false
				break
			}
		}

		if incomplete {
			line_score := 0
			for i := len(stack) - 1; i >= 0; i-- {
				line_score *= 5
				line_score += scores[paren_map[stack[i]]]
			}

			line_scores = append(line_scores, line_score)
		}
	}

	sort.Slice(line_scores, func(i, j int) bool {
		return line_scores[i] < line_scores[j]
	})
	return line_scores[len(line_scores)/2]
}
