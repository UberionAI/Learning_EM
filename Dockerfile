# building
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o app main.go

# starting
FROM alpine:3.19

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/app /app/app

CMD ["/app/app"]