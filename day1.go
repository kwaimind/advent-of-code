package main

import (
	"fmt"
	"kwaimind/aoc24/file"
	"sort"
)

func parseInput(input [][]int) ([]int, []int) {
	n := len(input)
	left := make([]int, 0, n)
	right := make([]int, 0, n)

	for _, row := range input {
		left = append(left, row[0])
		right = append(right, row[1])
	}

	return left, right
}

func DayOnePartOne() {
	input := file.ReadFile("input.txt")

	output := make([][]int, 0)

	left, right := parseInput(input)

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

func DayOnePartTwo() {
	input := file.ReadFile("input.txt")

	left, right := parseInput(input)

	count := 0

	for _, key := range left {
		for _, value := range right {
			match := 0
			if value == key {
				match++
			}
			count += match * key
		}
	}

	fmt.Println(count)

}
