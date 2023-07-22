// Package calculator reads a string and calculate based on input
package calculator

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// 1 + 2 - 3
// +-
// 83

// (1 + 2) - 2
// 2 + (3 - 1)
func Calculate(s string) int {
	reversed := ReverseString(strings.ReplaceAll(s, " ", ""))
	stackNumbers := NewIntStack()
	stackSign := NewStringStack()

	for _, n := range reversed {
		if string(n) == "+" || string(n) == "-" {
			stackSign.Push(string(n))
		} else {
			first, err := strconv.Atoi(string(n))
			if err != nil {
				log.Fatalln("error on converting on Calculate", err)
			}
			stackNumbers.Push(first)
		}
	}

	return ProcessCalculation(stackNumbers, stackSign)
}

func ProcessCalculation(numbers *Stack[int], signs *Stack[string]) int {
	for signs.Size() > 0 {
		pickedSign := signs.Top()
		signs.Pop()

		switch pickedSign {
		case "+":
			first, second := getOperands(numbers)
			total := first + second

			numbers.Push(total)
			fmt.Println("plus", numbers.String())

		case "-":
			first, second := getOperands(numbers)
			total := first - second

			numbers.Push(total)
			fmt.Println("sub", numbers.String())
		}
	}

	result := numbers.Top()
	numbers.Pop()

	return result
}

func getOperands(numbers *Stack[int]) (int, int) {
	first := numbers.Top()
	numbers.Pop()

	second := numbers.Top()
	numbers.Pop()

	return first, second
}

func ReverseString(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))

	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteRune(rune(s[i]))
	}

	return sb.String()
}
