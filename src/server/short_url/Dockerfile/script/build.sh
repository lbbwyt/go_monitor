#!/usr/bin/env bash
echo "start....."
ls -l /go/src/ |awk '/^d/ {print $NF}'
ls -l /go/src/app/ |awk '/^d/ {print $NF}'
ls -l /go/src/script/ |awk '/^d/ {print $NF}'
cd /go/src/app/ && ./main