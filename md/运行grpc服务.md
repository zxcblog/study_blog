# 获取go语言的grpc服务
```shell
go get -u google.golang.org/grpc
```

# 实现user服务的pb
创建 `internal/service/user/user.go` 文件, 内容如下， 为了实现user.UserServer服务
```go
package user

import (
	"context"
	"errors"
	"zxcblog/study_blog/pb/user"
)

type UserService struct {
}

func NewUserService() user.UserServer {
	return UserService{}
}

func (u UserService) Register(ctx context.Context, req *user.RegisterReq) (*user.UserAuthRes, error) {
	//TODO implement me
	return nil, errors.New("implement me")
}

func (u UserService) Login(ctx context.Context, req *user.LoginReq) (*user.UserAuthRes, error) {
	//TODO implement me
	return nil, errors.New("implement me")
}

```

# 启动grpc
创建`main.go`文件
```go
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	user2 "zxcblog/study_blog/internal/service/user"
	"zxcblog/study_blog/pb/user"
)

func main() {
	// 监听9090端口
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("监听端口失败, err:", err.Error())
		return
	}
    
	// grpc注册用户服务
	s := grpc.NewServer()
	user.RegisterUserServer(s, user2.NewUserService())
    
	// 使用协程监听端口， 让用户服务可以被调用
	go func() {
		if err = s.Serve(lis); err != nil {
			fmt.Println("服务启动失败, err:", err.Error())
			return
		}
		fmt.Println("服务启动成功， 监听端口为：9090")
	}()

	//	调用grpc服务
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("grpc调用服务启动失败", err.Error())
		return
	}
	defer conn.Close()
    
	// 调用注册接口
	c := user.NewUserClient(conn)
	_, err = c.Register(context.Background(), &user.RegisterReq{
		Account:         "",
		Nickname:        "",
		Password:        "",
		ConfirmPassword: "",
		Mobile:          "",
		MobileCache:     "",
		ImgCache:        "",
	})
	if err != nil {
		fmt.Println("用户注册失败", err.Error())
		return
	}
}
```

# 注册一元拦截器
添加一元拦截器 `internal/middleware/interceptor.go`
```go
// Interceptor grpc 服务器端拦截器
func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	// 前置逻辑处理
	log.Println("获取当前grpc的信息", info.FullMethod)

	m, err := handler(ctx, req)

	log.Println("查看请求后信息", m)
	return m, err
}
```

修改main.go文件
```go

// 注册拦截器
s := grpc.NewServer(grpc.UnaryInterceptor(middleware.Interceptor))
```

# 注册流拦截器
注册流拦截器 `internal/middleware/stream_interceptor.go`
```go
package middleware

import (
	"google.golang.org/grpc"
	"log"
	"time"
)

// StreamInterceptor 流拦截器
type StreamInterceptor struct {
	grpc.ServerStream
}

// RecvMsg 实现RecvMsg函数，用来处理流RPC所接收到的消息
func (s *StreamInterceptor) RecvMsg(m interface{}) error {
	log.Println("========= [server stream interceptor wrapper] receive a message (Type %T) at %s", m, time.Now().Format(time.RFC3339))
	return s.ServerStream.RecvMsg(m)
}

// SendMsg 实现SendMsg函数，处理流RPC所发送的消息
func (s *StreamInterceptor) SendMsg(m interface{}) error {
	log.Println("=== [server stream interceptor wrapper] send a message (Type %T) at %v", m, time.Now().Format(time.RFC3339))
	return s.ServerStream.SendMsg(m)
}

// newStreamInterceptor 创建新包装器流的实例
func newStreamInterceptor(s grpc.ServerStream) grpc.ServerStream {
	return &StreamInterceptor{
		ServerStream: s,
	}
}

// MiddlewareStreamInterceptor 流拦截器的实例
func MiddlewareStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, hanlder grpc.StreamHandler) error {
	// 前置处理阶段
	log.Println("===========   [server stream interceptor]", info.FullMethod)

	//
	err := hanlder(srv, newStreamInterceptor(ss))
	if err != nil {
		log.Println("RPC failed with error %v", err)
	}

	return err
}
```

修改`main.go`文件
```go
s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.Interceptor),
		grpc.StreamInterceptor(middleware.MiddlewareStreamInterceptor),
	)
```