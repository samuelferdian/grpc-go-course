package main

import (
	"fmt"
	pb "grpc-go-course/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}

		log.Printf("receiving: %v\n", req)
		res += fmt.Sprintf("Hello %s!", req.FirstName)
	}
}
