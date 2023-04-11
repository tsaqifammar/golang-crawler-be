package main

import (
	"github.com/tsaqifammar/url-crawler/utils"
)

// define crawler

func main() {
	f := utils.URLFetcher{}

	f.Fetch("https://google.com")
}
