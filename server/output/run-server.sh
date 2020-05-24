#!/bin/bash -e

DIR=$(readlink -f "$0") && DIR=$(dirname "$DIR") && cd "$DIR" || exit 1

. ./common.sh

"${DIR}/build-server.sh" "$TYPE"
echo 'done'

echo
"${DIR}/stop-server.sh" "$TYPE" || :
echo 'done'

mv "${DIR}/${BIN_NEXT}" "${DIR}/${BIN_FILE}"

echo
"${DIR}/start-server.sh" "$TYPE"
echo 'done'
