.PHONY: build run test clean proto

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server

test:
	go test ./...

proto:
	protoc \
  --go_out=pkg --go_opt=paths=source_relative \
  --go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
  proto/*.proto


clean:
	rm -rf bin/
