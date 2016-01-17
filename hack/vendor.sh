#!/usr/bin/env bash
set -e

cd "$(dirname "$BASH_SOURCE")/.."
rm -rf vendor/
source 'hack/.vendor-helpers.sh'

clone git github.com/docker/docker v1.10.0-rc1
clone git github.com/docker/go-units 651fc226e7441360384da338d0fd37f2440ffbe3
# FIXME use tags >v0.2.3 soon
clone git github.com/docker/engine-api 31b698aad77eafbef7e6a344a001038c46a68425
clone git github.com/docker/go-connections v0.1.2
# clone git github.com/docker/libcompose 4298aeee6e879e8efd75c048f43e5f000b2fae04
# clone git github.com/fsouza/go-dockerclient 39d9fefa6a7fd4ef5a4a02c5f566cb83b73c7293

clean && mv vendor/src/* vendor
