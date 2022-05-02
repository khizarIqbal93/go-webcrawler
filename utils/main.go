package utils

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func GetHtml(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	html := string(body)
	return html
}

func HtmlParser(htmlDoc string) ([]string, int) {
	var links []string
	doc, err := html.Parse(strings.NewReader(htmlDoc))
	if err != nil {
		panic(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links, len(links)
}

func ExtractLinks(link string) []string {
	htmlText := GetHtml(link)
	u, err := url.Parse(link)
	if err != nil {
		panic(err)
	}
	scheme := u.Scheme
	host := u.Hostname()
	linksFound, numOfLinks := HtmlParser(htmlText)
	var validLinks []string
	for i := 0; i < numOfLinks; i++ {
		if strings.HasPrefix(linksFound[i], scheme+"://"+host) {
			validLinks = append(validLinks, linksFound[i])
		} else if strings.HasPrefix(linksFound[i], "/") {
			fullLink := scheme + "://" + host + linksFound[i]
			validLinks = append(validLinks, fullLink)
		}
	}
	return validLinks
}

func GoCrawl(links []string) (map[string]bool, map[string]int) {
	visitedLinks := make(map[string]bool)
	linkFrequency := make(map[string]int)
	newLinks := []string{}
	for i := 0; i < len(links); i++ {
		if !visitedLinks[links[i]] {
			extractedLinks := ExtractLinks(links[i])
			visitedLinks[links[i]] = true
			newLinks = append(newLinks, extractedLinks...)
		}
	}
	for _, link := range newLinks {
		if linkFrequency[link] == 0 {
			linkFrequency[link] = 1
		} else {
			linkFrequency[link]++
		}
	}
	return visitedLinks, linkFrequency
}
