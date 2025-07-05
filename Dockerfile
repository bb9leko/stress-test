FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.mod ./
RUN go mod download

COPY . .

RUN go build -o stress-tester ./cmd/main.go

# Imagem final
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/stress-tester .

ENTRYPOINT ["./stress-tester"]