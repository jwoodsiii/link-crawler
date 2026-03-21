package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("usage: link-crawler <base-url> <max-concurrency> <max-pages>")
		os.Exit(1)
	}
	rawBaseURL := os.Args[1]
	maxConcurrency, _ := strconv.Atoi(os.Args[2])
	maxPages, _ := strconv.Atoi(os.Args[3])
	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Println("Failed to configure: %v", err)
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)
	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	if err = writeJSONReport(cfg.pages, "report.json"); err != nil {
		fmt.Println("Error writing json report:", err)
	}
}
