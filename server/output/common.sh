BIN="plutus-server"

TYPE="${1:-dev}"

BIN_FILE="${TYPE}/${BIN}"
BIN_NEXT="${BIN_FILE}-next"

PID_FILE="${BIN_FILE}.pid"
LOG_FILE="${TYPE}/server.log"

EXE="${DIR}/${BIN_FILE}"
