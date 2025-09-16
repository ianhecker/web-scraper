package job

import "strings"

type Post struct {
	Title    string
	Company  string
	Salary   string
	Date     string
	Location string
	IsRemote bool
	URL      string
}

func MakePost(
	title,
	company,
	salary,
	date,
	location string,
	isRemote bool,
	url string,
) Post {
	return Post{
		Title:    title,
		Company:  company,
		Salary:   salary,
		Date:     date,
		Location: location,
		IsRemote: isRemote,
		URL:      url,
	}
}

func MakePostFromDoc(
	title,
	company,
	salary,
	date,
	location string,
	isRemote bool,
	url string,
) Post {
	return MakePost(
		strings.TrimSpace(title),
		strings.TrimSpace(company),
		strings.TrimSpace(salary),
		strings.TrimSpace(date),
		strings.TrimSpace(location),
		isRemote,
		strings.TrimSpace(url),
	)
}
