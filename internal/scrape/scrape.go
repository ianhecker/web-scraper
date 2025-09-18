package scrape

import (
	"bytes"
	"fmt"

	"github.com/ianhecker/web-scraper/internal/job"
)

func ScrapePages(url string, pages int) (job.Jobs, error) {
	jobs := []job.Job{}

	for pageNumber := 1; pageNumber <= 5; pageNumber++ {

		body, _, err := Get(url, pageNumber)
		if err != nil {
			return nil, fmt.Errorf("error scraping page: %w", err)
		}

		reader := bytes.NewReader(body)
		rawJobs, err := FindRawJobs(reader)
		if err != nil {
			return nil, fmt.Errorf("error parsing job posts: %w", err)
		}

		pageJobs, err := job.MakeJobsFromRawJobs(rawJobs)
		if err != nil {
			return nil, fmt.Errorf("error converting jobs: %w", err)
		}
		jobs = append(jobs, pageJobs...)
	}
	return job.MakeJobs(jobs), nil
}
