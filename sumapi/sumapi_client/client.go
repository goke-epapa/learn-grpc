package main

import (
	"context"
	"log"

	"github.com/goke-epapa/learn-grpc/sumapi/sumapipb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := sumapipb.NewSumServiceClient(conn)
	doUnary(c)
}

func doUnary(c sumapipb.SumServiceClient) {
	req := &sumapipb.SumRequest{
		Calculation: &sumapipb.Sum{
			FirstNumber:  10,
			SecondNumber: 3,
		},
	}

	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Sum RPC %w", err)
	}

	log.Printf("Response from Sum API: %v", res)
}
