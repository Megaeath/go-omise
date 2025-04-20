# 🧧 Go-Tamboon: Digital Donation Processing System

**Go-Tamboon** is a modern, scalable donation processing system built with Go, Kafka, MongoDB, and Redis. Inspired by traditional Thai donation practices, this project simulates processing encrypted donation data, emphasizing concurrency, observability, and clean architecture principles.

---

## 📐 Architecture Overview

```plaintext
+-------------+       Kafka        +-------------+        MongoDB
|             |  ─────────────▶   |             |   ─────────────▶
|    CLI      |                   |   Worker    |               Log DB
|  (CSV Load) |◀─────────────┐    | (Consumer)  |◀────────┐
+-------------+              │    +-------------+         │
                             ▼                          ▲
                         Kafka Topic                MongoDB
                           "charge-topic"             (Charge Request Log)
                             ▲                          ▲
+-------------+              │                          │
|             |  REST API    │                          │
|   API       |─────────────▶|                          │
|   Server    |    Redis (Rate Limit)                   │
+-------------+                                          
```

---

## 🚀 Getting Started

### Prerequisites

Ensure you have the following installed:
- [Go](https://golang.org/dl/) (version 1.18 or later)
- [Kafka](https://kafka.apache.org/quickstart)
- [MongoDB](https://www.mongodb.com/try/download/community)
- [Redis](https://redis.io/download)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/go-tamboon.git
   cd go-tamboon
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up environment variables:
   Create a `.env` file in the root directory and configure the following:
   ```env
   KAFKA_BROKER=localhost:9092
   MONGO_URI=mongodb://localhost:27017
   REDIS_ADDR=localhost:6379
   ```

### Running the Project

1. Start Kafka, MongoDB, and Redis services.

2. Run the API server:
   ```bash
   go run cmd/api/main.go
   ```

3. Run the worker:
   ```bash
   go run cmd/worker/main.go
   ```

4. Load sample donation data via the CLI:
   ```bash
   go run cmd/cli/main.go ../data/fng.csv 
   ```

---

## 🧪 Testing

Run unit tests with coverage:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```
