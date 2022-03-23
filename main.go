package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"golang.org/x/net/html"
)

func main() {
	var url string
	fmt.Println("Enter a url")
	fmt.Scanln(&url)
	startTime := time.Now()
	// channel := make(chan bool)
	links:= extractLinks(url)
	// for 
	//  go link(chanel)
	color.Blue("%s links found!", strconv.Itoa(len(links)))
	visited, linkMap := goCrawl(links)
	color.Cyan("%s unique pages visited!", strconv.Itoa(len(visited)))
	fmt.Println(len(linkMap), "<<<<link map")
	fmt.Println(time.Since(startTime))
	// <- channel
}



// func baseDomain(url string) string {
// 	r, _ := regexp.Compile(`^(?:https?:\/\/)?(?:[^@\n]+@)?(?:www\.)?([^:\/\n?]+)`)
// 	domain := r.FindAllString(url, -1)
// 	return domain[0]
// }

func contains(links []string, link string) bool {
	for _, x := range links {
		if x == link {
			return true
		}
	}
	return false
}

func goCrawl(links []string) ([]string, map[string]int) {
	visitedLinks := []string{}
	linkFrequency := make(map[string]int)
	newLinks := []string{}
	for i := 0; i < len(links); i++ {
		if !contains(visitedLinks, links[i]) {
			extractedLinks := extractLinks(links[i])
			visitedLinks = append(visitedLinks, links[i])
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

func extractLinks(link string) []string  {
	htmlText := getHtml(link)
	u, err := url.Parse(link)
	if err != nil {
		panic(err)
	}
	scheme := u.Scheme
	host := u.Hostname()
	linksFound, numOfLinks := htmlParser(htmlText)
	var validLinks []string
	for i := 0; i < numOfLinks; i++ {
		if strings.HasPrefix(linksFound[i], scheme + "://" + host) {
			validLinks = append(validLinks, linksFound[i])
		} else if strings.HasPrefix(linksFound[i], "/") {
			fullLink := scheme + "://" + host + linksFound[i]
			validLinks = append(validLinks, fullLink)
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
