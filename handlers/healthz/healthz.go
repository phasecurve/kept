package healthz

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       "Healthy",
	}, nil
}
