package job

type Post struct {
	Title     string
	Company   string
	Salary    string
	Date      string
	URL       string
	Locations []string
}

func MakePost(title, company, salary, date, url string, location []string) Post {
	return Post{
		Title:     title,
		Company:   company,
		Salary:    salary,
		Date:      date,
		URL:       url,
		Locations: location,
	}
}
