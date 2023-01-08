package main

import (
	"context"
	"log"
	"time"

	pb "grpc-go-course/greet/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("GreetWithDeadline was invoked with: %v\n", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client cancel the request")
			return nil, status.Error(codes.Canceled, "The client cancel the request")
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}
