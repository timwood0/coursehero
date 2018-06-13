#!/bin/sh
# Start the Coursehero httpd.

bindir=`dirname $0`
$bindir/httpd $1 &
pid=$!

if [ -d /var/run ]; then
	echo $pid > /var/run/lock/httpd.pid
fi

