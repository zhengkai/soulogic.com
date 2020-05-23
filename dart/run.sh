#!/bin/bash -ex

DIR="$(dirname "$(readlink -f "$0")")" && cd "$DIR" || exit 1

. ./common.inc.sh

cd "$WWW" || exit 1

git checkout .
git clean -df
git pull --rebase

git checkout "$BRANCH"

PACKAGE_SUM="${DIR}/package-sum"
if ! md5sum -c "$PACKAGE_SUM" 2>/dev/null
then
	cd "$CLIENT"
	rm -rf node_modules || :
fi

cd "${WWW}/proto"
make client

"${WWW}/client/dist/build.sh" "$DOMAIN" "$NAME"

md5sum "${CLIENT}/package.json" > "$PACKAGE_SUM"
