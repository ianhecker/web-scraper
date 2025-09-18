package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/ianhecker/web-scraper/internal/job"
)

func WriteFile(file *os.File, jobs job.Jobs) error {
	if file == nil {
		return fmt.Errorf("given nil file")
	}

	err := file.Truncate(0)
	if err != nil {
		return fmt.Errorf("truncate failed: %w", err)
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("seek failed: %w", err)
	}

	w := csv.NewWriter(file)
	err = w.Write(CSV_HEADERS)
	if err != nil {
		return fmt.Errorf("error writing file headers: %w", err)
	}

	for _, job := range jobs {

		record := job.MarshalCSV()
		err := w.Write(record)
		if err != nil {
			return fmt.Errorf("error writing record: %w", err)
		}
	}
	w.Flush()

	if err := w.Error(); err != nil {
		return fmt.Errorf("flush error: %w", err)
	}
	return nil
}
