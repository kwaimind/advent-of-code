package main

import (
	"testing"
)

func TestDayThreePartOne(t *testing.T) {

	filename := "inputs/day3.test.txt"
	expected := 161

	result := DayThreePartOne(filename)

	if result != expected {
		t.Errorf("DayThreePartOne(%s) = %d; want %d", filename, result, expected)
	}
}

func TestDayThreePartTwo(t *testing.T) {

	filename := "inputs/day3.test.txt"
	expected := 48

	result := DayThreePartTwo(filename)

	if result != expected {
		t.Errorf("DayThreePartTwo(%s) = %d; want %d", filename, result, expected)
	}
}
