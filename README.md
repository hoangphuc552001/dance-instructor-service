# Go Dance Instructor Project

## 🚀 Overview
This is a simple Go project about a dance instructor. The project is built using Go and can be run locally or in a Docker container.

---

## 🛠️ Prerequisites
Ensure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.21 or later)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

---

## 🚢 Running the Project

### Option 1️⃣: Run Locally (Without Docker)

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd <project-directory>
   cp .env.example .env
   ```
2. **Install Go:**
   Follow the official [Go installation guide](https://golang.org/doc/install).
3. **Install dependencies:**
   ```bash
   go mod download + go mod tidy
   ```
4. **Run the app:**
   ```bash
   go run ./cmd/main.go
   ```
5. **Access the app:**
   The app will start at: [http://localhost:8080](http://localhost:8080)

---

### Option 2️⃣: Run with Docker

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd <project-directory>
   cp .env.example .env
   ```
2. **Build the Docker image:**
   ```bash
   docker-compose build
   ```
3. **Start the container:**
   ```bash
   docker-compose up -d
   ```
4. **Check logs:**
   ```bash
   docker-compose logs -f
   ```
5. **Stop the container:**
   ```bash
   docker-compose down
   ```

The app will be accessible at [http://localhost:8080](http://localhost:8080)

---

## ✅ Testing

To run tests, use the following command:
```bash
go test ./...
```

Or run tests in a specific folder (e.g., `test`):
```bash
go test ./test
```

---

## 🏗️ Building the Project

To build the Go binary:
```bash
go build -o main ./cmd/main.go
```

The binary will be generated in the project root and can be run directly:
```bash
./main
```

---

## 📚 Environment Variables
You can configure the app using environment variables. Modify the `docker-compose.yml` or create a `.env` file.

Example:
```dotenv
HOST=localhost
PORT=8080
INTERNAL_SECRET=token
```

---
## Running Swagger

To view and interact with the API documentation, start the server and open the following link in your browser:

[Swagger UI](http://localhost:8080/swagger/index.html#)

Ensure the server is running on port 8080. If you have a different port configuration, update the URL accordingly.

## Golangci-lint

`golangci-lint` is a fast, configurable, and extensible linter for Go. It helps catch bugs, style issues, and potential errors in your code.

### Installation

To install `golangci-lint`, run the following command:

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

Ensure the `$GOPATH/bin` directory is in your system's `PATH` so the `golangci-lint` command is accessible.

### Usage

To run the linter on your project, use:

```bash
golangci-lint run
```

This will scan your code for issues and display them in the terminal.

### Custom Configuration

You can customize the linting rules by creating a `.golangci.yml` file in the root of your project. Here’s a basic example:

```yaml
linters:
  enable:
    - errcheck
    - gosimple
    - govet
    - staticcheck
    - unused

run:
  timeout: 5m
```

For more configuration options, check the official documentation: [Golangci-lint Docs](https://golangci-lint.run/)

### Running with CI/CD

Integrate `golangci-lint` into your CI/CD pipeline to ensure consistent code quality. Add a step in your pipeline configuration to run:

```bash
golangci-lint run --timeout=5m
```

This ensures linting checks are enforced before merging code.

---

Need help tweaking the config or adding it to your CI setup? Let me know!
## 🔥 Contributing
1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

---

## 📜 License
This project is licensed under the MIT License.

---

## ✨ Contact
For any questions or feedback, please reach out via email or open an issue on this repo.

---

Let's build something awesome! 🚀

