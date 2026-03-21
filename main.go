package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := os.Args[1]
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		log.Printf("Failed to parse baseurl %s", rawBaseURL)
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)
	pages := make(map[string]PageData, 0)
	var mu sync.Mutex
	var wg sync.WaitGroup
	concurrencyControl := make(chan struct{}, 2)
	cfg := &config{
		pages:              pages,
		baseURL:            baseURL,
		mu:                 &mu,
		concurrencyControl: concurrencyControl,
		wg:                 &wg,
	}

	cfg.concurrentCrawlPage(rawBaseURL)
	cfg.wg.Wait()

	for k, v := range cfg.pages {
		fmt.Printf("%s: %s\n", k, v)
	}
}
