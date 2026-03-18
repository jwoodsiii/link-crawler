package main

import (
	"fmt"
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
