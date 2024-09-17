package deadlink

import (
	"fmt"
	"net/http"
	"os"
)

type Props struct{
	client *http.Client
	host string
	url string
	scraped *[]string
	broken *[]string
}

func Scraper(url string){
	client := http.Client{}
	scraped := []string{}
	broken := []string{}

	res, err := client.Get(url)
	if err != nil{
		panic("Failed to fetch url")
	}
	baseHost := fmt.Sprintf("%v://%v", res.Request.URL.Scheme, res.Request.Host)

	props := Props{
		client: &client,
		host:  baseHost,
		url: url,
		scraped: &scraped,
		broken: &broken,
	}
	FetchLink(props)

	const colorRed = "\033[0;31m"
	const colorNone = "\033[0m"
    fmt.Fprintf(os.Stdout, "\n%s BROKEN LINKS:\n", colorRed)
	for _, link := range broken{
		fmt.Fprintf(os.Stdout, "%s   %s\n%s\n", colorRed, link, colorNone)
	}
}