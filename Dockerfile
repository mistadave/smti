FROM golang:1.22-alpine AS builder

LABEL maintainer="David Stäheli <mistrdave@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o smti .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/smti .

CMD ["./smti"]