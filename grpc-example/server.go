package main

import (
	"context"
	"grpc-vs-rest/grpc-example/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type UserServer struct {
	pb.UnimplementedUserServer
}

func (s UserServer) Info(_ context.Context, r *pb.Request) (*pb.Response, error) {
	resp := pb.Response{
		ID:    r.ID,
		Name:  "foo",
		Email: "foo@localhost",
	}
	return &resp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServer(server, UserServer{})

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
