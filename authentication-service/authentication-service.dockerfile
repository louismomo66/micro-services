FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY authApp /app

RUN chmod +x /app/authApp

CMD ["/app/authApp"]
