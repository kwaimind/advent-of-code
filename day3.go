package main

import (
	"kwaimind/aoc24/utils"
	"regexp"
	"strconv"
)

func findValidFunctionCalls(line string) []string {
	compRegEx := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	return compRegEx.FindAllString(line, -1)
}

func findValidFunctionCallsWithStoppers(line string) []string {
	compRegEx := regexp.MustCompile(`mul\(\d{1,3}\,\d{1,3}\)|do\(\)|don't\(\)`)
	return compRegEx.FindAllString(line, -1)
}

func parseFunctionValues(functionCall string) []int {
	compRegEx := regexp.MustCompile(`\d{1,3}`)
	matches := compRegEx.FindAllString(functionCall, -1)

	if len(matches) != 2 {
		return []int{0, 0}
	}

	result := []int{}

	for _, match := range matches {
		first, _ := strconv.Atoi(match)
		result = append(result, first)
	}

	return result
}

func DayThreePartOne(filename string) int {
	input := utils.ReadFileAsString(filename)

	result := 0

	functionCalls := findValidFunctionCalls(input)
	for _, functionCall := range functionCalls {
		values := parseFunctionValues(functionCall)
		result += values[0] * values[1]
	}

	return result
}

func DayThreePartTwo(filename string) int {
	input := utils.ReadFileAsString(filename)

	result := 0

	functionCalls := findValidFunctionCallsWithStoppers(input)

	enabled := true

	for _, functionCall := range functionCalls {
		if functionCall == "don't()" {
			enabled = false
			continue
		}
		if functionCall == "do()" {
			enabled = true
			continue
		}
		if enabled {
			values := parseFunctionValues(functionCall)
			result += values[0] * values[1]
		}
	}

	return result
}
