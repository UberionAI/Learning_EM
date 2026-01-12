ARG APP_VERSION="1.0.0"

# building
FROM golang:1.24-alpine AS builder

ARG APP_VERSION
ENV APP_VERSION=$APP_VERSION

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