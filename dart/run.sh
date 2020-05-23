#!/bin/bash -ex

cd "$(dirname "$(readlink -f "$0")")" || exit 1

. ./common.inc.sh

NODEMODULE="${WWW}/client/node_modules"
if [ -d "$NODEMODULE" ]; then
	( cd "$NODEMODULE" && rm -rf ./*)
fi

cd "$WWW" || exit 1

git checkout .
git clean -df
git pull --rebase

git checkout "$BRANCH"

cd "${WWW}/client" && NG_CLI_ANALYTICS=ci npm i

"${WWW}/client/dist/build.sh" "$DOMAIN" "$NAME"
