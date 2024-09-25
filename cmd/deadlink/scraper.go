package deadlink

import (
	"fmt"
	"net/http"
	"sync"
)

type Link struct {
	page string
	url  string
}

type Props struct {
	mu             *sync.Mutex
	client         *http.Client
	host           string
	url            string
	found          *[]string
	scraped        *[]string
	scrapedChannel chan string
	broken         chan Link
	wg             *sync.WaitGroup
}

func Scraper(url string) {
	client := http.Client{}
	found := []string{url}
	scraped := []string{}
	broken := []Link{}
	scrapedChannel := make(chan string)
	brokenChannel := make(chan Link)

	var wg sync.WaitGroup
	success := make(chan bool, 1)

	res, err := client.Get(url)
	if err != nil {
		panic("Failed to fetch url")
	}
	baseHost := fmt.Sprintf("%v://%v", res.Request.URL.Scheme, res.Request.Host)

	props := Props{
		mu:             &sync.Mutex{},
		client:         &client,
		host:           baseHost,
		url:            url,
		found:          &found,
		scraped:        &scraped,
		scrapedChannel: scrapedChannel,
		broken:         brokenChannel,
		wg:             &wg,
	}
	props.wg.Add(1)
	go FetchLink(url, props)

	go func() {
		wg.Wait()
		success <- true
	}()

	for {
		select {
		case link := <-props.scrapedChannel:
			props.mu.Lock()
			scraped = append(scraped, link)
			props.mu.Unlock()
		case link := <-props.broken:
			props.mu.Lock()
			broken = append(broken, link)
			props.mu.Unlock()
		case <-success:
			defer close(props.broken)
			defer close(props.scrapedChannel)
			PrintLinks(&broken)
			fmt.Println("Broken links:", len(broken))
			return
		}
	}
}
