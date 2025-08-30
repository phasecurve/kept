.PHONY: setup test run build deploy deploy-staging deploy-prod sst-dev clean fmt prep

# Install required local tooling
setup:
	go install gotest.tools/gotestsum@latest

# Run test suite with coverage
test:
	gotestsum --format testname -- -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

# Start the service locally
run:
	go run ./cmd/keptd

build:
	cd functions/healthz && go build -o bootstrap .

deploy:
	npx sst@latest deploy --stage prod

deploy-staging:
	npx sst@latest deploy --stage staging

deploy-prod:
	npx sst@latest deploy --stage prod

sst-dev:
	npx sst@latest dev

# Clean build artifacts
clean:
	rm -f keptd coverage.out

# Format Go code
fmt:
	go fmt ./...

# Prepare code for commit
prep: fmt
	go vet ./...
	go mod tidy