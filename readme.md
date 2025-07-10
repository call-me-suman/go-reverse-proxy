# Reverse Proxy in Go 🚀

A simple and modular reverse proxy server written in Go, with route-based forwarding, YAML configuration, and basic rate limiting.

---

## ✨ Features

- 🔁 Reverse proxy for multiple backend services
- 🛣️ Dynamic routing based on `/endpoint` path
- ⚙️ Config-driven via `data/config.yaml` using Viper
- ⛔ Rate limiting middleware (requests per IP)
- 🔐 Authentication module ready
- 🏥 Health check endpoint included
- 📦 Docker test services included via Makefile
- 🧱 Clean and extensible architecture

---

## 📂 Project Structure

```
.
├── cmd/
│   └── main.go              # Entry point
├── data/
│   └── config.yaml          # Configuration file
├── internal/
│   ├── auth/
│   │   └── auth.go          # Authentication logic
│   ├── configs/
│   │   └── config.go        # Configuration loader
│   ├── ratelimiter/
│   │   └── ratelimiter.go   # Rate limiting middleware
│   └── server/
│       ├── healthcheck.go   # Health check endpoint
│       ├── proxyhandler.go  # Reverse proxy handler
│       └── server.go        # Main server logic
├── .gitignore
├── go.mod                   # Go module file
├── go.sum                   # Go dependencies
├── Makefile                 # Dev scripts
└── README.md
```

---

## 🧾 Sample `config.yaml`

Place this file in the `data/` directory:

```yaml
server:
  host: "localhost"
  listen_port: "8069"

resources:
  - name: server1
    destination_url: "http://localhost:9001"
    endpoint: /server1
  - name: server2
    destination_url: "http://localhost:9002"
    endpoint: /server2
  - name: server3
    destination_url: "http://localhost:9003"
    endpoint: /server3

Rate:
  reqs: 5
  persec: 10
```

---

## 🧪 Run Locally

1. **Start Docker-based test services**
   ```bash
   make run-containers
   ```

2. **Run the proxy server**
   ```bash
   go run cmd/main.go
   ```

3. **Test using curl or Postman:**
   ```bash
   curl http://localhost:8069/server1/get
   ```

---

## 🧰 Dev Commands

```bash
make run-containers      # Start httpbin test backends
make stop                # Stop all running containers
make run-proxy-server    # Run the Go reverse proxy
```

---

## 🔜 Roadmap (Planned Features)

- 🔐 JWT Authentication Middleware
- 📜 Request & Response Logging
- ♻️ Retry & Circuit Breaker on backend failures
- 📈 Prometheus Metrics + `/health` endpoint
- 🧠 Round-robin load balancing between backends
- ⚡ Config hot-reloading on file change

---

## 🛠 Technologies Used

- **Go 1.21+**
- **net/http**
- **httputil.ReverseProxy**
- **Viper** for config management
- **Docker** (for backend testing)

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- Make (optional, for convenience commands)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/call-me-suman/go-reverse-proxy
   cd go-reverse-proxy
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. **Configure your `data/config.yaml` file according to your needs**

4. Follow the "Run Locally" steps above

---



## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

