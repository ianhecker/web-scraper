package job

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type Job struct {
	ID       ID
	Title    string
	Company  string
	Salary   string
	Date     time.Time
	Location string
	IsRemote bool
	URL      url.URL
}

func MakeJob(title, company, salary, dateStr, location string, isRemote bool, href string) (Job, error) {
	date, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		return Job{}, fmt.Errorf("error parsing date: %w", err)
	}

	base, _ := url.Parse("https://gojobs.run")
	URL, err := url.Parse(href)
	if err != nil {
		return Job{}, fmt.Errorf("error parseing URL: %w", err)
	}
	fullURL := base.ResolveReference(URL)

	ID := MakeID(title, company, date, *URL)

	job := Job{
		ID:       ID,
		Title:    title,
		Company:  company,
		Salary:   salary,
		Date:     date,
		Location: location,
		IsRemote: isRemote,
		URL:      *fullURL,
	}
	return job, nil
}

func (j Job) MarshalCSV() []string {
	return []string{
		j.ID.String(),
		j.Title,
		j.Company,
		j.Salary,
		j.Date.Format("02-01-2006"),
		j.Location,
		fmt.Sprintf("%t", j.IsRemote),
		j.URL.String(),
	}
}

func (j *Job) UnmarshalCSV(record []string) error {
	if len(record) != 8 {
		return fmt.Errorf("invalid csv record length: %s", len(record))
	}

	j.Title = record[1]
	j.Company = record[2]
	j.Salary = record[3]
	j.Location = record[5]

	ID, err := MakeIDFromString(record[0])
	if err != nil {
		return fmt.Errorf("error creating ID: %w", err)
	}
	j.ID = ID

	date, err := time.Parse("02-01-2006", record[4])
	if err != nil {
		return fmt.Errorf("error parsing date: %w", err)
	}
	j.Date = date

	isRemote, err := strconv.ParseBool(record[6])
	if err != nil {
		return fmt.Errorf("error parsing bool: %w", err)
	}
	j.IsRemote = isRemote

	URL, err := url.Parse(record[7])
	if err != nil {
		return fmt.Errorf("error parsing URL: %w", err)
	}
	j.URL = *URL

	return nil
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
		ID:       j.ID.String(),
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
	var jobs = make(Jobs, len(rawJobs))
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
