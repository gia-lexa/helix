package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Structure to hold headers for JSON output
type Header struct {
	Key string   `json:"Key"`
	Val []string `json:"Val"`
}

func main() {
	// Command-line flags
	filterFlag := flag.String("filter", "", "comma-separated list of headers to be displayed (e.g., cache,vary)")
	jsonFlag := flag.Bool("json", false, "output is formatted into JSON for easy parsing")
	plainFlag := flag.Bool("plain", false, "output is formatted without any extraneous spacing or ANSI color codes")
	flag.Parse()

	// Ensure a URL is provided
	if flag.NArg() < 1 {
		fmt.Println("Usage: helix [options] <URL>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Extract URL
	url := flag.Arg(0)

	// Make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Fetch headers and status code
	headers := resp.Header
	statusCode := resp.Status

	// Apply filtering if specified
	filteredHeaders := headers
	if *filterFlag != "" {
		filteredHeaders = filterHeaders(headers, *filterFlag)
	}

	// Output results based on flags
	if *jsonFlag {
		outputJSON(filteredHeaders, statusCode)
	} else {
		outputFormatted(filteredHeaders, statusCode, *plainFlag)
	}
}

// Function to filter headers based on comma-separated filter list
func filterHeaders(headers http.Header, filter string) http.Header {
	filtered := make(http.Header)
	filters := strings.Split(filter, ",")
	for _, key := range filters {
		if values, exists := headers[http.CanonicalHeaderKey(strings.TrimSpace(key))]; exists {
			filtered[key] = values
		}
	}
	return filtered
}

// Function to output headers in JSON format
func outputJSON(headers http.Header, statusCode string) {
	var headerList []Header
	for key, values := range headers {
		headerList = append(headerList, Header{Key: key, Val: values})
	}

	// Include status code as a header
	headerList = append(headerList, Header{Key: "Status Code", Val: []string{statusCode}})

	jsonOutput, err := json.Marshal(headerList)
	if err != nil {
		fmt.Printf("Error formatting JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonOutput))
}

// Function to output headers in a formatted way
func outputFormatted(headers http.Header, statusCode string, plain bool) {
	for key, values := range headers {
		if plain {
			fmt.Printf("%s: %s\n", key, strings.Join(values, ", "))
		} else {
			fmt.Printf("\033[1;34m%s:\033[0m %s\n", key, strings.Join(values, ", "))
		}
	}
	// Output status code
	if plain {
		fmt.Printf("Status Code: %s\n", statusCode)
	} else {
		fmt.Printf("\033[1;32mStatus Code:\033[0m %s\n", statusCode)
	}
}
