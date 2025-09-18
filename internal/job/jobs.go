package job

type Jobs []Job

func MakeJobs(jobs []Job) Jobs {
	return Jobs(jobs)
}

func (jobs Jobs) Len() int {
	return len(jobs)
}

func (jobs Jobs) Swap(i, j int) {
	jobs[i], jobs[j] = jobs[j], jobs[i]
}

func (jobs Jobs) Less(i, j int) bool {
	d1, d2 := jobs[i].Date, jobs[j].Date

	y1, y2 := d1.Year(), d2.Year()
	if y1 != y2 {
		return y1 > y2
	}

	m1, m2 := d1.Month(), d2.Month()
	if m1 != m2 {
		return m1 > m2
	}

	day1, day2 := d1.Day(), d2.Day()
	if day1 != day2 {
		return day1 > day2
	}

	if jobs[i].Company != jobs[j].Company {
		return jobs[i].Company < jobs[j].Company
	}

	if jobs[i].URL.String() != jobs[j].URL.String() {
		return jobs[i].URL.String() < jobs[j].URL.String()
	}
	return jobs[i].ID < jobs[j].ID
}
