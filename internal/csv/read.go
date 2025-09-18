package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/ianhecker/web-scraper/internal/job"
)

func ReadFile(file *os.File) (job.JobsMap, error) {
	if file == nil {
		return nil, fmt.Errorf("given nil file")
	}

	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("error seeking to start: %w", err)
	}

	r := csv.NewReader(file)
	gotHeaders, err := r.Read()
	if err != nil {

		if err == io.EOF {
			return nil, fmt.Errorf("empty csv: missing headers")
		}
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	err = equalStrings(gotHeaders, CSV_HEADERS)
	if err != nil {
		return nil, fmt.Errorf("error with headers: %w", err)
	}

	jobs := job.MakeJobsMap()
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("error reading record: %w", err)
		}

		var j job.Job
		if err := j.UnmarshalCSV(record); err != nil {
			return nil, err
		}

		added := jobs.Add(j)
		if !added {
			return nil, fmt.Errorf("error adding duplicate ID: %s", j.ID)
		}
	}
	return jobs, nil
}

func equalStrings(a, b []string) error {
	if len(a) != len(b) {
		return fmt.Errorf("wrong length: got:%d want:%d", len(a), len(b))
	}
	for i := range a {
		if a[i] != b[i] {
			return fmt.Errorf("header mismatch: got:%s want:%s", a[i], b[i])
		}
	}
	return nil
}
