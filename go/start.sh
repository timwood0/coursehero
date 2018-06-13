#!/bin/sh
# Start the Coursehero httpd.

# See appspec.yml
: ${bindir:=/usr/local/coursehero/bin}
pidfile=/tmp/ch_httpd.pid

if [ ! -x "$bindir/httpd" ]; then
	echo "$bindir/httpd not found."; exit 1
fi

$bindir/httpd $1 > /dev/null 2>1 0>&1 &
pid=$!
echo $pid > $pidfile

