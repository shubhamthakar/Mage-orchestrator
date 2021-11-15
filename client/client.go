package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/shubhamthakar/Mage/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000"
)

func main() {

	fmt.Println("Enter Your First Name: ")
	var name string
	fmt.Scanln(&name)
	//connect with server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//creating new client
	c := pb.NewUserNameClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//calling the remote method
	
	r, err := c.GetUserByName(ctx, &pb.Username{Name: name})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf(`User Details: NAME: %s AGE: %s ID: %d`, r.GetName(), r.GetClass(), r.GetRoll())
	

}

