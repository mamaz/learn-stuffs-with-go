package insertinterval

func insertInterval(existingIntervals []Interval, newInterval Interval) []Interval {
	return nil
}

func insertAt[T any](slice []T, newObj T, index int) []T {
	newSlice := append(slice, *new(T))

	if index >= len(slice)-1 {
		newSlice[len(slice)] = newObj
		return newSlice
	}

	// shift right
	for i := len(newSlice) - 1; i > index; i-- {
		newSlice[i] = newSlice[i-1]
	}

	newSlice[index] = newObj

	return newSlice
}
