# Redis Docker Setup

This Docker Compose configuration sets up a Redis database using the Bitnami Redis 6.2 image.

## Prerequisites

- Docker
- Docker Compose

## Configuration

### Environment Variables

Redis is configured to run without a password for development convenience.

## Usage

### Start Redis

```bash
docker-compose up -d
```

### Stop Redis

```bash
docker-compose down
```

### View Logs

```bash
docker-compose logs redis
```

### Connect to Redis CLI

```bash
docker-compose exec redis redis-cli
```

## Connection Details

- **Host**: localhost
- **Port**: 6379
- **Password**: None (no authentication required)
- **Database**: 0 (default)

## Data Persistence

Redis data is persisted in Docker volumes:
- `redis_data`: Contains the Redis database files
- `redis_config`: Contains Redis configuration files

## Network

The Redis container runs on a custom network `taskqueue-network` for better isolation.

## Security Notes

- This configuration runs Redis without authentication for development
- For production, consider enabling password protection
- The Redis port (6379) is exposed to the host for development convenience 