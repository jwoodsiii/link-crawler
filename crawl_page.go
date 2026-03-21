package main

import (
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrURL string) {
	cfg.concurrencyControl <- struct{}{} // acquire slot
	defer func() {
		<-cfg.concurrencyControl // release slot
		cfg.wg.Done()
	}()
	if cfg.maxPagesReached() {
		log.Print("max pages reached, stopping crawl")
		return
	}
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
		cfg.wg.Add(1)
		go cfg.crawlPage(link)
	}
	log.Printf("crawling complete\n")
}

func (cfg *config) concurrentCrawlPage(rawCurrURL string) {
	cfg.wg.Add(1)
	cfg.concurrencyControl <- struct{}{}
	go cfg.crawlPage(rawCurrURL)
}

func (cfg *config) maxPagesReached() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages) >= cfg.maxPages
}
