FROM golang:1.22-alpine as builder
WORKDIR /app

COPY . .

RUN go build -ldflags="-w -s" -o stream-auth ./cmd/auth/main.go

FROM scratch
WORKDIR /app

COPY --from=builder /app/stream-auth .

CMD [ "./stream-auth" ]

EXPOSE 8080
