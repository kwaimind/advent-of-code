package main

import (
	"testing"
)

func TestDayTwoPartOne(t *testing.T) {

	filename := "inputs/day2.test.txt"
	expected := 2

	result := DayTwoPartOne(filename)

	if result != expected {
		t.Errorf("DayTwoPartOne(%s) = %d; want %d", filename, result, expected)
	}
}

func TestDayTwoPartTwo(t *testing.T) {
	t.Skip("Skipping TestDayTwoPartTwo as it is unsolved")

	filename := "inputs/day2.test.txt"
	expected := 4

	result := DayTwoPartTwo(filename)

	if result != expected {
		t.Errorf("DayTwoPartTwo(%s) = %d; want %d", filename, result, expected)
	}
}
