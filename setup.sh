#!/bin/bash
# setup.sh - Initialize Go modules and download dependencies

# Initialize Go module
go mod init github.com/yourusername/bookstore-api-gateway

# Get required packages
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get github.com/joho/godotenv
go get github.com/sirupsen/logrus
go get github.com/golang-jwt/jwt/v5

# Add all packages from source files
go mod tidy

echo "Go modules initialized and dependencies downloaded."
echo "You can now build the application with: go build -o api-gateway ./cmd/server"