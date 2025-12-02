package main

import (
	"os"
	"strconv"
	"strings"
)

func must(check bool, msg string) {
	if !check {
		panic(msg)
	}
}

func parseInput(filepath string) []string {
	res, err := os.ReadFile(filepath)
	must(err == nil, "failed to read file")
	data := strings.Split(string(res), ",")
	return data
}

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

func counter() (func() int, func(int)) {
	count := 0
	set := func(val int) {
		count += val
	}
	get := func() int {
		return count
	}

	return get, set
}

func day2(filepath string) string {
	puzzle := parseInput(filepath)

	getCount, setCount := counter()

	for _, row := range puzzle {
		start, end := ranges(row)

		diff := end - start
		must(diff >= 0, "end must be greater than start")

		for i := 0; i < diff+1; i++ {
			count := countDoubleDigit(i + start)
			setCount(count)
		}

	}

	sum := getCount()
	return strconv.Itoa(sum)
}
