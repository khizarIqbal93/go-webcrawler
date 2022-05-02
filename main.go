package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/KhizarIqbal93/go-webcrawler/utils"
	"github.com/fatih/color"
)

func main() {
	var url string
	fmt.Println("Enter a url")
	fmt.Scanln(&url)
	startTime := time.Now()
	// channel := make(chan bool)
	links := utils.ExtractLinks(url)
	// for
	//  go link(chanel)
	color.Blue("%s links found!", strconv.Itoa(len(links)))
	visited, linkMap := utils.GoCrawl(links)
	color.Cyan("%s unique pages visited!", strconv.Itoa(len(visited)))
	fmt.Println(len(linkMap), "<<<<link map")
	fmt.Println(time.Since(startTime))
	// <- channel
}
