# Base Go image
FROM golang:1.23.3-alpine as builder

RUN mkdir /app

WORKDIR /app

# Copy everything into the container
COPY . .

# Ensure we build relative to the go.mod location
WORKDIR /app/cmd/api

RUN CGO_ENABLED=0 go build -o brokerApp .

# Build a minimal image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/cmd/api/brokerApp /app

CMD ["/app/brokerApp"]
