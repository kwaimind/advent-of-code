package main

import (
	"os"
	"strings"
)

func counter() func(...int) int {
	count := 0
	return func(val ...int) int {
		if len(val) == 1 {
			count += val[0]
			return count
		}
		return count
	}
}

func must(check bool, msg string) {
	if !check {
		panic(msg)
	}
}

type SeparatorType int

const (
	Line SeparatorType = iota
	Comma
)

func getPuzzle(filepath string, sep SeparatorType) []string {
	separator := map[SeparatorType]string{
		Line:  "\n",
		Comma: ",",
	}[sep]

	data, err := os.ReadFile(filepath)
	must(err == nil, "failed to read file")
	return strings.Split(string(data), separator)
}
