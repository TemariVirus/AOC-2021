package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	input := string(unwrap(os.ReadFile("23.txt")))

	start := time.Now()
	fmt.Println(solution23Part2(input))
	fmt.Println("Time taken:", time.Since(start))
}

func unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// https://stackoverflow.com/a/109025
func popcount(x uint64) int {
	x = x - ((x >> 1) & 0x5555555555555555)                        // add pairs of bits
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333) // quads
	x = (x + (x >> 4)) & 0x0F0F0F0F0F0F0F0F                        // groups of 8
	return int((x * 0x0101010101010101) >> 56)                     // horizontal sum of bytes
}
