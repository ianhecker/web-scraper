package job

import "fmt"

type JobsMap map[ID]Job

func MakeJobs(jobs ...Job) JobsMap {
	m := make(map[ID]Job, len(jobs))
	for _, job := range jobs {
		m[job.ID] = job
	}
	return m
}

func (m JobsMap) Add(job Job) bool {
	_, exists := m[job.ID]
	if !exists {
		m[job.ID] = job
		return true
	}
	return false
}

func (m JobsMap) AddJobs(jobs ...Job) Jobs {
	added := Jobs{}
	for _, job := range jobs {
		if m.Add(job) {
			added = append(added, job)
		}
	}
	return added
}

func (m JobsMap) ToJobs() Jobs {
	jobs := make([]Job, len(m))
	i := 0
	for _, job := range m {
		jobs[i] = job
		i++
	}
	return Jobs(jobs)
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
