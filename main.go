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
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("监听端口失败, err:", err.Error())
		return
	}

	s := grpc.NewServer()
	user.RegisterUserServer(s, user2.NewUserService())

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
