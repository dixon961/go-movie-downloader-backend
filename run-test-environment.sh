#!/bin/bash

echo "Stopping existing containers..."
docker-compose down

echo "Rebuilding Go application image..."
docker-compose build movie-downloader-go

echo "Starting all services..."
docker-compose up -d