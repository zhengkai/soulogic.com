#!/bin/bash

DIR="$(dirname "$(readlink -f "$0")")" && cd "$DIR" || exit 1

sudo ln -sf "$DIR/nginx-dev.conf" /etc/nginx/vhost.d/200-soulogic-dev
sudo ln -sf "$DIR/nginx-test.conf" /etc/nginx/vhost.d/201-soulogic-test

mkdir -p /log/soulogic-dev
mkdir -p /log/soulogic-test
