# Golang Encryption and Key Management Service

## 🔐 Overview
This project implements a production-grade encryption and key management system written in Go. It offers:

- AES-256 encryption and decryption of data
- RESTful APIs for key lifecycle management
- In-memory key storage with thread-safe operations
- Swagger-based REST API documentation hosted on GitHub Pages
- Unit test coverage for critical components
- Docker support for containerized deployment

## 🧱 Architecture
```
+-------------+      POST/PUT/DELETE       +-------------------+
|             |  <----------------------> |                   |
|   Client    |                           | Golang REST API   |
|             |  ---------------------->  |                   |
+-------------+      Encrypt/Decrypt      +-------------------+
                                             |          |
                                             v          v
                                   +---------------+ +---------------+
                                   | Key Manager   | | Encryptor     |
                                   | (internal)    | | (internal)    |
                                   +---------------+ +---------------+
```

## 🗂 Project Structure
```
.
├── cmd/                        # Main application entrypoint
├── internal/
│   ├── api/                   # REST handlers and routes
│   ├── encryption/            # AES encryption logic
│   └── keymanager/            # Key lifecycle (create/delete/reset)
├── docs/
│   ├── index.html             # Swagger UI for GitHub Pages
│   └── openapi.yaml           # OpenAPI spec
├── test/                      # Unit tests
├── Dockerfile                 # Docker setup
├── go.mod / go.sum            # Go module definitions
```

## 📦 API Endpoints
| Method | Endpoint                  | Description           |
|--------|---------------------------|-----------------------|
| POST   | `/api/key`               | Create new key        |
| DELETE | `/api/key/{id}`          | Delete key by ID      |
| PUT    | `/api/key/{id}/reset`    | Reset key by ID       |
| POST   | `/api/encrypt`           | Encrypt data          |
| POST   | `/api/decrypt`           | Decrypt data          |

### 🔄 Payload Format (Encrypt/Decrypt)
```json
{
  "key_id": "abc123",
  "data": "Hello, World!"
}
```

## ✅ Running the Application
### Locally (with Go)
```bash
git clone https://github.com/aditya-bhuyan/go-encryption-service.git
cd go-encryption-service
go run cmd/main.go
```

### With Docker
```bash
docker build -t go-encryption-service .
docker run -p 8080:8080 go-encryption-service
```

## 🧪 Running Unit Tests
```bash
go test ./...
```

## 🛠️ Makefile Usage

This project includes a `Makefile` to streamline local development and CI integration:

| Command              | Description                           |
|----------------------|---------------------------------------|
| `make build`         | Build the server binary               |
| `make server`        | Run the server locally                |
| `make client`        | Run the Go-based API client           |
| `make test`          | Run all unit tests                    |
| `make coverage`      | Generate test coverage report         |
| `make docker-build`  | Build the Docker image                |
| `make docker-run`    | Run the Docker container (port 8080)  |

### 🧪 Test and Coverage

To run all unit tests with coverage:
```bash
make test
make coverage
````

You can also run manually:

```bash
go test -v -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

### 🐳 Docker Build & Deploy

To build and run the project in a container:

```bash
make docker-build
make docker-run
```

---




## 📚 API Documentation
- Swagger UI: [https://aditya-bhuyan.github.io/go-encryption-service](https://aditya-bhuyan.github.io/go-encryption-service)
- OpenAPI Spec: `docs/openapi.yaml`

## 📜 License
This project is licensed under the [GNU GPLv3 License](https://www.gnu.org/licenses/gpl-3.0.en.html).

## 👤 Author
Aditya Pratap Bhuyan  
GitHub: [@aditya-bhuyan](https://github.com/aditya-bhuyan)  
LinkedIn: [Aditya Bhuyan](https://linkedin.com/in/adityabhuyan)
