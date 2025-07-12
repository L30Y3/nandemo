FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

WORKDIR /app/merchant-service
RUN go mod download

RUN go build -o merchant-service ./cmd/main.go

FROM gcr.io/distroless/base-debian12

COPY --from=builder /app/merchant-service/merchant-service /merchant-service

ENTRYPOINT ["/merchant-service"]
