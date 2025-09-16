package job

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Job struct {
	Title    string
	Company  string
	Salary   string
	Date     time.Time
	Location string
	IsRemote bool
	URL      url.URL
}

func (j Job) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Title    string
		Company  string
		Salary   string
		Date     string
		Location string
		IsRemote bool
		URL      string
	}{
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

	job := Job{
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
