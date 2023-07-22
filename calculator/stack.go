package calculator

import (
	"fmt"
)

type Stack[T comparable] struct {
	s    []T
	size int
}

func NewStringStack() *Stack[string] {
	return &Stack[string]{
		s:    []string{},
		size: 0,
	}
}

func NewIntStack() *Stack[int] {
	return &Stack[int]{
		s:    []int{},
		size: 0,
	}
}

/*
Push() function  will push the value into the stack i.e. append

	the value to the array
*/
func (stack *Stack[T]) Push(v T) {
	stack.s = append(stack.s, v)
	stack.size++
}

// Pop function will pop the value from the top of the stack */
func (stack *Stack[T]) Pop() {
	if stack.Empty() {
		return
	}
	stack.s = stack.s[:stack.size-1]
	stack.size--
}

/*
Top returns the top value of the stack i.e. the value of the last

	index of the array
*/
func (stack *Stack[T]) Top() T {
	if stack.Empty() {
		return *new(T)
	}
	return stack.s[stack.size-1]
}

/*
	Empty function will return true if size of the array is

equal to zero or false in all other cases
*/
func (stack *Stack[T]) Empty() bool {
	return stack.size == 0
}

// Size function will return the size of the array
func (stack *Stack[T]) Size() int {
	return stack.size
}

// String converts stack to string
// String function is used for printing purposes
func (stack *Stack[T]) String() string {
	res := "["
	for i, num := range stack.s {
		res += fmt.Sprint(num)
		if i < len(stack.s)-1 {
			res += ", "
		}
	}
	res += "]"
	return res
}
