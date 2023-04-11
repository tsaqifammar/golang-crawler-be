package utils

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// todo, tambah cacher, tambah rate limiter juga
type URLFetcher struct{}

func (f *URLFetcher) Fetch(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return getUrlsFromHTMLPage(string(content)), nil
}

func getUrlsFromHTMLPage(htmlString string) []string {
	// Get all the anchor tags and returns the urls contained in the "href"s
	doc, _ := html.Parse(strings.NewReader(htmlString))

	urls := make([]string, 0)

	var f func(node *html.Node)
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					u, err := url.ParseRequestURI(a.Val)
					if err == nil && u.Scheme != "" && u.Host != "" {
						urls = append(urls, a.Val)
					}
				}
			}
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return urls
}
