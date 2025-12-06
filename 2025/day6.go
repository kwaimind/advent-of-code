package main

import (
	"fmt"
	"kwaimind/adventofcode2025/utils"
	"regexp"
	"strconv"
)

func unroll(puzzleRows []string) ([][]string, [][]string) {

	data := [][]string{}

	operations := [][]string{}
	for _, row := range puzzleRows {
		digit, err := regexp.Compile(`\d+`)
		utils.Must(err == nil, "Error running regex")
		symbol, err := regexp.Compile(`\+|\*`)
		utils.Must(err == nil, "Error running regex")

		isSymbol := symbol.MatchString(row)

		if isSymbol {
			operations = append(operations, symbol.FindAllString(row, -1))
		} else {
			data = append(data, digit.FindAllString(row, -1))
		}

	}

	return data, operations
}

func flip(data [][]string) [][]string {
	flipped := [][]string{}

	for i := 0; i < len(data[0]); i++ {
		newRow := []string{}
		for j := range data {
			newRow = append(newRow, data[j][i])
		}
		flipped = append(flipped, newRow)
	}

	return flipped
}

func getLatestOperation(operations []string, idx int) string {
	if idx >= len(operations) {
		return operations[len(operations)-1]
	}
	return operations[idx]
}

func calculate(data [][]string, operations [][]string) int {

	sum := 0

	for idx, v := range data {

		digits := utils.Map(v, func(n string) int {
			val, err := strconv.Atoi(n)
			utils.Must(err == nil, "Error converting string to int")
			return val
		})

		sumOfNums := utils.Reduce(digits, 0, func(acc, n, i int) int {
			op := getLatestOperation(operations[idx], i)
			if op == "+" {
				acc += n
			} else {
				if acc == 0 {
					acc = 1
				}
				acc *= n
			}
			return acc
		})

		sum += sumOfNums

	}

	return sum
}

func day6(filepath string) string {
	rows := utils.GetPuzzle(filepath, utils.Line)
	unrolledRows, unrolledOperations := unroll(rows)
	flippedRows := flip(unrolledRows)
	flippedOperations := flip(unrolledOperations)
	sum := calculate(flippedRows, flippedOperations)
	return fmt.Sprintf("%d", sum)
}
