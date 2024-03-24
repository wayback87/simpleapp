// Simple application that exposes a couple of endpoints

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// main
func main() {

	http.HandleFunc("/simpleapp/version", appVersion)
	http.HandleFunc("/simpleapp/isReady", isReady)
	http.HandleFunc("/simpleapp/isAlive", isAlive)

	http.ListenAndServe(":8080", nil)
}

// Application version endpoint
func appVersion(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["Version"] = "24.3.0"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("%s", err)
	}
	w.Write(jsonResp)
}

// Liveness probe endpoint
func isAlive(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["Liveness probe"] = "200"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("%s", err)
	}
	w.Write(jsonResp)
}

// Readiness probe endpoint
func isReady(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["Readiness probe"] = "200"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("%s", err)
	}
	w.Write(jsonResp)
}
