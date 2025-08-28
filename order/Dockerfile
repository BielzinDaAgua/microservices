FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o order-service ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/order-service .
COPY --from=builder /app/config/config.yaml ./config/config.yaml

EXPOSE 50051

CMD ["./order-service"]