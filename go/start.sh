#!/bin/sh
# Start the Coursehero httpd.

bindir=`dirname $0`

# Enable starting from unpack or installed locations
# See appspec.yml
if [ -f "$bindir/bin/httpd" ]; then
	chmod 755 "$bindir/bin/httpd"
	$bindir/bin/httpd $1 &
	pid=$!
elif [ -f "$bindir/httpd" ]; then
	chmod 755 "$bindir/httpd"
	$bindir/httpd $1 &
	pid=$!
else
	echo "Httpd not found."; exit 1
fi

if [ -d /var/run/lock ]; then
	echo $pid > /var/run/lock/httpd.pid
fi

