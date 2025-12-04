package main

import (
	"os"
	"strconv"
	"strings"
)

func ranges(data string) (int, int) {
	res := strings.Split(data, "-")
	must(len(res) == 2, "invalid range format")
	start, _ := strconv.Atoi(res[0])
	end, _ := strconv.Atoi(res[1])
	return start, end
}

func countDoubleDigit(num int) int {
	strNum := strconv.Itoa(num)
	half := len(strNum) / 2
	left := strNum[:half]
	right := strNum[half:]
	if left == right {
		return num
	}
	return 0
}

func day2(filepath string) string {
	res, err := os.ReadFile(filepath)
	must(err == nil, "failed to read file")
	puzzle := strings.Split(string(res), ",")

	_count := counter()

	for _, row := range puzzle {
		start, end := ranges(row)

		diff := end - start
		must(diff >= 0, "end must be greater than start")

		for i := 0; i < diff+1; i++ {
			count := countDoubleDigit(i + start)
			_count(count)
		}

	}

	sum := _count()
	return strconv.Itoa(sum)
}
