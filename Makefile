TARGET_PROTO_FILES=$(shell find protobuf -name '*.proto' | sort)
PACKAGE_ROOT=github.com/sqrt-7/goodhost
GO_VERSION_LOCK_PROTOBUF=latest
GO_VERSION_LOCK_GRPC_GW=v1.16.0
PROTOC_VERSION=3.17.3

PROTOC_INCLUDES= -I. \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway

install-protoc:
	curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-osx-x86_64.zip
	unzip protoc-${PROTOC_VERSION}-osx-x86_64.zip -d ${HOME}/.local
	rm -rf protoc-${PROTOC_VERSION}-osx-x86_64.zip

install-deps-go:
	cd $(shell mktemp -d) && \
	go mod init temp 2>&1 && \
	go get -v -ldflags="-s -w" \
		github.com/golang/protobuf/protoc-gen-go@${GO_VERSION_LOCK_PROTOBUF} \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@${GO_VERSION_LOCK_GRPC_GW} \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@${GO_VERSION_LOCK_GRPC_GW} \
		2>&1
	mkdir -p ${GOPATH}/src/github.com/grpc-ecosystem/
	cp -r ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@${GO_VERSION_LOCK_GRPC_GW} \
		${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway

autogen-go:
	for PROTOFILE in ${TARGET_PROTO_FILES}; do \
		protoc ${PROTOC_INCLUDES} \
			--go_out=plugins=grpc:. \
			--grpc-gateway_out=logtostderr=true:. \
			$$PROTOFILE \
		; done
	rm -rf ./pkg/autogen
	mv ${PACKAGE_ROOT}/pkg/autogen ./pkg/autogen
	rm -rf $(shell echo ${PACKAGE_ROOT} | tr "/" "\n" | head -n 1)

run-localserver:
	rm -rf bin
	go build -o ./bin/localserver ./cmd/localserver/main.go
	./bin/localserver