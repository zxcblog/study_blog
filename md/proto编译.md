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

# 更新要拉取的模块并锁定版本
docker run --rm -v "$(pwd)/proto:/workspace" --workdir /workspace zxc/buf:v1 mod update

# 将proto转换成go文件
docker run --rm -v "$(pwd)/proto:/workspace" --workdir /workspace zxc/buf:v1 generate
```

## buf.gen.yaml
是命令用于生成语言集成代码的配置文件 您的选择。此文件最常与模块一起使用（但可以与其他输入类型一起使用），并且通常放置在 Protobuf 文件根目录下

# 在服务中运行
```shell
# 打包服务中的proto文件并生成pb
docker run --rm -v "$(pwd):/workspace" --workdir /workspace/proto zxc/buf:v1 generate
```
