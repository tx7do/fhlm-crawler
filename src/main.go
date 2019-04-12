package main

import (
	"github.com/jakecoffman/cron"

	fhlm "./crawler"
)

func main() {
	spec := "10 * * * * ?"
	cronJob := cron.New()
	cronJob.AddFunc(spec, func() {
		fhlm.StartCrawler("jlffc")
	}, "crawl")
	cronJob.Start()

	defer cronJob.Stop()

	select {}
}
