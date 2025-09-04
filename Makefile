.PHONY: setup test deploy-staging deploy-prod sst-dev clean fmt prep tidy build

setup:
	go install gotest.tools/gotestsum@latest

test:
	gotestsum --format testname -- -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

deploy-staging:
	npx sst@latest deploy --stage staging

deploy-prod:
	npx sst@latest deploy --stage prod

sst-dev:
	npx sst@latest dev

clean:
	rm -f coverage.out

fmt:
	go fmt ./...

tidy:
	go mod tidy
	cd lambda/functions/healthz && go mod tidy

build:
	go build ./...

prep: fmt
	go vet ./...
	go mod tidy
