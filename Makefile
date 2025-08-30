.PHONY: setup test run build deploy deploy-staging deploy-prod sst-dev clean fmt prep

setup:
	go install gotest.tools/gotestsum@latest

test:
	gotestsum --format testname -- -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

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

clean:
	rm -f keptd coverage.out

fmt:
	go fmt ./...

prep: fmt
	go vet ./...
	go mod tidy
