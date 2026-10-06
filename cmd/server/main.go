package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var version = "v1.0.0"

type Response struct {
	Application string `json:"application"`
	Version     string `json:"version"`
	Hostname    string `json:"hostname"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	response := Response{
		Application: "go-k8s-demo",
		Version:     version,
		Hostname:    hostname,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, version)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/version", versionHandler)

	fmt.Println("server started :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
