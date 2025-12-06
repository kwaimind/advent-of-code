package utils

import (
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Counter() func(...int) int {
	count := 0
	return func(val ...int) int {
		if len(val) == 1 {
			count += val[0]
			return count
		}
		return count
	}
}

func Must(check bool, msg string) {
	if !check {
		panic(msg)
	}
}

type SeparatorType int

const (
	Line SeparatorType = iota
	Comma
)

func GetPuzzle(filepath string, sep SeparatorType) []string {
	separator := map[SeparatorType]string{
		Line:  "\n",
		Comma: ",",
	}[sep]

	data, err := os.ReadFile(filepath)
	Must(err == nil, "failed to read file")
	return strings.Split(string(data), separator)
}

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Reduce[T1, T2 any](s []T1, accumulator T2, f func(T2, T1, int) T2) T2 {
	r := accumulator
	for i, v := range s {
		r = f(r, v, i)
	}
	return r
}
