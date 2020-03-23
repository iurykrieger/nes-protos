
GOPATH:=$(shell go env GOPATH)

PROTO_PATH=/go/src/github.com/iurykrieger/nes-protos/protos

.PHONY: proto
proto: build 
	docker run \
		-v ${PWD}/protos:${PROTO_PATH} \
		nes-protos \
		protoc \
		--proto_path=${PROTO_PATH}/src:. \
		--go_out=. \
		--go_opt=paths=source_relative \
		--micro_out=/go/src \
		protos/*.proto

.PHONY: build
build:
	docker build -t nes-protos .
