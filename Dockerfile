FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o hla_finder ./cmd/app/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/hla_finder .

EXPOSE 8080

CMD ["./hla_finder"]