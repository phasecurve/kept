package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/sdk/metric"

	healthzhandler "github.com/phasecurve/kept/lambda/handlers/healthz"
)

func main() {
	ctx := context.Background()
	
	otlpEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	
	fmt.Printf("Using base OTLP endpoint: %s\n", otlpEndpoint)
	fmt.Printf("OTLP library will auto-append /v1/metrics\n")
	
	otlpExporter, err := otlpmetrichttp.New(ctx,
		otlpmetrichttp.WithEndpointURL(fmt.Sprintf("http://%s:4318", otlpEndpoint)),
	)
	if err != nil {
		panic(err)
	}
	
	mp := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(otlpExporter,
			metric.WithInterval(5*time.Second))),
	)
	otel.SetMeterProvider(mp)

	meter := otel.Meter("healthz")
	invokeCounter, _ := meter.Int64Counter("healthz_invokes")

	wrappedHandler := func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		invokeCounter.Add(ctx, 1)
		
		response, err := healthzhandler.Handler(ctx, request)
		
		fmt.Printf("Flushing metrics before Lambda exit\n")
		if err := mp.ForceFlush(ctx); err != nil {
			fmt.Printf("Error flushing metrics: %v\n", err)
		}
		
		return response, err
	}

	instrumentedHandler := otellambda.InstrumentHandler(wrappedHandler)

	lambda.Start(instrumentedHandler)
}
