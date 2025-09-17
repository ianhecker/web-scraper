package main

import (
	"bytes"
	"fmt"
	"sort"

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

	added := existingJobs.AddJobs(newJobs...)

	if len(added) > 0 {
		fmt.Printf("added jobs: %d\n", len(added))

		for _, job := range added {
			fmt.Printf("%s @ %s\n", job.Title, job.Company)
		}
	}

	jobs := existingJobs.ToJobs()
	sort.Sort(jobs)

	err = csv.WriteFile(FILENAME, jobs)
	if err != nil {
		panic(err)
	}
}
