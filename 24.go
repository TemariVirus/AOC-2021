package main

import (
	"strconv"
	"strings"
)

// First Advent of Code problem that really stumped me :feelsbadman:
// Adapted from premun's answer
// https://github.com/premun/advent-of-code/blob/main/src/2021/24/Program.cs

// This operation is being repeated for every digit
func monadDigitOp(input int, z int, parameters [2]int) int {
	x := z % 26

	if parameters[0] < 0 {
		z /= 26
	}

	// This is the if we are trying to avoid for negative first parameter
	if x != input-parameters[0] {
		z *= 26
		z += input + parameters[1]
	}

	return z
}

func findModelNumber(z int, params [][2]int, index int, agg int64, findHighest bool) (number int64, found bool) {
	if index == len(params) {
		if z == 0 {
			return agg, true
		}
		return 0, false
	}

	// Based on highest/lowest flag, either run for 1..9 or 9..1
	i := 1
	if findHighest {
		i = 9
	}
	di := 1
	if findHighest {
		di = -1
	}
	for ; (findHighest && i > 0) || (!findHighest && i <= 9); i += di {
		p := params[index]

		// When z can go down, let's make sure it does
		// When first parameters is < 0, z is divided so let's make sure we don't hit the if() that multiplies z later
		if p[0] > 0 || z%26 == i-p[0] {
			newZ := monadDigitOp(i, z, p)
			result, found := findModelNumber(newZ, params, index+1, agg*10+int64(i), findHighest)
			if found {
				return result, true
			}
		}
	}

	return 0, false
}

func parseInput24(input string) [][2]int {
	lines := strings.Split(input, "\n")
	var params = [][2]int{}
	for i := 0; i < len(lines); i += 18 {
		params = append(params, [2]int{
			unwrap(strconv.Atoi(strings.Split(lines[i+5], " ")[2])),
			unwrap(strconv.Atoi(strings.Split(lines[i+15], " ")[2])),
		})
	}
	return params
}

func solution24Part1(input string) int64 {
	number, found := findModelNumber(0, parseInput24(input), 0, 0, true)
	if !found {
		panic("No solution found")
	}
	return number
}

func solution24Part2(input string) int64 {
	number, found := findModelNumber(0, parseInput24(input), 0, 0, false)
	if !found {
		panic("No solution found")
	}
	return number
}
