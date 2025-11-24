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
