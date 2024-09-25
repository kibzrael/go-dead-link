package main

import (
	"fmt"
	"kibzrael/deadlink/cmd/deadlink"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Arguments are required")
		return
	}

	switch args[0] {
	case "--url":
		if len(args) < 2 {
			fmt.Println("Url argument is required")
			return
		}
		url := args[1]
		start := time.Now()
		deadlink.Scraper(url)
		fmt.Println("Time taken:", time.Since(start))
	default:
		fmt.Println("No Command Found")

	}
}
