package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parsePuzzle(filepath string) []string {
	data, err := os.ReadFile(filepath)
	must(err == nil, "failed to read file")
	return strings.Split(string(data), "\n")
}

func buildCombinations(row string) []string {
	combinations := []string{}

	for i := 0; i < len(row); i++ {
		root := row[i]
		rest := row[i+1:]
		for _, char := range rest {
			combinations = append(combinations, fmt.Sprintf("%c%c", root, char))
		}
	}

	return combinations
}

func countMaxVal(combinations []string) int {
	sort.Sort(sort.Reverse(sort.StringSlice(combinations))) // INFO: mutates the original slice
	val, err := strconv.Atoi(combinations[0])
	must(err == nil, "failed to convert to int")
	return val
}

func day3(filepath string) string {
	puzzle := parsePuzzle(filepath)

	_count := counter()

	for _, row := range puzzle {
		combinations := buildCombinations(row)
		maxVal := countMaxVal(combinations)
		_count(maxVal)
	}

	res := _count()
	return strconv.Itoa(res)
}
