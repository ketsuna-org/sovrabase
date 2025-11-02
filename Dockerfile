# Stage 1: Build
FROM golang:1.25.2-alpine AS builder

# Build arguments for multi-architecture support
ARG TARGETARCH

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
# CGO_ENABLED=0 for static binary
# -ldflags="-w -s" to strip debug info and reduce binary size
# TARGETARCH will be automatically set by Docker buildx (amd64 or arm64)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build \
    -ldflags="-w -s" \
    -o sovrabase \
    ./cmd/server/main.go

# Stage 2: Runtime
FROM scratch

# Copy CA certificates for HTTPS
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy timezone data
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy binary
COPY --from=builder /build/sovrabase /sovrabase

# Copy config file
COPY --from=builder /build/config.yaml /config.yaml

# Expose gRPC port (adjust based on your config)
EXPOSE 50051

# Run the application
ENTRYPOINT ["/sovrabase"]
