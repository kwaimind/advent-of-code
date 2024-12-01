package main

import (
	"fmt"
	"os"
)

func main() {
	day := os.Args[1]
	filename := os.Args[2]

	if len(filename) == 0 {
		panic("No filename argument provided")
	}

	if len(day) == 0 {
		panic("No day argument provided")
	}

	result := Solutions[day](filename)

	fmt.Println(result)

}
