package main

import (
	"fmt"
	"kwaimind/adventofcode2025/utils"
)

// Really struggled to understand this one, had to look up the solution
// and work through it to understand.
// @see https://www.reddit.com/r/adventofcode/comments/1pg9w66/comment/nsqpzvq
func day7(filepath string) string {

	puzzle := utils.GetPuzzle(filepath, utils.Line)

	sum := 0

	state := make([]int, len(puzzle[0]))

	for idx, row := range puzzle {

		for i, char := range row {

			fmt.Printf("Row: %d ", idx)
			fmt.Printf("Sum: %d ", sum)
			fmt.Println(state)

			if char == 'S' {
				state[i] = 1
			} else if char == '^' && state[i] > 0 {
				sum++
				state[i-1] += state[i]
				state[i+1] += state[i]
				state[i] = 0
			}

		}

	}

	return fmt.Sprintf("Result: %d", sum)
}
