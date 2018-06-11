#!/bin/sh -x

[ -n "$GOPATH" ] || { echo '$GOPATH not set.'; exit 1; }
cd $GOPATH/src/httpd
go install
cd main
go install
cd $GOPATH/bin
# Hack around package-name->executable-name rule. :/
mv main httpd
