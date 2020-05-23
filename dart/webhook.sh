#!/bin/bash

DIR="$(dirname "$(readlink -f "$0")")" && cd "$DIR" || exit 1

. ./common.inc.sh

CONF="${DIR}/webhook.json"

cp ./webhook-tpl.json "$CONF"

sed -i "s#%NAME%#${NAME}#g" "$CONF"
sed -i "s#%SCRIPT%#${DIR}/webhook-call.sh#g" "$CONF"
sed -i "s#%DIR%#${DIR}#g" "$CONF"

SECRET=$(cat ./webhook-secret)
sed -i "s#%SECRET%#${SECRET}#g" "$CONF"

webhook -hooks "${DIR}/webhook.json" \
	-urlprefix webhook \
	-ip 127.0.0.1 \
	--port "$WH_PORT" \
	> "${LOG}/webhook.log"
