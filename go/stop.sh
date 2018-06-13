#!/bin/sh
# Stop a Coursehero httpd.

pidfile=/var/run/lock/httpd.pid
if [ -f $pidfile ]; then
	kill -TERM `cat $pidfile` 2> /dev/null || echo "Httpd process `cat $pidfile` not found."
	rm -f $pidfile
else
	echo "No known httpd known running."
fi

