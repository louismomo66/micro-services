FROM golang:1.24.1-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o listenerApp .

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/listenerApp /app

CMD ["/app/listenerApp"]
