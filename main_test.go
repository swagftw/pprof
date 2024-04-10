package main

import (
	"strings"
	"testing"
)

func TestGetStats(t *testing.T) {
	response := getStats("Mozilla/5.0 (X11; Linux x86_64; rv:124.0) Gecko/20100101 Firefox/124.0")
	if response == "" {
		t.Error("Response is empty")
	}

	splits := strings.Split(response, "\n")
	// response should have 4 new line chars
	if len(splits) != 4 {
		t.Errorf("Response should have 4 new line chars, got %d", len(strings.Split(response, "\n")))
	} else {
		if !strings.HasPrefix(splits[0], "Hostname: ") {
			t.Errorf("Hostname should be 'Hostname: ', got %s", splits[0])
		}

		if !strings.HasPrefix(splits[1], "Browser: ") {
			t.Errorf("Browser should be 'Browser: ', got %s", splits[1])
		}

		if !strings.HasPrefix(splits[2], "OS: ") {
			t.Errorf("OS should be 'OS: ', got %s", splits[2])
		}

		if !strings.HasPrefix(splits[3], "Latency: ") {
			t.Errorf("Latency should be 'Latency: ', got %s", splits[3])
		}
	}
}

func BenchmarkGetStats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getStats("Mozilla/5.0 (X11; Linux x86_64; rv:124.0) Gecko/20100101 Firefox/124.0")
	}
}
