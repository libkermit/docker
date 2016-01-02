FROM golang:1.5.2

RUN apt-get update && apt-get install -y \
    build-essential \
    --no-install-recommends

# Install build dependencies
RUN go get golang.org/x/tools/cmd/cover
RUN go get github.com/golang/lint/golint
RUN go get golang.org/x/tools/cmd/vet

# enable GO15VENDOREXPERIMENT
ENV GO15VENDOREXPERIMENT 1

WORKDIR /go/src/github.com/vdemeester/libkermit

COPY . /go/src/github.com/vdemeester/libkermit
