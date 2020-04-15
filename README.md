# Dynafake

![CI](https://github.com/foxygoat/dynafake/workflows/ci/badge.svg?branch=master)
[![Godoc](https://img.shields.io/badge/godoc-ref-blue)](https://pkg.go.dev/foxygo.at/dynafake)
[![Slack chat](https://img.shields.io/badge/slack-gophers-795679?logo=slack)](https://app.slack.com/client/T029RQSE6/C011LKH21HC)

File-based
[DynamoDB](https://pkg.go.dev/github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface)
fake. Use in tests and local development.

### Development

-   Pre-requisites: [go](https://golang.org/doc/go1.14),
    [golangci-lint](https://github.com/golangci/golangci-lint/releases/tag/v1.24.0),
    GNU make
-   Build with `make`
-   View build options with `make help`