package main

import (
	"fmt"
	"kwaimind/adventofcode2025/utils"
	"os"
	"strconv"
	"strings"
)

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
	data, err := os.ReadFile(filename)
	utils.Check(err)
	puzzle := strings.Split(string(data), "\n")

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
