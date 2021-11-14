package main

import (
	"context"
	"errors"
	"log"
	"net"
	"strconv"

	pb "github.com/shubhamthakar/Mage/datamonk/proto"
	"google.golang.org/grpc"
)

const (
	port = ":10000"
)

type UserNameServer struct {
	pb.UnimplementedUserNameServer
}

func (s *UserNameServer) GetMockUserData(ctx context.Context, in *pb.Username) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	if len(in.GetName()) < 6{

		return nil, errors.New("length is short")
	}
	return &pb.User{Name: in.GetName(), Class: strconv.Itoa(len(in.GetName())), Roll: int64(len(in.GetName())*10)}, nil
	
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
