package deadlink

import (
	"fmt"
)


func FetchLink(page string, props Props){
	res, err := (*props.client).Get(props.url)
	if err != nil{
		fmt.Println("Failed to fetch url:", props.url)
		return
	}

	*props.scraped = append(*props.scraped, props.url)

	if res.StatusCode != 200{
		fmt.Println("*Link could not be fetched*")
		*props.broken = append(*props.broken, Link{url: props.url, page: page})
	}

	host := fmt.Sprintf("%v://%v", res.Request.URL.Scheme, res.Request.Host)
	if host == props.host{
		ParseHTML(res.Body, props)
	}
}