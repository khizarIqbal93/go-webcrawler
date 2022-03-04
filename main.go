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
	links, numOfLinks := aTagLinkMatcher(html)
	fmt.Printf("This page has %v links!\n", numOfLinks)
	fmt.Println("here are all the links\n>>>>>>\n", links, "\n>>>>>>>")
}

func aTagLinkMatcher(html string) ([]string, int) {
	r, _ := regexp.Compile(`<a (href=)?"https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&\/\/=]*)">`)
	var aTags []string = r.FindAllString(html, -1)
	links := []string{}
	for i := 0; i < len(aTags); i++ {
		links = append(links, aTags[i][9:len(aTags[i]) -2])
	}
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