package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ianhecker/web-scraper/job"
	"github.com/ianhecker/web-scraper/scrape"
)

func main() {
	url := "https://gojobs.run"

	body, _, err := scrape.Get(url)
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(body)
	rawJobs, err := scrape.FindRawJobs(reader)
	if err != nil {
		panic(err)
	}

	jobs, err := job.MakeJobsFromRawJobs(rawJobs)
	if err != nil {
		panic(err)
	}

	bytes, err := json.MarshalIndent(jobs, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
