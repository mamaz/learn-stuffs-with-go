package nongenerics

func AddInt(values []int) int {
	var result int

	for _, val := range values {
		result += val
	}

	return result
}
