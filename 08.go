package main

import (
	"strings"
)

type Digit = Set[rune]

func solution8Part1(input string) int {
	count := 0
	for _, line := range strings.Split(input, "\n") {
		output := strings.Split(line, " | ")[1]
		for _, digit := range strings.Fields(output) {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				count++
			}
		}
	}

	return count
}

func solution8Part2(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " | ")
		digits := apply(strings.Fields(parts[0]), func(x string) Digit {
			return makeSetFrom[rune]([]rune(x))
		})
		output_digits := apply(strings.Fields(parts[1]), func(x string) Digit {
			return makeSetFrom[rune]([]rune(x))
		})

		// Each digit appears exactly once
		found_digits := [10]Digit{}
		found_digits[8] = makeSetFrom([]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'})
		for _, d := range digits {
			if d.len() == 2 {
				found_digits[1] = d
			} else if d.len() == 3 {
				found_digits[7] = d
			} else if d.len() == 4 {
				found_digits[4] = d
			}
		}
		for _, d := range digits {
			if d.len() == 5 && d.intersect(found_digits[4]).len() == 2 {
				found_digits[2] = d
			}
			if d.len() == 5 && d.containsSet(found_digits[7]) {
				found_digits[3] = d
			}

			if d.len() == 6 && !d.containsSet(found_digits[7]) {
				found_digits[6] = d
			}
			if d.len() == 6 && d.containsSet(found_digits[4]) {
				found_digits[9] = d
			}
		}
		for _, d := range digits {
			if d.len() == 5 && d.intersect(found_digits[2]).len() == 3 {
				found_digits[5] = d
			}
		}
		for _, d := range digits {
			if d.len() == 6 && d.intersect(found_digits[5]).len() == 4 {
				found_digits[0] = d
			}
		}

		sum += 1000 * index(found_digits[:], output_digits[0])
		sum += 100 * index(found_digits[:], output_digits[1])
		sum += 10 * index(found_digits[:], output_digits[2])
		sum += 1 * index(found_digits[:], output_digits[3])
	}

	return sum
}
