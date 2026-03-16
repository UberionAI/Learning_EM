package main

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	pb "task_2/proto/proto"
)

func BenchmarkListUsers(b *testing.B) {
	conn, _ := grpc.Dial("localhost:8080", grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.ListUsers(context.Background(), &pb.ListUsersRequest{})
	}
}
