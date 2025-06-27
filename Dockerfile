FROM golang:1.22-alpine AS builder

LABEL maintainer="David Stäheli <mistrdave@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o smti .

FROM alpine:latest AS runtime

WORKDIR /app

COPY --from=builder /app/smti .

CMD ["./smti"]