#!/bin/bash
set -e

PRIVATE_KEY=$(cat config.toml | grep private_key | sed "s/^.*'\(.*\)'$/\1/g")
echo "private_key: $PRIVATE_KEY"

BEST=$(cat result.csv | head -2 | tail -1 | sed -E 's/^(.*?),.*%.*$/\1/g')
BEST_IP=$(echo $BEST | sed -E 's/^(.*):(.*)$/\1/g')
BEST_PORT=$(echo $BEST | sed -E 's/^(.*):(.*)$/\2/g')

echo "best endpoint: $BEST_IP:$BEST_PORT"

cat config.tmp.json \
| jq ".outbounds[0].server = \"$BEST_IP\"" \
| jq ".outbounds[0].server_port = $BEST_PORT" \
| jq ".outbounds[0].private_key = \"$PRIVATE_KEY\""  > config.json
chmod 666 config.json
