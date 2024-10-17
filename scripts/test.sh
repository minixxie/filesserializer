#!/bin/bash

set -e

echo "[/tmp/new/]:"
ls -ld /tmp/new || true
cd /tmp/old/
find . -type f | sort | xargs md5sum > /tmp/checksum1
cat /tmp/checksum1

/app

echo "[/tmp/new/](After Unmarshal):"
ls -lt /tmp/new/
cd /tmp/new/
find . -type f | sort | xargs md5sum > /tmp/checksum2
cat /tmp/checksum2


md5sum /tmp/checksum1
md5sum /tmp/checksum2

if [ "$(md5sum /tmp/checksum1 | awk '{print $1}')" == "$(md5sum /tmp/checksum2 | awk '{print $1}')" ]; then
	echo "SUCCESS"
else
	echo >&2 "FAILED"
	exit -1
fi
