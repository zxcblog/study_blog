package middleware

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

// Interceptor grpc 服务器端一元拦截器
// 实现 grpc.UnaryServerInterceptor 方法
func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	// 前置逻辑处理
	log.Println("获取当前grpc的信息", info.FullMethod)

	m, err := handler(ctx, req)

	log.Println("查看请求后信息", m)
	return m, err
}

//// SteamInterceptor 流拦截器
//// 实现 grpc.StreamServerInterceptor 方法
//func SteamInterceptor(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error {
//
//}
