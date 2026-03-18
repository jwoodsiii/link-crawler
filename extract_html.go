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
	var firstPara string
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return fmt.Errorf("failed to parse HTML: %w", err).Error()
	}
	doc.Find("main p").EachWithBreak(func(i int, s *goquery.Selection) bool {
		firstPara = s.Text()
		return false
	})
	if firstPara == "" {
		doc.Find("p").EachWithBreak(func(i int, s *goquery.Selection) bool {
			firstPara = s.Text()
			return false
		})
	}
	return firstPara
}
