# Here's how you can run the `Makefile` on **Windows**, **macOS**, and **Linux**:

---

### ✅ **For macOS & Linux**

These systems typically come with `make` pre-installed.

```bash
make build        # Compile the Go binary
make server       # Run the server
make client       # Run the test client
make test         # Run all unit tests
make coverage     # Generate coverage report
make docker-build # Build Docker image
make docker-run   # Run Docker container
```

---

### ✅ **For Windows**

#### Option 1: Use Git Bash (recommended)

If you have **Git for Windows** installed, open **Git Bash** and run:

```bash
make build
```

#### Option 2: Use Windows Subsystem for Linux (WSL)

1. Install [WSL](https://learn.microsoft.com/en-us/windows/wsl/install)
2. Open a WSL terminal
3. Navigate to your project folder
4. Run `make build`, `make server`, etc.

#### Option 3: Use Make with MinGW or Chocolatey

1. Install [Chocolatey](https://chocolatey.org/)
2. Run: `choco install make`
3. Open PowerShell or CMD and run:

```bash
make build
```

---

### 🧰 Alternative for Native Windows Without `make`

If you don't want to install `make`, just run the equivalent Go commands:

```bash
go build -o encryption-service ./cmd
go run ./cmd/main.go
go run ./client/client.go
go test ./...
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```


