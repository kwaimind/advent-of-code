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
