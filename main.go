package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	content, err := os.ReadFile("3.txt")
	if err != nil {
		panic(err)
	}
	input := string(content)

	start := time.Now()
	fmt.Println(solution_3_2(input))
	fmt.Println("Time taken:", time.Since(start))
}
