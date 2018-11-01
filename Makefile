.PHONE: build
build:
	go build --race -o bin/everify ./cmd/*.go

.PHONY: proto
proto:
	protoc --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. proto/e-verify.proto

.PHONY: test
test:
	go test -race ./...
