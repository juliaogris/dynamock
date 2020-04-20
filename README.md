# Dynamock

![CI](https://github.com/foxygoat/dynamock/workflows/ci/badge.svg?branch=master)
[![Godoc](https://img.shields.io/badge/godoc-ref-blue)](https://pkg.go.dev/foxygo.at/dynamock)
[![Slack chat](https://img.shields.io/badge/slack-gophers-795679?logo=slack)](https://gophers.slack.com/app_redirect?channel=foxygoat)

File-based
[DynamoDB](https://pkg.go.dev/github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface)
fake. Use in tests and local development.

### Development

-   Pre-requisites: [go](https://golang.org/doc/go1.14),
    [golangci-lint](https://github.com/golangci/golangci-lint/releases/tag/v1.24.0),
    GNU make
-   Build with `make`
-   View build options with `make help`
