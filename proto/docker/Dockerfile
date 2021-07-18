FROM golang:1.16

ARG PROTOC_VERSION=3.15.8
ARG PROTOC_GEN_DOC_VERSION=1.4.1

RUN apt-get update && \
    apt-get -y -q install unzip make

RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip && \
    unzip -o protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local bin/protoc && \
    unzip -o protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local include/* && \
    rm -rf protoc-${PROTOC_VERSION}-linux-x86_64.zip

RUN go get -u github.com/golang/protobuf/protoc-gen-go && \
    go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc && \
    go get -u github.com/envoyproxy/protoc-gen-validate

RUN cd ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate && \
    make build

WORKDIR /ts

RUN apt-get -y -q install nodejs npm

RUN npm install ts-protoc-gen

WORKDIR /workspace

ENTRYPOINT ["protoc"]
