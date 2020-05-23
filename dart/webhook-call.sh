#!/bin/bash

DIR="$(dirname "$(readlink -f "$0")")" && cd "$DIR" || exit 1

../singleton.sh "${DIR}/run.sh" "${LOG}/webhook-run.log"
