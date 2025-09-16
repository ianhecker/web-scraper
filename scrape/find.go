package scrape

import (
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
	"github.com/ianhecker/web-scraper/job"
)

func FindJobs(reader io.Reader) ([]job.Post, error) {
	posts := []job.Post{}

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("error creating document: %w", err)
	}

	doc.Find(".col-lg-12 > .job-item").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".job-title h5").Text()
		company := s.Find(".job-employer").Text()

		salary := s.Find(".job-salary-amount").Text()

		date := ""
		s.Find("ul.job-info-list li").Each(func(_ int, li *goquery.Selection) {
			if li.Find(".fe-clock").Length() > 0 {
				date = li.Text()
			}
		})

		locations := []string{}
		s.Find("ul.job-info-list li").Each(func(_ int, li *goquery.Selection) {
			if li.Find(".fe-map-pin").Length() > 0 {
				locations = append(locations, li.Text())
			}
		})

		url, _ := s.Find(".job-bottom a.theme-btn").Attr("href")

		post := job.MakePost(title, company, salary, date, url, locations)
		posts = append(posts, post)
	})
	return posts, nil
}
