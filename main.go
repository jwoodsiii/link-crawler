package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("usage: link-crawler <base-url> <max-concurrency> <max-pages>")
		os.Exit(1)
	}
	rawBaseURL := os.Args[1]
	maxConcurrency, _ := strconv.Atoi(os.Args[2])
	maxPages, _ := strconv.Atoi(os.Args[3])
	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Println("Failed to configure: %v", err)
		os.Exit(1)
	}
	cfg.concurrentCrawlPage(rawBaseURL)
	cfg.wg.Wait()

	for k, v := range cfg.pages {
		fmt.Printf("%s: %s\n", k, v)
	}
}
