package main

import "fmt"

/*
  NW  N  NE
   \  |  /
  W - @ - E
   /  |  \
  SW  S  SE

	 8  1  2
   \  |  /
  7 - @ - 3
   /  |  \
   6  5  4
*/

func day4(filepath string) string {

	puzzle := getPuzzle(filepath, Line)

	sum := 0

	for i := range puzzle {
		// INFO: handle the lines
		row := puzzle[i]

		// INFO: handle the chars
		for j := 0; j < len(row); j++ {

			store := []string{}

			if !isForklift(string(puzzle[i][j])) {
				continue
			}

			countAdjacent := getAdjacent(&store)

			// get all 8 directions
			countAdjacent(puzzle, i, j, -1, 0)  // above
			countAdjacent(puzzle, i, j, 1, 0)   // below
			countAdjacent(puzzle, i, j, 0, 1)   // right
			countAdjacent(puzzle, i, j, 0, -1)  // left
			countAdjacent(puzzle, i, j, -1, -1) // above-left
			countAdjacent(puzzle, i, j, -1, 1)  // above-right
			countAdjacent(puzzle, i, j, 1, -1)  // below-left
			countAdjacent(puzzle, i, j, 1, 1)   // below-right

			if len(store) < 4 {
				sum++
			}
		}

	}

	return fmt.Sprintf("%d", sum)
}

func isForklift(c string) bool {
	return c == "@"
}

func getAdjacent(store *[]string) func(puzzle []string, row, col, deltaRow, deltaCol int) {
	return func(puzzle []string, row, col, deltaRow, deltaCol int) {
		nextRow, nextCol := row+deltaRow, col+deltaCol
		if nextRow >= 0 && nextRow < len(puzzle) && nextCol >= 0 && nextCol < len(puzzle[nextRow]) {
			if val := string(puzzle[nextRow][nextCol]); isForklift(val) {
				*store = append(*store, val)
			}
		}
	}
}
