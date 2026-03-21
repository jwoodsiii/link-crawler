package main

import (
	"fmt"
	"log"
	"net/url"
	"sync"
)

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, ok := cfg.pages[normalizedURL]; ok {
		log.Printf("Already crawled currURL: %s\n", normalizedURL)
		return false
	} else {
		log.Print("Marking page as seen\n")
		cfg.pages[normalizedURL] = PageData{}
		return true
	}
}

func (cfg *config) setPageData(url string, data PageData) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	cfg.pages[url] = data
}

func configure(rawBaseURL string, maxConcurrency int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &config{
		pages:              make(map[string]PageData),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil
}
