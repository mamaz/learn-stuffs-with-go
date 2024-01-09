package functions

import (
	"fmt"
	"os"
)

func CountNumOfBytes(filename string) (int, error) {
	if filename == "" {
		return 0, fmt.Errorf("no filename is specified")
	}

	contents, err := os.ReadFile(filename)
	if err != nil {
		return 0, fmt.Errorf("error on opening file %w", err)
	}
	return len(contents), nil
}
