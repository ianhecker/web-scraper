package main

import (
	"bytes"
	"fmt"
	"log"
	"sort"

	"github.com/ianhecker/web-scraper/csv"
	"github.com/ianhecker/web-scraper/job"
	"github.com/ianhecker/web-scraper/scrape"
)

const JOBS_FILENAME = "jobs.csv"
const NEW_JOBS_FILENAME = "new.csv"
const URL = "https://gojobs.run/search?location=United+States&remote-jobs=on&sort-order="

func main() {
	jobsFile, err := csv.OpenOrCreate(JOBS_FILENAME)
	checkErr(err)
	defer jobsFile.Close()

	newJobsFile, err := csv.OpenOrCreate(NEW_JOBS_FILENAME)
	checkErr(err)
	defer newJobsFile.Close()

	jobs, err := csv.ReadFile(jobsFile)
	checkErr(err)

	var fetchedJobs []job.Job
	for pageNumber := 1; pageNumber <= 5; pageNumber++ {

		body, _, err := scrape.Get(URL, pageNumber)
		checkErr(err)

		reader := bytes.NewReader(body)
		rawJobs, err := scrape.FindRawJobs(reader)
		checkErr(err)

		jobs, err := job.MakeJobsFromRawJobs(rawJobs)
		checkErr(err)

		fetchedJobs = append(fetchedJobs, jobs...)
	}

	newJobs := jobs.AddJobs(fetchedJobs...)
	if len(newJobs) > 0 {
		fmt.Printf("added jobs: %d\n", len(newJobs))
	}

	err = csv.WriteFile(newJobsFile, newJobs)
	checkErr(err)

	sortedJobs := jobs.ToJobs()
	sort.Sort(sortedJobs)

	err = csv.WriteFile(jobsFile, sortedJobs)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
