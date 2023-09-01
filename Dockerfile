FROM golang:1.21.0-alpine3.18 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# 更新下载软件
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update \
    && apk add --no-cache ca-certificates git openssh-client make bash yarn curl \
    && rm -rf /var/cache/apk/* \
    && git config --global http.version HTTP/1.1 && git config --global http.postBuffer 524288000

RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.14.0 \
          github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.14.0 \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0 \
    && go install github.com/envoyproxy/protoc-gen-validate@v0.9.0

FROM alpine:3.18

COPY --from=builder /go/bin /usr/local/bin
COPY ./services/proto/buf-Linux-x86_64 /usr/local/bin/buf

RUN chmod +x "/usr/local/bin/buf"

ENTRYPOINT ["/usr/local/bin/buf"]


