package main

import (
	"log/slog"

	"github.com/aws/aws-lambda-go/lambda"
	healthzhandler "github.com/phasecurve/kept/handlers/healthz"
)

func main() {
	slog.Info("App started...")
	// http.HandleFunc("/healthz", healthz)
	// http.ListenAndServe(":8080", nil)
	lambda.Start(healthzhandler.Handler)
}
