package healthz

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHealthzShouldBeOk(t *testing.T) {
	ctx := context.Background()
	request := events.APIGatewayV2HTTPRequest{}

	resp, err := Handler(ctx, request)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected 200 but got %d", resp.StatusCode)
	}
}

func TestHandlerWritesHealthy(t *testing.T) {
	ctx := context.Background()
	request := events.APIGatewayV2HTTPRequest{}

	resp, err := Handler(ctx, request)
	if err != nil {
		t.Error(err)
	}

	if resp.Body != "Healthy" {
		t.Errorf("Expected 'Healthy' but got '%s'", resp.Body)
	}
}
