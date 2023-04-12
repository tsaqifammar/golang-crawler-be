package lib

import "log"

// TODO: add rate limiter
type Crawler struct {
	cache         *URLCacher
	rootUrl       string
	depth, maxUrl int
	hasCrawled    bool
}

func NewCrawler(url string, depth int, maxUrl int) Crawler {
	c := Crawler{
		rootUrl:    url,
		depth:      depth,
		maxUrl:     maxUrl,
		hasCrawled: false,
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
	if depth == c.depth || c.cache.GetUrlCount() >= c.maxUrl {
		return
	}

	urls, err := Fetch(url)

	if err != nil {
		return
	}

	for _, u := range urls {
		if !c.cache.HasVisited(u) {
			c.cache.MarkVisited(u, depth+1, url)
			go c.getUrls(u, depth+1)
		}
	}
}

// To get the the resulting urls after running Crawl.
func (c *Crawler) GetResults() {
	// Generate a tree based on information that is saved in cache
}
