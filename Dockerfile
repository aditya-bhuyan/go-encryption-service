FROM golang:1.23.4 AS builder

WORKDIR /app
COPY . .

# Ensure replace directive works before tidy
RUN go mod edit -replace=github.com/aditya-bhuyan/go-encryption-service=.
RUN go mod tidy
RUN go build -o encryption-service ./cmd

FROM alpine
WORKDIR /app
COPY --from=builder /app/encryption-service .
EXPOSE 8080
CMD ["./encryption-service"]
