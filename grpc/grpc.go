package grpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RegisterGRPC(port int) (*grpc.Server, error) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return nil, err
	}
	server := grpc.NewServer()
	// 反射接口支持查询
	reflection.Register(server)
	// 启动服务
	if err = server.Serve(listen); err != nil {
		log.Fatalf("failed to server: %v", err)
		return nil, err
	}
	return nil, err
}
