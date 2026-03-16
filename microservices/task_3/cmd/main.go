package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"task_2/internal/server/grpc_server"
)

func main() {
	lis, err := net.Listen("tcp", ":8080") //nolint:errcheck
	if err != nil {
		log.Fatal(err)
	}
	//s := grpc.NewServer()
	//grpc_server.RegisterGRPC(s)
	//fmt.Println("gRPC сервер на :8080")
	//s.Serve(lis)
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_server.AuthUnaryInterceptor),
	)

	grpc_server.RegisterGRPC(s)
	fmt.Println("gRPC сервер на :8080 (с JWT auth)")
	s.Serve(lis)
}
