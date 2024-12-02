package utils

// INFO: taken from https://stackoverflow.com/a/75435478/1958620
func All[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}