#!/bin/bash

DIR="$(dirname "$(readlink -f "$0")")" && cd "$DIR" || exit 1

. ./common.inc.sh

../singleton.sh "${DIR}/run.sh" "${LOG}/webhook-run.log" >> /log/singleton.txt 2>&1
