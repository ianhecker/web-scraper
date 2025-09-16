package job

type Raw struct {
	Title    string
	Company  string
	Salary   string
	Date     string
	Location string
	IsRemote bool
	URL      string
}

func MakeRaw(
	title,
	company,
	salary,
	date,
	location string,
	isRemote bool,
	url string,
) Raw {
	return Raw{
		Title:    title,
		Company:  company,
		Salary:   salary,
		Date:     date,
		Location: location,
		IsRemote: isRemote,
		URL:      url,
	}
}
