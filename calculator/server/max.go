package main

import (
	pb "grpc-go-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max func was invoked")

	var maximum int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while receiving request: %v\n", err)
		}

		if number := req.Number; number > maximum {
			maximum = number

			err := stream.Send(&pb.MaxResponse{
				Result: maximum,
			})

			if err != nil {
				log.Fatalf("error while sending data to client: %v\n", err)
			}
		}
	}
}
