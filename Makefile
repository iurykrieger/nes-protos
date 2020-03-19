
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	for proto in **/*.proto ; do \
        protoc \
			--proto_path=${GOPATH}/src:. \
			--micro_out=. \
			--go_out=plugins=grpc,paths=source_relative:. \
		$$proto ; \
    done
