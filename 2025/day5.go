package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isLineBreak(line string) bool {
	return line == ""
}

func isRange(line string) bool {
	return strings.Contains(line, "-")
}

func day5(filepath string) string {

	puzzle := getPuzzle(filepath, Line)

	count := 0

	ranges := []string{}
	keys := []string{}

	for _, v := range puzzle {
		if isLineBreak(v) {
			continue
		}
		if isRange(v) {
			ranges = append(ranges, v)
		} else {
			keys = append(keys, v)
		}
	}

	for _, key := range keys {
		val, _ := strconv.Atoi(key)
		fmt.Printf("val %d\n", val)
		for _, rangeRow := range ranges {
			res := strings.Split(rangeRow, "-")
			min, _ := strconv.Atoi(res[0])
			max, _ := strconv.Atoi(res[1])

			if val >= min && val <= max {
				count++
				break
			}

		}

	}

	return fmt.Sprintf("%d", count)

}
