package main

import (
	"context"
	"log"
	"time"

	pb "github.com/shubhamthakar/Mage/datamonk/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:10000"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserNameClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetMockUserData(ctx, &pb.Username{Name: "Shubham"})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf(`User Details: NAME: %s AGE: %s ID: %d`, r.GetName(), r.GetClass(), r.GetRoll())
	

}

