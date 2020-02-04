package main

import (
	"context"
	"log"
	"net"

	"github.com/goke-epapa/learn-grpc/sumapi/sumapipb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *sumapipb.SumRequest) (*sumapipb.SumResponse, error) {
	firstNumber := req.GetCalculation().GetFirstNumber()
	secondNumber := req.GetCalculation().GetSecondNumber()

	result := firstNumber + secondNumber

	res := &sumapipb.SumResponse{
		Result: int64(result),
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to list: %v", err)
	}

	s := grpc.NewServer()
	sumapipb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
