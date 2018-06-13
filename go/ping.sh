#!/bin/sh

port=4000
[ -n "$1" ] && port=$1
nc `hostname` $port < /dev/null

