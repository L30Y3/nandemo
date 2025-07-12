FROM golang:1.24 AS builder

# Set root context
WORKDIR /app

# Copy entire monorepo
COPY . .

# Go into subdir for this service
WORKDIR /app/gateway-service

# Use correct local go.mod (we're now in the right folder)
RUN go mod download
RUN go build -o gateway-service ./cmd/main.go

FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/gateway-service/gateway-service /gateway-service
ENTRYPOINT ["/gateway-service"]

