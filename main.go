package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	content, err := os.ReadFile("4.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)

	start := time.Now()
	fmt.Println(solution_4_2(input))
	fmt.Println("Time taken:", time.Since(start))
}
