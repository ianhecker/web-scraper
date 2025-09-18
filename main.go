package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/ianhecker/web-scraper/internal/csv"
	"github.com/ianhecker/web-scraper/internal/scrape"
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

	pageNumber := 5
	fetchedJobs, err := scrape.ScrapePages(URL, pageNumber)
	checkErr(err)

	newJobs := jobs.AddNewJobs(fetchedJobs...)
	if len(newJobs) > 0 {
		fmt.Printf("added jobs: %d\n", len(newJobs))
	}

	sort.Sort(newJobs)
	err = csv.WriteFile(newJobsFile, newJobs)
	checkErr(err)

	allJobs := jobs.ToJobs()
	sort.Sort(allJobs)
	err = csv.WriteFile(jobsFile, allJobs)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
