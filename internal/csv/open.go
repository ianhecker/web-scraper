package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func OpenOrCreate(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("error opening/creating file: %w", err)
	}

	fi, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("error stat'ing file: %w", err)
	}

	// If newly created or otherwise empty, write headers once
	if fi.Size() == 0 {
		w := csv.NewWriter(file)

		err = w.Write(CSV_HEADERS)
		if err != nil {
			file.Close()
			return nil, fmt.Errorf("error writing headers: %w", err)
		}

		w.Flush()
		err = w.Error()
		if err != nil {
			file.Close()
			return nil, fmt.Errorf("error flushing headers: %w", err)
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			file.Close()
			return nil, fmt.Errorf("error seeking to start: %w", err)
		}
	}

	return file, nil
}
