package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/KhizarIqbal93/go-webcrawler/utils"
	"github.com/fatih/color"
)

var output []string

func main() {
	var inputUrl string
	fmt.Println("Enter a url")
	fmt.Scanln(&inputUrl)
	startTime := time.Now()
	visited := make(map[string]int)
	links := utils.ExtractLinksFromPage(inputUrl)
	color.Blue("%s links found!", strconv.Itoa(len(links)))
	c := make(chan []string)
	for index, link := range links {
		go func(pageLink string, i int, channel chan []string) {
			newLinks := utils.ExtractLinksFromPage(pageLink)
			channel <- newLinks
		}(link, index, c)
	}

	for i := 0; i < len(links); i++ {
		for _, link := range <-c {
			visited[link]++
			if visited[link] == 1 {
				output = append(output, link)
			}
		}
	}
	fmt.Println(len(output))
	fmt.Println(time.Since(startTime))

}
