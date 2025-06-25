FROM golang:1.23.4 AS builder
WORKDIR /go/src/go-encryption-service
COPY . .
RUN go mod tidy
RUN go build -o encryption-service ./cmd

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/go-encryption-service/encryption-service .
EXPOSE 8080
CMD ["./encryption-service"]