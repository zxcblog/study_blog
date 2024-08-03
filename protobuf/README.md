# 什么是protobuf

Protocol Buffers（Protobuf） 是 Google 公司开发的一种数据描述语言，并于 2008 年对外开源。Protobuf 刚开源时的定位类似于 XML、JSON 等数据描述语言，
是一种与语言、平台无关，可扩展的序列化结构化数据的数据描述语言，常用于通信协议，数据存储等等，通过附带工具生成代码并实现将结构化数
据序列化的功能。但是我们更关注的是 Protobuf 作为接口规范的描述语言，可以作为设计安全的跨语言 PRC 接口的基础工具。

## 基本语法
```protobuf
syntax = "proto3";

package helloworld;

option go_package = "github.com/zxcblog/study_blog/helloworld;helloworld";

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```
- `syntax`: 声明使用 `proto3` 语法，
- `package`: 声明包名，用于区分其他包。其他文件中调用时，通过包名进行调用
- `option`: 设置编译选项，如 go_package，用于指定生成的代码的包名
- `service`: 定义名为 Greeter 的 RPC 服务，包含一个名为 SayHello 的 RPC 方法，入参为 HelloRequest，出参为 HelloReply
- `message`: 定义消息体，每一个消息体的字段包含3个属性， 类型、 字段名称、 字段编号

在 XML 或 JSON 等数据描述语言中，一般通过成员的名字来绑定对应的数据。但是 Protobuf 编码却是通过成员的唯一编号来绑定对应的数据，
因此 Protobuf 编码后数据的体积会比较小，但是也非常不便于人类查阅。我们目前并不关注 Protobuf 的编码技术，最终生成的 Go 结构体可以自由采用
JSON 或 gob 等编码格式，因此大家可以暂时忽略 Protobuf 的成员编码部分。

# 安装和编译
> 此处安装使用docker进行安装，统一打包信息

Protobuf 核心的工具集是 C++ 语言开发的，在官方的 protoc 编译器中并不支持 Go 语言。要想基于上面的 hello.proto 文件生成相应的 Go 代码，
需要安装相应的插件。

- 安装官方的 protoc 工具，可以从 https://github.com/google/protobuf/releases 下载。然后是安装针对 Go 语言的
  代码生成插件，可以通过 go get github.com/golang/protobuf/protoc-gen-go 命令安装。

- 使用 buf cli 插件进行打包，Buf CLI 是现代、快速和高效的 Protobuf API 管理的终极工具。可以参考文档 `https://buf.build/docs/tutorials/getting-started-with-buf-cli`

接下来的安装和编译都使用 buf cli

## 编译dockerfile文件
在使用 docker 进行打包时， 从 https://github.com/bufbuild/buf/releases 下载时间长，所以提前下载好，减少打包时间。使用的版本为 1.35.1

在编译protobuf时，根据个人需要进行对应的插件信息安装。 下方dockerfile文件中， 包含了 go, grpc, grpc-gateway, validate, openapiv2 等插件，
validate 是参数校验的插件，openapiv2 是生成swagger文档的插件

```dockerfile
# 使用golang 1.22-alpine3.18 做基础镜像来进行安装
FROM golang:1.22-alpine3.18 as builder

# 设置go环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# 更新下载软件
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update \
    && apk add --no-cache ca-certificates git openssh-client make bash yarn curl \
    && rm -rf /var/cache/apk/* \
    && git config --global http.version HTTP/1.1 && git config --global http.postBuffer 524288000

# 下载编译protobuf时需要使用到的插件
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.21.0 \
          github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.21.0 \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1 \
    && go install github.com/envoyproxy/protoc-gen-validate@v1.0.4

# 将下载好的文件转移到alpine镜像中,使镜像体积减小
FROM alpine:3.18

COPY --from=builder /go/bin /usr/local/bin
COPY buf-Linux-x86_64 /usr/local/bin/buf

RUN chmod +x "/usr/local/bin/buf"

ENTRYPOINT ["/usr/local/bin/buf"]
```

## 将 Dockerfile 文件打包成镜像
将proto中的dockerfile打包成容器，后续pb打包都通过该容器进行打包
```shell
cd proto

docker build -f ./Dockerfile -t "study-blog/buf:v1" .

docker run --rm study-blog/buf:v1 --version
```

# 构建项目
在使用proto 设计 rpc 接口的通信协议时，会有一些需要其他的依赖。我们使用buf.yaml文件进行配置

下方所用的是v2版本的语法， 对应的语法的文件夹为protov2, v1版本语法在proto文件夹中
## 初始化buf.yaml
```shell
docker run --rm -v "$(pwd)/:/workspace" --workdir /workspace/proto study-blog/buf:v1 mod init buf.build/zxcblog/study-blog
```

buf.yaml文件详解
```yaml
version: v2  # 指定 buf.yaml 文件本身的格式版本
name: buf.build/zxcblog/study-blog # 指定模块（module）的名称，遵循 remote/owner/repository 的格式
breaking: # 用于定义变更检测（Breaking Change Detection）规则，可以帮助在版本迭代时检测是否引入了破坏兼容性的更改。
  use:
    - FILE
deps: #列出项目的外部依赖项，每个依赖项都有 import_path 和 remote。import_path 是在 .proto 文件中导入的路径，remote 是远程仓库的地址。
  - import_path: googleapis/google/api
    remote: https://github.com/googleapis/googleapis.git
lint: # 配置 Buf Lint 的规则集和排除规则
  use:
    - DEFAULT
```

## 创建 buf.gen.yaml文件
是命令用于生成语言集成代码的配置文件您的选择。此文件最常与模块一起使用（但可以与其他输入类型一起使用），并且通常放置在 Protobuf 文件根目录下
```yaml
version: v2
plugins:
  - local: protoc-gen-go
    out: ../pb
    opt:
      - paths=source_relative
  - local: protoc-gen-validate
    out: ../pb
    opt:
      - lang=go
      - paths=source_relative
  - local: protoc-gen-go-grpc
    out: ../pb
    opt:
      - paths=source_relative
      - require_unimplemented_servers=false
  - local: protoc-gen-grpc-gateway
    out: ../pb
    opt:
      - paths=source_relative
      - allow_repeated_fields_in_body=true

  # 生成swagger文件
  - local: protoc-gen-openapiv2
    out: ../pb
    strategy: all
    opt:
      - allow_merge=true,merge_file_name=info # 将不同文件夹下的swagger整合生成到同一个文件中,文件名称为 openapi
```


## 创建proto
```protobuf
syntax = "proto3";

package base.v1;

option go_package = "rat-race/pb/source/base;base"; // 生成以后的go文件包名

import "google/protobuf/any.proto";

// PageReq 分页请求
message PageReq {
  int64 limit = 1;
  int64 offset = 2;
}

// PageRes 分页请求返回
message PageRes {
  int64 total = 1;
  int64 current_page = 2;
  int64 total_page = 3;
  int64 page_size = 4;
}

// Error 错误消息定义
message Error {
  int32 code = 1;
  string message = 2;
  google.protobuf.Any detail = 3;
}
```

## 更新并锁定依赖模块版本
在写好需要的proto文件后， 更新并锁定依赖版本
```shell
docker run --rm -v "$(pwd)/:/workspace" --workdir /workspace/proto study-blog/buf:v1 dep update
```

## 将proto转换成go文件
```shell
docker run --rm -v "$(pwd)/:/workspace" --workdir /workspace/proto study-blog/buf:v1 generate
```
