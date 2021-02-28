FROM golang:1.14 as go-builder

WORKDIR /app

COPY . .
ARG GOPROXY=https://goproxy.io

RUN go mod download \
    && pwd && ls \
    && go install github.com/gobuffalo/packr/packr && \
    CGO_ENABLED=0 packr build -o rpsw-game


FROM alpine:latest as prod
COPY --from=go-builder /app/rpsw-game /app/rpsw-game

WORKDIR /app

ENTRYPOINT ["/app/rpsw-game"]