//
// Show how to create a decorator function
// to wrap things to do proper side effect
//
package main

import "fmt"

type Person struct {
	ID   string
	Name string
}

func (p *Person) PrintDefault() {
	fmt.Println("DoThings")
}

func (p *Person) PrintWithArgs(args ...interface{}) {
	for i := 0; i < len(args); i++ {
		fmt.Printf("[%v]st arg %v\n", i, args[i])
	}

}

func Log(method func()) func() {
	return func() {
		fmt.Println("iki log")
		method()
	}
}

func LogWithArgs(method func(args ...interface{}), args ...interface{}) func(retargs ...interface{}) {
	return func(arguments ...interface{}) {
		fmt.Println("iki log")
		method(args...)
	}
}

func main() {
	p := Person{
		ID:   "123",
		Name: "Mamaz",
	}

	// decorate PrintDefault, function without arguments, with Log function
	PrintDefault := Log(p.PrintDefault)
	PrintDefault()

	// decorate PrintWithArgs, function with arguments, with Log function
	PrintWithArgs := LogWithArgs(p.PrintWithArgs, p.Name, p.ID)
	PrintWithArgs(p.Name)
}
