FROM alpine

RUN set -ex \
    && apk upgrade \
    && apk add bash tzdata ca-certificates curl jq \
    && rm -rf /var/cache/apk/*