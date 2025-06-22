FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o order-service ./cmd/main.go

FROM gcr.io/distroless/base-debian12

COPY --from=builder /app/order-service /order-service

ENTRYPOINT ["/order-service"]
