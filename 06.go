package main

import (
	"strconv"
	"strings"
)

func solution6Part1(input string) int {
	ages := apply(strings.Split(input, ","), func(s string) int {
		return unwrap(strconv.Atoi(s))
	})
	fishes := make([]int, 9)
	for _, age := range ages {
		fishes[age]++
	}

	count := len(ages)
	for i := 0; i < 80; i++ {
		fishes = append(fishes[1:], fishes[0])
		count += fishes[8]
		fishes[6] += fishes[8]
	}

	return count
}

func solution6Part2(input string) int {
	ages := apply(strings.Split(input, ","), func(s string) int {
		return unwrap(strconv.Atoi(s))
	})
	fishes := make([]int, 9)
	for _, age := range ages {
		fishes[age]++
	}

	count := len(ages)
	for i := 0; i < 256; i++ {
		fishes = append(fishes[1:], fishes[0])
		count += fishes[8]
		fishes[6] += fishes[8]
	}

	return count
}
