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

func createUpdate() (func() int, func(val int)) {
	count := 0
	set := func(val int) {
		if val == match {
			count += 1
		}
	}
	get := func() int {
		return count
	}

	return get, set
}

func day1(filename string) string {
	puzzle := getPuzzle(filename)

	get, update := createUpdate()

	position := start

	for _, row := range puzzle {
		dir, amt := split(row)

		if dir == "R" {
			diff := (position + amt) % 100
			position = diff
			update(diff)
			continue
		}

		diff := (position - amt + 100) % 100
		position = diff
		update(diff)
	}

	password := get()
	fmt.Print("password is ", password)
	return "password is " + strconv.Itoa(password)
}
