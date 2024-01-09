package functions

import (
	"bufio"
	"fmt"
	"os"
)

func CountNumberOfLines(filename string) (int, error) {
	if filename == "" {
		return 0, fmt.Errorf("no filename is specified")
	}

	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error on opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// counting num of lines
	numOfLines := 0
	for scanner.Scan() {
		numOfLines++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error on scanning: %w", err)
	}

	return numOfLines, nil
}
