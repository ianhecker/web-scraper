package scrape

import (
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/ianhecker/web-scraper/internal/job"
)

func FindRawJobs(reader io.Reader) ([]job.Raw, error) {
	rawJobs := []job.Raw{}

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

		var locations []string
		isRemote := false
		s.Find("ul.job-info-list li").Each(func(_ int, li *goquery.Selection) {
			if li.Find(".fe-map-pin").Length() > 0 {

				txt := strings.TrimSpace(li.Text())
				if strings.EqualFold(txt, "Remote") {
					isRemote = true
				} else {
					locations = append(locations, txt)
				}
			}
		})
		location := ""
		if len(locations) > 0 {
			location = strings.Join(locations, ",")
		}

		url, _ := s.Find(".job-bottom a.theme-btn").Attr("href")

		title = strings.TrimSpace(title)
		company = strings.TrimSpace(company)
		salary = strings.TrimSpace(salary)
		date = strings.TrimSpace(date)
		location = strings.TrimSpace(location)
		url = strings.TrimSpace(url)

		rawJob := job.MakeRaw(title, company, salary, date, location, isRemote, url)
		rawJobs = append(rawJobs, rawJob)
	})
	return rawJobs, nil
}
