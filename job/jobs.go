package job

type Jobs map[ID]Job

func MakeJobs(jobs ...Job) Jobs {
	m := make(map[ID]Job, len(jobs))
	for _, job := range jobs {
		m[job.ID] = job
	}
	return m
}

func (m Jobs) AddNewJobs(jobs ...Job) []Job {
	added := []Job{}
	for _, job := range jobs {
		_, exists := m[job.ID]
		if !exists {
			m[job.ID] = job
			added = append(added, job)
		}
	}
	return added
}
