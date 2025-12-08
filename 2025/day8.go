package main

import (
	"fmt"
	"kwaimind/adventofcode2025/utils"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X, Y, Z int
}

func distance(p1, p2 Point) float64 {
	dx := float64(p2.X - p1.X)
	dy := float64(p2.Y - p1.Y)
	dz := float64(p2.Z - p1.Z)

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func unpack[T comparable](s []T, vars ...*T) {
	for i, str := range s {
		*vars[i] = str
	}
}

func day8(filepath string) string {
	puzzle := utils.GetPuzzle(filepath, utils.Line)

	points := []Point{}

	for _, row := range puzzle {

		vals := strings.Split(row, ",")
		digits := utils.Map(vals, func(n string) float64 {
			val, err := strconv.ParseFloat(n, 64)
			utils.Must(err == nil, "Error converting string to float64")
			return val
		})
		var x, y, z float64
		unpack(digits, &x, &y, &z)
		point := Point{int(x), int(y), int(z)}

		points = append(points, point)
	}

	type pointPair struct {
		p1, p2 Point
		dist   float64
	}

	pairs := []pointPair{}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := distance(points[i], points[j])
			pairs = append(pairs, pointPair{points[i], points[j], dist})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].dist < pairs[j].dist
	})

	fmt.Printf("closest pair: %+v\n", pairs[0])

	return "Not implemented yet"
}
