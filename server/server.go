package main

import (
	"context"
	"log"
	"net"
	"time"

	pb1 "github.com/shubhamthakar/Mage/orchestrator/proto"
	pb "github.com/shubhamthakar/Mage/proto"
	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

type UserNameServer struct {
	pb.UnimplementedUserNameServer
}

//Implementing interface
func (s *UserNameServer) GetUserByName(ctx context.Context, in *pb.Username) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())

	const (
		address = "localhost:9001"
	)
	//Establishing connection with orchestrator
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb1.NewUserNameClient(conn)

	ctx1, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	//Data received from orchestrator
	r, err := c.GetUser(ctx1, &pb1.Username{Name: in.GetName()})
	if err != nil {
		return nil,err
	}
	//Returning data to client
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
		log.Printf("failed to serve: %v", err)
	}
}
