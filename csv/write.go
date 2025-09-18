package csv

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/ianhecker/web-scraper/job"
)

func WriteFile(file *os.File, jobs job.Jobs) error {
	if file == nil {
		return fmt.Errorf("given nil file")
	}

	w := csv.NewWriter(file)
	err := w.Write(CSV_HEADER)
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
	return w.Error()
}
