package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/ianhecker/web-scraper/job"
)

func ReadFile(file *os.File) (job.JobsMap, error) {
	if file == nil {
		return nil, fmt.Errorf("given nil file")
	}

	r := csv.NewReader(file)
	headers, err := r.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	err = equalStrings(headers, CSV_HEADER)
	if err != nil {
		return nil, fmt.Errorf("error with headers: %w", err)
	}

	jobs := job.MakeJobs()
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading record: %w", err)
		}

		var job job.Job
		err = job.UnmarshalCSV(record)
		if err != nil {
			return nil, err
		}
		added := jobs.Add(job)
		if !added {
			return nil, fmt.Errorf("error adding duplicate ID: %s", job.ID)
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
