FROM golang:1.6.1

RUN apt-get update && apt-get install -y \
    iptables build-essential \
    --no-install-recommends

# Install build dependencies
RUN go get golang.org/x/tools/cmd/cover \
    && go get github.com/golang/lint/golint \
    && go get github.com/Masterminds/glide 

# Which docker version to test on and what default one to use
ENV DOCKER_VERSIONS 1.10.3 1.11.0
ENV DEFAULT_DOCKER_VERSION 1.10.3

# Download docker
RUN set -e; \
    for v in $(echo ${DOCKER_VERSIONS} | cut -f1); do \
        if test "${v}" = "1.9.1" || test "${v}" = "1.10.3"; then \
           mkdir -p /usr/local/bin/docker-${v}/; \
           curl https://get.docker.com/builds/Linux/x86_64/docker-${v} -o /usr/local/bin/docker-${v}/docker; \
           chmod +x /usr/local/bin/docker-${v}/docker; \
        else \
             curl https://get.docker.com/builds/Linux/x86_64/docker-${v}.tgz -o docker-${v}.tgz; \
             tar xzf docker-${v}.tgz -C /usr/local/bin/; \
             mv /usr/local/bin/docker /usr/local/bin/docker-${v}; \
             rm docker-${v}.tgz; \
        fi \
    done

# Set the default Docker to be run
RUN ln -s /usr/local/bin/docker-${DEFAULT_DOCKER_VERSION} /usr/local/bin/docker
    
WORKDIR /go/src/github.com/vdemeester/libkermit

COPY glide.yaml glide.yaml
COPY glide.lock glide.lock
RUN glide install

# Wrap all commands in the "docker-in-docker" script to allow nested containers
ENTRYPOINT ["hack/dind"]

COPY . /go/src/github.com/vdemeester/libkermit
