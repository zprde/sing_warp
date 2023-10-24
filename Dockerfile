FROM golang:1.21.0-alpine AS builder

WORKDIR /src
COPY . .

RUN apk add --no-cache git && \
    go mod download && \
    CGO_ENABLED=0 go build -ldflags="-s -w" -o "sing_warp"

FROM alpine:3.18.3

WORKDIR /

COPY --from=builder "/src/sing_warp" "/"


ENTRYPOINT ["/sing_warp"]
