package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"task_2/internal/server/grpc_server"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	grpc_server.RegisterGRPC(s)
	fmt.Println("gRPC сервер на :8080")
	s.Serve(lis)
}
