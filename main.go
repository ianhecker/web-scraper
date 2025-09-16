package main

import (
	"bytes"
	"fmt"

	"github.com/ianhecker/web-scraper/scrape"
)

func main() {
	url := "https://gojobs.run"

	body, _, err := scrape.Get(url)
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(body)
	posts, err := scrape.FindJobs(reader)
	if err != nil {
		panic(err)
	}

	for _, post := range posts {
		fmt.Printf("post: %+v\n", post)
	}
}
