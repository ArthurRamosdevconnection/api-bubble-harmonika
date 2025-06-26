FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY src src
COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go

RUN go build -o api-bubble-harmonika .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/api-bubble-harmonika .

EXPOSE 3000

CMD ["./api-bubble-harmonika"]