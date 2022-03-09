package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	var url string
	fmt.Println("Enter a url")
	fmt.Scanln(&url)
	htmlText := getHtml(url)
	links, numOfLinks := htmlParser(htmlText)
	fmt.Printf("This page has %v links!\n", numOfLinks)
	fmt.Println(links)
	visitedLinks, newLinks := goVisit(links)
	fmt.Println(len(visitedLinks), newLinks)
}

func goVisit(links []string) ([]string, []string) {
	visitedLinks := []string{}
	newLinks := []string{}
	for i := 0; i < len(links); i++ {
		if len(links[i]) <= 9 {
			continue
		}
		htmlText := getHtml(links[i])
		visitedLinks = append(visitedLinks, links[i])
		links, _ := htmlParser(htmlText)
		newLinks = append(newLinks, links...)
	}
	return visitedLinks, newLinks
}


func htmlParser(htmlDoc string) ([]string, int)   {
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