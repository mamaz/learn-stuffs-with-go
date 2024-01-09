package main

import (
	"flag"
	"fmt"
	"my-wc/functions"
	"os"
)

func main() {
	filename, isC, isL, isW := HandleInputs()

	defaultValues := !isC && !isL && !isW

	if defaultValues {
		showAllResults(filename)
	}

	if isC {
		numOfBytes, err := functions.CountNumOfBytes(filename)
		if err != nil {
			fmt.Printf("error on counting num of bytes: %v", err)
			return
		}

		fmt.Printf("%v %v\n", numOfBytes, filename)
	}

	if isL {
		numOfLines, err := functions.CountNumberOfLines(filename)
		if err != nil {
			fmt.Printf("error on counting num of bytes: %v", err)
			return
		}

		fmt.Printf("%v %v", numOfLines, filename)
	}

	if isW {
		numOfWords, err := functions.CountNumOfWords(filename)
		if err != nil {
			fmt.Printf("error on counting num of bytes: %v", err)
			return
		}

		fmt.Printf("%v %v", numOfWords, filename)
	}
}

func HandleInputs() (filename string, isC bool, isL bool, isW bool) {
	if len(os.Args) == 1 {
		fmt.Println("not enough arguments, my-wc [options] [filename]")
		return
	}
	lenArg := len(os.Args)
	filename = os.Args[lenArg-1] // filename will be the last argument

	isC = *flag.Bool("c", false, "count number of bytes in a file")
	isL = *flag.Bool("l", false, "count number of lines in a file")
	isW = *flag.Bool("w", false, "count number of words in a file")

	flag.Parse()

	return
}

func showAllResults(filename string) {
	numOfBytes, err := functions.CountNumOfBytes(filename)
	if err != nil {
		fmt.Printf("error on counting num of bytes: %v", err)
		return
	}

	numOfLines, err := functions.CountNumberOfLines(filename)
	if err != nil {
		fmt.Printf("error on counting num of bytes: %v", err)
		return
	}

	numOfWords, err := functions.CountNumOfWords(filename)
	if err != nil {
		fmt.Printf("error on getting the result: %v", err)
		return
	}

	fmt.Printf("%v %v %v %v", numOfBytes, numOfLines, numOfWords, filename)
}
