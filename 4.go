package main

import (
	"slices"
	"strconv"
	"strings"
)

func parse_input(input string) ([]int, [][25]int) {
	parts := strings.Split(input, "\n\n")

	numbers := make([]int, 0)
	for _, n := range strings.Split(parts[0], ",") {
		num, _ := strconv.Atoi(n)
		numbers = append(numbers, num)
	}

	boards := make([][25]int, len(parts)-1)
	for i, board := range parts[1:] {
		for j, num := range strings.Fields(board) {
			boards[i][j], _ = strconv.Atoi(num)
		}
	}

	return numbers, boards
}

func solution_4_1(input string) int {
	numbers, boards := parse_input(input)

	var win_board [25]int
	var draws int
	rows := make([][5]int, len(boards))
	cols := make([][5]int, len(boards))
	for i, drawn := range numbers {
		for j, board := range boards {
			for k, num := range board {
				if num == drawn {
					rows[j][k/5]++
					cols[j][k%5]++
					if rows[j][k/5] == 5 || cols[j][k%5] == 5 {
						win_board = board
						draws = i + 1
						goto winner_found
					}
					break
				}

			}
		}
	}

winner_found:
	sum := 0
	for _, num := range win_board {
		if !slices.Contains(numbers[:draws], num) {
			sum += num
		}
	}
	return sum * numbers[draws-1]
}

func solution_4_2(input string) int {
	numbers, boards := parse_input(input)

	var last_board [25]int
	var draws int
	rows := make([][5]int, len(boards))
	cols := make([][5]int, len(boards))
	won := make([]bool, len(boards))
	won_count := 0
	for i, drawn := range numbers {
		for j, board := range boards {
			for k, num := range board {
				if num == drawn {
					rows[j][k/5]++
					cols[j][k%5]++
					if rows[j][k/5] == 5 || cols[j][k%5] == 5 {
						if !won[j] {
							won_count++
						}
						won[j] = true
					}
				}
			}
			if won_count == len(boards) {
				last_board = board
				draws = i + 1
				goto last_found
			}
		}
	}

last_found:
	sum := 0
	for _, num := range last_board {
		if !slices.Contains(numbers[:draws], num) {
			sum += num
		}
	}
	return sum * numbers[draws-1]
}
