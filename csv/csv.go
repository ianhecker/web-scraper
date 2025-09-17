package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/ianhecker/web-scraper/job"
)

var CSV_HEADER = []string{
	"id",
	"title",
	"company",
	"salary",
	"date (DD-MM-YYYY)",
	"location",
	"is_remote",
	"url",
}

func WriteFile(path string, jobs []job.Job) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error crearing file: %w", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	err = w.Write(CSV_HEADER)
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

func ReadFile(path string) (job.JobsMap, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

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
