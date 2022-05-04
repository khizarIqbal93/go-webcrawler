package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/KhizarIqbal93/go-webcrawler/models"
	"github.com/KhizarIqbal93/go-webcrawler/utils"
	"github.com/fatih/color"
)

var output models.Links

func main() {
	var inputUrl string
	wantsOutput := flag.Bool("o", false, "set to true if you want to generate a JSON output in your current working directory")
	flag.Parse()
	// visited := make(map[string][]models.Link)
	// prompt for input
	fmt.Println("Enter a url")
	fmt.Scanln(&inputUrl)
	output.EntryPoint = inputUrl

	startTime := time.Now()

	output.LinksFound= utils.ExtractLinksFromPage(inputUrl)

	color.Blue("%s links found in %s", strconv.Itoa(len(output.LinksFound)), inputUrl)
	c := make(chan []models.Link)
	for index, link := range output.LinksFound {
		go func(pageLink models.Link, i int, channel chan []models.Link) {
			newLinks := utils.ExtractLinksFromPage(pageLink.Url)
			channel <- newLinks
		}(link, index, c)
	}

	for i := 0; i < len(output.LinksFound); i++ {
		linksFromChannel := <- c
		parent := linksFromChannel[0].Parent
		indexOfEntry := FindIndexOfParent(output.LinksFound, parent)
		output.LinksFound[indexOfEntry].Links = linksFromChannel
	}
	if (*wantsOutput) {
		outputJson, _ := json.Marshal(output)
		writeFileErr := os.WriteFile("output.json", outputJson, 0666)
		if writeFileErr != nil {
			log.Fatal("Could not create output file")
		}
	}
	fmt.Println(time.Since(startTime))

}

func FindIndexOfParent(a []models.Link, x string) int {
    for i, n := range a {
        if x == n.Url {
            return i
        }
    }
    return len(a)
}
