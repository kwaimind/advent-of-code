package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getPuzzle(filepath string) []string {
	puzzle, err := os.ReadFile(filepath)
	check(err)
	return strings.Split(string(puzzle), "\n")
}

func split(s string) (string, int) {
	res := strings.SplitAfterN(s, "", 2)
	num, _ := strconv.Atoi(res[1])
	return res[0], num
}

const (
	start = 50
	match = 0
)

func update(val int, target *int) {
	if val == match {
		*target += 1
	}
}

func main() {
	filename := os.Args[1]
	if filename == "" {
		filename = "input.txt"
	}
	puzzle := getPuzzle(filename)

	position := start
	password := 0

	for _, row := range puzzle {
		dir, amt := split(row)

		if dir == "R" {
			diff := (position + amt) % 100
			position = diff
			update(diff, &password)
			continue
		}

		diff := (position - amt + 100) % 100
		position = diff
		update(diff, &password)
	}

	fmt.Print("password is ", password)
}
