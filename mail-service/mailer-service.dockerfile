FROM alpine:latest
# Build a minimal image
RUN mkdir /app
WORKDIR /app
COPY mailerApp /app
COPY templates ./templates

CMD ["/app/mailerApp"]

