# protoc-gen-go-mock
[![Go Reference](https://pkg.go.dev/badge/github.com/kw510/protoc-gen-go-mock.svg)](https://pkg.go.dev/github.com/kw510/protoc-gen-go-mock)
[![Go Report Card](https://goreportcard.com/badge/github.com/kw510/protoc-gen-go-mock)](https://goreportcard.com/report/github.com/kw510/protoc-gen-go-mock)

This tool generates Go language bindings of services in protobuf definition files for mocking gRPC.

## Installation

```
go install github.com/kw510/protoc-gen-go-mock@latest
```

Also required:

- [protoc](https://github.com/protocolbuffers/protobuf)
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)

## Usage

### With `protoc`

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc --go_out=.examples \
       --go-grpc_out=.examples \
       --go-mock_out=.examples \
       .examples/petstore.proto
```

## Building locally
```shell
mkdir .build
export PATH=$PATH:$(pwd)/.build
go build -o ./.build
```
