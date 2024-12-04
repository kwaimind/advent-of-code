package main

type Solution func(string) int

var Solutions = map[string]Solution{
	"day1part1": DayOnePartOne,
	"day1part2": DayOnePartTwo,
	"day2part1": DayTwoPartOne,
	"day3part1": DayThreePartOne,
	"day3part2": DayThreePartTwo,
}
