#!/usr/bin/env bash
set -e

cd "$(dirname "$BASH_SOURCE")/.."
rm -rf vendor/
source 'hack/.vendor-helpers.sh'

# Docker / Engine-api related
clone git github.com/docker/engine-api 6de18e18540cda038b00e71a1f2946d779e83f87
clone git github.com/docker/docker 9e530247e066ef7a32e35a1f0f818c1e4048ad54
# FIXME use tags >v0.2.3 soon
clone git github.com/docker/distribution db17a23b961978730892e12a0c6051d43a31aab3
clone git github.com/docker/go-units 651fc226e7441360384da338d0fd37f2440ffbe3
clone git github.com/docker/go-connections v0.2.0
clone git github.com/opencontainers/runc 3d8a20bb772defc28c355534d83486416d1719b4
clone git github.com/Sirupsen/logrus v0.9.0
clone git golang.org/x/net 47990a1ba55743e6ef1affd3a14e5bac8553615d https://github.com/golang/net.git
clone git github.com/vbatts/tar-split v0.9.11

# libcompose related
clone git github.com/docker/libcompose 8cd02056eaca6d1a21b7075ed5ce52ce32bb8366
clone git github.com/cloudfoundry-incubator/candiedyaml 55a459c2d9da2b078f0725e5fb324823b2c71702
clone git github.com/flynn/go-shlex 3f9db97f856818214da2e1057f8ad84803971cff
clone git github.com/gorilla/context 14f550f51a
clone git github.com/gorilla/mux e444e69cbd

# gocheck related
clone git github.com/go-check/check 4f90aeace3a26ad7021961c297b22c42160c7b25

clean && mv vendor/src/* vendor
