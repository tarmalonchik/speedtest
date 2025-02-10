# GRPC

## Proto

### Install dependencies

[Buf](https://docs.buf.build/installation)

#### Mac

```shell
brew install bufbuild/buf/buf
```

### Generate sdk

#### Install protoc dependencies

List of all actual dependencies in the file [internal/tools/tools.go](../../internal/tools/tools.go)

⚠️ Run command in project root:

```shell
go install \
  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc \
  google.golang.org/protobuf/cmd/protoc-gen-go \
  github.com/pseudomuto/protoc-gen-doc
```

#### Generate sdk

Run command in root:

```shell
make genProto
```

## GUI client

### BloomRPC

[BloomRPC](https://github.com/bloomrpc/bloomrpc)