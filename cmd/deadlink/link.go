package deadlink

import (
	"fmt"
)


func FetchLink(props Props){
	res, err := (*props.client).Get(props.url)
	if err != nil{
		fmt.Println("Failed to fetch url:", props.url)
	}

	*props.scraped = append(*props.scraped, props.url)

	if res.StatusCode != 200{
		fmt.Println("*Link could not be fetched*")
		*props.broken = append(*props.broken, props.url)
	}

	host := fmt.Sprintf("%v://%v", res.Request.URL.Scheme, res.Request.Host)
	if host == props.host{
		ParseHTML(res.Body, props)
	}
}