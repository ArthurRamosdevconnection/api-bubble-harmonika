FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o api-bubble-harmonika .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/api-bubble-harmonika .

EXPOSE 3000

CMD ["./api-bubble-harmonika"]