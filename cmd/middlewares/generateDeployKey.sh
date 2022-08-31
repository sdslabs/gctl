#! /bin/sh
keyPath="$1"

ssh-keygen -t rsa -N "" -f $keyPath