package lib

import "sync"

type UrlCacheInfo struct {
	Depth     int
	ParentUrl string
}

type URLCacher struct {
	mu sync.Mutex

	Info     map[string]UrlCacheInfo
	UrlCount int
}

func NewURLCacher() *URLCacher {
	uc := URLCacher{
		Info:     make(map[string]UrlCacheInfo),
		UrlCount: 0,
	}
	return &uc
}

func (c *URLCacher) MarkVisited(url string, depth int, parentUrl string) {
	c.mu.Lock()
	c.Info[url] = UrlCacheInfo{
		Depth:     depth,
		ParentUrl: parentUrl,
	}
	c.UrlCount = len(c.Info)
	c.mu.Unlock()
}

func (c *URLCacher) HasVisited(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.Info[url]
	return ok
}

func (c *URLCacher) GetUrlCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.UrlCount
}
