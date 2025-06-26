FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o merchant-service ./cmd/main.go

FROM gcr.io/distroless/base-debian12

COPY --from=builder /app/merchant-service /merchant-service

ENTRYPOINT ["/merchant-service"]
