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
