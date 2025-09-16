package main

import (
	"fmt"

	"github.com/ianhecker/web-scraper/scrape"
)

func main() {
	url := "https://gojobs.run"

	body, _, err := scrape.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
