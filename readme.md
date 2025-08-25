# Kept
## Purpose

Kept is a lightweight, API-first knowledge graph service.
This initial version exists only to prove the delivery pipeline: one endpoint deployed and verifiable.

## Contracts
### Application

 - Exposes GET /healthz.

 - eturns status 200 and body ok.

 - Runs on :${KEPT_PORT} (defaults to 8080).

 - Starts cleanly, logs “listening :<port>”, exits gracefully on SIGTERM.

### Tests

 - Tests are executed via gotestsum.

 - go test ./... must pass without network access.

 - At least one test asserts that /healthz returns 200 OK.

### Makefile

 - make setup installs required local tooling.

 - make test runs the test suite with coverage.

 - make run starts the service locally.

 - make deploy invokes SST to deploy.

### CI

 - Triggered on push and pull_request.

 - Runs with Go 1.25.x.

 - Executes gotestsum … ./... and go build ./cmd/keptd.

 - Passes only if all tests succeed.

### CD

 - Triggered on push to main.

 - Authenticates to AWS via GitHub OIDC.

 - Deploys to region eu-west-2 using SST.

 - Workflow logs emit the API URL.

 - Calling GET {ApiUrl}/healthz must return 200 ok.

### SST stack

 - App name: kept.

 - Region: eu-west-2.

 - One HTTP API resource routes GET /healthz to the Go handler.

 - CloudFormation outputs include ApiUrl.

### Configuration

 - No secrets stored in repo.

 - Local defaults allow running without configuration.

 - CI requires no secrets.

 - CD requires an AWS role ARN supplied via repo/org secrets.

## Verification

From a fresh clone:

1. make setup
2. make test → green.
3. make run and in another terminal: