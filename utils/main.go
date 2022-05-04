package utils

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/KhizarIqbal93/go-webcrawler/models"
	"golang.org/x/net/html"
)

// takes in a url string and returns html doc in string
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

// takes a html doc in string and returns a slice of links found in the html doc
func ATagLinksExtractor(htmlDoc string) map[string]int {
	var linksFound = make(map[string]int)
	doc, err := html.Parse(strings.NewReader(htmlDoc))
	if err != nil {
		panic(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					linksFound[a.Val]++
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return linksFound
}

func ExtractLinksFromPage(link string) []models.Link {
	var linksFound []models.Link
	htmlText := GetHtml(link)
	u, err := url.Parse(link)
	if err != nil {
		panic(err)
	}
	scheme := u.Scheme
	host := u.Hostname()
	links := ATagLinksExtractor(htmlText)

	for key, _ := range links {
		if strings.HasPrefix(key, "/") {
			fullLink := scheme + "://" + host + key
			if fullLink +"/" == link || link +"/" == fullLink {
				continue
			}
			linkObj := models.Link{Url: fullLink, Parent: link}
			linksFound = append(linksFound, linkObj)
		} else if strings.HasPrefix(key, scheme+"://"+host) {
			linkObj := models.Link{Url: key, Parent: link}
			linksFound = append(linksFound, linkObj)
		} else if key == link + "/" {
			continue
		}
	}
	return linksFound
}
