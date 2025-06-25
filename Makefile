# Makefile for Go Encryption and Key Management Service

BINARY_NAME=encryption-service
CLIENT=client/client.go

.PHONY: all build run server client docker-build docker-run clean

all: build

build:
	go build -o $(BINARY_NAME) ./cmd

run: build
	./$(BINARY_NAME)

server:
	go run ./cmd/main.go

client:
	go run $(CLIENT)

docker-build:
	docker build -t go-encryption-service .

docker-run:
	docker run -p 8080:8080 go-encryption-service

clean:
	rm -f $(BINARY_NAME)
