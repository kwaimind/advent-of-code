package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
)

type Solution func(filename string) string

var solutions = map[string]Solution{
	"day1": day1,
	"day2": day2,
	"day3": day3,
	"day4": day4,
	"day5": day5,
	"day6": day6,
}

var args struct {
	Filename string `arg:"-f,--file" default:"test.txt"`
	Day      int    `arg:"-d,--day,required"`
}

func main() {

	arg.MustParse(&args)
	fmt.Printf("Running day %d with %s\n", args.Day, args.Filename)

	key := "day" + fmt.Sprint(args.Day)

	if _, exists := solutions[key]; !exists {
		panic(fmt.Sprintf("No solution for %s", key))
	}

	solution := solutions[key]
	result := solution(args.Filename)
	fmt.Println(result)

}
