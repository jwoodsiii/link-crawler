package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

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
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	return extractLinks(htmlBody, baseURL, "img", "src")
}

func extractPageData(html string, pageURL *url.URL) PageData {
	heading := getHeadingFromHTML(html)
	firstPara := getFirstParagraphFromHTML(html)
	outgoingLinks, err := getURLsFromHTML(html, pageURL)
	if err != nil {
		log.Printf("Error getting URLs from HTML: %v", err)
	}
	imgURLs, err := getImagesFromHTML(html, pageURL)
	if err != nil {
		log.Printf("Error getting images from HTML: %v", err)
	}
	return PageData{
		URL:            pageURL.String(),
		Heading:        heading,
		FirstParagraph: firstPara,
		OutgoingLinks:  outgoingLinks,
		ImageURLs:      imgURLs,
	}
}
