FROM golang:1.22-alpine as builder
WORKDIR /app

COPY . .

RUN go build -o stream-auth ./cmd/auth/main.go

ENTRYPOINT [ "./stream-auth" ]
