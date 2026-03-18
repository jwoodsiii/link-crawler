package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(URL string) (string, error) {
	parts, err := url.Parse(URL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
	return strings.ToLower(fmt.Sprintf("%s%s", parts.Host, strings.TrimSuffix(parts.Path, "/"))), nil
}
