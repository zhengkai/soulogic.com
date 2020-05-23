#!/bin/bash -ex

DIR="$(dirname "$(readlink -f "$0")")" && cd "$DIR" || exit 1

. ./common.inc.sh

CONF="$DIR/nginx.conf"

cp ./nginx.tpl "$CONF"

sudo ln -sf "$CONF" "/etc/nginx/vhost.d/200-$NAME"

sed -i "s#%DOMAIN%#${DOMAIN}#g" "$CONF"
sed -i "s#%NAME%#${NAME}#g" "$CONF"
sed -i "s#%LOG%#${LOG}#g" "$CONF"
sed -i "s#%HOSTNAME%#${HOSTNAME,,}#g" "$CONF"
sed -i "s#%WH_PORT%#${WH_PORT}#g" "$CONF"
