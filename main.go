package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	input := string(unwrap(os.ReadFile("08.txt")))

	start := time.Now()
	fmt.Println(solution_8_2(input))
	fmt.Println("Time taken:", time.Since(start))
}

func unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
