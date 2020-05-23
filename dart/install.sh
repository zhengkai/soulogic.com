#!/bin/bash -ex

DIR="$(dirname "$(readlink -f "$0")")" && cd "$DIR" || exit 1

. ./common.inc.sh

mkdir -p "$LOG"

if [ ! -d "$WWW" ]; then
	git clone --depth 1 --b "$BRANCH" "$GIT" "$WWW"
fi

cd "$DIR" || exit 1
./nginx.sh

SECRET="$HOME/hide/github/webhook-soulogic"
if [ ! -f "$SECRET" ]; then
	>&2 echo secret not found: "$SECRET"
	exit 1
fi
cp "$SECRET" ./webhook-secret

if [ -d "$START" ]; then
	ln -sf "${DIR}/webhook.sh" "${START}/webhook-${NAME}.sh"
fi
