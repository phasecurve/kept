package main

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"

	healthzhandler "github.com/phasecurve/kept/handlers/healthz"
)

func main() {
	consoleExporter, _ := stdoutmetric.New()
	mp := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(consoleExporter,
			metric.WithInterval(5*time.Second))),
	)
	otel.SetMeterProvider(mp)

	meter := otel.Meter("healthz")
	invokeCounter, _ := meter.Int64Counter("healthz_invokes")

	wrappedHandler := func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		invokeCounter.Add(ctx, 1)
		return healthzhandler.Handler(ctx, request)
	}

	instrumentedHandler := otellambda.InstrumentHandler(wrappedHandler)

	lambda.Start(instrumentedHandler)
}
