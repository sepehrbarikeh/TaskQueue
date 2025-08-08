# TaskQueue Engine 🚀

A robust, high-performance task queue system built with Go, featuring Redis for job storage, PostgreSQL for logging, and a RESTful API for job management. Perfect for microservices architecture and distributed systems.

[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](docker-compose.yml)

## ✨ Features

- **🎯 Job Queue Management**: Enqueue, process, and track jobs with retry mechanisms
- **⚡ Redis Storage**: Fast and reliable job storage using Redis
- **📊 PostgreSQL Logging**: Persistent job execution logs with detailed analytics
- **🌐 RESTful API**: Simple HTTP API for job management
- **🛡️ Graceful Shutdown**: Proper cleanup on application termination
- **👥 Worker Pool**: Configurable number of worker goroutines
- **🔄 Retry Logic**: Exponential backoff for failed jobs
- **⏰ Scheduled Jobs**: Support for delayed job execution
- **📈 Monitoring**: Built-in health checks and metrics
- **🐳 Docker Ready**: Complete Docker setup with multi-stage builds
- **🔒 Security**: Input validation and error handling
- **📝 Logging**: Comprehensive logging for debugging

## Architecture

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   Client    │───▶│   Server    │───▶│   Redis     │
│   (HTTP)    │    │   (Fiber)   │    │   (Queue)   │
└─────────────┘    └─────────────┘    └─────────────┘
                           │
                           ▼
                   ┌─────────────┐
                   │  Dispatcher │
                   │  (Workers)  │
                   └─────────────┘
                           │
                           ▼
                   ┌─────────────┐
                   │ PostgreSQL  │
                   │   (Logs)    │
                   └─────────────┘
```

## 📋 Prerequisites

- **Go 1.24+** - Programming language
- **Redis Server** - Job queue storage
- **PostgreSQL Database** - Job execution logs
- **Docker & Docker Compose** (optional) - Containerization

## 🚀 Quick Start

### Option 1: Docker (Recommended)

```bash
# Clone the repository
git clone <repository-url>
cd TaskQueue-Engine

# Start with Docker
docker-compose -f docker-compose.dev.yml up -d

# Test the API
curl http://localhost:4000/health
```

### Option 2: Local Development

1. **Clone the repository:**
```bash
git clone <repository-url>
cd TaskQueue-Engine
```

2. **Install dependencies:**
```bash
go mod tidy
```

3. **Configure the application:**
Edit `config.yml` with your settings:
```yaml
name: "TaskQueue Engine"
version: 1.0

server:
  host: "localhost"
  port: 4000

db:
  host: "localhost"
  user: "postgres"
  pass: "your_password"
  name: "taskqueue"
  port: 5432

redis:
  port: 6379
  host: localhost
  db: 0
  password: ""
```

4. **Build and run:**
```bash
go build -o taskqueue .
./taskqueue
```

## 📡 API Reference

### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Basic health check |
| `GET` | `/api/v1/health` | Detailed health check |
| `POST` | `/api/v1/jobs` | Enqueue a new job |

### Enqueue a Job

**Endpoint:** `POST /api/v1/jobs`

**Request Body:**
```json
{
  "payload": "Hello World",
  "queue": "send_email",
  "type": "send_email",
  "run_at": "2024-01-01T10:00:00Z",
  "max_retries": 3
}
```

**Response:**
```json
{
  "message": "job enqueued successfully",
  "id": "uuid-string",
  "queue": "send_email"
}
```

### Health Check

**Endpoint:** `GET /api/v1/health`

**Response:**
```json
{
  "status": "healthy",
  "time": "2024-01-01T10:00:00Z"
}
```

### Error Responses

```json
{
  "error": "queue name is required"
}
```

```json
{
  "error": "payload is required"
}
```

```json
{
  "error": "job type is required"
}
```

## 🎯 Job Types

The system supports the following job types:

| Type | Description | Default Retries |
|------|-------------|-----------------|
| `send_email` | Simulates sending an email | 3 |
| `process_image` | Simulates image processing | 3 |
| `write_log` | Simulates log writing | 3 |

### Job Properties

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `payload` | string | ✅ | Job data |
| `queue` | string | ✅ | Queue name |
| `type` | string | ✅ | Job type |
| `run_at` | datetime | ❌ | Scheduled execution time |
| `max_retries` | integer | ❌ | Maximum retry attempts |

## ⚙️ Configuration

The application uses a YAML configuration file (`config.yml`) with the following sections:

### Server Configuration
```yaml
server:
  host: "0.0.0.0"  # Server host
  port: 4000        # Server port
```

### Database Configuration
```yaml
db:
  host: "postgres"      # Database host
  user: "postgres"      # Database user
  pass: "postgres123"   # Database password
  name: "taskqueue"     # Database name
  port: 5432           # Database port
```

### Redis Configuration
```yaml
redis:
  host: "redis"        # Redis host
  port: 6379          # Redis port
  db: 0               # Redis database number
  password: "redis123" # Redis password
```

## 🐳 Docker Support

The project includes comprehensive Docker support with multiple deployment options:

### Quick Start with Docker
```bash
# Development setup
docker-compose -f docker-compose.dev.yml up -d

# Production setup
docker-compose up -d

# With monitoring tools
docker-compose --profile monitoring up -d
```

### Available Services
- **taskqueue**: Main application
- **redis**: Job queue storage
- **postgres**: Job execution logs
- **redis-commander**: Redis web UI (optional)
- **pgadmin**: PostgreSQL web UI (optional)

### Access Points
- **API**: http://localhost:4000
- **Redis Commander**: http://localhost:8081
- **pgAdmin**: http://localhost:8080

For detailed Docker documentation, see [README-Docker.md](README-Docker.md).

## 🛠️ Development

### Project Structure
```
TaskQueue Engine/
├── config/              # Configuration management
├── pkg/                # Core packages
│   ├── queue/          # Job queue models
│   └── worker/         # Worker and dispatcher logic
├── repository/         # Data access layer
│   ├── postgres/       # PostgreSQL operations
│   └── redis/          # Redis operations
├── server/             # HTTP server and API
├── docker-compose.yml  # Production Docker setup
├── docker-compose.dev.yml # Development Docker setup
├── Dockerfile          # Multi-stage Docker build
├── main.go             # Application entry point
├── config.yml          # Configuration file
├── env.example         # Environment variables template
└── README-Docker.md    # Docker documentation
```

### Development Commands

```bash
# Build the application
go build -o taskqueue .

# Run tests
go test ./...

# Run with hot reload (if using air)
air

# Check code quality
go vet ./...
go fmt ./...

# Generate documentation
godoc -http=:6060
```

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test ./pkg/worker -v
```

### Code Quality

```bash
# Format code
go fmt ./...

# Check for issues
go vet ./...

# Run linter (if using golangci-lint)
golangci-lint run
```

## 📊 Monitoring & Logs

### Health Checks
- **Application**: `GET /health` - Basic health check
- **API**: `GET /api/v1/health` - Detailed health status

### Logs
The application provides comprehensive logging:
- **Application logs**: Job processing, errors, and system events
- **Database logs**: Job execution history and analytics
- **Redis logs**: Queue operations and performance metrics

### Metrics
- Job success/failure rates
- Queue processing times
- Worker utilization
- Retry statistics

## 🔧 Troubleshooting

### Common Issues

1. **Connection refused to Redis/PostgreSQL**
   - Check if services are running
   - Verify connection settings in `config.yml`
   - Check network connectivity

2. **Jobs not being processed**
   - Verify queue name matches dispatcher configuration
   - Check worker count and job channel capacity
   - Review application logs for errors

3. **High memory usage**
   - Monitor Redis memory usage
   - Check for memory leaks in job processing
   - Adjust worker count if necessary

### Debug Commands

```bash
# Check application status
curl http://localhost:4000/health

# View Redis queues
docker-compose exec redis redis-cli KEYS queue:*

# Check PostgreSQL logs
docker-compose exec postgres psql -U postgres -d taskqueue -c "SELECT * FROM job_logs ORDER BY created_at DESC LIMIT 10;"
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines
- Follow Go coding standards
- Add tests for new features
- Update documentation
- Use conventional commit messages

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Go Fiber](https://gofiber.io/) - Web framework
- [Redis](https://redis.io/) - In-memory data store
- [PostgreSQL](https://www.postgresql.org/) - Database
- [GORM](https://gorm.io/) - ORM library

## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/your-repo/issues)
- **Documentation**: [README-Docker.md](README-Docker.md)
- **Email**: support@taskqueue.com
