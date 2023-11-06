FROM golang:alpine AS builder
ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /build
COPY . .
COPY ./configs/config.toml /build/config.toml

RUN go build -o app ./cmd/server/main.go
FROM alpine:3.18
WORKDIR /build

COPY --from=builder /build/app /build/app
COPY --from=builder /build/config.toml /build/config.toml


CMD ["/build/app", "--config-path=/build/config.toml"]