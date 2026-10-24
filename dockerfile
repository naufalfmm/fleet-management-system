FROM golang:1.27-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/migrate migration/main.go

RUN go build -o /app/api cmd/internal/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/migrate .
COPY --from=builder /app/api .

COPY --from=builder /app/.env .env
COPY --from=builder /app/migration/sql ./migration/sql

CMD ["sh", "-c", "./migrate migrate && ./api"]