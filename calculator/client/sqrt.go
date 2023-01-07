package main

import (
	"context"
	pb "grpc-go-course/calculator/proto"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("soSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %v\n", e.Message())
			log.Printf("Error code from server: %v\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println(" probably send a negative number")
				return
			}
		} else {
			log.Fatalf("A non grpc error: %v\n", err)
		}
	}

	log.Printf("Sqrt: %f\n", res.Result)
}
