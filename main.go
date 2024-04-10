package main

import (
	"bytes"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"regexp"
	"strconv"
	"sync"
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
func extractUserAgentInfo(userAgent string, buffer *bytes.Buffer) {
	for _, regex := range browserRegex {
		match := regex.FindStringSubmatch(userAgent)

		if len(match) > 2 {
			// browser = match[1] + " " + match[2]
			buffer.WriteString("\nBrowser: ")
			buffer.WriteString(match[1])
			buffer.WriteString(" ")
			buffer.WriteString(match[2])

			break
		}
	}

	// Extract OS information
	for key, regex := range osRegex {
		match := regex.FindStringSubmatch(userAgent)

		if len(match) > 0 {
			buffer.WriteString("\nOS: ")

			if key == "Linux" {
				buffer.WriteString("Linux")
			} else {
				buffer.WriteString(match[1])
				buffer.WriteString(" ")
				buffer.WriteString(match[2])
			}

			break
		}
	}
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

var bufferPool = sync.Pool{
	New: func() any {
		return &bytes.Buffer{}
	},
}

func getStats(ua string) string {
	buffer := bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()

	// stats := make([]string, 0, 4)

	timeNow := time.Now()

	// Add the hostname to the stats
	// stats = append(stats, "Hostname: "+hostname)
	buffer.WriteString("Hostname: ")
	buffer.WriteString(hostname)

	// Extract browser and OS information from the user-agent string
	extractUserAgentInfo(ua, buffer)

	latency := strconv.Itoa(int(time.Since(timeNow).Microseconds()))

	// Add the latency to the stats
	// stats = append(stats, "Latency: "+latency)
	buffer.WriteString("\nLatency: ")
	buffer.WriteString(latency)

	// Join the information strings
	// response := strings.Join(stats, "\n")
	bufferPool.Put(buffer)

	return buffer.String()
}

func main() {
	http.HandleFunc("/stats", handleRequest)
	fmt.Println("Server listening on port 9090...")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
