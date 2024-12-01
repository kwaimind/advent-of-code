package file

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

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