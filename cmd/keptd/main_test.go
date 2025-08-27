package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthzShouldBeOk(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(healthz))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 but got %d", resp.StatusCode)
	}
}

func TestHandlerWritesHealthy(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(healthz))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	contentBytes, err := io.ReadAll(resp.Body)
	content := string(contentBytes)
	if err != nil {
		t.Error(err)
	}
	if content != "Healthy." {
		t.Errorf("Expected 'Healthy.' but got '%s'", content)
	}

}
