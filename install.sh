#!/bin/sh

BINARY='/usr/local/bin'

echo "Building s2t"
go build s2t.go

echo "Installing s2t to $BINARY"
install -v s2t $BINARY

echo "Removing the build"
rm s2t
