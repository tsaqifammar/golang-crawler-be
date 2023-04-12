package lib

import (
	"fmt"
	"log"
	"time"
)

type Crawler struct {
	cache         *URLCacher
	ticker        <-chan time.Time
	rootUrl       string
	depth, maxUrl int
	hasCrawled    bool
}

func NewCrawler(url string, depth int, maxUrl int) Crawler {
	ticker := time.NewTicker(500 * time.Millisecond)
	c := Crawler{
		rootUrl:    url,
		depth:      depth - 1,
		maxUrl:     maxUrl,
		hasCrawled: false,
		ticker:     ticker.C,
	}
	return c
}

// To crawl for urls starting from a root url.
// Proceed with running GetResults after to get the results.
func (c *Crawler) Crawl() {
	if c.rootUrl == "" {
		log.Println("Crawl is used but crawler has not been setup yet")
		return
	}
	c.cache = NewURLCacher()
	c.getUrls(c.rootUrl, 0)
}

func (c *Crawler) getUrls(url string, depth int) {
	<-c.ticker
	if depth == c.depth {
		return
	}

	urls, err := Fetch(url)

	if err != nil {
		return
	}

	for _, u := range urls {
		if !c.cache.HasVisited(u) && c.cache.GetUrlCount() < c.maxUrl {
			c.cache.MarkVisited(u, depth+1, url)
			go c.getUrls(u, depth+1)
		}
	}
}

// To get the the resulting urls after running Crawl.
func (c *Crawler) GetResults() {
	// Generate a tree based on information that is saved in cache
	res := c.cache.Info
	fmt.Println(len(res))

	adj := make(map[string][]string)
	for url, info := range res {
		adj[info.ParentUrl] = append(adj[info.ParentUrl], url)
	}

	var dfs func(cur string, depth int)
	dfs = func(cur string, depth int) {
		for i := 0; i < depth; i++ {
			fmt.Print("|")
		}
		fmt.Println(cur)

		for _, c := range adj[cur] {
			if c == cur {
				log.Fatal("tf")
			}
			dfs(c, depth+1)
		}
	}

	dfs(c.rootUrl, 0)
}
