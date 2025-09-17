package job

type Jobs []Job

func (jobs Jobs) Len() int {
	return len(jobs)
}

func (jobs Jobs) Swap(i, j int) {
	jobs[i], jobs[j] = jobs[j], jobs[i]
}

func (jobs Jobs) Less(i, j int) bool {
	date1 := jobs[i].Date
	date2 := jobs[j].Date

	y1, y2 := date1.Year(), date2.Year()
	if y1 == y2 {

		m1, m2 := date1.Month(), date2.Month()
		if m1 == m2 {

			d1, d2 := date1.Day(), date2.Day()
			if d1 == d2 {

				c1, c2 := jobs[i].Company, jobs[j].Company
				if c1 == c2 {

					url1, url2 := jobs[i].URL.String(), jobs[j].URL.String()
					if url1 == url2 {

						return jobs[i].ID < jobs[j].ID
					}
					return url1 < url2
				}
				return c1 < c2
			}
			return d1 < d2
		}
		return m1 < m2
	}
	return y1 < y2
}
