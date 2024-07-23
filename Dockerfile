# Build stage
FROM golang:1.21-alpine3.19 AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o main main.go
RUN apk add --no-cache curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.19
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/main /app/migrate ./
COPY app.env start.sh wait-for.sh ./
COPY db/migration ./migration

EXPOSE 8080
ENTRYPOINT ["/app/start.sh"]
CMD ["/app/main"]
