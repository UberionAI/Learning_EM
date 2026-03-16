package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cre, _ := client.CreateUser(ctx, &proto.CreateUserRequest{Name: "Бобик"})
	fmt.Printf("Создан: %+v\n", cre.User)

	list, _ := client.ListUsers(ctx, &proto.ListUsersRequest{})
	fmt.Println("Список:", list.Users)

	upd, _ := client.UpdateUser(ctx, &proto.UpdateUserRequest{Id: "1", Name: "Бобик2"})
	fmt.Printf("Обновлён: %+v\n", upd.User)

	client.DeleteUser(ctx, &proto.DeleteUserRequest{Id: "1"})
}
