package main

import (
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("App started...")
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":8080", nil)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Healthy."))
}
