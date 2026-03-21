package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	report := make([]PageData, 0, len(keys))
	for _, k := range keys {
		report = append(report, pages[k])
	}
	jsonOutput, err := json.MarshalIndent(report, "", " ")
	if err != nil {
		return fmt.Errorf("Error marshalling JSON:", err)

	}
	err = os.WriteFile(filename, jsonOutput, 0644)
	if err != nil {
		return fmt.Errorf("Error writing JSON file:", err)
	}
	return nil
}
