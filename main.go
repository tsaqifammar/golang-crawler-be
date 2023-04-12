package main

import (
	"fmt"

	"github.com/tsaqifammar/url-crawler/lib"
)

// define crawler

func main() {
	fmt.Println("running")
	x, _ := lib.Fetch("https://google.com")
	for idx, url := range x {
		fmt.Println(idx, url)
	}
}
