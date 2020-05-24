#!/bin/bash -e

DIR=$(readlink -f "$0") && DIR=$(dirname "$DIR") && cd "$DIR"

. ./common.sh

echo building '"'$BIN'"'

DATE=$(TZ='Asia/Shanghai' date '+%Y-%m-%d %H:%M:%S')
GO_VERSION=$(go version)
GIT_VERSION=$(./git-hash.sh)

LDFLAGS="-X 'main.buildGoVersion=${GO_VERSION}' \
	-X 'main.buildTime=${DATE}' \
	-X 'main.buildType=${TYPE}' \
	-X 'main.buildHost=${HOSTNAME}' \
	-X 'main.buildGit=${GIT_VERSION}'"

cd ..
go build \
	-ldflags "$LDFLAGS" \
	-o "${DIR}/${BIN_NEXT}" \
	./*.go \
	2> >(while read -r line; do echo -e "\e[38;2;255;45;45;48;2;10;10;10m$line\e[0m" >&2; done)
