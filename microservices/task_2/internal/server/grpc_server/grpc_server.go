package grpc_server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
	"task_2/internal/handlers"
	"task_2/internal/models"
	pb "task_2/proto/proto"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := handlers.User{ID: "", Name: req.Name}
	id := strconv.Itoa(len(handlers.Users) + 1)
	user.ID = id
	handlers.Users[id] = models.User(user)

	return &pb.CreateUserResponse{User: &pb.User{Id: id, Name: req.Name}}, nil
}

func (s *UserServer) ListUsers(ctx context.Context, _ *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	resp := &pb.ListUsersResponse{}
	for _, u := range handlers.Users {
		resp.Users = append(resp.Users, &pb.User{Id: u.ID, Name: u.Name})
	}
	return resp, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	if u, ok := handlers.Users[req.Id]; ok {
		u.Name = req.Name
		handlers.Users[req.Id] = u
		return &pb.UpdateUserResponse{User: &pb.User{Id: req.Id, Name: req.Name}}, nil
	}
	return nil, status.Error(codes.NotFound, "user not found")
}

func (s *UserServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	delete(handlers.Users, req.Id)
	return &emptypb.Empty{}, nil
}

func RegisterGRPC(s *grpc.Server) {
	pb.RegisterUserServiceServer(s, &UserServer{})
}
