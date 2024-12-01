package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func OpenFile(filename string) (*os.File, func()) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	cleanup := func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}



	return file, cleanup
}

func ReadIntegerPairs(filename string) [][]int {
	file, cleanup := OpenFile(filename)
	defer cleanup()

	var result [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 2 {
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			result = append(result, []int{a, b})
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}