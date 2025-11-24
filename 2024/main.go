package main

import (
	"fmt"
	"os"
)

func main() {
	day := os.Args[1]

	if len(day) == 0 {
		panic("No day argument provided")
	}

	result := Solutions[day]("input.txt")

	fmt.Println(result)

}
