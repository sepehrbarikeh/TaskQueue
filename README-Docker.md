# TaskQueue Engine - Docker Setup

## 🐳 Docker Deployment Guide

### Prerequisites

- Docker Engine 20.10+
- Docker Compose 2.0+
- At least 2GB RAM available

### Quick Start

1. **Clone and navigate to the project:**
```bash
git clone <repository-url>
cd TaskQueue-Engine
```

2. **Copy environment variables:**
```bash
cp env.example .env
```

3. **Start the application:**
```bash
# For development (simple setup)
docker-compose -f docker-compose.dev.yml up -d

# For production (with monitoring)
docker-compose up -d

# For production with monitoring tools
docker-compose --profile monitoring up -d
```

### 🚀 Available Services

#### Core Services
- **taskqueue**: Main application (port 4000)
- **redis**: Redis database (port 6379)
- **postgres**: PostgreSQL database (port 5432)

#### Monitoring Services (optional)
- **redis-commander**: Redis web UI (port 8081)
- **pgadmin**: PostgreSQL web UI (port 8080)

### 📊 Access Points

| Service | URL | Username | Password |
|---------|-----|----------|----------|
| TaskQueue API | http://localhost:4000 | - | - |
| Redis Commander | http://localhost:8081 | - | - |
| pgAdmin | http://localhost:8080 | admin@taskqueue.com | admin123 |

### 🔧 Configuration

#### Environment Variables
Edit `.env` file to customize:

```bash
# PostgreSQL
POSTGRES_DB=taskqueue
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres123

# Redis
REDIS_PASSWORD=redis123

# pgAdmin
PGADMIN_EMAIL=admin@taskqueue.com
PGADMIN_PASSWORD=admin123
```

#### Application Config
Edit `config.yml` for application settings:

```yaml
server:
  host: "0.0.0.0"
  port: 4000

db:
  host: "postgres"
  user: "postgres"
  pass: "postgres123"
  name: "taskqueue"
  port: 5432

redis:
  host: "redis"
  port: 6379
  password: "redis123"
```

### 🛠️ Management Commands

#### Start Services
```bash
# Start all services
docker-compose up -d

# Start with specific profile
docker-compose --profile monitoring up -d

# Start development version
docker-compose -f docker-compose.dev.yml up -d
```

#### Stop Services
```bash
# Stop all services
docker-compose down

# Stop and remove volumes
docker-compose down -v
```

#### View Logs
```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f taskqueue
docker-compose logs -f redis
docker-compose logs -f postgres
```

#### Access Containers
```bash
# Access application container
docker-compose exec taskqueue sh

# Access Redis CLI
docker-compose exec redis redis-cli

# Access PostgreSQL
docker-compose exec postgres psql -U postgres -d taskqueue
```

### 🧪 Testing

#### Test API
```bash
# Health check
curl http://localhost:4000/health

# Enqueue a job
curl -X POST http://localhost:4000/api/v1/jobs \
  -H "Content-Type: application/json" \
  -d '{
    "payload": "Hello World",
    "queue": "send_email",
    "type": "send_email",
    "max_retries": 3
  }'
```

#### Test Redis
```bash
# Connect to Redis
docker-compose exec redis redis-cli

# List queues
KEYS queue:*

# Monitor Redis
MONITOR
```

#### Test PostgreSQL
```bash
# Connect to database
docker-compose exec postgres psql -U postgres -d taskqueue

# View job logs
SELECT * FROM job_logs ORDER BY created_at DESC LIMIT 10;

# View failed jobs
SELECT * FROM failed_jobs LIMIT 10;
```

### 📈 Monitoring

#### Redis Commander
- URL: http://localhost:8081
- Features: View Redis keys, monitor operations

#### pgAdmin
- URL: http://localhost:8080
- Username: admin@taskqueue.com
- Password: admin123
- Features: Database management, query execution

### 🔒 Security

#### Production Considerations
1. **Change default passwords** in `.env`
2. **Use secrets management** for sensitive data
3. **Enable SSL/TLS** for external access
4. **Restrict network access** using Docker networks
5. **Regular backups** of Redis and PostgreSQL data

#### Security Checklist
- [ ] Change default passwords
- [ ] Use strong passwords
- [ ] Enable Redis authentication
- [ ] Configure PostgreSQL access control
- [ ] Set up SSL certificates
- [ ] Configure firewall rules

### 🗄️ Data Persistence

#### Volumes
- `redis-data`: Redis persistent data
- `postgres-data`: PostgreSQL database files
- `pgadmin-data`: pgAdmin configuration

#### Backup
```bash
# Backup PostgreSQL
docker-compose exec postgres pg_dump -U postgres taskqueue > backup.sql

# Backup Redis
docker-compose exec redis redis-cli BGSAVE
```

### 🚨 Troubleshooting

#### Common Issues

1. **Port conflicts**
```bash
# Check what's using the port
lsof -i :4000
```

2. **Database connection issues**
```bash
# Check PostgreSQL logs
docker-compose logs postgres

# Test connection
docker-compose exec postgres pg_isready -U postgres
```

3. **Redis connection issues**
```bash
# Check Redis logs
docker-compose logs redis

# Test Redis
docker-compose exec redis redis-cli ping
```

4. **Application startup issues**
```bash
# Check application logs
docker-compose logs taskqueue

# Rebuild application
docker-compose build taskqueue
```

### 📝 Logs

#### Application Logs
```bash
# Real-time logs
docker-compose logs -f taskqueue

# Last 100 lines
docker-compose logs --tail=100 taskqueue
```

#### Database Logs
```bash
# PostgreSQL logs
docker-compose logs postgres

# Redis logs
docker-compose logs redis
```

### 🔄 Updates

#### Update Application
```bash
# Pull latest changes
git pull

# Rebuild and restart
docker-compose build taskqueue
docker-compose up -d taskqueue
```

#### Update Dependencies
```bash
# Update all services
docker-compose pull
docker-compose up -d
```

### 📊 Performance

#### Resource Limits
- **TaskQueue**: 512MB RAM, 0.5 CPU
- **Redis**: 256MB RAM, 0.25 CPU
- **PostgreSQL**: 512MB RAM, 0.5 CPU

#### Monitoring
```bash
# Resource usage
docker stats

# Container details
docker-compose ps
```

### 🎯 Best Practices

1. **Use specific image tags** instead of `latest`
2. **Implement health checks** for all services
3. **Use secrets** for sensitive data
4. **Regular backups** of persistent data
5. **Monitor resource usage**
6. **Implement logging** and monitoring
7. **Use production-ready** configurations
8. **Test thoroughly** before deployment

### 📚 Additional Resources

- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [Redis Documentation](https://redis.io/documentation)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Go Fiber Documentation](https://docs.gofiber.io/) 