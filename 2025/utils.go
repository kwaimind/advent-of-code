package main

func counter() func(...int) int {
	count := 0
	return func(val ...int) int {
		if len(val) == 1 {
			count += val[0]
			return count
		}
		return count
	}
}
