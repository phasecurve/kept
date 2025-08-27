.PHONY: setup test run build deploy clean fmt prep

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

# Compile keptd binary
build:
	go build -o keptd ./cmd/keptd

# Deploy via SST
deploy:
	npx sst deploy --stage prod

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