# Kept
## Purpose

Kept is a lightweight, API-first knowledge graph service.
This initial version exists only to prove the delivery pipeline: one endpoint deployed and verifiable.

## Contracts
### Application

 - Lambda function exposes GET /healthz via API Gateway.

 - Returns status 200 and body "Healthy".

### Tests

 - Tests are executed via gotestsum.

 - make test must pass without network access.

 - At least one test asserts that /healthz returns 200 OK.

### Makefile

 - make setup installs required local tooling.

 - make test runs the test suite with coverage.

 - make deploy-staging deploys to staging environment.
 
 - make deploy-prod deploys to production environment.

### CI

 - Triggered on push and pull_request.

 - Runs with Go 1.24.x.

 - Executes make test.

 - Passes only if all tests succeed.

### CD

 - Triggered on push to main.

 - Authenticates to AWS via GitHub OIDC.

 - Deploys to region eu-west-2 using SST.

 - Workflow logs emit the API URL.

 - Calling GET {ApiUrl}/healthz via deployed Lambda must return 200 ok.

### SST stack

 - App name: kept.

 - Region: eu-west-2.

 - One HTTP API resource routes GET /healthz to the Go Lambda handler.

 - CloudFormation outputs include ApiUrl.

### Configuration

 - No secrets stored in repo.

 - CI requires no secrets.

 - CD requires an AWS role ARN supplied via repo/org secrets.

## Verification

From a fresh clone:

1. make setup
2. make test → green.
3. make deploy-staging → deploys and tests staging environment.