# Web Scraper

Scrapes jobs from https://gojobs.run

## Usage

```bash
go run main.go
```

## Outputs

### jobs.csv

This list contains all scraped jobs

All scraped jobs are merged with any existing jobs

### new.csv

This list contains all newly scraped jobs not found in `jobs.csv`
