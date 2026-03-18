package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	var header string

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return fmt.Errorf("failed to parse HTML: %w", err).Error()
	}
	if doc.Find("h1").Text() != "" {
		header = doc.Find("h1").Text()
	} else if doc.Find("h2").Text() != "" {
		header = doc.Find("h2").Text()
	} else {
		header = ""
	}
	return header
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return fmt.Errorf("failed to parse HTML: %w", err).Error()
	}
	main := doc.Find("main")
	var p string
	if main.Length() > 0 {
		p = main.Find("p").First().Text()
	} else {
		p = doc.Find("p").First().Text()
	}

	return strings.TrimSpace(p)
}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	return extractLinks(htmlBody, baseURL, "a[href]", "href")
	// doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to parse HTMl: %w", err)
	// }
	// var urls []string
	// var fullURL string
	// doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
	// 	link, _ := s.Attr("href")
	// 	parsedLink, err := url.Parse(link)
	// 	if err != nil {
	// 		return
	// 	}
	// 	// fmt.Printf("Base url: %s\nselection values: %s\n", baseURL.String(), parsedLink.String())
	// 	if parsedLink.Scheme == "" {
	// 		fullURL = baseURL.String() + parsedLink.String()
	// 	} else {
	// 		fullURL = parsedLink.String()
	// 	}
	// 	urls = append(urls, fullURL)
	// })
	// return urls, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	return extractLinks(htmlBody, baseURL, "img", "src")
}
