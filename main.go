package main

import (
	"time"

	"github.com/tsaqifammar/url-crawler/lib"
)

// define crawler

func main() {
	c := lib.NewCrawler("https://google.com", 3, 100)
	c.Crawl()
	time.Sleep(5 * time.Second)
	c.GetResults()
}
