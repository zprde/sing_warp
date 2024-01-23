FROM golang:1.21.0-alpine AS builder

WORKDIR /src
COPY . .

RUN apk add --no-cache git build-base
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags="-s -w" \
    -tags "with_gvisor,with_quic,with_dhcp,with_wireguard,with_ech,with_utls,with_reality_server,with_acme,with_clash_api" \
    -o "/src/target/sing_warp"

FROM alpine:3.18.3

WORKDIR /
RUN set -ex \
    && apk upgrade \
    && apk add bash tzdata ca-certificates \
    && rm -rf /var/cache/apk/*
COPY --from=builder "/src/target/sing_warp" "/"


ENTRYPOINT ["/sing_warp"]
