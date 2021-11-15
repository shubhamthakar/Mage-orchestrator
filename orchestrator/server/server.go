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
	//Establishing connection with datamonk server at 10000
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb1.NewUserNameClient(conn)

	ctx1, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// r is data received from datamonk server
	r, err := c.GetMockUserData(ctx1, &pb1.Username{Name: in.GetName()})
	if err != nil {
		return nil,err
	}
	//returning data to orchestrator 1(server)
	return &pb.User{Name: r.Name, Class: r.Class, Roll: r.Roll}, err

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
