package main

import (
	"log/slog"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/stats", HealthWithStatsHandler)

	slog.Info("starting http server on :8080")

	_ = http.ListenAndServe(":8080", nil)
}

func HealthWithStatsHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}

// func HealthWithStatsHandler(writer http.ResponseWriter, request *http.Request) {
// 	stats, err := generateStats()
// 	if err != nil {
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		_, _ = writer.Write([]byte(stats))
//
// 		return
// 	}
//
// 	writer.WriteHeader(http.StatusOK)
//
// 	_, _ = writer.Write([]byte("OK"))
// }
//
// var data = make(map[string]string)
//
// func init() {
// 	var dataBytes, err = os.ReadFile("./stats.json")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	err = json.Unmarshal(dataBytes, &data)
// 	if err != nil {
// 		panic(err)
// 	}
// }
//
// var statsPool = sync.Pool{
// 	New: func() any {
// 		return new(bytes.Buffer)
// 	},
// }
//
// func generateStats() (string, error) {
// 	stats := statsPool.Get().(*bytes.Buffer)
// 	stats.Reset()
//
// 	for user, ua := range data {
// 		stats.WriteString(user)
// 		stats.WriteString("|")
// 		getOSAndBrowser(stats, ua)
// 	}
//
// 	statsPool.Put(stats)
// 	return stats.String(), nil
// }
//
// func getOSAndBrowser(stats *bytes.Buffer, ua string) {
// 	if strings.Contains(ua, "Android") {
// 		stats.WriteString("Android")
// 		stats.WriteString("|")
// 	} else if strings.Contains(ua, "Chrome") {
// 		stats.WriteString("Chrome")
// 		stats.WriteString("|")
// 	} else if strings.Contains(ua, "Safari") {
// 		stats.WriteString("Safari")
// 		stats.WriteString("|")
// 	} else if strings.Contains(ua, "Opera") {
// 		stats.WriteString("Opera")
// 		stats.WriteString("|")
// 	} else if strings.Contains(ua, "Firefox") {
// 		stats.WriteString("Firefox")
// 		stats.WriteString("|")
// 	} else if strings.Contains(ua, "Edge") {
// 		stats.WriteString("Edge")
// 		stats.WriteString("|")
// 	}
//
// 	if strings.Contains(ua, "Windows") {
// 		stats.WriteString("Windows")
// 	} else if strings.Contains(ua, "Mac OS X") {
// 		stats.WriteString("macOS")
// 	} else if strings.Contains(ua, "Linux") {
// 		stats.WriteString("Linux")
// 	}
//
// 	stats.WriteString("\n")
// }
