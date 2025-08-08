#!/bin/bash

echo "Testing TaskQueue Engine API..."

# Test health endpoint
echo "Testing health endpoint..."
curl -s http://localhost:4000/health
echo -e "\n"

# Test API health endpoint
echo "Testing API health endpoint..."
curl -s http://localhost:4000/api/v1/health
echo -e "\n"

# Test job enqueue
echo "Testing job enqueue..."
curl -s -X POST http://localhost:4000/api/v1/jobs \
  -H "Content-Type: application/json" \
  -d '{
    "payload": "Test email content",
    "queue": "email",
    "type": "send_email",
    "max_retries": 3
  }'
echo -e "\n"

echo "API tests completed!"
