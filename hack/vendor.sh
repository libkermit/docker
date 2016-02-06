#!/usr/bin/env bash
set -e

cd "$(dirname "$BASH_SOURCE")/.."
rm -rf vendor/
source 'hack/.vendor-helpers.sh'

clone git github.com/docker/docker v1.10.0
# FIXME use tags >v0.2.3 soon
clone git github.com/docker/distribution c301f8ab27f4913c968b8d73a38e5dda79b9d3d7
clone git github.com/docker/engine-api 9a940e4ead265e18d4feb9e3c515428966a08278
clone git github.com/docker/go-connections 3ef1e92d1b93e587895dfa466e1682ae63448801
clone git github.com/docker/go-units 651fc226e7441360384da338d0fd37f2440ffbe3
clone git github.com/opencontainers/runc 3d8a20bb772defc28c355534d83486416d1719b4
# clone git github.com/docker/libcompose 4298aeee6e879e8efd75c048f43e5f000b2fae04
# clone git github.com/fsouza/go-dockerclient 39d9fefa6a7fd4ef5a4a02c5f566cb83b73c7293
clone git github.com/Sirupsen/logrus v0.9.0
clone git golang.org/x/net 47990a1ba55743e6ef1affd3a14e5bac8553615d https://github.com/golang/net.git
clone git github.com/vbatts/tar-split v0.9.11

clean && mv vendor/src/* vendor
