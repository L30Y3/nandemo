FROM golang:1.24 AS builder

# Set root context
WORKDIR /app

# Copy entire monorepo
COPY . .

# Go into subdir for this service
WORKDIR /app/order-service

# Use correct local go.mod (we're now in the right folder)
RUN go mod download
RUN go build -o order-service ./cmd/main.go

FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/order-service/order-service /order-service
ENTRYPOINT ["/order-service"]

