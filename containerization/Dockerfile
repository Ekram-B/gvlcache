FROM ubuntu:18.04

MAINTAINER Ekram Bhuiyan <ebhuiyan@ezoic.com>
WORKDIR /root/

# Update system on which packages can be installed
RUN apt-get update

# Use package manager to install necessary packages
RUN apt-get -y --fix-missing install curl

# Set variable indicating which golang version to retreive - same as the one set up within the container
ENV GOLANG_VERSION 1.13.7
ENV GOROOT="/usr/local/go"
ENV GOPATH="/var/go/"

# Get GO supported by Linux on top of a AMD64 bit architecture and then unpack the tarball
RUN curl -sSL https://storage.googleapis.com/golang/go$GOLANG_VERSION.linux-amd64.tar.gz \
		| tar -v -C /usr/local -xz
# Add the directory containing binary to path so it can be run anywhere from the command line
ENV PATH /usr/local/go/bin:$PATH

