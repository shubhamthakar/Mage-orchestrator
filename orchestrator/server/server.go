package main

import (
	"context"
	"log"
	"net"
	"time"

	pb1 "github.com/shubhamthakar/Mage/datamonk/proto"
	pb "github.com/shubhamthakar/Mage/orchestrator/proto"
	"google.golang.org/grpc"
)

const (
	port = ":9001"
)

type UserNameServer struct {
	pb.UnimplementedUserNameServer
}

func (s *UserNameServer) GetUser(ctx context.Context, in *pb.Username) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())

	const (
		address = "localhost:10000"
	)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb1.NewUserNameClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetMockUserData(ctx, &pb1.Username{Name: in.GetName()})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &pb.User{Name: r.Name, Class: r.Class, Roll: r.Roll}, nil


	//return nil, errors.New("not implemented yet. Shubham will implement me")
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
