#!/bin/sh
# Start the Coursehero httpd.

bindir=`dirname $0`
if [ -x "$bindir/bin/httpd" ]; then
	$bindir/bin/httpd $1 &
	pid=$!
elif [ -x "$bindir/httpd" ]; then
	$bindir/httpd $1 &
	pid=$!
else
	echo "Httpd not found."; exit 1
fi

if [ -d /var/run ]; then
	echo $pid > /var/run/lock/httpd.pid
fi

