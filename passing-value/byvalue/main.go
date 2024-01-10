package main

import "fmt"

func updateValue(x int) {
	x = x + 1
}

func main() {
	num := 5
	updateValue(num)
	fmt.Println(num) // Output: 5 (unchanged)
}
