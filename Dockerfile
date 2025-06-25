FROM golang:1.23.4 AS builder
WORKDIR /app
COPY . .
RUN go build -o encryption-service ./cmd

FROM alpine
WORKDIR /app
COPY --from=builder /app/encryption-service .
EXPOSE 8080
CMD ["./encryption-service"]