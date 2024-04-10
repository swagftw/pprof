package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// UserAgentInfo represents information extracted from user-agent strings.
type UserAgentInfo struct {
	Browser string
	OS      string
}

// userAgentRegex maps regular expressions to extract browser and OS information.
var browserRegex = map[string]*regexp.Regexp{
	"Chrome":  regexp.MustCompile(`(Chrome)/(\d+\.\d+)`),
	"Firefox": regexp.MustCompile(`(Firefox)/(\d+\.\d+)`),
}

var osRegex = map[string]*regexp.Regexp{
	"Windows": regexp.MustCompile(`(Windows NT) (\d+\.\d+)`),
	"MacOS":   regexp.MustCompile(`(Mac OS X) (\d+\_\d+)`),
	"Linux":   regexp.MustCompile(`Linux`),
}

// extractUserAgentInfo extracts browser and OS information from a user-agent string.
func extractUserAgentInfo(userAgent string) (string, string) {
	var browser, os string

	// Extract browser information
	for _, regex := range browserRegex {
		match := regex.FindStringSubmatch(userAgent)
		if len(match) > 2 {
			browser = match[1] + " " + match[2]
			break
		}
	}

	// Extract OS information
	for key, regex := range osRegex {
		match := regex.FindStringSubmatch(userAgent)
		if len(match) > 0 {
			if key == "Linux" {
				os = "Linux"
			} else {
				os = match[1] + " " + match[2]
			}

			break
		}
	}

	return browser, os
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	response := getStats(r.UserAgent())

	// Send the response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

var hostname = getHostname()

func getHostname() string {
	host, _ := os.Hostname()

	return host
}

func getStats(ua string) string {
	stats := make([]string, 0, 4)

	timeNow := time.Now()

	// Add the hostname to the stats
	stats = append(stats, "Hostname: "+hostname)

	// Extract browser and OS information from the user-agent string
	browser, os := extractUserAgentInfo(ua)

	// Format the extracted information
	if browser != "" {
		stats = append(stats, "Browser: "+browser)
	}
	if os != "" {
		stats = append(stats, "OS: "+os)
	}

	latency := strconv.Itoa(int(time.Since(timeNow).Microseconds()))

	// Add the latency to the stats
	stats = append(stats, "Latency: "+latency)

	// Join the information strings
	response := strings.Join(stats, "\n")

	return response
}

func main() {
	http.HandleFunc("/stats", handleRequest)
	fmt.Println("Server listening on port 9090...")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
