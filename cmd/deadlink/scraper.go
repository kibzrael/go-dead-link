package deadlink

import (
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"
)

type Link struct{
	page string
	url string
}

type Props struct{
	client *http.Client
	host string
	url string
	scraped *[]string
	broken *[]Link
}

func Scraper(url string){
	client := http.Client{}
	scraped := []string{}
	broken := []Link{}

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
	FetchLink(url, props)

	const colorRed = "\033[0;31m"
	const colorNone = "\033[0m"
    fmt.Print("\nBROKEN LINKS:\n\n")

	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 4, ' ', 0)
	defer writer.Flush()
	fmt.Fprintf(writer, "%s   %s\t%s %s\n", colorNone, "Page", colorRed, "Link")
	for _, link := range broken{
		fmt.Fprintf(writer, "%s   %s\t%s %s\n", colorNone, link.page, colorRed, link.url)
	}
	fmt.Fprintf(writer, "\n%s", colorNone)
}