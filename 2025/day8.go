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

// Another day where I didn't fully understand the problem.
// Made with love from Claude Code.
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

	// Union-Find data structure
	parent := make(map[Point]Point)
	size := make(map[Point]int)

	// Initialize each point as its own parent
	for _, p := range points {
		parent[p] = p
		size[p] = 1
	}

	// Find root with path compression
	var find func(Point) Point
	find = func(p Point) Point {
		if parent[p] != p {
			parent[p] = find(parent[p])
		}
		return parent[p]
	}

	// Union two sets
	union := func(p1, p2 Point) {
		root1 := find(p1)
		root2 := find(p2)

		if root1 == root2 {
			return // already in same circuit
		}

		// Union by size
		if size[root1] < size[root2] {
			parent[root1] = root2
			size[root2] += size[root1]
		} else {
			parent[root2] = root1
			size[root1] += size[root2]
		}
	}

	// Connect 1000 closest pairs
	for i := 0; i < 1000 && i < len(pairs); i++ {
		union(pairs[i].p1, pairs[i].p2)
	}

	// Count circuit sizes
	circuitSizes := make(map[Point]int)
	for _, p := range points {
		root := find(p)
		circuitSizes[root] = size[root]
	}

	// Get all sizes and sort descending
	sizes := []int{}
	for _, s := range circuitSizes {
		sizes = append(sizes, s)
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	// Multiply three largest (or fewer if less than 3 circuits)
	result := 1
	for i := 0; i < 3 && i < len(sizes); i++ {
		result *= sizes[i]
	}

	return fmt.Sprintf("%d", result)
}
