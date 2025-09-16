package main

import (
	"bytes"
	"fmt"

	"github.com/ianhecker/web-scraper/csv"
	"github.com/ianhecker/web-scraper/job"
	"github.com/ianhecker/web-scraper/scrape"
)

const FILENAME = "jobs.csv"
const URL = "https://gojobs.run"

func main() {
	existingJobs, err := csv.ReadFile(FILENAME)
	if err != nil {
		panic(err)
	}

	body, _, err := scrape.Get(URL)
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(body)
	rawJobs, err := scrape.FindRawJobs(reader)
	if err != nil {
		panic(err)
	}

	newJobs, err := job.MakeJobsFromRawJobs(rawJobs)
	if err != nil {
		panic(err)
	}

	m := make(map[string]job.Job, len(existingJobs))
	for _, job := range existingJobs {
		m[job.ID] = job
	}

	addedJobs := []job.Job{}
	for _, job := range newJobs {
		_, exists := m[job.ID]
		if !exists {
			existingJobs = append(existingJobs, job)
			addedJobs = append(addedJobs, job)
		}
	}

	if len(addedJobs) > 0 {
		fmt.Printf("added jobs: %d\n", len(addedJobs))
		for _, job := range addedJobs {
			fmt.Printf("Title: %s\nCompany: %s\n", job.Title, job.Company)
		}
	}
}
