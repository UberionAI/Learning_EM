package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"task_2/internal/server/grpc_server"
	"task_2/proto/proto"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewUserServiceClient(conn)

	token, err := grpc_server.GenerateToken("user1")
	if err != nil {
		log.Fatal("Token error:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	md := metadata.Pairs("authorization", "Bearer "+token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	cre, err := client.CreateUser(ctx, &proto.CreateUserRequest{Name: "Бобик"})
	if err != nil {
		log.Fatal("CreateUser error:", err)
	}
	fmt.Printf("Создан: %+v\n", cre.User)

	list, err := client.ListUsers(ctx, &proto.ListUsersRequest{})
	if err != nil {
		log.Fatal("ListUsers error:", err)
	}
	fmt.Println("Список:", list.Users)

	upd, err := client.UpdateUser(ctx, &proto.UpdateUserRequest{Id: "1", Name: "Бобик2"})
	if err != nil {
		log.Fatal("UpdateUser error:", err)
	}
	fmt.Printf("Обновлён: %+v\n", upd.User)

	_, err = client.DeleteUser(ctx, &proto.DeleteUserRequest{Id: "1"})
	if err != nil {
		log.Fatal("DeleteUser error:", err)
	}
}
