package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func main() {
	var url string
	fmt.Println("Enter a url")
	fmt.Scanln(&url)
	html := getHtml(url)
	links, numOfLinks := htmlLinkMatcher(html)
	fmt.Printf("This page has %v links\n", numOfLinks)
	fmt.Println("here are all the links>>>>>>\n", links)
}

func htmlLinkMatcher(html string) ([]string, int) {
	r, _ := regexp.Compile(`(href=)?"https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)"`)
	var links []string = r.FindAllString(html, -1)
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
