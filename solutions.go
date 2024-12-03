package main

type Solution func(string) int

var Solutions = map[string]Solution{
	"day1part1": DayOnePartOne,
	"day1part2": DayOnePartTwo,
	"day2part1":  DayTwoPartOne,
	"day2part2": DayTwoPartTwo,
	"day3part1":  DayThreePartOne,
}
