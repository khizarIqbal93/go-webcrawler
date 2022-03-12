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
	domain := baseDomain(url)
	htmlText := getHtml(url)
	links, numOfLinks := htmlParser(htmlText)
	fmt.Printf("This page has %v links!\n", numOfLinks)
	fmt.Println(links)
	visitedLinks, newLinks := goVisit(links, domain)
	fmt.Println(len(visitedLinks), newLinks)
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

func goVisit(links []string, domain string) ([]string, []string) {
	visitedLinks := []string{}
	newLinks := []string{}
	for i := 0; i < len(links); i++ {
		if strings.HasPrefix(links[i], domain) || strings.HasPrefix(links[i], "/") {
			if !contains(visitedLinks, links[i]) {
				fmt.Println(links[i], i, "<<<<<<Bruh>>>>>>>>")
				// htmlText := getHtml(links[i]) // get html
				// visitedLinks = append(visitedLinks, links[i])
				// links, _ := htmlParser(htmlText)
				// newLinks = append(newLinks, links...)
			}
		} else {
			continue
		}
	}
	return visitedLinks, newLinks
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
