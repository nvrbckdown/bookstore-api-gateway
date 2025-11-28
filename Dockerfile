FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy the source code
COPY . .

# Remove any existing go.mod and go.sum if they exist
RUN rm -f go.mod go.sum

# Initialize Go modules and download dependencies
RUN go mod init github.com/nvrbckdown/bookstore-api-gateway
RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-contrib/cors@v1.4.0
RUN go get github.com/joho/godotenv
RUN go get github.com/sirupsen/logrus
RUN go get github.com/golang-jwt/jwt/v5
RUN go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api-gateway ./cmd/server

# Use a small image for the final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/api-gateway .

# Create a default .env file if one doesn't exist
# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./api-gateway"]
