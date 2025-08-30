.PHONY: setup test deploy-staging deploy-prod sst-dev clean fmt prep

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

prep: fmt
	go vet ./...
	go mod tidy
