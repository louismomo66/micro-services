# # Base Go image
# FROM golang:1.23.3-alpine as builder

# RUN mkdir /app

# WORKDIR /app

# # Copy everything into the container
# COPY . .

# # Ensure we build relative to the go.mod location
# WORKDIR /app/cmd/api

# RUN CGO_ENABLED=0 go build -o brokerApp .

# # Build a minimal image
# FROM alpine:latest

# RUN mkdir /app
# RUN mv /app/cmd/api/brokerApp /app
# COPY --from=builder /app/cmd/api/brokerApp /app
# RUN ls -la /app/
# CMD ["/app/brokerApp"]
# Base Go image
# Base go image
# --- Builder Stage ---
   # broker-service/Dockerfile
# --- Stage 1: Build ---
  # Builder
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# STATIC BUILD HERE ⬇️
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o brokerApp ./cmd/api

# Runtime
FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY brokerApp /app

RUN chmod +x /app/brokerApp

EXPOSE 8083
CMD ["/app/brokerApp"]
