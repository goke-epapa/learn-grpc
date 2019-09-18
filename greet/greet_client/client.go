package main

import (
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
	fmt.Println("Created client %f", c)
}
