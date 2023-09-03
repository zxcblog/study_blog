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
