package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"time"
	"zxcblog/study_blog/internal/middleware"
	user2 "zxcblog/study_blog/internal/service/user"
	"zxcblog/study_blog/pb/base"
	"zxcblog/study_blog/pb/user"
	"zxcblog/study_blog/pkg/errcode"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("监听端口失败, err:", err.Error())
		return
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.Interceptor),
		grpc.StreamInterceptor(middleware.MiddlewareStreamInterceptor),
	)
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

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Minute))
	defer cancel()
	_, err = c.Register(ctx, &user.RegisterReq{
		Account:         "1",
		Nickname:        "1",
		Password:        "1",
		ConfirmPassword: "1",
		Mobile:          "1",
		MobileCache:     "1",
		ImgCache:        "1",
	})
	if err != nil {

		status := errcode.FromError(err)
		fmt.Println("用户注册失败", status.Code())
		fmt.Println("用户注册失败", status.Message())

		for _, v := range status.Details() {
			if val, ok := v.(*base.Error); ok {
				fmt.Println("循环", val.Code, val.Message, val.Detail)
			}
		}
		return
	}
}
