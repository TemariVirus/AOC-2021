package main

import (
	"strconv"
	"strings"
)

type SFNumber struct {
	IsRegular   bool
	Value       int
	Left, Right *SFNumber
}

func parseSFNumber(input string) *SFNumber {
	if input[0] != '[' {
		return &SFNumber{IsRegular: true, Value: unwrap(strconv.Atoi(input))}
	}

	input = input[1 : len(input)-1]
	var left_str, right_str string
	depth := 0
	for i, c := range input {
		if depth == 0 && c == ',' {
			left_str = input[:i]
			right_str = input[i+1:]
			break
		}
		if c == '[' {
			depth++
		}
		if c == ']' {
			depth--
		}
	}

	return &SFNumber{
		IsRegular: false,
		Left:      parseSFNumber(left_str),
		Right:     parseSFNumber(right_str),
	}
}

func (n *SFNumber) magnitude() int {
	if n.IsRegular {
		return n.Value
	}
	return 3*n.Left.magnitude() + 2*n.Right.magnitude()
}

func (n *SFNumber) add(other *SFNumber) *SFNumber {
	sum := &SFNumber{
		IsRegular: false,
		Left:      n.copy(),
		Right:     other.copy(),
	}
	sum.reduce()
	return sum
}

func (n *SFNumber) reduce() {
	for {
		if exploded, _, _ := n.tryExplode(0); exploded {
			continue
		}
		if n.trySplit() {
			continue
		}
		break
	}
}

func (n *SFNumber) tryExplode(depth int) (exploded bool, left, right int) {
	if n.IsRegular {
		return false, 0, 0
	}

	if depth < 4 {
		if exploded, left, right := n.Left.tryExplode(depth + 1); exploded {
			if right == 0 {
				return true, left, right
			}
			n.Right.addFirst(right)
			return true, left, 0
		}
		if exploded, left, right := n.Right.tryExplode(depth + 1); exploded {
			if left == 0 {
				return true, left, right
			}
			n.Left.addLast(left)
			return true, 0, right
		}
		return false, 0, 0
	}

	left, right = n.Left.Value, n.Right.Value
	*n = SFNumber{
		IsRegular: true,
		Value:     0,
	}
	return true, left, right
}

func (n *SFNumber) trySplit() bool {
	if n.IsRegular {
		if n.Value < 10 {
			return false
		}
		left := &SFNumber{
			IsRegular: true,
			Value:     n.Value / 2,
		}
		right := &SFNumber{
			IsRegular: true,
			Value:     n.Value - left.Value,
		}
		*n = SFNumber{
			IsRegular: false,
			Left:      left,
			Right:     right,
		}
		return true
	}

	if n.Left.trySplit() {
		return true
	}
	if n.Right.trySplit() {
		return true
	}
	return false
}

func (n *SFNumber) addFirst(number int) {
	for !n.IsRegular {
		n = n.Left
	}
	n.Value += number
}

func (n *SFNumber) addLast(number int) {
	for !n.IsRegular {
		n = n.Right
	}
	n.Value += number
}

func (n *SFNumber) copy() *SFNumber {
	if n.IsRegular {
		return &SFNumber{
			IsRegular: true,
			Value:     n.Value,
		}
	}
	return &SFNumber{
		IsRegular: false,
		Left:      n.Left.copy(),
		Right:     n.Right.copy(),
	}
}

func (n *SFNumber) toString() string {
	if n.IsRegular {
		return strconv.Itoa(n.Value)
	}
	return "[" + n.Left.toString() + "," + n.Right.toString() + "]"
}

func solution_18_1(input string) int {
	numbers := apply(strings.Split(input, "\n"), func(s string) *SFNumber {
		return parseSFNumber(s)
	})
	sum := aggregate(numbers[1:], numbers[0], func(agg, n *SFNumber, _ int) *SFNumber {
		return agg.add(n)
	})

	return sum.magnitude()
}

func solution_18_2(input string) int {
	numbers := apply(strings.Split(input, "\n"), func(s string) *SFNumber {
		return parseSFNumber(s)
	})

	max_mag := 0
	for i, a := range numbers {
		for j, b := range numbers {
			if i == j {
				continue
			}
			max_mag = max(max_mag, a.add(b).magnitude())
		}
	}

	return max_mag
}
