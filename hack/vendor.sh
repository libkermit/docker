#!/usr/bin/env bash
set -e

cd "$(dirname "$BASH_SOURCE")/.."
rm -rf vendor/
source 'hack/.vendor-helpers.sh'

# Docker / Engine-api related
clone git github.com/docker/engine-api 9a940e4ead265e18d4feb9e3c515428966a08278
clone git github.com/docker/docker e18eb6ef394522fa44bce7a3b9bb244d45ce9b56
# FIXME use tags >v0.2.3 soon
clone git github.com/docker/distribution c301f8ab27f4913c968b8d73a38e5dda79b9d3d7
clone git github.com/docker/go-units 651fc226e7441360384da338d0fd37f2440ffbe3
clone git github.com/docker/go-connections v0.1.3
clone git github.com/docker/libtrust 9cbd2a1374f46905c68a4eb3694a130610adc62a
clone git github.com/opencontainers/runc 3d8a20bb772defc28c355534d83486416d1719b4
clone git github.com/Sirupsen/logrus v0.9.0
clone git golang.org/x/net 47990a1ba55743e6ef1affd3a14e5bac8553615d https://github.com/golang/net.git
clone git github.com/vbatts/tar-split v0.9.11

# libcompose related
clone git github.com/docker/libcompose 3ca15215f36154fbf64f15bfa305bfb0cebb6ca7
clone git github.com/cloudfoundry-incubator/candiedyaml 55a459c2d9da2b078f0725e5fb324823b2c71702
clone git github.com/flynn/go-shlex 3f9db97f856818214da2e1057f8ad84803971cff
clone git github.com/gorilla/context 14f550f51a
clone git github.com/gorilla/mux e444e69cbd

clean && mv vendor/src/* vendor
