#!/bin/bash -x

TZ='Asia/Shanghai'
TIME=$(date '+%Y-%m-%d %H:%M:%S')

echo "start $TIME"
echo "$TIME" >> /log/webhook.txt

# singleton.sh "脚本名" "文件锁名"

# 确保最后一次更新的 webhook 总能被触发
# 可能会多跑一次，但绝不会漏跑

SCRIPT="$1"
if [ ! -x "$SCRIPT" ]; then
	>&2 echo "can not run script $SCRIPT"
	exit 1
fi

DIR=$(dirname "$SCRIPT")
LOCK="${DIR}/lock"
LOCK_MAIN="${LOCK}.main"
LOCK_ADD="${LOCK}.add"

echo
echo "        script: $SCRIPT"
echo "lock file main: $LOCK_MAIN"
echo "           add: $LOCK_ADD"
echo

exec 100>"$LOCK_MAIN"
if ! flock -w 1 100; then
	echo "locked"
	touch "$LOCK_ADD"
	exit
fi

# ---   main    ---

"$SCRIPT" > "${2:-${DIR}/log}" 2>&1

# --- check add ---

flock -u 100
echo 'unlocked'
sleep 2
if [ -f "$LOCK_ADD" ]; then
	rm "$LOCK_ADD" || :
	echo
	echo add found, run again
	echo
	"$0" "$SCRIPT"
fi

echo "  end $TIME"
