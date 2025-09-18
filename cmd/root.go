package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

const JOBS_FILENAME = "jobs.csv"
const NEW_JOBS_FILENAME = "new.csv"
const URL = "https://gojobs.run/search?location=United+States&remote-jobs=on&sort-order="

var rootCmd = &cobra.Command{}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(scrapeCmd)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
