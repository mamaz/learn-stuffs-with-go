package main

import (
	"fmt"
	"reflect"
)

type Order struct {
	ID    string
	Items []string
}

func main() {

	order := Order{ID: "12", Items: []string{"sabun", "odol"}}

	fmt.Println("iterative")
	printFields(order)

	fmt.Println("recursive")
	printFieldsRec(order)
}

func printFields(object interface{}) {
	kind := reflect.TypeOf(object).Kind()
	if kind == reflect.Struct {
		value := reflect.ValueOf(object)

		for index := 0; index < value.NumField(); index++ {
			fmt.Println("key: ", value.Type().Field(index).Name, "value: ", value.Field(index))
		}
	}
}

func printFieldsRec(object interface{}) {
	doPrintFields(object, reflect.ValueOf(object).NumField()-1)
}

func doPrintFields(object interface{}, numfield int) {
	if numfield == -1 {
		return
	}

	value := reflect.ValueOf(object)

	fmt.Println("key: ", value.Type().Field(numfield).Name, "value: ", value.Field(numfield))

	doPrintFields(object, numfield-1)
}
