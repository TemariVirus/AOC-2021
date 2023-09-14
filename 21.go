package main

import (
	"strconv"
	"strings"
)

func nextRollDeterministic(roll *int) int {
	result := *roll
	*roll = (*roll + 1) % 100
	return result + 1
}

func solution21Part1(input string) int {
	lines := strings.Split(input, "\n")

	p1_pos, p2_pos := unwrap(strconv.Atoi(lines[0][28:])), unwrap(strconv.Atoi(lines[1][28:]))
	p1_pos, p2_pos = p1_pos-1, p2_pos-1
	p1_score, p2_score := 0, 0
	roll_count := 0
	die := 0
	for {
		p1_pos = (p1_pos +
			nextRollDeterministic(&die) +
			nextRollDeterministic(&die) +
			nextRollDeterministic(&die)) % 10
		p1_score += p1_pos + 1
		roll_count += 3
		if p1_score >= 1000 {
			return p2_score * roll_count
		}

		p2_pos = (p2_pos +
			nextRollDeterministic(&die) +
			nextRollDeterministic(&die) +
			nextRollDeterministic(&die)) % 10
		p2_score += p2_pos + 1
		roll_count += 3
		if p2_score >= 1000 {
			return p1_score * roll_count
		}
	}

	return 0
}

func solution21Part2(input string) int64 {
	splits := map[int]int64{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}

	lines := strings.Split(input, "\n")
	p1_pos, p2_pos := unwrap(strconv.Atoi(lines[0][28:])), unwrap(strconv.Atoi(lines[1][28:]))

	p1_wins, p2_wins := int64(0), int64(0)
	universes := map[[5]int]int64{{p1_pos - 1, p2_pos - 1, 0, 0, 0}: 1}
	for len(universes) > 0 {
		new_uni := map[[5]int]int64{}
		for uni, count := range universes {
			p1_pos, p2_pos, p1_score, p2_score, turn := uni[0], uni[1], uni[2], uni[3], uni[4]
			if turn == 0 {
				for roll, split := range splits {
					new_p1_pos := (p1_pos + roll) % 10
					new_p1_score := p1_score + new_p1_pos + 1
					new_count := count * split
					if new_p1_score >= 21 {
						p1_wins += new_count
					} else {
						new_uni[[5]int{new_p1_pos, p2_pos, new_p1_score, p2_score, 1}] += new_count
					}
				}
			} else {
				for roll, split := range splits {
					new_p2_pos := (p2_pos + roll) % 10
					new_p2_score := p2_score + new_p2_pos + 1
					new_count := count * split
					if new_p2_score >= 21 {
						p2_wins += new_count
					} else {
						new_uni[[5]int{p1_pos, new_p2_pos, p1_score, new_p2_score, 0}] += new_count
					}
				}
			}
		}

		universes = new_uni
	}

	return max(p1_wins, p2_wins)
}
