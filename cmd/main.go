package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"time"
)

func main() {
	var port = flag.String("port", "", "Port to listen on")
	var color = flag.String("color", "", "Color value to return in responses")

	flag.Parse()

	if *port == "" || *color == "" {
		panic("port and color are required")
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "ok"})
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{
			"ts":    time.Now().Format("2006-01-02T15:04:05.000Z07:00"),
			"color": *color,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	http.ListenAndServe(":"+*port, nil)
}
