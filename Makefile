
GOPATH:=$(shell go env GOPATH)

WORKDIR=/go/src/github.com/iurykrieger/nes-protos

.PHONY: proto
proto: build 
	docker run \
		-v ${PWD}:${WORKDIR} \
		nes-protos \
		protoc \
		--proto_path=${WORKDIR} \
		--go_out=/go/src \
		--micro_out=/go/src \
		protos/*.proto

.PHONY: build
build: proto
	docker build -t nes-protos .
