package main

import (
	"strconv"
	"strings"
)

func solution2Part1(input string) int {
	words := strings.Fields(input)
	x, z := 0, 0
	for i := 0; i < len(words); i += 2 {
		num, _ := strconv.Atoi(words[i+1])
		switch words[i] {
		case "forward":
			x += num
		case "down":
			z += num
		case "up":
			z -= num
		}
	}
	return x * z
}

func solution2Part2(input string) int {
	words := strings.Fields(input)
	x, depth, aim := 0, 0, 0
	for i := 0; i < len(words); i += 2 {
		num, _ := strconv.Atoi(words[i+1])
		switch words[i] {
		case "forward":
			x += num
			depth += aim * num
		case "down":
			aim += num
		case "up":
			aim -= num
		}
	}
	return x * depth
}
