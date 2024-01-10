package main

import "fmt"

func updateValueWithReference(x *int) {
	*x = *x + 1
}

func main() {
	num := 5
	updateValueWithReference(&num)
	fmt.Println(num) // Output: 6 (modified)
}
