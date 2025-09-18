package csv

import (
	"fmt"
	"os"
)

func OpenOrCreate(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("error opening or creating file: %w", err)
	}
	return file, nil
}
