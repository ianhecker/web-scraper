package cmd

import (
	"fmt"
	"sort"

	"github.com/ianhecker/web-scraper/internal/csv"
	"github.com/ianhecker/web-scraper/internal/scrape"
	"github.com/spf13/cobra"
)

var pages int

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape pages",
	Run: func(cmd *cobra.Command, args []string) {
		scrapePages(URL, pages)
	},
}

func scrapePages(url string, pages int) {
	jobsFile, err := csv.OpenOrCreate(JOBS_FILENAME)
	checkErr(err)
	defer jobsFile.Close()

	newJobsFile, err := csv.OpenOrCreate(NEW_JOBS_FILENAME)
	checkErr(err)
	defer newJobsFile.Close()

	jobs, err := csv.ReadFile(jobsFile)
	checkErr(err)

	pageNumber := 5
	fetchedJobs, err := scrape.ScrapePages(URL, pageNumber)
	checkErr(err)

	newJobs := jobs.AddNewJobs(fetchedJobs...)
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
	scrapeCmd.Flags().IntVarP(&pages, "pages", "p", 1, "number of pages")
}
