#!/bin/bash

go run /var/go/src/github.com/ezoic/gvlcache/
service memcached start -D FOREGROUND