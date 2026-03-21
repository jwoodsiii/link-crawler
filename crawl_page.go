package main

import (
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrURL string) {
	defer cfg.wg.Done()
	<-cfg.concurrencyControl
	curr, err := url.Parse(rawCurrURL)
	if err != nil {
		log.Fatalf("Invalid URL %s", rawCurrURL)
	}
	if cfg.baseURL.Hostname() != curr.Hostname() && curr.Host != "" {
		log.Printf("Base domain %s dne curr url domain %s", cfg.baseURL.Host, curr.Host)
		return
	}
	currURL, err := normalizeURL(curr.String())
	if err != nil {
		log.Fatalf("Error normalizing url %s", curr.String())
	}
	if isFirstVisit := cfg.addPageVisit(currURL); !isFirstVisit {
		return
	}
	log.Printf("Getting HTML for currURL: %s", rawCurrURL)
	htmlData, err := getHTML(rawCurrURL)
	if err != nil {
		log.Printf("Error getting html data for %s\nError: %v", rawCurrURL, err)
		return
	}
	pgData := extractPageData(htmlData, rawCurrURL)
	cfg.setPageData(currURL, pgData)
	for _, link := range pgData.OutgoingLinks {
		log.Printf("Spawing goroutine to crawl %s", link)
		cfg.concurrentCrawlPage(link)
	}
	log.Printf("crawling complete\n")
}

func (cfg *config) concurrentCrawlPage(rawCurrURL string) {
	cfg.wg.Add(1)
	var empty struct{}
	cfg.concurrencyControl <- empty
	go cfg.crawlPage(rawCurrURL)
}
