package main

import (
	"sort"
	"strings"
)

func parse_input_14(input string) (map[[2]rune]int64, map[[2]rune]rune) {
	parts := strings.Split(input, "\n\n")

	polymer := []rune(parts[0])
	pair_counts := make(map[[2]rune]int64)
	for i := 1; i < len(polymer); i++ {
		pair := [2]rune{polymer[i-1], polymer[i]}
		pair_counts[pair]++
	}

	rules := make(map[[2]rune]rune)
	for _, line := range strings.Split(parts[1], "\n") {
		runes := []rune(line)
		rules[[2]rune{runes[0], runes[1]}] = runes[6]
	}

	return pair_counts, rules
}

func process_polymer(polymer map[[2]rune]int64, rules map[[2]rune]rune) map[[2]rune]int64 {
	new_polymer := make(map[[2]rune]int64)
	for pair, count := range polymer {
		if insertion, ok := rules[pair]; ok {
			new_polymer[[2]rune{pair[0], insertion}] += count
			new_polymer[[2]rune{insertion, pair[1]}] += count
		} else {
			new_polymer[pair] += count
		}
	}
	return new_polymer
}

func count_polymer_elements(polymer map[[2]rune]int64, first rune, last rune) []int64 {
	counts := make(map[rune]int64)
	for pair, count := range polymer {
		counts[pair[0]] += count
		counts[pair[1]] += count
	}
	counts[first]++
	counts[last]++
	for pair := range counts {
		counts[pair] /= 2
	}

	sorted_counts := make([]int64, 0, len(counts))
	for _, count := range counts {
		sorted_counts = append(sorted_counts, count)
	}
	sort.Slice(sorted_counts, func(i, j int) bool {
		return sorted_counts[i] < sorted_counts[j]
	})
	return sorted_counts
}

func solution_14_1(input string) int64 {
	polymer, rules := parse_input_14(input)
	for i := 0; i < 10; i++ {
		polymer = process_polymer(polymer, rules)
	}

	polymer_str := []rune(strings.Split(input, "\n\n")[0])
	counts := count_polymer_elements(polymer, polymer_str[0], polymer_str[len(polymer_str)-1])
	return counts[len(counts)-1] - counts[0]
}

func solution_14_2(input string) int64 {
	polymer, rules := parse_input_14(input)
	for i := 0; i < 40; i++ {
		polymer = process_polymer(polymer, rules)
	}

	polymer_str := []rune(strings.Split(input, "\n\n")[0])
	counts := count_polymer_elements(polymer, polymer_str[0], polymer_str[len(polymer_str)-1])
	return counts[len(counts)-1] - counts[0]
}
