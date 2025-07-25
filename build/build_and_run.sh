#!/bin/sh

IMAGE_NAME=gogolook-assignment
CONTAINER_NAME=gogolook-assignment-container
DOCKERFILE=build/Dockerfile
PORT=8080

echo "Building Docker image..."
docker build -f $DOCKERFILE -t $IMAGE_NAME .

echo "Removing existing container (if any)..."
docker rm -f $CONTAINER_NAME 2>/dev/null || true

echo "Running Docker container..."
docker run -d -p $PORT:8080 --name $CONTAINER_NAME $IMAGE_NAME

echo "Container is running. Access: http://localhost:$PORT/" 