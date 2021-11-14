package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/shubhamthakar/Mage/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserNameServer struct {
	pb.UnimplementedUserNameServer
}

func (s *UserNameServer) GetUserByName(ctx context.Context, in *pb.Username) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	return nil, errors.New("not implemented yet. Shubham will implement me")
	//return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserNameServer(s, &UserNameServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
