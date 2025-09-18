package cmd

import (
	"fmt"
	"sort"

	"github.com/ianhecker/web-scraper/internal/csv"
	"github.com/ianhecker/web-scraper/internal/job"
	"github.com/spf13/cobra"
)

var (
	title    string
	company  string
	salary   string
	date     string
	location string
	isRemote bool
	url      string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add job to list",
	Run: func(cmd *cobra.Command, args []string) {
		addJob(title, company, salary, date, location, isRemote, url)
	},
}

func addJob(title, company, salary, date, location string, isRemote bool, url string) {
	jobsFile, err := csv.OpenOrCreate(JOBS_FILENAME)
	checkErr(err)
	defer jobsFile.Close()

	newJobsFile, err := csv.OpenOrCreate(NEW_JOBS_FILENAME)
	checkErr(err)
	defer newJobsFile.Close()

	jobs, err := csv.ReadFile(jobsFile)
	checkErr(err)

	manualJob, err := job.MakeJob(title, company, salary, date, location, isRemote, url)
	checkErr(err)

	newJobs := jobs.AddNewJobs(manualJob)
	if len(newJobs) > 0 {
		fmt.Printf("added jobs: %d\n", len(newJobs))
	}

	sort.Sort(newJobs)
	err = csv.WriteFile(newJobsFile, newJobs)
	checkErr(err)

	allJobs := jobs.ToJobs()
	sort.Sort(allJobs)
	err = csv.WriteFile(jobsFile, allJobs)
	checkErr(err)
}

func init() {
	addCmd.Flags().StringVarP(&title, "title", "t", "", "title of job")
	addCmd.Flags().StringVarP(&company, "company", "c", "", "company of job")
	addCmd.Flags().StringVarP(&salary, "salary", "s", "", "salary of job ($100,000 - $120,000)")
	addCmd.Flags().StringVarP(&date, "date", "d", "", "date (DD-MM-YYYY)")
	addCmd.Flags().StringVarP(&location, "location", "l", "", "location of job")
	addCmd.Flags().BoolVarP(&isRemote, "isRemote", "r", false, "is remote")
	addCmd.Flags().StringVarP(&url, "url", "u", "", "url of job post (gojobs.run/...)")

	addCmd.MarkFlagRequired("title")
	addCmd.MarkFlagRequired("company")
	addCmd.MarkFlagRequired("salary")
	addCmd.MarkFlagRequired("date")
	addCmd.MarkFlagRequired("location")
	addCmd.MarkFlagRequired("isRemote")
	addCmd.MarkFlagRequired("url")
}
