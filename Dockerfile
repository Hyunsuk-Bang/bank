#Build stage
FROM golang:1.20.1-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz --output file.tar.gz
RUN tar xf file.tar.gz

# Run Stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY wait-for.sh . 
COPY startup.sh .
COPY db/migration ./migration

EXPOSE 8080
ENTRYPOINT [ "/app/startup.sh" ]