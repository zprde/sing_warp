#!/bin/bash
set -e

PRIVATE_KEY=$(cat config.toml | grep private_key | sed "s/^.*'\(.*\)'$/\1/g")
echo "private_key: $PRIVATE_KEY"
echo "" > config.toml

BEST=$(cat result.csv | head -2 | tail -1 | sed -E 's/^(.*?),.*%.*$/\1/g')
BEST_IP=$(echo $BEST | sed -E 's/^(.*):(.*)$/\1/g')
BEST_PORT=$(echo $BEST | sed -E 's/^(.*):(.*)$/\2/g')

echo "best endpoint: $BEST_IP:$BEST_PORT"

cat config.tmp.json \
| sed "s/SERVER_IP/$BEST_IP/g" \
| sed "s/SERVER_PORT/$BEST_PORT/g" \
| sed "s/PRIVATE_KEY/$PRIVATE_KEY/g" \
> config.json
chmod 666 config.json
