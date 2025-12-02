package main

import (
	"fmt"
	"os"
)

var Solutions = map[string]func(filename string) string{
	"day1": day1,
	"day2": day2,
}

func main() {

	filename := os.Args[1]
	if filename == "" {
		filename = "input.txt"
	}

	day := os.Args[2]

	if len(day) == 0 {
		panic("No day argument provided")
	}

	key := "day" + day

	if _, exists := Solutions[key]; !exists {
		panic(fmt.Sprintf("No solution for %s", key))
	}

	solution := Solutions[key]
	result := solution(filename)
	fmt.Println(result)

}
