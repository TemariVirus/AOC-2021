package main

import (
	"slices"
	"strconv"
	"strings"
)

func solution_3_1(input string) int {
	words := strings.Fields(input)

	bit_counts := make([]int, len(words[0]))
	for _, word := range words {
		for i, c := range word {
			if c == '1' {
				bit_counts[i]++
			}
		}
	}

	gamma := 0
	for i, count := range bit_counts {
		if count > len(words)/2 {
			gamma |= 1 << (len(words[0]) - i - 1)
		}
	}

	elipson := gamma ^ ((1 << len(words[0])) - 1)
	return gamma * elipson
}

func solution_3_2(input string) int64 {
	words := strings.Fields(input)
	slices.Sort(words)

	slice := words[:]
	for i := 0; len(slice) > 1; i++ {
		split := find_split(slice, i)
		if split > len(slice)/2 {
			slice = slice[:split]
		} else {
			slice = slice[split:]
		}
	}
	o2, _ := strconv.ParseInt(slice[0], 2, 64)

	slice = words[:]
	for i := 0; len(slice) > 1; i++ {
		split := find_split(slice, i)
		if split > len(slice)/2 {
			slice = slice[split:]
		} else {
			slice = slice[:split]
		}
	}
	co2, _ := strconv.ParseInt(slice[0], 2, 64)

	return o2 * co2
}

// Binary search
func find_split(nums []string, index int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid][index] == '0' {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left][index] == '0' {
		return left + 1
	}
	return left
}
