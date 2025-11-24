package main

import (
	"kwaimind/aoc24/utils"
	"sort"
)

func DayOnePartOne(filename string) int {
	input := utils.ReadIntegerPairs(filename)

	output := make([][]int, 0)

	left, right := utils.UnzipPairs(input)

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

	return distance
}

func DayOnePartTwo(filename string) int {
	input := utils.ReadIntegerPairs(filename)

	left, right := utils.UnzipPairs(input)

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

	return count
}
