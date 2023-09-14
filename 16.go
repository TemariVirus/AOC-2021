package main

func parseInput16(input string) BitArray {
	result := makeBitArray(0, len(input)*4)
	for _, c := range input {
		var value int
		if c >= '0' && c <= '9' {
			value = int(c - '0')
		} else if c >= 'A' && c <= 'F' {
			value = int(c - 'A' + 10)
		} else {
			panic("invalid input")
		}
		for i := 3; i >= 0; i-- {
			result.append((value>>i)&1 == 1)
		}
	}

	return result
}

func showBits(bits BitArray) {
	println(string(apply(bits.toSlice(), func(value bool) rune {
		if value {
			return '1'
		}
		return '0'
	})))
}

func readBitsPacket(bits BitArray, i *int) (int, int) {
	type_id := bits.getRange(*i+3, *i+6)
	if type_id == 4 {
		return readBitsLiteral(bits, i)
	} else {
		return readBitsOperator(bits, i)
	}
}

func readBitsLiteral(bits BitArray, i *int) (int, int) {
	version := bits.getRange(*i, *i+3)

	result := 0
	for *i += 6; *i < bits.Length; {
		stop := !bits.get(*i)
		value := bits.getRange(*i+1, *i+5)
		*i += 5
		result <<= 4
		result |= int(value)
		if stop {
			break
		}
	}
	return result, int(version)
}

func readBitsOperator(bits BitArray, i *int) (int, int) {
	if bits.get(*i + 6) {
		return readBitsOperatorBySubpackets(bits, i)
	} else {
		return readBitsOperatorByLength(bits, i)
	}
}

func readBitsOperatorByLength(bits BitArray, i *int) (int, int) {
	version_sum := int(bits.getRange(*i, *i+3))
	type_id := int(bits.getRange(*i+3, *i+6))
	*i += 7
	length := bits.getRange(*i, *i+15)
	*i += 15

	old_i := *i
	packet_values := make([]int, 0)
	for *i < old_i+int(length) {
		value, version := readBitsPacket(bits, i)
		packet_values = append(packet_values, value)
		version_sum += version
	}

	return evalBitsOperator(type_id, packet_values), version_sum
}

func readBitsOperatorBySubpackets(bits BitArray, i *int) (int, int) {
	version_sum := int(bits.getRange(*i, *i+3))
	type_id := int(bits.getRange(*i+3, *i+6))
	*i += 7
	packets := bits.getRange(*i, *i+11)
	*i += 11

	packet_values := make([]int, 0, packets)
	for j := 0; j < int(packets); j++ {
		value, version := readBitsPacket(bits, i)
		packet_values = append(packet_values, value)
		version_sum += version
	}

	return evalBitsOperator(type_id, packet_values), version_sum
}

func evalBitsOperator(type_id int, values []int) int {
	switch type_id {
	case 0:
		return aggregate(values[1:], values[0], func(agg, value, _ int) int {
			return agg + value
		})
	case 1:
		return aggregate(values[1:], values[0], func(agg, value, _ int) int {
			return agg * value
		})
	case 2:
		return aggregate(values[1:], values[0], func(agg, value, _ int) int {
			return min(agg, value)
		})
	case 3:
		return aggregate(values[1:], values[0], func(agg, value, _ int) int {
			return max(agg, value)
		})
	case 5:
		if values[0] > values[1] {
			return 1
		}
		return 0
	case 6:
		if values[0] < values[1] {
			return 1
		}
		return 0
	case 7:
		if values[0] == values[1] {
			return 1
		}
		return 0
	}

	panic("invalid operator")
}

func solution16Part1(input string) int {
	bits := parseInput16(input)
	i := 0
	_, version_sum := readBitsPacket(bits, &i)
	return version_sum
}

func solution16Part2(input string) int {
	bits := parseInput16(input)
	i := 0
	value, _ := readBitsPacket(bits, &i)
	return value
}
