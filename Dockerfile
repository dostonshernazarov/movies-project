# Build stage
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o movies-app .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/movies-app .

COPY --from=builder /app/.env .

EXPOSE 8060

CMD ["./movies-app"]