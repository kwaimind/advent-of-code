package main

import (
	"testing"
)

func TestDayOnePartOne(t *testing.T) {

	filename := "test.txt"
	expected := 11

	result := DayOnePartOne(filename)

	if result != expected {
		t.Errorf("DayOnePartOne(%s) = %d; want %d", filename, result, expected)
	}
}

func TestDayOnePartTwo(t *testing.T) {

	filename := "test.txt"
	expected := 31

	result := DayOnePartTwo(filename)

	if result != expected {
		t.Errorf("DayOnePartTwo(%s) = %d; want %d", filename, result, expected)
	}
}
