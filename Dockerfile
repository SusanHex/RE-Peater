FROM golang:alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

COPY ./src/go.mod .
COPY ./src/go.sum .
RUN go mod download

COPY ./src .

RUN go build -trimpath -ldflags="-s -w" -o repeater .

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /app/repeater /app/repeater

ENTRYPOINT ["/app/repeater"]