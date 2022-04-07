package generics

func Add[T string | int | float64](values []T) T {
	var result T

	for _, val := range values {
		result += val
	}
	return result
}

// IndexOf returns the first
func IndexOf[T comparable](slice []T, value T) int {
	for index, val := range slice {
		if val == value {
			return index
		}
	}
	return -1
}
