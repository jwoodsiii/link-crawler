package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	client := http.Client{}
	req.Header.Set("User-Agent", "LinkCrawler/1.0")
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error while fetching html: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return "", fmt.Errorf("error while fetching html: status code %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("got non-HTML response: %s", contentType)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error while fetching html: %w", err)
	}
	return string(body), nil
}
