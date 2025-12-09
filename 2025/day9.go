package main

import (
	"fmt"
	"kwaimind/adventofcode2025/utils"
	"strconv"
	"strings"
)

type PointPair struct {
	x, y int
}

func calculateArea(p1, p2 PointPair) int {
	width := abs(p2.x-p1.x) + 1
	height := abs(p2.y-p1.y) + 1
	return width * height
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Used Claude Code to walk me through this
func day9(filepath string) string {

	puzzle := utils.GetPuzzle(filepath, utils.Line)

	redTiles := []PointPair{}

	for _, tiles := range puzzle {

		vals := strings.Split(tiles, ",")

		digits := utils.Map(vals, func(n string) int {
			val, err := strconv.Atoi(n)
			utils.Must(err == nil, "Error converting string to int")
			return val
		})

		var x, y int
		utils.Unpack(digits, &x, &y)

		redTiles = append(redTiles, PointPair{x, y})

	}

	maxArea := 0

	for i, tile := range redTiles {
		for j := i + 1; j < len(redTiles); j++ {
			tileB := redTiles[j]
			area := calculateArea(tile, tileB)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return fmt.Sprintf("%d", maxArea)
}
