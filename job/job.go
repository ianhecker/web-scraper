package job

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type Job struct {
	ID       string
	Title    string
	Company  string
	Salary   string
	Date     time.Time
	Location string
	IsRemote bool
	URL      url.URL
}

func MakeJob(
	title,
	company,
	salary,
	dateStr,
	location string,
	isRemote bool,
	urlStr string,
) (Job, error) {
	date, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		return Job{}, fmt.Errorf("error parsing date: %w", err)
	}

	URL, err := url.Parse(urlStr)
	if err != nil {
		return Job{}, fmt.Errorf("error parseing URL: %w", err)
	}

	ID := MakeID(title, company, date, *URL)

	job := Job{
		ID:       ID,
		Title:    title,
		Company:  company,
		Salary:   salary,
		Date:     date,
		Location: location,
		IsRemote: isRemote,
		URL:      *URL,
	}
	return job, nil
}

func MakeID(
	title,
	company string,
	date time.Time,
	url url.URL,
) string {
	seed := title +
		"|" + company +
		"|" + date.Format("02-01-2006") +
		"|" + url.String()
	base := strings.ToLower(seed)
	sum := sha1.Sum([]byte(base))
	return hex.EncodeToString(sum[:])[:7]
}

func (j Job) MarshalJSON() ([]byte, error) {
	tmp := struct {
		ID       string
		Title    string
		Company  string
		Salary   string
		Date     string
		Location string
		IsRemote bool
		URL      string
	}{
		ID:       j.ID,
		Title:    j.Title,
		Company:  j.Company,
		Salary:   j.Salary,
		Date:     j.Date.Format("02-01-2006"),
		Location: j.Location,
		IsRemote: j.IsRemote,
		URL:      j.URL.String(),
	}
	bytes, err := json.Marshal(tmp)
	if err != nil {
		return nil, fmt.Errorf("error marshaling job: %w", err)
	}
	return bytes, nil
}

func MakeJobsFromRawJobs(rawJobs []Raw) ([]Job, error) {
	var jobs = make([]Job, len(rawJobs))
	for i, rawJob := range rawJobs {
		job, err := MakeJob(
			rawJob.Title,
			rawJob.Company,
			rawJob.Salary,
			rawJob.Date,
			rawJob.Location,
			rawJob.IsRemote,
			rawJob.URL,
		)
		if err != nil {
			return nil, fmt.Errorf("error creating job: %w", err)
		}
		jobs[i] = job
	}
	return jobs, nil
}
