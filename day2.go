package main

import (
	"fmt"
	"kwaimind/aoc24/utils"
)

func isSafe(diff int) bool {
	return diff <= 3 && diff >= -3 && diff != 0
}

func isRowSorted(row []int) bool {
	var temp []int
	for i := range row {
		if i+1 < len(row) {
			diff := row[i] - row[i+1]
			temp = append(temp, diff)
		}
	}

	var allArePositive = func(n int) bool {
		return n <= 0
	}

	var allAreNegative = func(n int) bool {
		return n >= 0
	}

	return utils.All(temp, allArePositive) || utils.All(temp, allAreNegative)
}

func DayTwoPartOne(filename string) int {
	input := utils.ReadIntegerArray(filename)

	safeCount := 0

	for _, row := range input {

		safe := true

		if !isRowSorted(row) {
			safe = false
		}

		for i := range row {
			if i+1 < len(row) {
				diff := row[i] - row[i+1]

				if !isSafe(diff) {
					safe = false
					break
				}
			}
			if !safe {
				break
			}
		}
		if safe {
			safeCount++
		}
	}

	return safeCount
}

// INFO: Unsolved
func DayTwoPartTwo(filename string) int {
	input := utils.ReadIntegerArray(filename)

	safeCount := 0

	for _, row := range input {

		safe := true
		var tmp []int

		if !isRowSorted(row) {
			safe = false
		}

		for i := range row {
			if i+1 < len(row) {
				diff := row[i] - row[i+1]

				if !isSafe(diff) {
					safe = false
					tmp = append(tmp, row[i])
					break
				}
			}

			if !safe {
				break
			}
		}

		fmt.Println(tmp)

		if safe {
			safeCount++
		}
	}

	return safeCount
}
