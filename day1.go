package main

import (
	"fmt"
	"kwaimind/aoc24/file"
	"sort"
)

func DayOne() {
	input := file.ReadFile("input.txt")

	n := len(input)
	left := make([]int, 0, n)
	right := make([]int, 0, n)
	output := make([][]int, 0, n)

	for _, row := range input {
		left = append(left, row[0])
		right = append(right, row[1])
	}

	sort.Ints(left)
	sort.Ints(right)

	for i := 0; i < len(left); i++ {
		output = append(output, []int{left[i], right[i]})
	}

	distance := 0

	for i := 0; i < len(output); i++ {
		copy := output[i]
		sort.Ints(copy)
		distance += copy[1] - copy[0]
	}

	fmt.Println(distance)
}
