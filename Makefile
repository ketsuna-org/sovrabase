.PHONY: build run test clean proto docker-build docker-run docker-stop docker-clean docker-logs

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server

test:
	go test ./...

test-orchestrator:
	go test ./internal/orchestrator -v

proto:
	protoc \
  --go_out=pkg --go_opt=paths=source_relative \
  --go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
  proto/*.proto

clean:
	rm -rf bin/

# Docker commands
docker-build:
	docker build -t sovrabase:latest .

docker-run:
	docker compose up -d

docker-stop:
	docker compose stop

docker-down:
	docker compose down

docker-clean:
	docker compose down -v
	docker rmi sovrabase:latest || true

docker-logs:
	docker compose logs -f sovrabase

docker-shell:
	docker compose exec sovrabase sh

# Setup config from example
setup-config:
	@if [ ! -f config.yaml ]; then \
		cp config.example.yaml config.yaml; \
		echo "✅ config.yaml created from example. Please edit it before running!"; \
	else \
		echo "⚠️  config.yaml already exists"; \
	fi

# Quick start (setup + build + run)
start: setup-config docker-build docker-run
	@echo "🚀 Sovrabase is starting..."
	@echo "📝 Check logs with: make docker-logs"
	@echo "🔍 Health check: curl http://localhost:8080/health"
