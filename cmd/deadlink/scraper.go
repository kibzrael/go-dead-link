package deadlink

import (
	"fmt"
	"net/http"
	"time"
)

type Link struct{
	page string
	url string
}

type Props struct{
	client *http.Client
	host string
	url string
	found *[]string
	scraped *[]string
	scrapedChannel chan string
	broken chan Link
}

func Scraper(url string){
	client := http.Client{}
	found := []string{url}
	scraped := []string{}
	broken := []Link{}
	scrapedChannel := make(chan string)
	brokenChannel := make(chan Link)

	res, err := client.Get(url)
	if err != nil{
		panic("Failed to fetch url")
	}
	baseHost := fmt.Sprintf("%v://%v", res.Request.URL.Scheme, res.Request.Host)

	props := Props{
		client: &client,
		host:  baseHost,
		url: url,
		found: &found,
		scraped: &scraped,
		scrapedChannel: scrapedChannel,
		broken: brokenChannel,
	}
	go FetchLink(url, props)

	for{
		select{
		case link  := <- props.scrapedChannel:
			scraped = append(scraped, link)
		case link  := <- props.broken:
			broken = append(broken, link)
		case <- time.After(1 * time.Second):
			if len(found) == len(scraped){
				defer close(props.broken)
				defer close(props.scrapedChannel)
				PrintLinks(&broken)
				fmt.Println("Broken links:", len(broken))
				return
			}
		}
	}
}