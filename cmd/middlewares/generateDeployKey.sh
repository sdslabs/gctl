#! /bin/sh
appName="$1"

ssh-keygen -t rsa -N "" -f $HOME/.ssh/$appName