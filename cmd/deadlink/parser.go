package deadlink

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)


func ParseHTML(body io.Reader, props Props){
	page := props.url
	tokenizer := html.NewTokenizer(body)
	for {
		tokenType := tokenizer.Next()
		switch tokenType{
		case html.StartTagToken, html.EndTagToken:
			token := tokenizer.Token()
			if token.Data == "a"{
				for _, attr := range token.Attr{
					if attr.Key == "href"{
						link := attr.Val
						if strings.Split(link, "")[0] == "/" {
							link = fmt.Sprintf("%v%v", props.host, link)
						}
						fmt.Println("Checking Link:", link)
						scraped := Contains(props.scraped, link)
						if !scraped{
							props.url = link
							FetchLink(page, props)
						} else{
							fmt.Println("Already Checked: Skipping...")
						}
					}
				}
			}
		case html.ErrorToken:
			return
		}
	}
}