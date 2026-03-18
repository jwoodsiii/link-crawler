package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func extractLinks(htmlBody string, baseURL *url.URL, tag string, attr string) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTMl: %w", err)
	}

	urls := []string{}
	doc.Find(tag).Each(func(_ int, s *goquery.Selection) {
		link, exists := s.Attr(attr)
		if !exists {
			return
		}
		parsedLink, err := url.Parse(link)
		if err != nil {
			return
		}
		resolved := baseURL.ResolveReference(parsedLink)
		urls = append(urls, resolved.String())
	})
	return urls, nil
}
