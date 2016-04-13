FROM golang:1.6.1

RUN apt-get update && apt-get install -y \
    build-essential \
    --no-install-recommends

# Install build dependencies
RUN go get golang.org/x/tools/cmd/cover \
    && go get github.com/golang/lint/golint \
    && go get github.com/Masterminds/glide 

# Which docker version to test on
ENV DOCKER_VERSION 1.10.3

# enable GO15VENDOREXPERIMENT
ENV GO15VENDOREXPERIMENT 1

# Download docker
RUN set -ex; \
    curl https://get.docker.com/builds/Linux/x86_64/docker-${DOCKER_VERSION} -o /usr/local/bin/docker; \
    chmod +x /usr/local/bin/docker
    
WORKDIR /go/src/github.com/vdemeester/libkermit

COPY glide.yaml glide.yaml
COPY glide.lock glide.lock
RUN glide install

COPY . /go/src/github.com/vdemeester/libkermit
