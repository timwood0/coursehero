#!/bin/sh
# Start the Coursehero httpd.

# See appspec.yml
bindir=/usr/local/coursehero/bin

if [ ! -x "$bindir/httpd" ]; then
	echo "$bindir/httpd not found."; exit 1
fi

$bindir/httpd $1 &
pid=$!
if [ -d /var/run/lock ]; then
	echo $pid > /var/run/lock/httpd.pid
fi

