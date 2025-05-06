# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy modules first to leverage Docker cache
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application binary statically
# Ensure it builds from the correct directory if main.go is in cmd/api
RUN cd cmd/api && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/authApp .

# Stage 2: Create the final minimal image
FROM alpine:latest

RUN mkdir /app
WORKDIR /app

# Copy the schema file and the built binary from the builder stage
COPY --from=builder /app/cmd/api/db.sql /app/
COPY --from=builder /app/authApp /app/

# Ensure the binary is executable
RUN chmod +x /app/authApp

# Command to run the application
CMD ["/app/authApp"]
