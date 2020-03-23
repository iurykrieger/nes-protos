# Based on https://developers.google.com/protocol-buffers/docs/gotutorial
FROM golang:buster

LABEL maintainer=iurykrieger96@gmail.com

RUN apt-get update

# Install protobuf libs (https://github.com/protocolbuffers/protobuf/blob/master/src/README.md)
RUN apt-get install -y \
    autoconf \
    automake \
    libtool \
    curl \
    make \
    g++ \
    unzip

WORKDIR /usr/local

# Clone the protobuf release
RUN wget -c https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protobuf-cpp-3.6.1.tar.gz -O - | tar -xz


WORKDIR /usr/local/protobuf-3.6.1

# Installs protobuf with "protoc" binary
RUN ./configure --prefix=/usr; \
    make; \
    make check; \
    make install; \
    ldconfig

# Sets repository workdir
WORKDIR /go/src/github.com/iurykrieger/nes-protos

COPY go.mod .

# Install protobuf Golang generator plugin
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go

# Install protobuf Golang Micro generator plugin (https://github.com/micro/protoc-gen-micro)
RUN go install github.com/micro/protoc-gen-micro/v2
