# protoc 安装
protoc 是 Protobuf的编译器， 是用C++编写的， 主要功能是用于编译 `.proto` 文件。因为下载有时会消耗很长时间，所以提前下载好对应的编译文件，

[点击跳转 `protoc github` 地址](https://github.com/protocolbuffers/protobuf)

使用 [buf](https://github.com/bufbuild/buf) 代替 protoc 进行进行打包


# 将 Dockerfile 文件打包成镜像
```shell
docker build -f ./Dockerfile -t "zxc/buf:v1" .
```

查看镜像是否编译成功
```shell
docker run --rm zxc/buf:v1 -v
```

# 使用打包后的镜像编译proto文件
```shell
# 初始化proto模块
docker run --rm -v "$(pwd)/proto:/workspace" --workdir /workspace zxc/buf:v1 mod init
```

