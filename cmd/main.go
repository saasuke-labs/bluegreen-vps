package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "ok"})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		color := os.Getenv("COLOR")
		resp := map[string]string{
			"ts":    time.Now().Format("2006-01-02T15:04:05.000Z07:00"),
			"color": color,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":" + port, nil)
}