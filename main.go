package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	var url string
	fmt.Println("Enter a url")
	fmt.Scanln(&url)
	links:= extractLinks(url)
	fmt.Println(links)
	fmt.Println(len(links))
}

func baseDomain(url string) string {
	r, _ := regexp.Compile(`^(?:https?:\/\/)?(?:[^@\n]+@)?(?:www\.)?([^:\/\n?]+)`)
	domain := r.FindAllString(url, -1)
	return domain[0]
}

func contains(links []string, link string) bool {
	for _, x := range links {
		if x == link {
			return true
		}
	}
	return false
}

func goCrawl(links []string) ([]string, []string) {
	visitedLinks := []string{}
	newLinks := []string{}
	
	return visitedLinks, newLinks
}

func extractLinks(url string) []string  {
	htmlText := getHtml(url)
	domain := baseDomain(url)
	linksFound, numOfLinks := htmlParser(htmlText)
	var validLinks []string
	for i := 0; i < numOfLinks; i++ {
		if strings.HasPrefix(linksFound[i], domain) || strings.HasPrefix(linksFound[i], "/") {
			validLinks = append(validLinks, linksFound[i])
		}
	}
	return validLinks
}

func htmlParser(htmlDoc string) ([]string, int) {
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

func getHtml(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() //close connection
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	html := string(body)
	return html
}
