package lib

import (
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Fetch for the urls contained in a webpage of a given url
func Fetch(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	urls, err := getUrlsFromHTMLPage(string(content))
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func getUrlsFromHTMLPage(htmlString string) ([]string, error) {
	// Get all the anchor tags and return the urls contained in the "href"s
	doc, err := html.Parse(strings.NewReader(htmlString))

	if err != nil {
		return nil, err
	}

	urls := make([]string, 0)

	var f func(node *html.Node)
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" && IsUrl(a.Val) {
					urls = append(urls, a.Val)
				}
			}
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return urls, nil
}
