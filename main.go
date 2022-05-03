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
	links := utils.ExtractLinksFromPage(inputUrl)
	color.Blue("%s links found!", strconv.Itoa(len(links)))
	c := make(chan bool)
	for index, link := range links {
		go func(pageLink string, i int, channel chan bool) {
			newLinks := utils.ExtractLinksFromPage(pageLink)
			output = append(output, newLinks...)
			fmt.Println(i, output)
			channel <- true

		}(link, index, c)
	}
	<-c
	fmt.Println(len(output))
	fmt.Println(time.Since(startTime))

}
