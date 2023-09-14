package main

import (
	"strings"
)

func parseInput20(input string) (BitArray, []BitArray) {
	parts := strings.Split(input, "\n\n")

	algo := makeBitArray(0, 512)
	for _, c := range parts[0] {
		algo.append(c == '#')
	}

	image := []BitArray{}
	for _, line := range strings.Split(parts[1], "\n") {
		row := makeBitArray(0, 0)
		for _, c := range line {
			row.append(c == '#')
		}
		image = append(image, row)
	}

	return algo, image
}

func addImageBoarder(image []BitArray, width int, is_bg_white bool) []BitArray {
	new_image := []BitArray{}
	for i := 0; i < width; i++ {
		new_row := makeBitArray(image[0].Length+2*width, -1)
		if is_bg_white {
			for j := 0; j < new_row.Length; j++ {
				new_row.set(j)
			}
		}
		new_image = append(new_image, new_row)
	}
	for _, row := range image {
		new_row := makeBitArray(0, image[0].Length+2*width)
		for i := 0; i < width; i++ {
			new_row.append(is_bg_white)
		}
		for i := 0; i < row.Length; i++ {
			new_row.append(row.get(i))
		}
		for i := 0; i < width; i++ {
			new_row.append(is_bg_white)
		}
		new_image = append(new_image, new_row)
	}
	for i := 0; i < width; i++ {
		new_row := makeBitArray(image[0].Length+2*width, -1)
		if is_bg_white {
			for j := 0; j < new_row.Length; j++ {
				new_row.set(j)
			}
		}
		new_image = append(new_image, new_row)
	}

	return new_image
}

func enhanceImage(image []BitArray, algo BitArray, is_bg_white bool) []BitArray {
	new_image := []BitArray{}

	// Add top boarder
	new_image = append(new_image, makeBitArray(0, image[0].Length+2))
	for j := 0; j < image[0].Length+2; j++ {
		new_image[0].append(is_bg_white)
	}
	new_image = append(new_image, new_image[0])

	// Enhance image
	for i := 1; i < len(image)-1; i++ {
		row := makeBitArray(image[0].Length+2, -1)
		for j := 1; j < image[0].Length-1; j++ {
			index := int(image[i-1].getRange(j-1, j+2)<<6 +
				image[i].getRange(j-1, j+2)<<3 +
				image[i+1].getRange(j-1, j+2),
			)
			if algo.get(index) {
				row.set(j + 1)
			}
		}
		if is_bg_white {
			row.set(0)
			row.set(1)
			row.set(row.Length - 2)
			row.set(row.Length - 1)
		}
		new_image = append(new_image, row)
	}

	// Add bottom boarder
	new_image = append(new_image, new_image[0])
	new_image = append(new_image, new_image[0])

	return new_image
}

func solution20Part1(input string) int {
	algo, image := parseInput20(input)

	is_bg_white := false
	image = addImageBoarder(image, 2, is_bg_white)
	for i := 0; i < 2; i++ {
		is_bg_white = (!is_bg_white && algo.get(0)) || (is_bg_white && algo.get(511))
		image = enhanceImage(image, algo, is_bg_white)
	}

	if is_bg_white {
		panic("Infinity")
	}
	return aggregate(image, 0, func(agg int, value BitArray, _ int) int {
		return agg + value.setCount()
	})
}

func solution20Part2(input string) int {
	algo, image := parseInput20(input)

	is_bg_white := false
	image = addImageBoarder(image, 2, is_bg_white)
	for i := 0; i < 50; i++ {
		is_bg_white = (!is_bg_white && algo.get(0)) || (is_bg_white && algo.get(511))
		image = enhanceImage(image, algo, is_bg_white)
	}

	if is_bg_white {
		panic("Infinity")
	}
	return aggregate(image, 0, func(agg int, value BitArray, _ int) int {
		return agg + value.setCount()
	})
}
