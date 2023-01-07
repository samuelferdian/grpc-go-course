package main

import (
	pb "grpc-go-course/calculator/proto"
	"log"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("primes function was invoked with %v\v", in)

	number := in.Number
	divider := int64(2)

	for number > 1 {
		if number%divider == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divider,
			})

			number /= divider
		} else {
			divider++
		}
	}

	return nil
}
