package main

import (
	"context"
	"fmt"
	"log"

	"github.com/goke-epapa/learn-grpc/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // No SSL certs
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetingRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Goke",
			LastName:  "Obasa",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Greet RPC %w", err)
	}

	log.Printf("Response from Greet: %v", res)
}
