# protoc-gen-go-mock

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
protoc --go_out=. \
       --go-grpc_out=. \
       --go-mock_out=. --go-mock_opt=framework=testify \
       petstore.proto
```
