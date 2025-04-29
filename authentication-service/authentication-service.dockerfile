# Stage 1: Build your authApp
FROM golang:1.23.3-alpine AS builder

WORKDIR /app
COPY . .   
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o authApp ./cmd/api

# Stage 2: Create a minimal final image
FROM alpine:latest

RUN mkdir /app
COPY --from=builder /app/authApp /app/

# Run the compiled binary by default
CMD ["/app/authApp"]
