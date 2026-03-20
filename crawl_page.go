package main

import (
	"fmt"
	"log"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrURL string, pages map[string]int) {
	base, err := url.Parse(rawBaseURL)
	if err != nil {
		log.Fatalf("Invalid url %s", rawBaseURL)
	}
	curr, err := url.Parse(rawCurrURL)
	if err != nil {
		log.Fatalf("Invalid url %s", rawCurrURL)
	}
	if base.Host != curr.Host && curr.Host != "" {
		log.Printf("Base domain %s dne curr url domain %s", base.Host, curr.Host)
		return
	}
	currURL, err := normalizeURL(curr.String())
	if err != nil {
		log.Fatalf("Error normalizing url: %s", curr.String())
	}
	if _, visited := pages[currURL]; visited {
		log.Printf("Already crawled currURL: %s\n Number of crawls %d", currURL, pages[currURL])
		pages[currURL] += 1
		return
	} else {
		log.Printf("Adding currURL: %s to pages map", currURL)
		pages[currURL] = 1
	}
	log.Printf("Getting HTML for currURL: %s", rawCurrURL)
	htmlData, err := getHTML(rawCurrURL)
	if err != nil {
		log.Printf("Error getting html data for %s\nError: %v", rawCurrURL, err)
		return
	}
	pgData := extractPageData(htmlData, rawCurrURL)
	for _, url := range pgData.OutgoingLinks {
		log.Printf("Recursively crawling %s", url)
		crawlPage(rawBaseURL, url, pages)
	}
	log.Print("Finished recursive crawl")
	for k, v := range pages {
		fmt.Printf("%s: %d\n", k, v)
	}
}
